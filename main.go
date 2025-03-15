package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/illionillion/go-calendar-cli/internal/calendar"
	"github.com/illionillion/go-calendar-cli/internal/input"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Println("Failed to initialize tcell:", err)
		os.Exit(1)
	}

	if err := screen.Init(); err != nil {
		fmt.Println("Failed to initialize screen:", err)
		os.Exit(1)
	}
	defer screen.Fini()

	now := time.Now()
	year, month := now.Year(), now.Month()

	// メインループ
	for {
		screen.Clear()
		calendar.DrawCalendar(screen, year, month)
		screen.Show()

		if !input.HandleInput(screen, &year, &month) {
			return
		}
	}
}
