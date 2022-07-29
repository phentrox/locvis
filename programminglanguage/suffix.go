package programminglanguage

import "os"

func GetProgrammingLanguageSuffix(programmingLanguage string) string {
	programmingLanguageSuffixes := getProgrammingLanguageSuffixes()
	var suffix string = programmingLanguageSuffixes[programmingLanguage]
	if suffix == "" {
		println("Programming Language not supported: ", programmingLanguage)
		os.Exit(1)
	}
	return suffix
}

func getProgrammingLanguageSuffixes() map[string]string {
	programmingLanguageSuffixes := map[string]string{
		"java":   ".java",
		"go":     ".go",
		"golang": ".go",
	}
	return programmingLanguageSuffixes
}
