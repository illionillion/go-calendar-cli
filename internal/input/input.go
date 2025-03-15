package input

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

/*
キー入力処理
*/
func HandleInput(screen tcell.Screen, year *int, month *time.Month) bool {
	ev := screen.PollEvent()
	switch event := ev.(type) {
	case *tcell.EventKey:
		switch event.Key() {
		case tcell.KeyRight, tcell.KeyUp:
			*month++
			if *month > 12 {
				*month = 1
				*year++
			}
		case tcell.KeyLeft, tcell.KeyDown:
			*month--
			if *month < 1 {
				*month = 12
				*year--
			}
		case tcell.KeyCtrlC:
			return false // 終了
		}
	}
	return true
}
