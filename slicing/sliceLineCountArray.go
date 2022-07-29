package slicing

import "locvis/entities"

func GetTopSlice(lineCounts []entities.LineCount, top int) []entities.LineCount {
	if len(lineCounts) <= top {
		return lineCounts
	}
	return lineCounts[0:top]
}
