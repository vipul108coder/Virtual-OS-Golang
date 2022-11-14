package main

import (
	"fmt"
	"image/color"
	"math"
	"strconv"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showBmiCalc() {
	// a := app.New()
	w := myApp.NewWindow("BMI Calculator")

	w.Resize(fyne.NewSize(400, 300))

	label := canvas.NewText("Enter Your Height & Weight", color.Black)
	label.Alignment = fyne.TextAlignCenter
	label.TextSize = 20

	result := canvas.NewText("", color.Black)
	result.Alignment = fyne.TextAlignCenter
	result.TextSize = 16

	inputH := widget.NewEntry()
	inputH.SetPlaceHolder("Enter height in Centimeter..")

	inputW := widget.NewEntry()
	inputW.SetPlaceHolder("Enter weight in Kilogram...")

	btn1 := widget.NewButton("Calculate BMI", func() {
		h, _ := strconv.ParseFloat(inputH.Text, 64)
		w, _ := strconv.ParseFloat(inputW.Text, 64)

		result.Text, label.Text = calculateBMI(h/100, w)
		result.Refresh()
		label.Refresh()
	})

	btn2 := widget.NewButton("Clear", func() {
		label.Refresh()
		label.Text = "Enter Your Height & Weight"
		result.Text = ""
		inputH.SetText("")
		inputW.SetText("")
	})

	w.SetContent(container.NewBorder(nil, nil, nil, nil, container.NewVBox(
		label,
		result,
		inputH,
		inputW,
		btn1,
		btn2,
	)))

	w.Show()
}

func calculateBMI(height, weight float64) (string, string) {
	var BMI float64 = weight / math.Pow(height, 2)

	if BMI < 18.5 {
		return fmt.Sprintf("BMI: %.2f", BMI), "You are Underweight. Eat More!"
	} else if BMI >= 18.5 && BMI < 25 {
		return fmt.Sprintf("BMI: %.2f", BMI), "You have Normal weight. Carry on.."
	} else if BMI >= 25 && BMI < 30 {
		return fmt.Sprintf("BMI: %.2f", BMI), "You are Overweight. Eat Less!"
	} else {
		return fmt.Sprintf("BMI: %.2f", BMI), "You are Obese. Don't Eat!"
	}
}
