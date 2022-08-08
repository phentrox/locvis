package linecount

import "locvis/entities"

func countTotalLinesOfCode(lineCounts []entities.LineCount) int {
	var totalLinesOfCode int = 0

	for _, lineCount := range lineCounts {
		totalLinesOfCode += lineCount.LineCount
	}

	return totalLinesOfCode
}
