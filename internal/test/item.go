package test

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type Item struct {
	item unstructured.Unstructured
}

func NewItem() Item {
	item := Item{
		item: unstructured.Unstructured{},
	}
	return item.SetName("bongo").
		SetAPIVersion("v1").
		SetKind("Pod").
		SetNamespace("default")
}

func (i Item) SetName(name string) Item {
	i.item.SetName(name)
	return i
}

func (i Item) SetNamespace(name string) Item {
	i.item.SetNamespace(name)
	return i
}

func (i Item) SetKind(name string) Item {
	i.item.SetKind(name)
	return i
}

func (i Item) SetAPIVersion(name string) Item {
	i.item.SetAPIVersion(name)
	return i
}

func (i Item) SetAnnotations(annos map[string]string) Item {
	i.item.SetAnnotations(annos)
	return i
}

func (i Item) SetLabels(labels map[string]string) Item {
	i.item.SetLabels(labels)
	return i
}

func (i Item) Item() unstructured.Unstructured {
	return i.item
}
