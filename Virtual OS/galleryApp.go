package main

import (
	//"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func showGalleryApp(w fyne.Window) {
	// myApp := app.New()
	// w := myApp.NewWindow("Image Gallery")
	// w.Resize(fyne.NewSize(600, 400))
	imgSource := "" //Image URL here

	files, err := ioutil.ReadDir(imgSource)
	if err != nil {
		log.Fatal(err)
	}

	var arr []string

	for _, file := range files {
		//fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir() {
			extension := strings.Split(file.Name(), ".")[1]
			//fmt.Println(extension)
			if extension == "png" || extension == "jpg" {
				arr = append(arr, imgSource+"\\"+file.Name())

			}
		}
	}
	tabs := container.NewAppTabs(container.NewTabItem("Image", canvas.NewImageFromFile(arr[0])))

	for i := 1; i < len(arr); i++ {
		// image := canvas.NewImageFromFile(arr[i])
		// image.FillMode = canvas.ImageFillOriginal
		tabs.Append(container.NewTabItem("Image", canvas.NewImageFromFile(arr[i])))
	}

	tabs.SetTabLocation(container.TabLocationTrailing)

	w.SetContent(container.NewBorder(panelContent, nil, nil, nil, tabs))
	w.Show()
}
