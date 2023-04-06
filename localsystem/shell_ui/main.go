package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {

	err := ui.Init()

	defer ui.Close()

	if err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	// Create a new paragraph
	p := widgets.NewParagraph()
	p.Text = "This is the beginning of Gaxiom"
	p.SetRect(0, 0, 50, 3)

	ui.Render(p)

	for e := range ui.PollEvents() {

		// if the user presses any button, quit
		if e.Type == ui.KeyboardEvent {
			break
		}
	}

}
