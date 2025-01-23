package program

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/oleksandrcherevkov/typer/internal/typer"
)

type Program struct {
	typer tea.Model
}

func New(line string) Program {
	return Program{
		typer: typer.New(line),
	}
}

func (p Program) Init() tea.Cmd {
	return tea.Batch(tea.ClearScreen, p.typer.Init())
}

func (p Program) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			cmd = tea.Quit
		default:
			p.typer, cmd = p.typer.Update(msg)
		}
	}
	return p, cmd
}

func (p Program) View() string {
	return p.typer.View()
}
