package main

import (
	"subtitleFinder/subtitle"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myWindow fyne.Window

func main() {

	myApp := app.New()
	var content *fyne.Container
	// myApp.Settings().SetTheme(&myTheme{})
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

	content.Add(widget.NewButton("Add more items", func() {
		subtitle.GetSubtitles(movieFilePathEntry.Text, movieUrlEntry.Text, addLog)
	}))
	myWindow.SetContent(addLog("A", "B"))
	myWindow.ShowAndRun()
}

func addLog(dirName string, subName string) {

	// content.Add(widget.NewLabel(dirName))
	idDisplay := canvas.NewText("loading", theme.ErrorColor())
	idDisplay.TextStyle.Monospace = true
	idDisplay.TextSize = 32

	myWindow.SetContent(idDisplay)

	// idDisplay.Alignment = fyne.Text

	// a := widget.NewLabelWithStyle(dirName, fyne.TextAlignLeading, fyne.TextStyle{Italic: true, Monospace: true})
	// content.Add(createLabel())
}

// func openGoogle() {
// 	url := "http://google.com"
// 	cmd := exec.Command("open", url)
// 	cmd.Start()
// }
