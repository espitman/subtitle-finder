package main

import (
	"image/color"

	"fyne.io/fyne/canvas"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
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
	return 32
}

func main() {

	// myApp := app.New()
	// // myApp.Settings().SetTheme(&myTheme{})
	// myWindow := myApp.NewWindow("Entry Widget")
	// // myWindow.Resize(fyne.Size{Width: 800, Height: 600})
	// // myWindow.SetPadded(true)
	// // myWindow.CenterOnScreen()

	// // movieFilePathEntry := widget.NewEntry()
	// // movieUrlEntry := widget.NewEntry()

	// // content = container.NewVBox(
	// // 	widget.NewLabel("Enter movie file path"),
	// // 	movieFilePathEntry,
	// // 	widget.NewLabel("Enter subscene movie url"),
	// // 	movieUrlEntry,
	// // )

	// // content.Add(widget.NewButton("Add more items", func() {
	// // 	subtitle.GetSubtitles(movieFilePathEntry.Text, movieUrlEntry.Text, addLog)
	// // }))
	// rect := canvas.NewRectangle(color.White)
	// myWindow.SetContent(rect)
	// myWindow.ShowAndRun()

	myApp := app.New()
	w := myApp.NewWindow("Rectangle")

	rect := canvas.NewRectangle(color.White)
	w.SetContent(rect)

	w.Resize(fyne.NewSize(150, 100))
	w.ShowAndRun()
}

func addLog(dirName string, subName string) *canvas.Rectangle {

	rect := canvas.NewRectangle(color.White)
	return rect
	// content.Add(widget.NewLabel(dirName))
	// idDisplay := canvas.NewText("loading", theme.ErrorColor())
	// idDisplay.TextStyle.Monospace = true
	// idDisplay.TextSize = 32

	// return idDisplay

	// idDisplay.Alignment = fyne.Text

	// a := widget.NewLabelWithStyle(dirName, fyne.TextAlignLeading, fyne.TextStyle{Italic: true, Monospace: true})
	// content.Add(createLabel())
}

// func openGoogle() {
// 	url := "http://google.com"
// 	cmd := exec.Command("open", url)
// 	cmd.Start()
// }
