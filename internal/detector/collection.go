package detector

import (
	"sync"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type Collection struct {
	Items []unstructured.Unstructured
	m     *sync.RWMutex
}

func NewCollection() *Collection {
	return &Collection{
		Items: []unstructured.Unstructured{},
		m:     &sync.RWMutex{},
	}
}

func (c *Collection) Push(item unstructured.Unstructured) {
	c.m.Lock()
	defer c.m.Unlock()

	c.Items = append(c.Items, item)
}
