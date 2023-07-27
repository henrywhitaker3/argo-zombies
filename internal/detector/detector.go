package detector

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/utils/strings/slices"
)

type Detector struct {
	client    *kubernetes.Clientset
	dynClient *dynamic.DynamicClient
}

func NewDetector(client *kubernetes.Clientset, dynClient *dynamic.DynamicClient) *Detector {
	return &Detector{
		client:    client,
		dynClient: dynClient,
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
	collection := NewCollection()

	filters := BuildFilters()

	_, list, err := d.client.ServerGroupsAndResources()
	if err != nil {
		return nil, err
	}

	for _, group := range list {
		gv, err := schema.ParseGroupVersion(group.GroupVersion)
		if err != nil {
			return nil, err
		}

	rl:
		for _, resource := range group.APIResources {
			gvr := schema.GroupVersionResource{
				Group:    gv.Group,
				Version:  gv.Version,
				Resource: resource.Name,
			}

			for _, skip := range getSkipList() {
				if gvr == skip {
					// fmt.Printf("skipping %s/%s/%s", gvr.Group, gvr.Version, gvr.Resource)
					continue rl
				}
			}

			if !slices.Contains(resource.Verbs, "list") {
				continue
			}
			// fmt.Printf("processing resource %s/%s/%s\n", gv.Group, gv.Version, resource.Name)

			restAPI := d.dynClient.Resource(gvr).Namespace("")

			gvrList, err := restAPI.List(ctx, metav1.ListOptions{})
			if err != nil {
				continue
			}

		il:
			for _, item := range gvrList.Items {
				for _, filter := range filters {
					if filter(item) {
						continue il
					}
				}
				collection.Push(item)
			}
		}
	}

	return collection, nil
}
