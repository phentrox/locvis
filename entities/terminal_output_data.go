package entities

type TerminalOutputData struct {
	LineCounts       []LineCount
	TotalLinesOfCode int
	NumberOfFiles    int
	ShowPaths        bool
}
