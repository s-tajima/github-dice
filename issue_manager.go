package main

import (
	"errors"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"regexp"
	"strings"
)

type IssueManager struct {
	Client       *github.Client
	Organization string
	Team         string
	Repository   string
	DryRun       bool
}

type Users []*github.User

func NewIssueManager(organization string, repository string, team string, token string, dryRun bool) *IssueManager {

	tc := oauth2.NewClient(oauth2.NoContext, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}))

	im := &IssueManager{}
	im.Client = github.NewClient(tc)
	im.Organization = organization
	im.Repository = repository
	im.Team = team
	im.DryRun = dryRun

	return im
}

func (im *IssueManager) FindIssues(spec string) ([]*github.Issue, error) {
	members, err := im.findUsersByTeamName(im.Team)
	if err != nil {
		return nil, err
	}
	queryString := im.buildQuery(spec)
	searchResult, _, err := im.Client.Search.Issues(queryString, &github.SearchOptions{})
	if err != nil {
		return nil, err
	}

	var targets []*github.Issue
Loop:
	for i, issue := range searchResult.Issues {
		for _, member := range members {
			if *issue.User.Login == *member.Login {
				targets = append(targets, &searchResult.Issues[i])
				continue Loop
			}
		}
	}

	return targets, nil
}

func (im *IssueManager) FindCandidatesOfReviewers(issue *github.Issue) ([]string, error) {
	var users Users
	users, err := im.findUsersByTeamName(im.Team)
	if err != nil {
		return nil, err
	}
	candidates := Users(users.removeUser(issue.User))

	return candidates.GetLoginNames(), nil
}

func (im *IssueManager) Assign(issue *github.Issue, assignee string, assignAuthor bool) (*github.Issue, error) {
	assignees := []string{assignee}
	if assignAuthor {
		assignees = append(assignees, *issue.User.Login)
	}
	if im.DryRun {
		return issue, nil
	}
	i, _, err := im.Client.Issues.AddAssignees(im.Organization, im.Repository, *issue.Number, assignees)

	return i, err
}

func (im *IssueManager) IsAlreadyAssignedExpectAuthor(issue *github.Issue) bool {
	assignUsersExpectAuthor := im.getAssignUsersExpectAuthor(issue)

	return len(assignUsersExpectAuthor) > 0
}

func (im *IssueManager) UnassignUsersExpectAuthor(issue *github.Issue) (*github.Issue, error) {
	users := im.getAssignUsersExpectAuthor(issue)
	var asignees []string
	for _, u := range users {
		asignees = append(asignees, *u.Login)
	}
	i, _, err := im.Client.Issues.RemoveAssignees(im.Organization, im.Repository, *issue.Number, asignees)
	return i, err
}

func (im *IssueManager) getAssignUsersExpectAuthor(issue *github.Issue) []*github.User {
	var assignees Users
	assignees = issue.Assignees
	return assignees.removeUser(issue.User)
}

func (im *IssueManager) Comment(issue *github.Issue, comment string) bool {
	ic := &github.IssueComment{Body: &comment}
	if im.DryRun {
		return true
	}
	_, _, err := im.Client.Issues.CreateComment(im.Organization, im.Repository, *issue.Number, ic)

	return err != nil
}

func (im *IssueManager) findUsersByTeamName(name string) ([]*github.User, error) {
	teams, _, err := im.Client.Repositories.ListTeams(im.Organization, im.Repository, nil)
	if err != nil {
		return nil, err
	}

	for _, t := range teams {
		if *t.Name == name {
			users, _, err := im.Client.Organizations.ListTeamMembers(*t.ID, &github.OrganizationListTeamMembersOptions{})
			return users, err
		}
	}

	return nil, errors.New("team not found")
}

func (users Users) removeUser(user *github.User) []*github.User {
	var candidates []*github.User
	for _, u := range users {
		if *u.Login != *user.Login {
			candidates = append(candidates, u)
		}
	}
	return candidates
}

func (users Users) GetLoginNames() []string {
	var names []string
	for _, u := range users {
		names = append(names, *u.Login)
	}

	return names
}

func (im *IssueManager) buildQuery(spec string) string {
	queries := regexp.MustCompile(" +").Split(spec, -1)
	queries = append(queries, "repo:"+im.Organization+"/"+im.Repository)

	return strings.Join(queries, " ")
}
