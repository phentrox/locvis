package statistics

import (
	"github.com/montanaflynn/stats"
	"locvis/entities"
	"locvis/terminaloutput"
)

func createStatisticsOutputData(lineCounts []entities.LineCount) entities.StatisticsOutputData {
	var float64LineCounts []float64 = ConvertLineCountsToFloats(lineCounts)
	median, _ := stats.Median(float64LineCounts)

	statisticsOutputData := entities.StatisticsOutputData{
		Median: median,
	}
	return statisticsOutputData
}

func CreateDataAndPrintToTerminal(lineCounts []entities.LineCount) {
	statisticsOutputData := createStatisticsOutputData(lineCounts)
	terminaloutput.PrintStatistics(statisticsOutputData)
}
