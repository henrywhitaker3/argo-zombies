package detect

import (
	"github.com/fatih/color"
	"github.com/henrywhitaker3/argo-zombies/internal/detector"
	"github.com/rodaine/table"
)

func NewTable(col *detector.Collection) table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Name", "Namespace", "Kind", "APIVersion")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, item := range col.Items {
		tbl.AddRow(item.GetName(), item.GetNamespace(), item.GetKind(), item.GetAPIVersion())
	}

	return tbl
}
