package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
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
		drawCalendar(screen, year, month)
		screen.Show()

		ev := screen.PollEvent()
		switch event := ev.(type) {
		case *tcell.EventKey:
			switch event.Key() {
			case tcell.KeyRight, tcell.KeyUp:
				month++
				if month > 12 {
					month = 1
					year++
				}
			case tcell.KeyLeft, tcell.KeyDown:
				month--
				if month < 1 {
					month = 12
					year--
				}
			case tcell.KeyCtrlC:
				return
			}
		}
	}
}

func drawCalendar(screen tcell.Screen, year int, month time.Month) {
	// キャプション
	headerCaption := "カレンダー（十字キーで移動・Ctrl+Cで停止）"
	printAt(screen, 1, 0, headerCaption, tcell.ColorYellow)

	// ヘッダー
	header := fmt.Sprintf("   %d年 %d月   ", year, month)
	printAt(screen, 2, 2, header, tcell.ColorGreen)

	// 曜日
	weekDays := []string{"日", "月", "火", "水", "木", "金", "土"}
	for i, day := range weekDays {
		printAt(screen, i*4+2, 4, day, tcell.NewHexColor(0x00FFFF))
	}

	// カレンダーの日付を取得
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	startWeekDay := int(firstDay.Weekday())
	daysInMonth := daysInMonth(year, month)

	// 日付を描画
	x, y := startWeekDay*4+2, 6
	for day := 1; day <= daysInMonth; day++ {
		printAt(screen, x, y, fmt.Sprintf("%2d", day), tcell.ColorWhite)
		x += 4
		if x > 6*4+2 {
			x = 2
			y++
		}
	}
}

func daysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.Local).Day()
}

func printAt(screen tcell.Screen, x, y int, msg string, color tcell.Color) {
	style := tcell.StyleDefault.Foreground(color).Background(tcell.ColorDefault)
	for i, r := range msg {
		screen.SetContent(x+i, y, r, nil, style)
	}
}
