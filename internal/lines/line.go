package lines

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var lineStyle = lipgloss.NewStyle().
	Align(lipgloss.Left)

type Line struct {
	position int
	lineRaw  string
	line     []*Character
}

var _ (tea.Model) = (*Line)(nil)

func New(text string) *Line {
	line := &Line{
		lineRaw: text,
	}
	line.parseLine()
	return line
}

func (l *Line) Init() tea.Cmd {
	return nil
}

func (l *Line) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "backspace":
			return l.deleteChar()

		case "enter":
			return l.checkChar('\n')

		default:
			if len(msg.Runes) != 1 {
				return l, nil
			}
			char := msg.Runes[0]
			return l.checkChar(char)
		}
	}
	return l, nil
}

func (l *Line) View() string {
	var sb strings.Builder
	for _, ch := range l.line {
		sb.WriteString(ch.String())
	}
	line := lineStyle.Render(sb.String())

	return line
}

func (l *Line) parseLine() (tea.Model, tea.Cmd) {
	l.position = 0
	l.line = BrakeString(l.lineRaw)
	return l, nil
}

func (l *Line) deleteChar() (tea.Model, tea.Cmd) {
	if l.IsOverEdge() {
		return l, nil
	}

	current := l.line[l.position]
	current.Unselect()

	l.position--

	l.selectCurrent()

	return l, nil
}

func (l *Line) checkChar(char rune) (tea.Model, tea.Cmd) {

	if l.IsOverEdge() {
		return l, nil
	}

	current := l.line[l.position]
	current.Check(char)
	current.Pass()

	l.position++

	l.selectCurrent()

	return l, nil
}

func (l *Line) IsOverEdge() bool {
	return l.isCaretBeforeLine() || l.isCaretAfterLine()
}

func (l *Line) isCaretBeforeLine() bool {
	return l.position < 0
}

func (l *Line) isCaretAfterLine() bool {
	return l.position >= len(l.line)
}

func (l *Line) ReturnToEdge() {
	if l.isCaretBeforeLine() {
		l.position++
		current := l.line[l.position]
		current.Select()
		return
	}
	if l.isCaretAfterLine() {
		l.position--
		current := l.line[l.position]
		current.Select()
		return
	}
	current := l.line[l.position]
	current.Select()
}

func (l *Line) selectCurrent() {
	if l.IsOverEdge() {
		return
	}

	current := l.line[l.position]
	current.Select()
}
