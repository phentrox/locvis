package main

import (
	"fmt"
	"locvis/entities"
	"locvis/filewalk"
	"locvis/linecount"
	"locvis/programminglanguage"
	"locvis/userinput"
)

func main() {
	var programmingLanguage string = userinput.GetProgrammingLanguage()
	fmt.Println()

	var dirsToBeSkipped = []string{".git", ".idea"}

	var programmingLanguageSuffix string = programminglanguage.GetProgrammingLanguageSuffix(programmingLanguage)
	if programmingLanguageSuffix == "" {
		println("Programming Language not supported")
		main()
	}
	var files []string = filewalk.GetFilesInDir(dirsToBeSkipped, programmingLanguageSuffix)
	var lineCounts []entities.LineCount = linecount.CountLinesFromArrayWithPaths(files)
	for _, lineCount := range lineCounts {
		fmt.Println(lineCount.LineCount, " ", lineCount.Path)
	}
}
