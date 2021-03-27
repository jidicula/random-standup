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
		os.Exit(1)
	}

	now := time.Now()
	rand.Seed(now.UnixNano())

	fmt.Printf("# %s\n", now.Format("2006-01-02"))

	subteams := roster.Keys()
	// Tree key order is not guaranteed, so slice of keys has to be sorted
	// by key position in TOML
	sort.Slice(subteams, func(i, j int) bool {
		return roster.GetPosition(subteams[i]).Line < roster.GetPosition(subteams[j]).Line
	})
	for i, subteam := range subteams {
		members := roster.GetArray(subteam + ".members")
		if members == nil {
			continue
		}
		printShuffledList(members.([]string), subteam)
		if i != len(subteams)-1 {
			fmt.Println()
		}
	}
}

// printShuffledList accepts a team's member list and name and prints a team's
// name and a randomized list of its members to stdout.
func printShuffledList(teamMembers []string, teamName string) {

	rand.Shuffle(len(teamMembers), func(i, j int) {
		teamMembers[i], teamMembers[j] = teamMembers[j], teamMembers[i]
	})

	fmt.Printf("## %s\n", teamName)

	for _, name := range teamMembers {
		fmt.Println(name)
	}
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
