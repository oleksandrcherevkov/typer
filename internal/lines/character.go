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
	if ch.position == Future {
		return string(ch.Expected)
	}

	if ch.position == Current {
		return selectedStyle.Render(string(ch.Expected))
	}

	switch ch.CheckStatus {
	case Idle:
		return string(ch.Expected)
	case Correct:
		return correctStyle.Render(string(ch.Typed))
	case Wrong:
		return wrongStyle.Render(string(ch.Expected))
	case Corrected:
		return correctedStyle.Render(string(ch.Typed))
	}

	return string(ch.Expected)
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
