package detector

import (
	"context"
	"fmt"
	"runtime"
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/utils/strings/slices"
)

type dynClientFunc func() (*dynamic.DynamicClient, error)

type Detector struct {
	client    *kubernetes.Clientset
	dynClient dynClientFunc
}

type DetectorOpts struct {
	Client    *kubernetes.Clientset
	DynClient dynClientFunc
}

func NewDetector(opts DetectorOpts) *Detector {
	return &Detector{
		client:    opts.Client,
		dynClient: opts.DynClient,
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

	groupCh := make(chan *metav1.APIResourceList, 100)

	wg := &sync.WaitGroup{}
	if err := d.setupHandlers(ctx, wg, collection, groupCh); err != nil {
		return nil, fmt.Errorf("setup detector handler: %w", err)
	}

	_, list, err := d.client.ServerGroupsAndResources()
	if err != nil {
		return nil, err
	}

	for _, group := range list {
		wg.Add(1)
		groupCh <- group
	}

	wg.Wait()

	return collection, nil
}

func (d *Detector) setupHandlers(
	ctx context.Context,
	wg *sync.WaitGroup,
	col *Collection,
	ch <-chan *metav1.APIResourceList,
) error {
	for range runtime.NumCPU() {
		dynClient, err := d.dynClient()
		if err != nil {
			return fmt.Errorf("build detector dynamic client: %w", err)
		}
		go d.handleGroups(ctx, handleGroupOpts{
			dynClient: dynClient,
			col:       col,
			wg:        wg,
			ch:        ch,
		})
	}
	return nil
}

type handleGroupOpts struct {
	dynClient *dynamic.DynamicClient
	col       *Collection
	wg        *sync.WaitGroup
	ch        <-chan *metav1.APIResourceList
}

func (d *Detector) handleGroups(ctx context.Context, opts handleGroupOpts) {
	for {
		select {
		case <-ctx.Done():
			return
		case group := <-opts.ch:
			d.processGroup(ctx, processGroupOpts{
				dynClient: opts.dynClient,
				col:       opts.col,
				group:     group,
			})
			opts.wg.Done()
		}
	}
}

type processGroupOpts struct {
	dynClient *dynamic.DynamicClient
	col       *Collection
	group     *metav1.APIResourceList
}

func (d *Detector) processGroup(ctx context.Context, opts processGroupOpts) {
	filters := BuildFilters()

rl:
	for _, resource := range opts.group.APIResources {
		gv, err := schema.ParseGroupVersion(opts.group.GroupVersion)
		if err != nil {
			return
		}

		gvr := schema.GroupVersionResource{
			Group:    gv.Group,
			Version:  gv.Version,
			Resource: resource.Name,
		}

		for _, skip := range getSkipList() {
			if gvr == skip {
				continue rl
			}
		}

		if !slices.Contains(resource.Verbs, "list") {
			continue
		}

		restAPI := opts.dynClient.Resource(gvr).Namespace("")

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
			opts.col.Push(item)
		}
	}
}
