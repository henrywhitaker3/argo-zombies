package dashboard

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGithubCreateOrUpdateDashboard(t *testing.T) {
	t.Run("creates an issue when no dashboard exists", func(t *testing.T) {
		var createRequest github.IssueRequest
		dashboard := newGithubDashboard(t, func(w http.ResponseWriter, r *http.Request) {
			switch {
			case r.Method == http.MethodGet && r.URL.Path == "/repos/owner/repo/issues":
				writeGithubJSON(t, w, []github.Issue{})
			case r.Method == http.MethodPost && r.URL.Path == "/repos/owner/repo/issues":
				require.NoError(t, json.NewDecoder(r.Body).Decode(&createRequest))
				w.WriteHeader(http.StatusCreated)
				writeGithubJSON(t, w, github.Issue{})
			default:
				http.NotFound(w, r)
			}
		})

		require.NoError(t, dashboard.CreateOrUpdateDashboard("body"))
		require.NotNil(t, createRequest.Title)
		assert.Equal(t, "zombies", *createRequest.Title)
		assert.Equal(t, "body", *createRequest.Body)
		assert.Equal(t, labels, *createRequest.Labels)
	})

	t.Run("updates an existing dashboard", func(t *testing.T) {
		var updateRequest github.IssueRequest
		dashboard := newGithubDashboard(t, func(w http.ResponseWriter, r *http.Request) {
			switch {
			case r.Method == http.MethodGet && r.URL.Path == "/repos/owner/repo/issues":
				writeGithubJSON(
					t,
					w,
					[]github.Issue{
						{
							Number: github.Int(7),
							Title:  github.String("zombies"),
							State:  github.String("open"),
						},
					},
				)
			case r.Method == http.MethodPatch && r.URL.Path == "/repos/owner/repo/issues/7":
				require.NoError(t, json.NewDecoder(r.Body).Decode(&updateRequest))
				writeGithubJSON(t, w, github.Issue{})
			default:
				http.NotFound(w, r)
			}
		})

		require.NoError(t, dashboard.CreateOrUpdateDashboard("updated body"))
		require.NotNil(t, updateRequest.Body)
		assert.Equal(t, "updated body", *updateRequest.Body)
		assert.Equal(t, labels, *updateRequest.Labels)
	})

	t.Run("finds a dashboard on a later page", func(t *testing.T) {
		listRequests := 0
		updated := false
		dashboard := newGithubDashboard(t, func(w http.ResponseWriter, r *http.Request) {
			switch {
			case r.Method == http.MethodGet && r.URL.Path == "/repos/owner/repo/issues" && r.URL.Query().Get("page") == "":
				listRequests++
				w.Header().Set("Link", "<http://"+r.Host+"/repos/owner/repo/issues?page=2>; rel=\"next\"")
				writeGithubJSON(t, w, []github.Issue{{Number: github.Int(1), Title: github.String("another issue")}})
			case r.Method == http.MethodGet && r.URL.Path == "/repos/owner/repo/issues" && r.URL.Query().Get("page") == "2":
				listRequests++
				writeGithubJSON(t, w, []github.Issue{{Number: github.Int(7), Title: github.String("zombies"), State: github.String("open")}})
			case r.Method == http.MethodPatch && r.URL.Path == "/repos/owner/repo/issues/7":
				updated = true
				writeGithubJSON(t, w, github.Issue{})
			default:
				http.NotFound(w, r)
			}
		})

		require.NoError(t, dashboard.CreateOrUpdateDashboard("updated body"))
		assert.Equal(t, 2, listRequests)
		assert.True(t, updated)
	})

	t.Run("returns issue list errors", func(t *testing.T) {
		dashboard := newGithubDashboard(t, func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "unavailable", http.StatusServiceUnavailable)
		})

		assert.Error(t, dashboard.CreateOrUpdateDashboard("body"))
	})
}

func TestGithubOptsAuthType(t *testing.T) {
	assert.Equal(t, githubAuthToken, GithubOpts{Token: "token"}.authType())
	assert.Equal(
		t,
		githubAuthApp,
		GithubOpts{AppClientID: "id", AppInstallationID: 1, AppPrivateKey: "key"}.authType(),
	)
	assert.Empty(t, GithubOpts{AppClientID: "id"}.authType())
}

func newGithubDashboard(t *testing.T, handler http.HandlerFunc) *Github {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	client, err := github.NewEnterpriseClient(server.URL+"/", server.URL+"/", server.Client())
	require.NoError(t, err)

	return &Github{
		ctx:    context.Background(),
		client: client,
		owner:  "owner",
		repo:   "repo",
		title:  "zombies",
	}
}

func writeGithubJSON(t *testing.T, w http.ResponseWriter, value any) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	require.NoError(t, json.NewEncoder(w).Encode(value))
}
