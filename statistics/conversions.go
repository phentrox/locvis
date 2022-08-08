package statistics

import "locvis/entities"

func ConvertLineCountsToFloats(lineCounts []entities.LineCount) []float64 {
	var floats []float64
	for _, lineCount := range lineCounts {
		var floatNumber float64 = float64(lineCount.LineCount)
		floats = append(floats, floatNumber)
	}
	return floats
}
