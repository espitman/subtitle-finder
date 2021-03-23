package main

import (
	"subtitleFinder/subtitle"

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
		subtitle.GetSubtitles(movieFilePathEntry.Text, movieUrlEntry.Text, addLog)
	}))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func leftColumn() *fyne.Container {
	return container.NewGridWithColumns(1,
		container.NewBorder(
			widget.NewLabel("Input:"),
			nil,
			nil,
			nil,
			input,
		),
		container.NewBorder(
			widget.NewLabel("Output:"),
			nil,
			nil,
			nil,
			container.NewMax(output),
		),
	)
}

func addLog(dirName string, subName string) {

	// content.Add(widget.NewLabel(dirName))
	// c := content.Add(canvas)
	// t := canvas.NewText("0", color.White)
	// t.Text = "0"
	// content.Add(c)

	// a := widget.NewLabelWithStyle(dirName, fyne.TextAlignLeading, fyne.TextStyle{Italic: true, Monospace: true})
	// content.Add(a)
}

// func openGoogle() {
// 	url := "http://google.com"
// 	cmd := exec.Command("open", url)
// 	cmd.Start()
// }
