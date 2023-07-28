package dashboard

import (
	"context"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Github struct {
	ctx    context.Context
	client *github.Client
	owner  string
	repo   string
}

func NewGithub(ctx context.Context, repo string, token string) *Github {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	s := strings.SplitN(repo, "/", 2)

	return &Github{
		ctx:    ctx,
		client: client,
		owner:  s[0],
		repo:   s[1],
	}
}

func (g *Github) CreateOrUpdateDashboard() error {
	list, _, err := g.client.Issues.ListByRepo(g.ctx, g.owner, g.repo, &github.IssueListByRepoOptions{
		Labels: labels,
	})
	if err != nil {
		return err
	}

	for _, issue := range list {
		if *issue.Title == title {
			return g.updateIssue(issue)
		}
	}

	return g.createIssue()
}

func (g *Github) createIssue() error {
	issue := &github.IssueRequest{
		Title:  &title,
		Labels: &labels,
	}
	_, _, err := g.client.Issues.Create(g.ctx, g.owner, g.repo, issue)
	return err
}

func (g *Github) updateIssue(issue *github.Issue) error {
	body := "bongo"
	ir := &github.IssueRequest{
		Title:  issue.Title,
		Body:   &body,
		State:  issue.State,
		Labels: &labels,
	}

	_, _, err := g.client.Issues.Edit(g.ctx, g.owner, g.repo, *issue.Number, ir)
	return err
}
