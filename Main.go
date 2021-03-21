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

// printShuffledList prints a team's name and a randomized list of its members.
func printShuffledList(teamMembers []string, teamName string) {

	rand.Shuffle(len(teamMembers), func(i, j int) {
		teamMembers[i], teamMembers[j] = teamMembers[j], teamMembers[i]
	})

	fmt.Printf("## %s\n", teamName)

	for _, name := range teamMembers {
		fmt.Println(name)
	}
}
