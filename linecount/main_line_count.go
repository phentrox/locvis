package linecount

import (
	"locvis/entities"
	"locvis/slicing"
	"locvis/sorting"
	"locvis/terminaloutput"
)

func createLineCountData(lineCounts []entities.LineCount, configVar entities.Config) entities.TerminalOutputData {
	var sortedLineCounts = sorting.Sort(lineCounts)
	var totalLinesOfCode int = countTotalLinesOfCode(lineCounts)
	var topLineCounts []entities.LineCount = slicing.GetTopSlice(sortedLineCounts, configVar.Top)

	terminalOutputData := entities.TerminalOutputData{
		LineCounts:       topLineCounts,
		TotalLinesOfCode: totalLinesOfCode,
		NumberOfFiles:    len(lineCounts),
		ShowPaths:        configVar.Paths,
	}

	return terminalOutputData
}

func CreateLineCountDataAndPrintToTerminal(lineCounts []entities.LineCount, configVar entities.Config) {
	terminalOutputData := createLineCountData(lineCounts, configVar)
	terminaloutput.PrintLineCount(terminalOutputData)
}
