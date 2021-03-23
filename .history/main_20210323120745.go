package main

import (
	"os/exec"
	"subtitleFinder/subtitle"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")
	myWindow.Resize(fyne.Size{Width: 800, Height: 600})
	myWindow.SetPadded(true)
	myWindow.CenterOnScreen()

	content := container.NewVBox(
		widget.NewLabel("Enter movie file path"),
		widget.NewEntry(),
		widget.NewLabel("Enter subscene movie url"),
		widget.NewEntry(),
	)

	content.Add(widget.NewButton("Add more items", func() {
		subtitle.GetSubtitles()
	}))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func openGoogle() {
	url := "http://google.com"
	cmd := exec.Command("open", url)
	cmd.Start()
}
