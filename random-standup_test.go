package main

import (
	"fmt"
	"testing"
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
