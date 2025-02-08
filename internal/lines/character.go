package lines

import (
	"github.com/charmbracelet/lipgloss"
)

type Check int

const (
	Idle Check = iota
	Correct
	Wrong
	Corrected
)

type Position int

const (
	Future Position = iota
	Current
	Passed
)

var (
	defaultStyle   = lipgloss.NewStyle()
	correctStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#28ee81")).Background(lipgloss.Color("#197d6e"))
	wrongStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#8B0000")).Background(lipgloss.Color("#fc785e"))
	correctedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#af6400")).Background(lipgloss.Color("#ffdd60"))
)
var (
	selectedStyle = lipgloss.NewStyle().Background(lipgloss.Color("#444444"))
)

type Character struct {
	Expected    rune
	Typed       rune
	CheckStatus Check
	position    Position
}

func BrakeString(s string) []*Character {
	res := make([]*Character, 0)
	for _, c := range s {
		res = append(res, &Character{Expected: c, CheckStatus: Idle})
	}
	return res
}

func (ch *Character) String() string {
	style, char := ch.render()
	if char == '\n' {
		char = '↵'
	}
	return style.Render(string(char))
}

func (ch *Character) Check(char rune) bool {
	ch.Typed = char

	same := ch.Expected == char
	wasWrong := ch.CheckStatus == Wrong
	wasCorrected := ch.CheckStatus == Corrected

	if same && (wasWrong || wasCorrected) {
		ch.CheckStatus = Corrected
		return true
	}

	if same {
		ch.CheckStatus = Correct
		return true
	}

	ch.CheckStatus = Wrong
	return false
}

func (ch *Character) Select() {
	ch.position = Current
}

func (ch *Character) Unselect() {
	ch.position = Future
}

func (ch *Character) Pass() {
	ch.position = Passed
}

func (ch *Character) render() (lipgloss.Style, rune) {
	if ch.position == Future {
		return defaultStyle, ch.Expected
	}

	if ch.position == Current {
		return selectedStyle, ch.Expected
	}

	switch ch.CheckStatus {
	case Idle:
		return defaultStyle, ch.Expected
	case Correct:
		return correctStyle, ch.Typed
	case Wrong:
		return wrongStyle, ch.Expected
	case Corrected:
		return correctedStyle, ch.Typed
	}

	return defaultStyle, ch.Expected
}
