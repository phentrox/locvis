package slicing

import "locvis/entities"

func GetTopTenSlice(lineCounts []entities.LineCount) []entities.LineCount {
	if len(lineCounts) < 10 {
		return lineCounts
	}
	return lineCounts[0:9]
}
