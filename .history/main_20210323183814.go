package main

import (
	"image/color"
	"subtitleFinder/subtitle"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var myWindow fyne.Window
var content *fyne.Container

func main() {

	myApp := app.New()
	myWindow = myApp.NewWindow("Entry Widget")
	myWindow.Resize(fyne.Size{Width: 800, Height: 600})
	myWindow.SetPadded(true)
	myWindow.CenterOnScreen()

	movieFilePathEntry := widget.NewEntry()
	movieUrlEntry := widget.NewEntry()

	content = container.NewVBox(
		widget.NewLabel("Enter movie file path"),
		movieFilePathEntry,
		widget.NewLabel("Enter subscene movie url"),
		movieUrlEntry,
	)

	content.Add(widget.NewButton("Get Subtitles", func() {
		subtitle.GetSubtitles(movieFilePathEntry.Text, movieUrlEntry.Text, log)
	}))
	addLogTopPadding()
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func addLogTopPadding() {
	r1 := canvas.NewRectangle(color.RGBA{0, 0, 0, 0})
	r1.SetMinSize(fyne.NewSize(1024, 10))
	content.Add(r1)

}

func log(text string) {
	logText := canvas.NewText(text, color.RGBA{0, 255, 0, 1})
	logText.TextStyle.Monospace = true
	logText.TextSize = 8
	content.Add(logText)
	content.Refresh()
}

// func openGoogle() {
// 	url := "http://google.com"
// 	cmd := exec.Command("open", url)
// 	cmd.Start()
// }
