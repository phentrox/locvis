package terminaloutput

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"locvis/entities"
	"os"
	"path"
)

func PrintLineCount(lineCounts []entities.LineCount, showPaths bool) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	if showPaths {
		addRowsWithPaths(lineCounts, t)
	} else {
		addRowsWithoutPaths(lineCounts, t)
	}

	t.Render()
}

func addRowsWithoutPaths(lineCounts []entities.LineCount, t table.Writer) {
	t.AppendHeader(table.Row{"LoC", "Filename"})
	for _, lineCount := range lineCounts {
		var filename string = path.Base(lineCount.Path)
		t.AppendRow(table.Row{lineCount.LineCount, filename})
	}
}

func addRowsWithPaths(lineCounts []entities.LineCount, t table.Writer) {
	t.AppendHeader(table.Row{"LoC", "Filename", "Path"})
	for _, lineCount := range lineCounts {
		var filename string = path.Base(lineCount.Path)
		t.AppendRow(table.Row{lineCount.LineCount, filename, lineCount.Path})
	}
}
