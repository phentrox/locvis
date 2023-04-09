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
	// Only one for each language!
	programmingLanguageSuffixes := map[string]string{
		"java":   ".java",
		"go":     ".go",
		"c":      ".c",
		"js":     ".js",
		"ts":     ".ts",
		"rust":   ".rs",
		"c#":     ".cs",
		"python": ".py",
	}
	return programmingLanguageSuffixes
}
