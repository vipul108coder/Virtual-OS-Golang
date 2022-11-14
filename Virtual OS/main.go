package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()
var myWindow fyne.Window = myApp.NewWindow("My OS")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var btn5 fyne.Widget
var btn6 fyne.Widget

var img fyne.CanvasObject
var panelContent *fyne.Container
var desktopBtn fyne.Widget

func main() {
	// myApp.Settings().SetTheme(theme.LightTheme())
	img = canvas.NewImageFromFile("images\\deskWall.jpg")

	btn1 = widget.NewButtonWithIcon("Weather App", theme.InfoIcon(), func() {
		showWeatherApp(myWindow)
	})
	btn2 = widget.NewButtonWithIcon("Calculator", theme.CancelIcon(), func() {
		showCalculator()
	})
	// btn3 = widget.NewButtonWithIcon("Gallery App", theme.MediaPhotoIcon(), func() {
	// 	showGalleryApp(myWindow)
	// })
	btn4 = widget.NewButtonWithIcon("Text Editor", theme.DocumentIcon(), func() {
		showTextEditor()
	})
	// btn5 = widget.NewButtonWithIcon("BMI Calc", theme.FyneLogo(), func() {
	// 	showBmiCalc()
	// })
	btn6 = widget.NewButtonWithIcon("Music Player", theme.MediaMusicIcon(), func() {
		showMusicPlayr()
	})

	desktopBtn = widget.NewButtonWithIcon("Desktop", theme.HomeIcon(), func() {
		myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	})

	panelContent = container.NewVBox(container.NewGridWithColumns(7, desktopBtn, btn1, btn2, btn3, btn4, btn5, btn6))

	myWindow.Resize(fyne.NewSize(1080, 720))
	myWindow.CenterOnScreen()
	myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))

	myWindow.ShowAndRun()
}
