package lines

import "github.com/charmbracelet/lipgloss"

type Check int

const (
	None Check = iota
	Correct
	Wrong
)

var (
	correctStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#28ee81")).Background(lipgloss.Color("#197d6e"))
	wrongStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#8B0000")).Background(lipgloss.Color("#fc785e"))
)

type Character struct {
	Expected    rune
	Typed       rune
	CheckStatus Check
}

func BrakeString(s string) []Character {
	res := make([]Character, 0)
	for _, c := range s {
		res = append(res, Character{Expected: c, CheckStatus: None})
	}
	return res
}

func (ch *Character) String() string {
	switch ch.CheckStatus {
	case None:
		return string(ch.Expected)
	case Correct:
		return correctStyle.Render(string(ch.Typed))
	case Wrong:
		return wrongStyle.Render(string(ch.Typed))
	}
	return string(ch.Expected)
}
