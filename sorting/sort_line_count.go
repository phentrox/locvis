package sorting

import (
	"locvis/entities"
	"sort"
)

func Sort(lineCounts []entities.LineCount) []entities.LineCount {
	sort.SliceStable(lineCounts, func(i, j int) bool {
		return lineCounts[i].LineCount > lineCounts[j].LineCount
	})
	return lineCounts
}
