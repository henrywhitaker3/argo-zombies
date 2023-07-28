package dashboard

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Github struct {
	client *github.Client
}

func NewGithub(ctx context.Context, token string) *Github {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return &Github{
		client: client,
	}
}
