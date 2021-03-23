package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var content *fyne.Container

func main() {
	a := app.New()
	w := a.NewWindow("Tappable")
	w.SetContent(newTappableIcon(theme.FyneLogo()))
	w.ShowAndRun()
}

func createLabel() fyne.CanvasObject {

	fmt.Println(theme.TextSize())

	label := widget.NewLabel("XxxxxxxxX")
	label.Resize(fyne.NewSize(100, 20))

	// label.SetMinSize(fyne.NewSize(100, 20))
	label.TextStyle.Bold = true
	return label
}

func addLog(dirName string, subName string) {

	// content.Add(widget.NewLabel(dirName))
	// c := content.Add(canvas)
	// t := canvas.NewText("0", color.White)
	// t.Text = "0"
	// content.Add(c)

	// a := widget.NewLabelWithStyle(dirName, fyne.TextAlignLeading, fyne.TextStyle{Italic: true, Monospace: true})
	content.Add(createLabel())
}

// func openGoogle() {
// 	url := "http://google.com"
// 	cmd := exec.Command("open", url)
// 	cmd.Start()
// }
