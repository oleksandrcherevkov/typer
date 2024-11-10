package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/oleksandrcherevkov/typer/internal/program"
)

func main() {
	filePath := getFilePath()
	p := tea.NewProgram(program.New(filePath))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func getFilePath() string {
	args := os.Args[1:]
	if len(args) < 1 {
		panic("no file path in arguments")
	}
	fmt.Println(args)
	return args[0]
}
