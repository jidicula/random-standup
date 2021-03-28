package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/pelletier/go-toml"
)

func TestShuffleTeam(t *testing.T) {

	var tests = []struct {
		teamMembers []string
		teamName    string
		want        string
	}{
		{[]string{"Alice", "Bob", "Carol", "David"}, "Subteam 1", "## Subteam 1\nAlice\nBob\nDavid\nCarol\n"},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("(%s, %s)", tt.teamMembers, tt.teamName)
		t.Run(testname, func(t *testing.T) {
			output := shuffleTeam(tt.teamMembers, tt.teamName)
			if output != tt.want {
				t.Errorf("got %s, want %s", output, tt.want)
			}
		})
	}
}

func TestGetSortedKeys(t *testing.T) {

	emptySubteams, _ := toml.Load(`
["subteam 1"]

[subteam-2]

["subteam 3"]`)

	mixedSubteams, _ := toml.Load(`
["subteam 1"]
members = []
[subteam-2]
members = ["Alice", "Bob"]
["subteam 3"]`)

	var tests = []struct {
		roster *toml.Tree
		want   []string
	}{

		{emptySubteams, []string{"subteam 1", "subteam-2", "subteam 3"}},
		{mixedSubteams, []string{"subteam 1", "subteam-2", "subteam 3"}},
	}

	for _, tt := range tests {
		testname := tt.roster.String()
		t.Run(testname, func(t *testing.T) {
			output := getSortedKeys(tt.roster)
			if !reflect.DeepEqual(output, tt.want) {
				t.Errorf("got %s, want %s", output, tt.want)
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

	var tests = []struct {
		roster string
		want   string
	}{
		{singleSubteam, `## Subteam-1
Alice
Carol
David
Bob
`},
		{subteamPair, `## Subteam-1
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
		{lastSubteamEmpty, `## Subteam-1
Carol
David
Bob
Alice

## Subteam 2
Heidi
Grace
Erin
Frank

`}, // trailing newline is workaround
		{middleSubteamEmpty, `## Subteam-1
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
	}

	for _, tt := range tests {
		testname := tt.roster

		t.Run(testname, func(t *testing.T) {
			rosterTree, _ := toml.Load(tt.roster)
			output := standupOrder(rosterTree)
			if output != tt.want {
				t.Errorf("got %s, want %s", output, tt.want)
			}
		})
	}

}
