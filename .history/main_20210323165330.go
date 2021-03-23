package main

import (
	"fmt"
	"image/color"
	"subtitleFinder/subtitle"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var content *fyne.Container

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

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
