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
	width          int
	rawText        string
	lines          []*lines.Line
	currentLine    int
	linesWindow    int
	linesWindowTop int
}

var _ (tea.Model) = (*Model)(nil)
var _ (model.Sized) = (*Model)(nil)
var _ (model.SizedModel) = (*Model)(nil)

func New(text string, width int, linesWindow int, linesWindowTop int) *Model {
	return &Model{
		width:          width,
		rawText:        text,
		lines:          make([]*lines.Line, 0),
		linesWindow:    linesWindow,
		linesWindowTop: linesWindowTop,
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
	for i, line := range p.visibleLines() {
		sb.WriteString(line.View())
		notLastLine := i < p.linesWindow-1
		if notLastLine {
			sb.WriteRune('\n')
		}
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
	return m.updateLine(msg, true)
}

func (m *Model) checkChar(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.updateLine(msg, false)
}

func (m *Model) updateLine(msg tea.Msg, delete bool) (tea.Model, tea.Cmd) {
	line := m.getCurrentLine()
	line.Update(msg)

	if line.IsOverEdge() {
		m.moveLine(delete)
	}

	return m, nil
}

func (m *Model) moveLine(delete bool) {
	if delete {
		m.prevLine()
	} else {
		m.nextLine()
	}
	line := m.getCurrentLine()
	line.ReturnToEdge()
}

func (m *Model) getCurrentLine() *lines.Line {
	return m.lines[m.currentLine]
}

func (m *Model) nextLine() {
	lastLine := m.currentLine == len(m.lines)
	if lastLine {
		return
	}
	m.currentLine++
}

func (m *Model) prevLine() {
	firstLine := m.currentLine == 0
	if firstLine {
		return
	}
	m.currentLine--
}

func (m *Model) visibleLines() []*lines.Line {
	startLine := m.currentLine - m.linesWindowTop
	if m.currentLine < m.linesWindowTop {
		startLine = 0
	}
	endLine := startLine + m.linesWindow
	if endLine > len(m.lines) {
		endLine = len(m.lines)
		startLine = endLine - m.linesWindow
	}
	return m.lines[startLine:endLine]
}
