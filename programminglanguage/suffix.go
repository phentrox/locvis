package programminglanguage

func GetProgrammingLanguageSuffix(programmingLanguage string) string {
	programmingLanguageSuffixes := getProgrammingLanguageSuffixes()
	return programmingLanguageSuffixes[programmingLanguage]
}

func getProgrammingLanguageSuffixes() map[string]string {
	programmingLanguageSuffixes := map[string]string{
		"java": ".java",
		"go":   ".go",
	}
	return programmingLanguageSuffixes
}
