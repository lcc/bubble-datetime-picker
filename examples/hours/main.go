package main

import (
	"fmt"

	dtPicker "github.com/lcc/bubble-datetime-picker"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	model := dtPicker.NewHourSelectorModel()
	p := tea.NewProgram(&model)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("selected: %s\n", model.String())
	return
}
