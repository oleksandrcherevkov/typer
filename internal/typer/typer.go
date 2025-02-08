package typer

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/oleksandrcherevkov/typer/internal/lines"
	"github.com/oleksandrcherevkov/typer/internal/model"
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
	lineStyle = lipgloss.NewStyle().
			Align(lipgloss.Left)
)

type Model struct {
	position  int
	lineRaw   string
	line      []*lines.Character
	lineWidth int
	width     int
}

var _ (tea.Model) = (*Model)(nil)
var _ (model.Sized) = (*Model)(nil)
var _ (model.SizedModel) = (*Model)(nil)

func New(line string, lineWidth int) *Model {
	return &Model{
		lineRaw: line,
	}
}

func (p *Model) Init() tea.Cmd {
	p.parseLine()
	return nil
}

func (p *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "backspace":
			return p.deleteChar()

		default:
			return p.checkChar(msg)
		}
	}
	return p, nil
}

func (p *Model) View() string {
	typerStyle = typerStyle.Width(p.width)
	lineStyle = lineStyle.Width(p.lineWidth)

	var sb strings.Builder
	for _, ch := range p.line {
		sb.WriteString(ch.String())
	}
	line := lineStyle.Render(sb.String())
	box := boxStyle.Render(line)
	program := typerStyle.Render(box)

	return program
}

func (p *Model) Size(width int, height int) {
	p.width = width
}

func (p *Model) parseLine() (tea.Model, tea.Cmd) {
	p.position = 0
	p.line = lines.BrakeString(p.lineRaw)
	current := p.line[p.position]
	current.Select()
	return p, nil
}

func (p *Model) deleteChar() (tea.Model, tea.Cmd) {
	if p.position == 0 {
		current := p.line[p.position]
		current.Select()
		return p, nil
	}
	current := p.line[p.position]
	current.Unselect()
	p.position--
	current = p.line[p.position]
	current.Select()
	return p, nil
}

func (p *Model) checkChar(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if len(msg.Runes) != 1 {
		return p, nil
	}

	if p.position >= len(p.line) {
		return p.parseLine()
	}

	char := msg.Runes[0]
	current := p.line[p.position]
	current.Check(char)
	current.Pass()

	p.position++

	if p.position >= len(p.line) {
		return p.parseLine()
	}

	current = p.line[p.position]
	current.Select()

	return p, nil
}
