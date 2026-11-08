package detector

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/henrywhitaker3/argo-zombies/internal/config"
	"github.com/henrywhitaker3/argo-zombies/internal/exclusions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

func TestProcessGroupCollectsOnlyUnmanagedResources(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/v1/pods":
			writeJSON(t, w, map[string]any{
				"apiVersion": "v1",
				"kind":       "PodList",
				"items": []map[string]any{
					{
						"apiVersion": "v1",
						"kind":       "Pod",
						"metadata": map[string]any{
							"name":      "managed",
							"namespace": "default",
							"labels": map[string]string{
								"argocd.argoproj.io/instance": "app",
							},
						},
					},
					{
						"apiVersion": "v1",
						"kind":       "Pod",
						"metadata": map[string]any{
							"name":      "unmanaged",
							"namespace": "default",
						},
					},
				},
			})
		case "/api/v1/services":
			http.Error(w, "unavailable", http.StatusServiceUnavailable)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	config.Cfg = &config.Config{Exclusions: exclusions.Exclusions{}}
	t.Cleanup(func() { config.Cfg = nil })

	dynamicClient, err := dynamic.NewForConfig(&rest.Config{Host: server.URL})
	require.NoError(t, err)

	collection := NewCollection()
	(&Detector{}).processGroup(context.Background(), processGroupOpts{
		dynClient: dynamicClient,
		col:       collection,
		group: &metav1.APIResourceList{
			GroupVersion: "v1",
			APIResources: []metav1.APIResource{
				{Name: "pods", Verbs: []string{"list"}},
				{Name: "services", Verbs: []string{"list"}},
				{Name: "configmaps", Verbs: []string{"get"}},
				{Name: "nodes", Verbs: []string{"list"}},
			},
		},
	})

	require.Len(t, collection.Items, 1)
	assert.Equal(t, "unmanaged", collection.Items[0].GetName())
}

func TestCollectionPushIsSafeForConcurrentUse(t *testing.T) {
	collection := NewCollection()
	var waitGroup sync.WaitGroup

	for range 100 {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			collection.Push(baseTestItem())
		}()
	}

	waitGroup.Wait()
	assert.Len(t, collection.Items, 100)
}

func writeJSON(t *testing.T, w http.ResponseWriter, value any) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	require.NoError(t, json.NewEncoder(w).Encode(value))
}

func baseTestItem() unstructured.Unstructured {
	return unstructured.Unstructured{}
}
