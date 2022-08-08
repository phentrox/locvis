package main

import (
	"locvis/config"
	"locvis/entities"
	"locvis/linecount"
	"locvis/localconfig"
	"locvis/statistics"
)

func main() {
	var localConfigVar entities.LocalConfig = localconfig.GetLocalConfig()
	var configVar entities.Config = config.CreateConfig(localConfigVar)

	var lineCounts []entities.LineCount = linecount.GetLineCountsFromFiles(configVar)
	linecount.CreateLineCountDataAndPrintToTerminal(lineCounts, configVar)
	statistics.CreateDataAndPrintToTerminal(lineCounts)
}
