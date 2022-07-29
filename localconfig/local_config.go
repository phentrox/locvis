package localconfig

import (
	"gopkg.in/yaml.v3"
	"locvis/entities"
	"log"
	"os"
)

func GetLocalConfig() entities.LocalConfig {
	const YamlFileName string = ".locvis.yaml"
	const YmlFileName string = ".locvis.yml"

	if fileExists(YamlFileName) {
		return parseYamlFile(YamlFileName)
	}
	if fileExists(YmlFileName) {
		return parseYamlFile(YmlFileName)
	}

	showPaths := true
	defaultConfig := entities.LocalConfig{
		Top:         10,
		Paths:       &showPaths,
		IgnoredDirs: []string{},
	}
	return defaultConfig
}

func parseYamlFile(fileName string) entities.LocalConfig {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Could not open ", fileName)
	}
	var localConfig entities.LocalConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&localConfig)
	if err != nil {
		log.Fatalln("Could not decode ", fileName)
	}
	f.Close()
	return localConfig
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
