package main

import (
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

	var programmingLanguage string = localconfig.GetLang(localConfigVar)

	var dirsToBeSkipped = []string{".git", ".idea", ".mvn"}

	var programmingLanguageSuffix string = programminglanguage.GetProgrammingLanguageSuffix(programmingLanguage)

	var showPaths bool = localconfig.GetShowPaths(localConfigVar)

	var files []string = filewalk.GetFilesInDir(dirsToBeSkipped, programmingLanguageSuffix)
	var lineCounts []entities.LineCount = linecount.CountLinesFromArrayWithPaths(files)
	lineCounts = sorting.Sort(lineCounts)

	var topNumber int = localconfig.GetTop(localConfigVar)
	var topTenHighestLineCounts []entities.LineCount = slicing.GetTopSlice(lineCounts, topNumber)

	terminaloutput.PrintLineCount(topTenHighestLineCounts, showPaths)
}
