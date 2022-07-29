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

	var programmingLanguage string = programminglanguage.Get(localConfigVar)

	var dirsToBeSkipped = []string{".git", ".idea", ".mvn"}

	var programmingLanguageSuffix string = programminglanguage.GetProgrammingLanguageSuffix(programmingLanguage)
	if programmingLanguageSuffix == "" {
		println("Programming Language not supported")
		main()
	}

	var showPaths bool = localconfig.GetShowPaths(localConfigVar)

	var files []string = filewalk.GetFilesInDir(dirsToBeSkipped, programmingLanguageSuffix)
	var lineCounts []entities.LineCount = linecount.CountLinesFromArrayWithPaths(files)
	lineCounts = sorting.Sort(lineCounts)
	var topTenHighestLineCounts []entities.LineCount = slicing.GetTopTenSlice(lineCounts)
	terminaloutput.PrintLineCount(topTenHighestLineCounts, showPaths)
}
