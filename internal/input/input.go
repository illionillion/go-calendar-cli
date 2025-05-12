package input

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

/*
キー入力処理
*/
func HandleInput(screen tcell.Screen, year *int, month *time.Month, day *int) bool {
	ev := screen.PollEvent()
	switch event := ev.(type) {
	case *tcell.EventKey:
		switch event.Key() {
		case tcell.KeyUp:
			*month++
			if *month > 12 {
				*month = 1
				*year++
			}
		case tcell.KeyDown:
			*month--
			if *month < 1 {
				*month = 12
				*year--
			}
		case tcell.KeyRight:
			// 日付を進める
			*day++
			// 月末を超えた場合
			daysInMonth := time.Date(*year, *month+1, 0, 0, 0, 0, 0, time.Local).Day()
			if *day > daysInMonth {
				*day = 1
				*month++
				if *month > 12 {
					*month = 1
					*year++
				}
			}
		case tcell.KeyLeft:
			// 日付を戻す
			*day--
			// 月初を下回った場合
			if *day < 1 {
				*month--
				if *month < 1 {
					*month = 12
					*year--
				}
				daysInMonth := time.Date(*year, *month+1, 0, 0, 0, 0, 0, time.Local).Day()
				*day = daysInMonth
			}
		case tcell.KeyCtrlC:
			return false // 終了
		}
	}
	return true
}
