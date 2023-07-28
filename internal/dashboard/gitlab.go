package dashboard

import (
	"context"
	"fmt"
	"strings"

	"github.com/xanzy/go-gitlab"
)

type Gitlab struct {
	ctx    context.Context
	client *gitlab.Client
	owner  string
	repo   string
}

func NewGitlab(ctx context.Context, repo string, token string) (*Gitlab, error) {
	client, err := gitlab.NewClient(token)
	if err != nil {
		return nil, err
	}

	s := strings.SplitN(repo, "/", 2)

	return &Gitlab{
		ctx:    ctx,
		client: client,
		owner:  s[0],
		repo:   s[1],
	}, nil
}

func (g *Gitlab) CreateOrUpdateDashboard(body string) error {
	var proj *gitlab.Project
	projs, _, err := g.client.Projects.ListProjects(&gitlab.ListProjectsOptions{
		Membership: gitlab.Bool(true),
	})
	if err != nil {
		return err
	}

	for _, project := range projs {
		if project.Name == g.repo {
			proj = project
		}
	}

	if proj == nil {
		return fmt.Errorf("could not access repo %s", g.repo)
	}

	issues, _, err := g.client.Issues.ListProjectIssues(proj.ID, &gitlab.ListProjectIssuesOptions{})
	if err != nil {
		return err
	}
	for _, issue := range issues {
		if issue.Title == title {
			return g.updateIssue(issue, body)
		}
	}

	return g.createIssue(proj, body)
}

func (g *Gitlab) createIssue(proj *gitlab.Project, body string) error {
	_, _, err := g.client.Issues.CreateIssue(proj.ID, &gitlab.CreateIssueOptions{
		Title:       &title,
		Labels:      (*gitlab.Labels)(&labels),
		Description: &body,
	})
	return err
}

func (g *Gitlab) updateIssue(issue *gitlab.Issue, body string) error {
	return nil
}
