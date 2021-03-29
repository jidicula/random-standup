package main

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/pelletier/go-toml"
)

func TestShuffleTeam(t *testing.T) {

	tests := map[string]struct {
		teamMembers []string
		teamName    string
		want        string
	}{
		"four names": {[]string{"Alice", "Bob", "Carol", "David"}, "Subteam 1", "## Subteam 1\nAlice\nBob\nDavid\nCarol\n"},
	}

	for name, tt := range tests {

		t.Run(name, func(t *testing.T) {
			rand.Seed(0)
			got := shuffleTeam(tt.teamMembers, tt.teamName)
			if got != tt.want {
				t.Errorf("%s: got %s, want %s", name, got, tt.want)
			}
		})
	}
}

func TestGetSortedKeysWithMembers(t *testing.T) {

	emptySubteams, _ := toml.Load(`
["subteam 1"]

[subteam-2]

["subteam 3"]`)

	var emptyStringSlice []string

	mixedSubteams, _ := toml.Load(`
["subteam 1"]
members = []
[subteam-2]
members = ["Alice", "Bob"]
["subteam 3"]`)

	tests := map[string]struct {
		roster *toml.Tree
		want   []string
	}{

		"empty subteams": {emptySubteams, emptyStringSlice},
		"empty memberlist, full subteam, empty subteam": {mixedSubteams, []string{"subteam 1", "subteam-2"}},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := getSortedKeysWithMembers(tt.roster)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s: got %s, want %s", name, got, tt.want)
			}
		})
	}
}

func TestStandupOrder(t *testing.T) {
	singleSubteam := `
[Subteam-1]
members = [
        "Alice",
        "Bob",
        "Carol",
        "David"
        ]
`
	subteamPair := singleSubteam + "\n" + `["Subteam 2"]
members = ["Erin", "Frank", "Grace", "Heidi"]`

	lastSubteamEmpty := subteamPair + "\n" + `

["Empty Subteam"]
`
	middleSubteamEmpty := singleSubteam + "\n" + `

["Empty Subteam"]
` + "\n" + `["Subteam 2"]
members = ["Erin", "Frank", "Grace", "Heidi"]`

	onlyEmptySubteam := `["Empty Subteam"]`

	tests := map[string]struct {
		roster string
		want   string
	}{
		"1 subteam": {singleSubteam, `## Subteam-1
Alice
Carol
David
Bob
`},
		"2 subteams": {subteamPair, `## Subteam-1
Bob
Carol
David
Alice

## Subteam 2
Erin
Grace
Heidi
Frank
`},
		"2 full subteams, empty subteam": {lastSubteamEmpty, `## Subteam-1
Carol
David
Bob
Alice

## Subteam 2
Heidi
Grace
Erin
Frank
`},
		"1 empty subteam between 2 full subteams": {middleSubteamEmpty, `## Subteam-1
Bob
David
Alice
Carol

## Subteam 2
Erin
Grace
Heidi
Frank
`},
		"1 empty subteam": {onlyEmptySubteam, ``},
	}

	for name, tt := range tests {
		rand.Seed(0)
		t.Run(name, func(t *testing.T) {
			rosterTree, _ := toml.Load(tt.roster)
			got := standupOrder(rosterTree)
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}

}
