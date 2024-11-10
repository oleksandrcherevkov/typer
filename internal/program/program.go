package program

import (
	"bufio"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
	"github.com/oleksandrcherevkov/typer/internal/lines"
)

var programStyle = lipgloss.NewStyle().
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

type Program struct {
	text        *bufio.Scanner
	textClosing func()
	position    int
	lineRaw     string
	line        []lines.Character
}

func New(filePath string) *Program {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("file can not be opened")
	}
	fileClose := func() {
		file.Close()
	}

	scan := bufio.NewScanner(file)

	return &Program{
		text:        scan,
		textClosing: fileClose,
	}
}

func (p *Program) Init() tea.Cmd {
	return tea.ClearScreen
}

func (p *Program) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			return p.exitProgram()

		case "down":
			return p.nextLine()

		case "backspace":
			return p.deleteChar()

		default:
			return p.checkChar(msg)
		}
	}
	return p, nil
}

func (p *Program) View() string {
	physicalWight, _, _ := term.GetSize(os.Stdout.Fd())
	programStyle = programStyle.Width(physicalWight)
	var sb strings.Builder
	for _, ch := range p.line {
		sb.WriteString(ch.String())
	}
	line := lineStyle.Render(sb.String())
	box := boxStyle.Render(line)
	program := programStyle.Render(box)
	return program
}

func (p *Program) exitProgram() (tea.Model, tea.Cmd) {
	p.textClosing()
	return p.teaQuit()
}

func (p *Program) teaQuit() (tea.Model, tea.Cmd) {
	return p, tea.Quit
}

func (p *Program) nextLine() (tea.Model, tea.Cmd) {
	s := p.text.Scan()
	if !s {
		return p, tea.Quit
	}
	p.lineRaw = p.text.Text()
	p.lineRaw = strings.TrimSpace(p.lineRaw)
	p.position = 0
	p.line = lines.BrakeString(p.lineRaw)
	return p, nil
}

func (p *Program) deleteChar() (tea.Model, tea.Cmd) {
	p.position--
	p.line[p.position].CheckStatus = lines.None
	return p, nil
}

func (p *Program) checkChar(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if len(msg.Runes) != 1 {
		return p, nil
	}

	if p.position == len(p.line) {
		return p.nextLine()
	}

	char := msg.Runes[0]
	expected := p.line[p.position].Expected
	p.line[p.position].Typed = char
	if char == expected {
		p.line[p.position].CheckStatus = lines.Correct
	} else {
		p.line[p.position].CheckStatus = lines.Wrong
	}
	p.position++
	return p, nil
}
