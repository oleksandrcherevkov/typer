package typer

import (
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
	"github.com/oleksandrcherevkov/typer/internal/lines"
)

var typerStyle = lipgloss.NewStyle().
	Align(lipgloss.Center)

var boxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#874BFD")).
	Padding(0, 1).
	BorderTop(true).
	BorderLeft(true).
	BorderRight(true).
	BorderBottom(true)

var lineStyle = lipgloss.NewStyle().
	Align(lipgloss.Left).
	Width(80)

type Model struct {
	position int
	lineRaw  string
	line     []*lines.Character
}

var _ (tea.Model) = (*Model)(nil)

func New(line string) *Model {
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
	physicalWight, _, _ := term.GetSize(os.Stdout.Fd())
	typerStyle = typerStyle.Width(physicalWight)
	var sb strings.Builder
	for _, ch := range p.line {
		sb.WriteString(ch.String())
	}
	line := lineStyle.Render(sb.String())
	box := boxStyle.Render(line)
	program := typerStyle.Render(box)
	return program
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
