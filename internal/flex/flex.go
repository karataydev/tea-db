package flex

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
)

// Flexible Element Interface
type Flex interface {
	tea.Model
	resize(int, int)
	Priority() int
}

type FlexBase struct {
	focus bool
	height int
	width int
	priority int
}

func (f *FlexBase) resize(width int, height int) {
	f.width = width
	f.height = height
}

func (f FlexBase) Priority() int {
	return f.priority
}


