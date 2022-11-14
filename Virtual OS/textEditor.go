package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func showTextEditor() {
	// myApp := app.New()
	myWin := myApp.NewWindow("Text-Editor")
	myWin.Resize(fyne.NewSize(600, 500))

	count := 1

	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Manipulate text-files here.."),
		),
	)

	content.Add(widget.NewButton("Add New File", func() {
		content.Add(widget.NewLabel("New File" + strconv.Itoa(count)))
		count++
	}))

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")

	input.MultiLine = true

	saveBtn := widget.NewButton("Save File", func() {
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text)

				uc.Write(textData)
			},
			myWin)
		saveFileDialog.SetFileName("New File" + strconv.Itoa(count-1) + ".txt")
		saveFileDialog.Show()
	})

	openBtn := widget.NewButton("Open File", func() {
		openFileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				readData, _ := ioutil.ReadAll(r)

				output := fyne.NewStaticResource("File", readData)

				viewFile := widget.NewMultiLineEntry()

				viewFile.SetText(string(output.StaticContent))

				win := fyne.CurrentApp().NewWindow(string(output.StaticName))

				saveEditBtn := widget.NewButton("Save File", func() {
					saveEditFileDialog := dialog.NewFileSave(
						func(uc fyne.URIWriteCloser, _ error) {
							textData := []byte(viewFile.Text)
							uc.Write(textData)
						}, win)
					saveEditFileDialog.Show()
				})

				win.SetContent(container.NewVBox(container.NewScroll(viewFile), saveEditBtn))
				win.Resize(fyne.NewSize(400, 500))
				win.Show()
			}, myWin,
		)

		openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		openFileDialog.Show()
	})

	textEditorCont := container.NewVBox(
		content,
		input,
		container.NewHBox(
			saveBtn,
			openBtn,
		),
	)

	myWin.SetContent(container.NewBorder(nil, nil, nil, nil, textEditorCont))

	myWin.Show()
}
