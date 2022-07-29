package linecount

import (
	"bytes"
	"io"
	"locvis/entities"
	"log"
	"os"
)

func CountLinesFromArrayWithPaths(arrayWithPaths []string) []entities.LineCount {
	var lineCounts []entities.LineCount

	for _, path := range arrayWithPaths {
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}

		lineCount, err := countLines(file)
		if err != nil {
			panic(err.Error())
		}

		err = file.Close()
		if err != nil {
			log.Fatalln("Could not close file ", file.Name())
		}

		lineCountEntity := entities.LineCount{
			Path:      path,
			LineCount: lineCount,
		}

		lineCounts = append(lineCounts, lineCountEntity)
	}

	return lineCounts
}

func countLines(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
