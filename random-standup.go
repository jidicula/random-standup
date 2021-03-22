// random-standup
//     Copyright (C) 2021  Johanan Idicula
//
//     This program is free software: you can redistribute it and/or modify
//     it under the terms of the GNU Affero General Public License as published
//     by the Free Software Foundation, either version 3 of the License, or
//     (at your option) any later version.
//
//     This program is distributed in the hope that it will be useful,
//     but WITHOUT ANY WARRANTY; without even the implied warranty of
//     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//     GNU Affero General Public License for more details.
//
//     You should have received a copy of the GNU Affero General Public License
//     along with this program.  If not, see <https://www.gnu.org/licenses/>.

/*

random-standup is a tool for randomizing the order of team member updates in a
standup meeting.

Usage:

    random-standup <roster>

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

2. Run the command on the roster file:
    $ random-standup example-roster.toml
    # 2021-03-21
    ## Subteam-1
    David
    Bob
    Carol
    Alice

    ## Subteam 2
    Frank
    Erin
    Heidi
    Grace
*/
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/pelletier/go-toml"
)

func main() {
	file := os.Args[1]

	roster, err := toml.LoadFile(file)
	if err != nil {
		fmt.Printf("Error %s\n", err.Error())
		os.Exit(1)
	}

	now := time.Now()
	rand.Seed(now.UnixNano())

	fmt.Printf("# %s\n", now.Format("2006-01-02"))

	subteams := roster.Keys()
	for i, subteam := range subteams {
		members := roster.GetArray(subteam + ".members").([]string)
		printShuffledList([]string(members), subteam)
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
