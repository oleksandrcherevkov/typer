package typer

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/oleksandrcherevkov/typer/internal/lines"
	"github.com/oleksandrcherevkov/typer/internal/model"
	"github.com/oleksandrcherevkov/typer/internal/text"
)

var (
	typerStyle = lipgloss.NewStyle().
			Align(lipgloss.Center)
	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(0, 1).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)
)

type Model struct {
	width       int
	rawText     string
	lines       []*lines.Line
	currentLine int
}

var _ (tea.Model) = (*Model)(nil)
var _ (model.Sized) = (*Model)(nil)
var _ (model.SizedModel) = (*Model)(nil)

func New(text string, width int) *Model {
	return &Model{
		width:   width,
		rawText: text,
		lines:   make([]*lines.Line, 0),
	}
}

func (p *Model) Init() tea.Cmd {
	p.breakText()
	p.lines[0].ReturnToEdge()
	return nil
}

func (p *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "backspace":
			return p.deleteChar(msg)

		default:
			return p.checkChar(msg)
		}
	}
	return p, nil
}

func (p *Model) View() string {
	typerStyle = typerStyle.Width(p.width)

	var sb strings.Builder
	for _, line := range p.lines {
		sb.WriteString(line.View())
		sb.WriteRune('\n')
	}
	box := boxStyle.Render(sb.String())
	program := typerStyle.Render(box)

	return program
}

func (p *Model) Size(width int, height int) {
	p.width = width
}

func (p *Model) breakText() (tea.Model, tea.Cmd) {
	textLines := text.Lines(p.rawText, p.width-4)
	for _, line := range textLines {
		p.lines = append(p.lines, lines.New(line))
	}
	return p, nil
}

func (m *Model) deleteChar(msg tea.Msg) (tea.Model, tea.Cmd) {
	currentLine := m.lines[m.currentLine]
	currentLine.Update(msg)
	if currentLine.IsOverEdge() {
		m.currentLine--
		// TODO: check line
		currentLine := m.lines[m.currentLine]
		currentLine.ReturnToEdge()
	}
	return m, nil
}
func (m *Model) checkChar(msg tea.Msg) (tea.Model, tea.Cmd) {
	currentLine := m.lines[m.currentLine]
	currentLine.Update(msg)
	if currentLine.IsOverEdge() {
		m.currentLine++
		// TODO: check line
		currentLine := m.lines[m.currentLine]
		currentLine.ReturnToEdge()
	}
	return m, nil
}
