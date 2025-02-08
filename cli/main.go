package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/oleksandrcherevkov/typer/internal/program"
	"github.com/oleksandrcherevkov/typer/internal/text"
)

func main() {
	text := getText()
	p := tea.NewProgram(program.New(text))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func getText() string {
	filePath := getFilePath()
	text, err := text.GetText(filePath)
	if err != nil {
		panic(err.Error())
	}

	return text
}

func getFilePath() string {
	args := os.Args[1:]
	if len(args) < 1 {
		panic("no file path in arguments")
	}
	fmt.Println(args)
	return args[0]
}
