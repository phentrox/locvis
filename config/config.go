package config

import (
	"locvis/entities"
	"locvis/skip"
)

func CreateConfig(localConfig entities.LocalConfig) entities.Config {
	var lang string = GetLang(localConfig)
	var top int = GetTop(localConfig)
	var showPaths bool = GetShowPaths(localConfig)
	var dirsToBeSkipped []string = skip.CreateDirsToBeSkipped(localConfig)

	config := entities.Config{
		Language:   lang,
		Top:        top,
		Paths:      showPaths,
		IgnoreDirs: dirsToBeSkipped,
	}

	return config
}
