package main

import (
	"fmt"
	"os/exec"

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

	movieFilePathEntry := widget.NewEntry()
	movieUrlEntry := widget.NewEntry()

	content := container.NewVBox(
		widget.NewLabel("Enter movie file path"),
		movieFilePathEntry,
		widget.NewLabel("Enter subscene movie url"),
		movieUrlEntry,
	)

	content.Add(widget.NewButton("Add more items", func() {
		fmt.Println(movieFilePathEntry.Text)
		fmt.Println(movieUrlEntry.Text)
		// subtitle.GetSubtitles()
	}))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func openGoogle() {
	url := "http://google.com"
	cmd := exec.Command("open", url)
	cmd.Start()
}
