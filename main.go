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
	lineCounts = sorting.Sort(lineCounts)

	var topTenHighestLineCounts []entities.LineCount = slicing.GetTopSlice(lineCounts, configVar.Top)

	terminaloutput.PrintLineCount(topTenHighestLineCounts, configVar.Paths)
}
