package main

import (
	"image/color"
	"os/exec"
	"subtitleFinder/subtitle"

	"fyne.io/fyne/canvas"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var content *fyne.Container

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")
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

	content.Add(widget.NewButton("Add more items", func() {
		subtitle.GetSubtitles(movieFilePathEntry.Text, movieUrlEntry.Text, addDirsResult)
	}))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func addDirsResult(dirName string, subName string) {
	// content.Add(widget.NewLabel(dirName))
	// content.Add(widget.NewLabel(subName))
	text := canvas.NewText("Text Object", color.White)
	text.Alignment = fyne.TextAlignTrailing
	text.TextStyle = fyne.TextStyle{Italic: true}
	w.SetContent(text)
}

func openGoogle() {
	url := "http://google.com"
	cmd := exec.Command("open", url)
	cmd.Start()
}
