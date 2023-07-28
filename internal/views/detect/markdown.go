package detect

import (
	"bytes"
	"text/template"

	"github.com/henrywhitaker3/argo-zombies/internal/detector"
)

var issueTemplate string = `# Argo Zombies

Here is a list of resources not managed by ArgoCD:

| Name | Namespace | Kind | APIVersion |
| --- | --- | --- | --- |
{{- with .Collection }}
{{- range . }}
| {{ .Name }} | {{ .Namespace }} | {{ .Kind }} | {{ .APIVersion }} |
{{- end }}
{{- end }}
`

type item struct {
	Name       string
	Namespace  string
	Kind       string
	APIVersion string
}

func NewMarkdown(col *detector.Collection) (string, error) {
	tpl, err := template.New("issue").Parse(issueTemplate)
	if err != nil {
		return "", err
	}

	c := []item{}

	for _, i := range col.Items {
		c = append(c, item{
			Name:       i.GetName(),
			Namespace:  i.GetNamespace(),
			Kind:       i.GetKind(),
			APIVersion: i.GetAPIVersion(),
		})
	}

	buf := new(bytes.Buffer)
	d := map[string]any{
		"Collection": c,
	}

	if err := tpl.Execute(buf, d); err != nil {
		return "", err
	}

	return buf.String(), nil
}
