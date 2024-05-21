package flex

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FlexElement interface {
	tea.Model
	setSize(int, int)
	resizeChildren()
	Priority() int
}

type flex struct {
	focus bool
	height int
	width int
	priority int
	pSum int
	children *[]FlexElement
}

type column struct {
	flex
}

type row struct {
	flex
}

func (f *flex) setSize(width int, height int) {
	f.width = width
	f.height = height
}

func (f flex) Priority() int {
	return f.priority
}

func (f *flex) resizeChildren() {
	
}

func newFlex(width, height, priority int, children *[]FlexElement) *flex {
	pSum := 0
	for _, e := range *children {
		pSum += e.Priority()
	}

	return &flex {
			width: width,
			height: height,
			pSum: pSum,
			children: children,
			priority: priority,
	}
}

func (c *column) resizeChildren() {
	for _, flex := range *c.children {
		flex.setSize(c.width, c.height * (flex.Priority() / c.pSum))
	}
}

func (r *row) resizeChildren() {
	for _, flex := range *r.children {
		flex.setSize(r.width * (flex.Priority() / r.pSum), r.height)
	}
}


func NewColumn(width, height, priority int, children *[]FlexElement) *column {
	f := newFlex(width, height, priority, children)
	c := &column {
		flex: *f,
	}
	c.resizeChildren()
	return c
}

func NewRow(width, height, priority int, children *[]FlexElement) *row {
	f := newFlex(width, height, priority, children)
	r := &row {
		flex: *f,
	}
	r.resizeChildren()
	return r
}

func (c column) Init() tea.Cmd {
	return nil
}


func (c column) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.setSize(msg.Width, msg.Height)
		c.resizeChildren()
	case tea.KeyMsg:
		switch msg.String() {
			case "ctrl+c", "q":
				return c, tea.Quit
		}
	}
	return c, cmd
}


func (c column) View() string {
	childStrs := []string{} 
	for _, flex := range *c.children {
		childStrs = append(childStrs, flex.View())
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		childStrs...,
	)
}

func (c row) Init() tea.Cmd {
	return nil
}


func (c row) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.setSize(msg.Width, msg.Height)
		c.resizeChildren()
	case tea.KeyMsg:
		switch msg.String() {
			case "ctrl+c", "q":
				return c, tea.Quit
		}
	}
	return c, cmd
}


func (c row) View() string {
	childStrs := []string{} 
	for _, flex := range *c.children {
		childStrs = append(childStrs, flex.View())
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		childStrs...,
	)
}
