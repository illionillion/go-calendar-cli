package calendar

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/illionillion/go-calendar-cli/internal/render"
)

/*
カレンダーを描画
*/
func DrawCalendar(screen tcell.Screen, year int, month time.Month, currentDay int) {
	// キャプション
	headerCaption := "カレンダー（十字キーで移動・Ctrl+Cで停止）"
	render.PrintAt(screen, 1, 0, headerCaption, tcell.ColorYellow)

	// ヘッダー
	header := fmt.Sprintf("   %d年 %d月 %d日   ", year, month, currentDay)
	render.PrintAt(screen, 2, 2, header, tcell.ColorGreen)

	// 曜日
	weekDays := []string{"日", "月", "火", "水", "木", "金", "土"}
	for i, day := range weekDays {
		render.PrintAt(screen, i*4+2, 4, day, tcell.NewHexColor(0x00FFFF))
	}

	// カレンダーの日付を取得
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	startWeekDay := int(firstDay.Weekday())
	daysInMonth := DaysInMonth(year, month)

	// 日付を描画
	x, y := startWeekDay*4+2, 6
	for day := 1; day <= daysInMonth; day++ {
		color := tcell.ColorWhite
		if day == currentDay {
			color = tcell.ColorRed
		}
		render.PrintAt(screen, x, y, fmt.Sprintf("%2d", day), color)
		x += 4
		if x > 6*4+2 {
			x = 2
			y++
		}
	}
}

/*
指定した年・月の日数を取得
*/
func DaysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.Local).Day()
}
