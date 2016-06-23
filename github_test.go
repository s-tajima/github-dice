package main

import (
	"github.com/google/go-github/github"
	"testing"
)

func TestfindTeamByName_found(t *testing.T) {
	teams := []*github.Team{}

	team1ID := 1
	team2ID := 2
	team1Name := "team1"
	team2Name := "team2"

	team1 := github.Team{ID: &team1ID, Name: &team1Name}
	team2 := github.Team{ID: &team2ID, Name: &team2Name}

	teams = append(teams, &team1)
	teams = append(teams, &team2)

	r := findTeamByName("team1", teams)
	if *r.ID != 1 {
		t.Error("team1 should have ID:1")
	}
}

func TestfindTeamByName_notfound(t *testing.T) {
	teams := []*github.Team{}

	team1ID := 1
	team2ID := 2
	team1Name := "team1"
	team2Name := "team2"

	team1 := github.Team{ID: &team1ID, Name: &team1Name}
	team2 := github.Team{ID: &team2ID, Name: &team2Name}

	teams = append(teams, &team1)
	teams = append(teams, &team2)

	r := findTeamByName("team3", teams)
	if r != nil {
		t.Error("no team should be found.")
	}
}
