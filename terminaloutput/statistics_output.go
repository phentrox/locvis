package terminaloutput

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"locvis/entities"
	"os"
)

func PrintStatistics(statisticsOutputData entities.StatisticsOutputData) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	addStats(statisticsOutputData, t)

	t.Render()
}

func addStats(statisticsOutputData entities.StatisticsOutputData, t table.Writer) {
	t.AppendHeader(table.Row{"Median"})
	t.AppendRow(table.Row{statisticsOutputData.Median})
}
