package detector

import (
	"context"
	"fmt"
	"sync"

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

func (d *Detector) Detect(ctx context.Context) error {
	col, err := d.getResources(ctx)
	if err != nil {
		return err
	}

	for name := range col.Items {
		fmt.Println(name)
	}

	return nil
}

func (d *Detector) getResources(ctx context.Context) (*Collection, error) {
	collection := NewCollection()
	wg := sync.WaitGroup{}

	filters := BuildFilters()

	_, list, err := d.client.ServerGroupsAndResources()
	if err != nil {
		return nil, err
	}

	for _, group := range list {
		wg.Add(1)

		go func(group *metav1.APIResourceList) error {
			gv, err := schema.ParseGroupVersion(group.GroupVersion)
			if err != nil {
				return err
			}

			wg.Add(1)
			go func() {
			rl:
				for _, resource := range group.APIResources {
					gvr := schema.GroupVersionResource{
						Group:    gv.Group,
						Version:  gv.Version,
						Resource: resource.Name,
					}

					for _, skip := range getSkipList() {
						if gvr == skip {
							fmt.Printf("skipping %s/%s/%s", gvr.Group, gvr.Version, gvr.Resource)
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
						// TODO: find out how to ignore reource when it's from a newer one
						for _, filter := range filters {
							if filter(item) {
								// fmt.Printf("filtering item %s/%s/%s\n", item.GetAPIVersion(), item.GetKind(), item.GetName())
								continue il
							}
						}
						// fmt.Printf("zombie item %s/%s %s/%s/%s\n", gv.Group, gv.Version, item.GetAPIVersion(), item.GetKind(), item.GetName())
						collection.Push(&item)
					}
				}
				wg.Done()
			}()

			wg.Done()

			return nil
		}(group)
	}

	wg.Wait()

	return collection, nil
}
