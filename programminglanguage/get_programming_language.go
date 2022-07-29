package programminglanguage

import (
	"locvis/entities"
	"locvis/userinput"
)

func Get(localConfig entities.LocalConfig) string {
	var programmingLanguage string

	if localConfig.Language != "" {
		programmingLanguage = localConfig.Language
	} else {
		programmingLanguage = userinput.GetProgrammingLanguage()
	}

	return programmingLanguage
}
