package main

import (
	"fmt"
	"os"
	"github.com/karataymarufemre/tea-db/internal/flex"
	tea "github.com/charmbracelet/bubbletea"
)



func main() {

	childeren := &[]flex.FlexElement{}
	var row flex.FlexElement = flex.NewRow(100, 100, 1, childeren)
	rows := []flex.FlexElement{}	
	rows = append(rows, row)
	c := flex.NewColumn(200, 200, 1, &rows)
	p := tea.NewProgram(c)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
