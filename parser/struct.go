package parser

type Task struct {
	Name         string
	Requirements map[string]int
	Products     map[string]int
	NbCycles        int
}
