package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type Options struct {
	Query         string `short:"q" long:"query" default:"type:pr is:open" description:"Query strings. For search Issues/Pull Requests."`
	Comment       string `short:"c" long:"comment" default:":game_die:" description:"Comment. Would be posted before assigned."`
	DryRun        bool   `short:"n" long:"dry-run" description:"If true, show candidates and list Issues, without assign."`
	Force         bool   `short:"f" long:"force" description:"If true, reassign even if already assigned."`
	RunOnce       bool   `short:"o" long:"run-once" description:"If true, assign assign only one Issue."`
	AssignAuthor  bool   `short:"a" long:"assign-author" description:"If true, Issue's author also assigns."`
	Limit         int    `short:"l" long:"limit" default:"0" description:"A maximum number of assign Issues."`
	ExemptedUsers string `short:"e" long:"exempted-users" default:"" description:"User names separated by comma who exempt assignee."`
	Debug         bool   `short:"d" long:"debug"`
}

func main() {
	godotenv.Load()
	org := os.Getenv("GITHUB_ORGANIZATION")
	repo := os.Getenv("GITHUB_REPO")
	team := os.Getenv("GITHUB_TEAM")
	token := os.Getenv("GITHUB_ACCESS_TOKEN")

	var opts Options
	p := flags.NewParser(&opts, flags.Default)
	_, err := p.ParseArgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	d := NewDice(strings.Split(opts.ExemptedUsers, ","))
	im := NewIssueManager(org, repo, team, token, opts.DryRun)

	issues, err := im.FindIssues(opts.Query)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	assinedNumber := 0
	for _, issue := range issues {
		if im.IsAlreadyAssignedExpectAuthor(issue) {
			continue
		}
		reviewers, err := im.FindCandidatesOfReviewers(issue)
		if err != nil {
			log.Fatal(err)
			continue
		}
		if opts.Debug {
			log.Println(fmt.Sprintf("Candidates are %s", reviewers))
		}
		assignee := d.Throw(reviewers)

		_, err = im.Assign(issue, assignee, opts.AssignAuthor)
		if err != nil {
			log.Fatal(err)
			continue
		}

		if len(opts.Comment) > 0 && opts.DryRun == false {
			im.Comment(issue, opts.Comment+"@"+assignee)
		}

		assinedNumber++
		if opts.Debug {
			log.Println(fmt.Sprintf("#%d %s %s => author:%s assigned:%s", *issue.Number, *issue.HTMLURL, *issue.Title, *issue.User.Login, assignee))
		}
		if opts.Limit > 0 && opts.Limit <= assinedNumber {
			break
		}
	}
}
