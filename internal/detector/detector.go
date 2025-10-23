package detector

import (
	"context"

	"github.com/henrywhitaker3/argo-zombies/internal/config"
	"github.com/henrywhitaker3/argo-zombies/internal/logger"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/utils/strings/slices"
)

type Detector struct {
	client    *kubernetes.Clientset
	dynClient *dynamic.DynamicClient
	cfg       *config.Config
}

func NewDetector(client *kubernetes.Clientset, dynClient *dynamic.DynamicClient, cfg *config.Config) *Detector {
	return &Detector{
		client:    client,
		dynClient: dynClient,
		cfg:       cfg,
	}
}

func (d *Detector) Detect(ctx context.Context) (*Collection, error) {
	col, err := d.getResources(ctx)
	if err != nil {
		return nil, err
	}

	return col, nil
}

func (d *Detector) getResources(ctx context.Context) (*Collection, error) {
	logger := logger.Logger(ctx)

	collection := NewCollection()

	filters := BuildFilters(d.cfg)

	_, list, err := d.client.ServerGroupsAndResources()
	if err != nil {
		return nil, err
	}

	for _, group := range list {
		gv, err := schema.ParseGroupVersion(group.GroupVersion)
		if err != nil {
			return nil, err
		}

		gl := logger.With("group", gv.Group, "version", gv.Version)
		gl.Debugw("processing group")

	rl:
		for _, resource := range group.APIResources {
			gvr := schema.GroupVersionResource{
				Group:    gv.Group,
				Version:  gv.Version,
				Resource: resource.Name,
			}
			gvrl := gl.With("resource", gvr.Resource)

			for _, skip := range getSkipList(d.cfg) {
				if gvr == skip {
					gvrl.Debug("skipping")
					continue rl
				}
			}

			if !slices.Contains(resource.Verbs, "list") {
				continue
			}
			gvrl.Debugw("processing resource")

			restAPI := d.dynClient.Resource(gvr).Namespace("")

			gvrList, err := restAPI.List(ctx, metav1.ListOptions{})
			if err != nil {
				gvrl.Errorw("failed to list resources", "error", err)
				continue
			}

		il:
			for _, item := range gvrList.Items {
				for _, filter := range filters {
					if filter(item) {
						gvrl.Debugw("item filtered out", "name", item.GetName(), "namespace", item.GetNamespace())
						continue il
					}
				}
				collection.Push(item)
			}
		}
	}

	return collection, nil
}
