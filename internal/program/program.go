package program

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
	"github.com/oleksandrcherevkov/typer/internal/model"
	"github.com/oleksandrcherevkov/typer/internal/typer"
)

type Program struct {
	typer model.SizedModel
}

func New(text string) Program {
	return Program{
		typer: typer.New(text, 80, 5, 2),
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
			_, cmd = p.typer.Update(msg)
		}
	}
	return p, cmd
}

func (p Program) View() string {
	physicalWight, _, _ := term.GetSize(os.Stdout.Fd())
	p.typer.Size(physicalWight, 0)
	return p.typer.View()
}
