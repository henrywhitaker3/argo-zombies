package dashboard

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gitlab "gitlab.com/gitlab-org/api/client-go/v3"
)

func TestGitlabCreateOrUpdateDashboard(t *testing.T) {
	t.Run("creates an issue when no dashboard exists", func(t *testing.T) {
		created := false
		dashboard := newGitlabDashboard(t, func(w http.ResponseWriter, r *http.Request) {
			switch {
			case r.Method == http.MethodGet && r.URL.Path == "/api/v4/projects":
				writeGitlabJSON(t, w, []map[string]any{{"id": 7, "name": "repo"}})
			case r.Method == http.MethodGet && r.URL.Path == "/api/v4/projects/7/issues":
				writeGitlabJSON(t, w, []map[string]any{})
			case r.Method == http.MethodPost && r.URL.Path == "/api/v4/projects/7/issues":
				created = true
				writeGitlabJSON(t, w, map[string]any{"id": 1})
			default:
				http.NotFound(w, r)
			}
		})

		require.NoError(t, dashboard.CreateOrUpdateDashboard("body"))
		assert.True(t, created)
	})

	t.Run("updates an existing open dashboard", func(t *testing.T) {
		updated := false
		dashboard := newGitlabDashboard(t, func(w http.ResponseWriter, r *http.Request) {
			switch {
			case r.Method == http.MethodGet && r.URL.Path == "/api/v4/projects":
				writeGitlabJSON(t, w, []map[string]any{{"id": 7, "name": "repo"}})
			case r.Method == http.MethodGet && r.URL.Path == "/api/v4/projects/7/issues":
				writeGitlabJSON(
					t,
					w,
					[]map[string]any{
						{"id": 1, "iid": 3, "project_id": 7, "title": "zombies", "state": "opened"},
					},
				)
			case r.Method == http.MethodPut && r.URL.Path == "/api/v4/projects/7/issues/3":
				updated = true
				writeGitlabJSON(t, w, map[string]any{"id": 1})
			default:
				http.NotFound(w, r)
			}
		})

		require.NoError(t, dashboard.CreateOrUpdateDashboard("updated body"))
		assert.True(t, updated)
	})

	t.Run("returns an error when the project is unavailable", func(t *testing.T) {
		dashboard := newGitlabDashboard(t, func(w http.ResponseWriter, r *http.Request) {
			writeGitlabJSON(t, w, []map[string]any{})
		})

		assert.EqualError(
			t,
			dashboard.CreateOrUpdateDashboard("body"),
			"could not access repo repo",
		)
	})
}

func newGitlabDashboard(t *testing.T, handler http.HandlerFunc) *Gitlab {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	client, err := gitlab.NewClient(
		"token",
		gitlab.WithBaseURL(server.URL),
		gitlab.WithHTTPClient(server.Client()),
	)
	require.NoError(t, err)

	return &Gitlab{
		ctx:    context.Background(),
		client: client,
		repo:   "repo",
		title:  "zombies",
	}
}

func writeGitlabJSON(t *testing.T, w http.ResponseWriter, value any) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	require.NoError(t, json.NewEncoder(w).Encode(value))
}
