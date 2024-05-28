package flex

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
)

type Row struct {
	FlexParent
}


func (r *Row) resize(width int, height int) {
	r.FlexBase.resize(width, height)
	for _, flex := range *r.children {
		newWidth := int(float32(r.width) *r.coefficient(flex))
		flex.resize(newWidth, r.height)
	}
}
