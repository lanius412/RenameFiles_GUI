package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"github.com/sqweek/dialog"

	"log"


	//"Rename/rename"
	"github.com/ry0-suke/RenameFiles/rename"
)

func main() {
	app := app.New()
	window := app.NewWindow("Rename")
	window.Resize(fyne.NewSize(400, 80))

	setDirPath := widget.NewEntry() //Entry for Selected Directory
	setDirPath.SetText("          Select a Directory          ")
	setDirPath.Disable()

	var dir string
	var err error

	browseBtn := widget.NewButton("Browse", func() { //Browse Directory
		dir, err = dialog.Directory().Title("Open Dialog").Browse()
		log.Print("Browse...")
		if err != nil {
			log.Println("Directory Select Cancelled")
		}
		setDirPath.SetText(dir)

		log.Println(dir)
	})


	execBtn := widget.NewButton("Exec", func() { //Execute Rename
		if dir == "" {
			log.Println("Directory Not Selected")
		} else {
			log.Print("Rename...")
			rename.Rename(dir)
			window.Close()
		}
	})

	window.SetContent( //Layout Contents
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			fyne.NewContainerWithLayout(layout.NewHBoxLayout(), setDirPath, layout.NewSpacer(),fyne.NewContainerWithLayout(
					layout.NewVBoxLayout(), browseBtn,execBtn,
				),
			),
		),
	)

	window.ShowAndRun()
}
