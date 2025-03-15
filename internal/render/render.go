package render

import (
	"github.com/gdamore/tcell/v2"
)

/*
画面に文字を表示する
*/
func PrintAt(screen tcell.Screen, x, y int, msg string, color tcell.Color) {
	style := tcell.StyleDefault.Foreground(color).Background(tcell.ColorDefault)
	for i, r := range msg {
		screen.SetContent(x+i, y, r, nil, style)
	}
}
