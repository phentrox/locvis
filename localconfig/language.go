package localconfig

import (
	"locvis/entities"
	"locvis/userinput"
)

func GetLang(localConfig entities.LocalConfig) string {
	var programmingLanguage string

	if localConfig.Language != "" {
		programmingLanguage = localConfig.Language
	} else {
		programmingLanguage = userinput.GetProgrammingLanguage()
	}

	return programmingLanguage
}
