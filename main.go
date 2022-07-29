package main

import (
	"locvis/entities"
	"locvis/filewalk"
	"locvis/linecount"
	"locvis/programminglanguage"
	"locvis/slicing"
	"locvis/sorting"
	"locvis/terminaloutput"
	"locvis/userinput"
)

func main() {
	var programmingLanguage string = userinput.GetProgrammingLanguage()

	var dirsToBeSkipped = []string{".git", ".idea", ".mvn"}

	var programmingLanguageSuffix string = programminglanguage.GetProgrammingLanguageSuffix(programmingLanguage)
	if programmingLanguageSuffix == "" {
		println("Programming Language not supported")
		main()
	}

	var files []string = filewalk.GetFilesInDir(dirsToBeSkipped, programmingLanguageSuffix)
	var lineCounts []entities.LineCount = linecount.CountLinesFromArrayWithPaths(files)
	lineCounts = sorting.Sort(lineCounts)
	var topTenHighestLineCounts []entities.LineCount = slicing.GetTopTenSlice(lineCounts)
	terminaloutput.PrintLineCount(topTenHighestLineCounts)
}
