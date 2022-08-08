package main

import (
	"locvis/config"
	"locvis/entities"
	"locvis/linecount"
	"locvis/localconfig"
	"locvis/slicing"
	"locvis/sorting"
	"locvis/statistics"
	"locvis/terminaloutput"
)

func main() {
	var localConfigVar entities.LocalConfig = localconfig.GetLocalConfig()
	var configVar entities.Config = config.CreateConfig(localConfigVar)

	lineCounts := linecount.GetLineCountsFromFiles(configVar)
	var sortedLineCounts = sorting.Sort(lineCounts)
	var totalLinesOfCode int = linecount.CountTotalLinesOfCode(lineCounts)
	var topLineCounts []entities.LineCount = slicing.GetTopSlice(sortedLineCounts, configVar.Top)

	terminalOutputData := entities.TerminalOutputData{
		LineCounts:       topLineCounts,
		TotalLinesOfCode: totalLinesOfCode,
		NumberOfFiles:    len(lineCounts),
		ShowPaths:        configVar.Paths,
	}

	terminaloutput.PrintLineCount(terminalOutputData)

	statistics.CreateDataAndPrintToTerminal(lineCounts)
}
