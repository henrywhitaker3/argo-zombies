package detector

import (
	"fmt"
	"sync"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type Collection struct {
	Items map[string]*unstructured.Unstructured
	m     *sync.RWMutex
}

func NewCollection() *Collection {
	return &Collection{
		Items: map[string]*unstructured.Unstructured{},
		m:     &sync.RWMutex{},
	}
}

func (c *Collection) Push(item *unstructured.Unstructured) {
	c.m.Lock()
	defer c.m.Unlock()

	c.Items[fmt.Sprintf("%s/%s/%s", item.GetAPIVersion(), item.GetKind(), item.GetName())] = item
}
