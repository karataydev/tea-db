package flex

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
)

type Column struct {
	FlexParent
}

func (c *Column) resize(width int, height int)  {
	c.FlexBase.resize(width, height)
	for _, flex := range *c.children {
		newHeight := int(float32(c.height) * c.coefficient(flex))
		flex.resize(c.width, newHeight)
	}
}
