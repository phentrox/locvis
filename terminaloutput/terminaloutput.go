package terminaloutput

import (
	"fmt"
	"locvis/entities"
)

func PrintLineCount(lineCounts []entities.LineCount) {
	for _, lineCount := range lineCounts {
		fmt.Println(lineCount.LineCount, " ", lineCount.Path)
	}
}
