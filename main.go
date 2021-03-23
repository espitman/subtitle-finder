package main

import (
	"fmt"
	"image/color"
	"os/exec"
	"strings"
	"subtitleFinder/subtitle"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var myWindow fyne.Window
var content *fyne.Container
var logContent *fyne.Container
var movieFilePathEntry *widget.Entry

func main() {

	myApp := app.New()
	myWindow = myApp.NewWindow("Entry Widget")
	myWindow.Resize(fyne.Size{Width: 800, Height: 600})
	myWindow.SetPadded(true)
	myWindow.CenterOnScreen()

	movieFilePathEntry = widget.NewEntry()
	movieUrlEntry := widget.NewEntry()

	content = container.NewVBox(
		widget.NewLabel("Enter movie file path"),
		movieFilePathEntry,
		widget.NewLabel("Enter subscene movie url"),
		movieUrlEntry,
	)

	logContent = container.NewVBox()

	content.Add(widget.NewButton("Get Subtitles", func() {
		subtitle.GetSubtitles(movieFilePathEntry.Text, movieUrlEntry.Text, log, addCheckButtons)
	}))
	addPaddingTop(content)
	addPaddingTop(logContent)

	myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, logContent))
	myWindow.ShowAndRun()

}

func addPaddingTop(c *fyne.Container) {
	r1 := canvas.NewRectangle(color.RGBA{0, 0, 0, 0})
	r1.SetMinSize(fyne.NewSize(1024, 10))
	c.Add(r1)
}

func log(text string) {
	logText := canvas.NewText(text, color.RGBA{0, 255, 0, 1})
	logText.TextStyle.Monospace = true
	logText.TextSize = 8
	logContent.Add(logText)
	logContent.Refresh()
}

func addCheckButtons(count int) {
	var selectedIndex string
	var subs []string
	for i := 0; i < count; i++ {
		subs = append(subs, "Subtitle - "+fmt.Sprint(i))
	}
	label := widget.NewLabel("Choose subtitle")
	combo := widget.NewSelect(subs,
		func(value string) {
			selectedIndex = value
		})
	button := widget.NewButton("Check Subtitles", func() {
		selectedIndex = strings.Replace(selectedIndex, "Subtitle - ", "", 1)
		fmt.Println("Checks:" + selectedIndex)
		subtitle.CheckSub(selectedIndex, log)
		openMovie(movieFilePathEntry.Text)

	})
	content.Add(label)
	grid := container.New(layout.NewGridLayout(2), combo, button)
	content.Add(grid)
}

func openMovie(path string) {
	// url := "http://google.com"
	cmd := exec.Command("open", path)
	cmd.Start()
}
