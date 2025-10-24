package dashboard

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-github/github"
	"github.com/jferrl/go-githubauth/v2"
	"golang.org/x/oauth2"
)

type Github struct {
	ctx    context.Context
	client *github.Client
	owner  string
	repo   string
}

type GithubOpts struct {
	Repo              string
	Token             string
	AppClientID       string
	AppInstallationID int
	AppPrivateKey     string
}

func (g GithubOpts) authType() string {
	if g.Token != "" {
		return githubAuthToken
	}
	if g.AppClientID != "" && g.AppInstallationID != 0 && g.AppPrivateKey != "" {
		return githubAuthApp
	}
	return ""
}

const (
	githubAuthToken = "token"
	githubAuthApp   = "app"
)

func NewGithub(ctx context.Context, opts GithubOpts) (*Github, error) {
	var httpClient *http.Client
	switch opts.authType() {
	case githubAuthToken:
		httpClient = githubTokenClient(ctx, opts.Token)
	case githubAuthApp:
		c, err := githubAppClient(ctx, opts.AppClientID, opts.AppInstallationID, opts.AppPrivateKey)
		if err != nil {
			return nil, fmt.Errorf("build app client: %w", err)
		}
		httpClient = c
	default:
		return nil, fmt.Errorf("invalid github auth config")
	}

	client := github.NewClient(httpClient)
	s := strings.SplitN(opts.Repo, "/", 2)

	return &Github{
		ctx:    ctx,
		client: client,
		owner:  s[0],
		repo:   s[1],
	}, nil
}

func githubAppClient(
	ctx context.Context,
	clientID string,
	installationID int,
	privateKey string,
) (*http.Client, error) {
	appTokenSource, err := githubauth.NewApplicationTokenSource(clientID, []byte(privateKey))
	if err != nil {
		return nil, fmt.Errorf("create applicaiton token source: %w", err)
	}

	installationTokenSource := githubauth.NewInstallationTokenSource(
		int64(installationID),
		appTokenSource,
	)

	return oauth2.NewClient(ctx, installationTokenSource), nil
}

func githubTokenClient(ctx context.Context, token string) *http.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	return oauth2.NewClient(ctx, ts)
}

func (g *Github) CreateOrUpdateDashboard(body string) error {
	list, _, err := g.client.Issues.ListByRepo(
		g.ctx,
		g.owner,
		g.repo,
		&github.IssueListByRepoOptions{
			Labels: labels,
		},
	)
	if err != nil {
		return err
	}

	for _, issue := range list {
		if *issue.Title == title {
			return g.updateIssue(issue, body)
		}
	}

	return g.createIssue(body)
}

func (g *Github) createIssue(body string) error {
	issue := &github.IssueRequest{
		Title:  &title,
		Labels: &labels,
		Body:   &body,
	}
	_, _, err := g.client.Issues.Create(g.ctx, g.owner, g.repo, issue)
	return err
}

func (g *Github) updateIssue(issue *github.Issue, body string) error {
	ir := &github.IssueRequest{
		Title:  issue.Title,
		Body:   &body,
		State:  issue.State,
		Labels: &labels,
	}

	_, _, err := g.client.Issues.Edit(g.ctx, g.owner, g.repo, *issue.Number, ir)
	return err
}
