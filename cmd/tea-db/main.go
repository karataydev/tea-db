package main

import (
	"fmt"
	"os"
	"github.com/karataymarufemre/tea-db/internal/flex"
	tea "github.com/charmbracelet/bubbletea"
)



func main() {

	childeren := &[]flex.FlexElement{flex.NewFlexStr(100, 100, "RRRRR")}
	var row flex.FlexElement = flex.NewRow(100, 100, 1, childeren)
	rows := []flex.FlexElement{flex.NewFlexStr(100, 100, "CCCC")}	
	rows = append(rows, row)
	c := flex.NewColumn(200, 200, 1, &rows)

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()
	p := tea.NewProgram(c)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
