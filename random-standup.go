// random-standup is a tool for randomizing the order of team member updates in
// a standup meeting.
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/pelletier/go-toml"
)

var usage = `Usage: random-standup [options...] <roster TOML>

random-standup is a tool for randomizing the order of team member updates in a
standup meeting.

Example:

1. Create a team roster in a TOML file:
    # example-roster.toml
    [Subteam-1]
    members = [
            "Alice",                # TOML spec allows whitespace to break arrays
            "Bob",
            "Carol",
            "David"
            ]

    ["Subteam 2"]                   # Keys can have whitespace in quoted strings
    members = ["Erin", "Frank", "Grace", "Heidi"]

    ["Empty Subteam"]               # Subteam with 0 members won't be printed

    ["Subteam 3"]
    members = [
            "Ivan",
            "Judy",
            "Mallory",
            "Niaj"
    ]

2. Run the command on the roster file:
    $ random-standup example-roster.toml
    # 2021-03-27
    ## Subteam-1
    Alice
    David
    Bob
    Carol

    ## Subteam 2
    Grace
    Heidi
    Frank
    Erin

    ## Subteam 3
    Judy
    Niaj
    Ivan
    Mallory

Options:
  --help (-h) Prints this message.
`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s\n", usage)

	}
	flag.Parse()
	if flag.NArg() < 1 {
		usageAndExit("")
	}

	file := flag.Arg(0)

	roster, err := toml.LoadFile(file)
	if err != nil {
		fmt.Printf("Error %s\n", err.Error())
		os.Exit(2)
	}

	now := time.Now()
	rand.Seed(now.UnixNano())

	fmt.Printf("# %s\n", now.Format("2006-01-02"))
	fmt.Printf("%s", standupOrder(roster))
}

// standupOrder returns the randomized standup order from a toml.Tree.
func standupOrder(roster *toml.Tree) string {
	var order string
	subteams := getSortedKeysWithMembers(roster)
	for i, subteam := range subteams {
		members := roster.GetArray(subteam + ".members")
		shuffledTeam := shuffleTeam(members.([]string), subteam)
		order += shuffledTeam

		if i != len(subteams)-1 {
			order += "\n"
		}
	}
	return order
}

// getSortedKeysWithMembers returns a slice of keys with members subkey from
// the TOML sorted by their position in the TOML.
func getSortedKeysWithMembers(roster *toml.Tree) []string {
	subteams := roster.Keys()
	// Tree key order is not guaranteed, so slice of keys has to be
	// explicitly sorted
	sort.Slice(subteams, func(i, j int) bool {
		return roster.GetPosition(subteams[i]).Line <
			roster.GetPosition(subteams[j]).Line
	})
	var cleanSubteams []string
	for _, name := range subteams {
		if roster.GetArray(name+".members") != nil {
			cleanSubteams = append(cleanSubteams, name)
		}
	}
	return cleanSubteams
}

// shuffleTeam accepts a team's member list and name and returns the
// shuffled, stringified team list beginning with the team name.
func shuffleTeam(teamMembers []string, teamName string) string {
	list := ""
	rand.Shuffle(len(teamMembers), func(i, j int) {
		teamMembers[i], teamMembers[j] = teamMembers[j], teamMembers[i]
	})
	list = fmt.Sprintf("## %s\n", teamName)

	for _, name := range teamMembers {
		list += name + "\n"
	}
	return list
}

// usageAndExit prints usage string and exits with nonzero code.
func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, "%s\n", msg)
		fmt.Fprintf(os.Stderr, "\n\n")
	}

	flag.Usage()
	// fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}
