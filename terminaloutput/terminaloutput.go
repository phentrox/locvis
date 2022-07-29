package terminaloutput

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"locvis/entities"
	"os"
	"path"
)

func PrintLineCount(lineCounts []entities.LineCount) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"LoC", "Filename", "Path"})

	for _, lineCount := range lineCounts {
		var filename string = path.Base(lineCount.Path)
		t.AppendRow(table.Row{lineCount.LineCount, filename, lineCount.Path})
		//fmt.Println(lineCount.LineCount, " ", filename, " ", lineCount.Path)
	}
	t.Render()
}
