package model

import tea "github.com/charmbracelet/bubbletea"

type Sized interface {
	Size(width int, hight int)
}

type SizedModel interface {
	Sized
	tea.Model
}
