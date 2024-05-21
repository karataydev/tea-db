package flex

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
)

type FlexElement interface {
	tea.Model
	setSize(int, int)
	resizeChildren()
	Priority() int
}

type flexString struct {
	flex
	str string
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
		cof := float32(flex.Priority()) / float32(c.pSum)
		newHeight := int(float32(c.height) * cof)
		flex.setSize(c.width, newHeight)
		flex.resizeChildren()
	}
}

func (r *row) resizeChildren() {
	for _, flex := range *r.children {
		cof := float32(flex.Priority()) / float32(r.pSum)
		newWidth := int(float32(r.width) * cof)
		flex.setSize(newWidth, r.height)
		flex.resizeChildren()
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

func NewFlexStr(width, height int, str string) *flexString {
	f := newFlex(width, height, 1, &[]FlexElement{})
	r := &flexString {
		flex: *f,
		str: str,
	}
	return r
}

func (f *flexString) setSize(width int, height int) {
	f.width = width
	f.height = height
}



func (c flexString) Init() tea.Cmd {
	return nil
}


func (c flexString) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd	
	return c, cmd
}


func (c flexString) View() string {
	return lipgloss.NewStyle().
		Height(c.height).
		Width(c.width).
		Render(c.str)
}


func (c column) Init() tea.Cmd {
	return nil
}


func (c column) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		log.Printf("%d - %d", msg.Width, msg.Height)
		c.setSize(msg.Width-2, msg.Height-3)
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

	return lipgloss.NewStyle().
		Height(c.height-10).
		Width(c.width).
		//Padding(1,1).
		Border(lipgloss.RoundedBorder()).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				childStrs...,
			),
		)
}

func (c row) Init() tea.Cmd {
	return nil
}


func (c row) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
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

	return lipgloss.NewStyle().
		Height(c.height).
		Width(c.width-2).
		//Padding(1,1).
		Border(lipgloss.RoundedBorder()).
		Render(
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				childStrs...,
			),
		)
}
