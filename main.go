package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		fmt.Println("Failed to initialize termbox:", err)
		os.Exit(1)
	}
	defer termbox.Close()

	now := time.Now()
	year, month := now.Year(), now.Month()

	// メインループ
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		drawCalendar(year, month)
		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowRight, termbox.KeyArrowUp:
				month++
				if month > 12 {
					month = 1
					year++
				}
			case termbox.KeyArrowLeft, termbox.KeyArrowDown:
				month--
				if month < 1 {
					month = 12
					year--
				}
			case termbox.KeyCtrlC:
				return
			}
		}
	}
}

func drawCalendar(year int, month time.Month) {
	// キャプション
	headerCaption := "カレンダー（十字キーで移動・Ctrl+Cで停止）"
	printAt(1, 0, headerCaption, termbox.ColorYellow)

	header := fmt.Sprintf("   %d年 %d月   ", year, month)
	printAt(2, 2, header, termbox.ColorGreen)

	// 曜日
	weekDays := []string{"日", "月", "火", "水", "木", "金", "土"}
	for i, day := range weekDays {
		printAt(i*4+2, 4, day, termbox.ColorCyan)
	}

	// カレンダーの日付を取得
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	startWeekDay := int(firstDay.Weekday())
	daysInMonth := daysInMonth(year, month)

	// 日付を描画
	x, y := startWeekDay*4+2, 6
	for day := 1; day <= daysInMonth; day++ {
		printAt(x, y, fmt.Sprintf("%2d", day), termbox.ColorWhite)
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

func printAt(x, y int, msg string, color termbox.Attribute) {
	for i, r := range msg {
		termbox.SetCell(x+i, y, r, color, termbox.ColorDefault)
	}
}
