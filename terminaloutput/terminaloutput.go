package terminaloutput

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"locvis/entities"
	"os"
	"path"
)

func PrintLineCount(terminalOutputData entities.TerminalOutputData) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	if terminalOutputData.ShowPaths {
		addRowsWithPaths(terminalOutputData, t)
	} else {
		addRowsWithoutPaths(terminalOutputData, t)
	}

	t.Render()
}

func addRowsWithoutPaths(terminalOutputData entities.TerminalOutputData, t table.Writer) {
	t.AppendHeader(table.Row{"LoC", "Filename"})
	for _, lineCount := range terminalOutputData.LineCounts {
		var filename string = path.Base(lineCount.Path)
		t.AppendRow(table.Row{lineCount.LineCount, filename})
	}
	t.AppendFooter(table.Row{terminalOutputData.TotalLinesOfCode, terminalOutputData.NumberOfFiles})
}

func addRowsWithPaths(terminalOutputData entities.TerminalOutputData, t table.Writer) {
	t.AppendHeader(table.Row{"LoC", "Filename", "Path"})
	for _, lineCount := range terminalOutputData.LineCounts {
		var filename string = path.Base(lineCount.Path)
		t.AppendRow(table.Row{lineCount.LineCount, filename, lineCount.Path})
	}
	t.AppendFooter(table.Row{terminalOutputData.TotalLinesOfCode, terminalOutputData.NumberOfFiles, ""})
}
