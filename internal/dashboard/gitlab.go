package dashboard

import (
	"context"
	"fmt"
	"strings"

	gitlab "gitlab.com/gitlab-org/api/client-go"
)

type Gitlab struct {
	ctx    context.Context
	client *gitlab.Client
	owner  string
	repo   string
	title  string
}

type GitlabOpts struct {
	Title string
	Repo  string
	Token string
}

func NewGitlab(ctx context.Context, opts GitlabOpts) (*Gitlab, error) {
	client, err := gitlab.NewClient(opts.Token)
	if err != nil {
		return nil, err
	}

	s := strings.SplitN(opts.Repo, "/", 2)

	return &Gitlab{
		ctx:    ctx,
		client: client,
		owner:  s[0],
		repo:   s[1],
		title:  opts.Title,
	}, nil
}

func (g *Gitlab) CreateOrUpdateDashboard(body string) error {
	var proj *gitlab.Project
	projs, _, err := g.client.Projects.ListProjects(&gitlab.ListProjectsOptions{
		Membership: gitlab.Ptr(true),
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
		if issue.Title == g.title && issue.State != "closed" {
			return g.updateIssue(issue, body)
		}
	}

	return g.createIssue(proj, body)
}

func (g *Gitlab) createIssue(proj *gitlab.Project, body string) error {
	_, _, err := g.client.Issues.CreateIssue(proj.ID, &gitlab.CreateIssueOptions{
		Title:       &g.title,
		Labels:      (*gitlab.LabelOptions)(&labels),
		Description: &body,
	})
	return err
}

func (g *Gitlab) updateIssue(issue *gitlab.Issue, body string) error {
	_, _, err := g.client.Issues.UpdateIssue(issue.ProjectID, issue.IID, &gitlab.UpdateIssueOptions{
		Title:       &g.title,
		Description: &body,
		Labels:      (*gitlab.LabelOptions)(&labels),
	})
	return err
}
