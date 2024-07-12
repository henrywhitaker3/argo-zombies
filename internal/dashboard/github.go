package dashboard

import (
	"context"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Github struct {
	client *github.Client
	owner  string
	repo   string
}

func NewGithub(repo string, token string) *Github {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)

	s := strings.SplitN(repo, "/", 2)

	return &Github{
		client: client,
		owner:  s[0],
		repo:   s[1],
	}
}

func (g *Github) CreateOrUpdateDashboard(ctx context.Context, body string) error {
	list, _, err := g.client.Issues.ListByRepo(ctx, g.owner, g.repo, &github.IssueListByRepoOptions{
		Labels: labels,
	})
	if err != nil {
		return err
	}

	for _, issue := range list {
		if *issue.Title == title {
			return g.updateIssue(ctx, issue, body)
		}
	}

	return g.createIssue(ctx, body)
}

func (g *Github) createIssue(ctx context.Context, body string) error {
	issue := &github.IssueRequest{
		Title:  &title,
		Labels: &labels,
		Body:   &body,
	}
	_, _, err := g.client.Issues.Create(ctx, g.owner, g.repo, issue)
	return err
}

func (g *Github) updateIssue(ctx context.Context, issue *github.Issue, body string) error {
	ir := &github.IssueRequest{
		Title:  issue.Title,
		Body:   &body,
		State:  issue.State,
		Labels: &labels,
	}

	_, _, err := g.client.Issues.Edit(ctx, g.owner, g.repo, *issue.Number, ir)
	return err
}
