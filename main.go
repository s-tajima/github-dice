package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

type Dice struct {
	Opts Options

	githubToken        string
	githubOrganization string
	githubTeam         string
	githubRepo         string
}

type Options struct {
	Query   string `short:"q" long:"query" default:"is:issue" description:"query strings for search issue/pull-request."`
	DryRun  bool   `short:"n" long:"dry-run" description:"show candidates and list issues, without assign."`
	Force   bool   `short:"f" long:"force" description:"if true, reassign even if already assigned."`
	RunOnce bool   `short:"o" long:"run-once" description:"if true, assign just once issue."`
	Debug   bool   `short:"d" long:"debug"`
}

func (d *Dice) initialize(args []string) {
	godotenv.Load()

	p := flags.NewParser(&d.Opts, flags.Default)

	_, err := p.ParseArgs(args)
	if err != nil {
		os.Exit(1)
	}

	d.githubToken = os.Getenv("GITHUB_ACCESS_TOKEN")
	d.githubOrganization = os.Getenv("GITHUB_ORGANIZATION")
	d.githubTeam = os.Getenv("GITHUB_TEAM")
	d.githubRepo = os.Getenv("GITHUB_REPO")
}

func (d *Dice) run() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: d.githubToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	t, _, err := client.Organizations.ListTeams(d.githubOrganization, &github.ListOptions{PerPage: 100})
	if err != nil {
		d.log("List teams failed.")
		d.log(err.Error())
		os.Exit(1)
	}

	tt := findTeamByName(d.githubTeam, t)

	members, _, err := client.Organizations.ListTeamMembers(*tt.ID, &github.OrganizationListTeamMembersOptions{})
	if err != nil {
		d.log("List team members failed.")
		d.log(err.Error())
		os.Exit(1)
	}
	if members == nil {
		d.log("No body in this team.")
		os.Exit(2)
	}
	d.log("Candidates are " + joinUsers(members))

	addQuery := d.Opts.Query
	q := "repo:%s/%s " + addQuery
	r, _, err := client.Search.Issues(fmt.Sprintf(q, d.githubOrganization, d.githubRepo), &github.SearchOptions{})
	if err != nil {
		d.log("Search issue failed.")
		os.Exit(1)
	}
	if r == nil {
		d.log("No issue matched.")
		os.Exit(2)
	}

	for _, i := range r.Issues {
		if !d.Opts.Force && i.Assignee != nil {
			d.log(fmt.Sprintf("Skip already assigned isssue #%d (%s) ", *i.Number, *i.Title))
			continue
		}

		u := selectMember(members)

		if !d.Opts.DryRun {
			client.Issues.Edit(d.githubOrganization, d.githubRepo, *i.Number, &github.IssueRequest{Assignee: u.Login})
		}

		dryrun := ""
		if d.Opts.DryRun {
			dryrun = "(dryrun) "
		}

		d.log(fmt.Sprintf("%sAssigned %s on #%d (%s)", dryrun, *u.Login, *i.Number, *i.Title))

		if d.Opts.RunOnce {
			break
		}
	}
}

func (d *Dice) log(str string) {
	log.Println(str)
}

func main() {
	d := &Dice{}
	d.initialize(os.Args[1:])
	d.run()
}
