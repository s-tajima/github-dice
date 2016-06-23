package main

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/go-github/github"
)

func findTeamByName(name string, t []*github.Team) *github.Team {
	for i, v := range t {
		if name == *v.Name {
			return t[i]
		}
	}
	return new(github.Team)
}

func selectMember(members []*github.User) *github.User {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(members))
	return members[r]
}

func joinUsers(users []*github.User) (s string) {
	var ss []string
	for _, v := range users {
		ss = append(ss, *v.Login)
	}
	s = strings.Join(ss, ", ")
	return
}
