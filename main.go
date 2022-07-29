package main

import (
	"locvis/config"
	"locvis/entities"
	"locvis/filewalk"
	"locvis/linecount"
	"locvis/localconfig"
	"locvis/programminglanguage"
	"locvis/slicing"
	"locvis/sorting"
	"locvis/terminaloutput"
)

func main() {
	var localConfigVar entities.LocalConfig = localconfig.GetLocalConfig()
	var configVar entities.Config = config.CreateConfig(localConfigVar)

	var programmingLanguageSuffix string = programminglanguage.GetProgrammingLanguageSuffix(configVar.Language)
	var files []string = filewalk.GetFilesInDir(configVar.IgnoreDirs, programmingLanguageSuffix)
	var lineCounts []entities.LineCount = linecount.CountLinesFromArrayWithPaths(files)
	var totalLinesOfCode int = linecount.CountTotalLinesOfCode(lineCounts)
	lineCounts = sorting.Sort(lineCounts)
	var topLineCounts []entities.LineCount = slicing.GetTopSlice(lineCounts, configVar.Top)

	terminalOutputData := entities.TerminalOutputData{
		LineCounts:       topLineCounts,
		TotalLinesOfCode: totalLinesOfCode,
		NumberOfFiles:    len(lineCounts),
		ShowPaths:        configVar.Paths,
	}

	terminaloutput.PrintLineCount(terminalOutputData)
}
