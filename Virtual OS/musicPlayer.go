package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var format beep.Format
var pause bool = false
var streamer beep.StreamSeekCloser

func showMusicPlayr() {
	var myWin fyne.Window = myApp.NewWindow("Music Player")
	myWin.Resize(fyne.NewSize(400, 400))

	image := canvas.NewImageFromFile("images\\clipart326945.png")
	image.FillMode = canvas.ImageFillOriginal

	var label1 *canvas.Text = canvas.NewText("Music Player", color.Black)
	label1.Alignment = fyne.TextAlignCenter
	label1.TextStyle = fyne.TextStyle{Bold: true}
	label1.TextSize = 18

	var label2 *canvas.Text = canvas.NewText("File Name", color.Black)
	label2.Alignment = fyne.TextAlignCenter

	var browseBtn fyne.Widget = widget.NewButton("Browse Files", func() {
		fd := dialog.NewFileOpen(func(uc fyne.URIReadCloser, e error) {
			streamer, format, _ = mp3.Decode(uc)
			label2.Text = uc.URI().Name()
			label2.Refresh()
		}, myWin)
		fd.Show()
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))
	})

	controls := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			speaker.Play(streamer)
		}),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			if !pause {
				pause = true
				speaker.Lock()
			} else if pause {
				pause = false
				speaker.Unlock()
			}
		}),
		widget.NewToolbarAction(theme.MediaStopIcon(), func() {
			speaker.Clear()
			label2.Text = "File Name"
			label2.Refresh()
		}),
		widget.NewToolbarSpacer(),
	)

	var musicContainer = container.NewVBox(
		label1,
		image,
		browseBtn,
		label2,
		controls,
	)

	myWin.SetContent(container.NewBorder(nil, nil, nil, nil, musicContainer))

	myWin.Show()
}
