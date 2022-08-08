package linecount

import (
	"locvis/entities"
	"locvis/filewalk"
	"locvis/programminglanguage"
)

func GetLineCountsFromFiles(configVar entities.Config) []entities.LineCount {
	var programmingLanguageSuffix string = programminglanguage.GetProgrammingLanguageSuffix(configVar.Language)
	var files []string = filewalk.GetFilesInDir(configVar.IgnoreDirs, programmingLanguageSuffix)
	var lineCounts []entities.LineCount = countLinesFromArrayWithPaths(files)
	return lineCounts
}
