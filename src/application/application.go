package application

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/iwahing/image-layout-checker/src/checker"
)

type Application struct {
	app     fyne.App
	mainWin fyne.Window

	templateFile *widget.Entry
	teamFolder   *widget.Entry
	status       *widget.Entry

	template string
	team     string
}

func (a *Application) Init() {
	a.app = app.New()

	a.mainWin = a.app.NewWindow("Image Layout Checker")

	a.templateFile = widget.NewEntry()
	a.templateFile.SetPlaceHolder("Template File...")
	a.templateFile.Disable()
	a.teamFolder = widget.NewEntry()
	a.teamFolder.SetPlaceHolder("Team Folder...")
	a.teamFolder.Disable()

	a.status = widget.NewEntry()
	a.status.MultiLine = true
	a.status.Disable()
	// a.app.Settings().SetTheme(&MyTheme{})

	a.template = ""
	a.team = ""

	a.mainWin.SetContent(a.makeGUI())
	a.mainWin.Resize(fyne.NewSize(700, 550))
	a.mainWin.ShowAndRun()
}

func (a *Application) makeGUI() fyne.CanvasObject {
	main := container.NewVBox(
		a.templateFile,
		widget.NewButton("Select Template File", a.openTemplateFile),
		a.teamFolder,
		widget.NewButton("Select Team Folder", a.openTeamFolder),
		widget.NewButton("Check", a.checkFiles),
		a.status,
	)

	return main
}

func (a *Application) openTemplateFile() {
	fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			// Handle error
			fmt.Println("Error")
			fmt.Println(err)
			return
		}
		// Read file contents from reader

		if reader.URI() != nil {
			a.template = reader.URI().Path()
			a.templateFile.SetText(reader.URI().Path())
		}

	}, a.mainWin)

	fileDialog.Resize(fyne.NewSize(650, 500))
	fileDialog.Show()
}

func (a *Application) openTeamFolder() {
	fileDialog := dialog.NewFolderOpen(func(list fyne.ListableURI, err error) {
		if err != nil {
			// Handle error
			fmt.Println("Error")
			fmt.Println(err)
			return
		}
		// Read file contents from reader
		a.team = list.Path()
		a.teamFolder.SetText(list.Path())
	}, a.mainWin)

	fileDialog.Resize(fyne.NewSize(650, 500))
	fileDialog.Show()
}

func (a *Application) checkFiles() {

	if a.template == "" {
		a.app.SendNotification(fyne.NewNotification("Template File empty", "Template File has not been selected"))
		a.status.SetText("Template File has not been selected")
		a.status.TextStyle.Bold = true
		a.status.TextStyle.Italic = true
		a.status.Refresh()
		return
	}

	if a.team == "" {
		a.app.SendNotification(fyne.NewNotification("Team Folder empty", "Team Folder has not been selected"))
		a.status.SetText("Team Folder has not been selected")
		a.status.TextStyle.Bold = true
		a.status.TextStyle.Italic = true
		a.status.Refresh()
		return
	}

	c := checker.Checker{}
	c.Init(a.template, a.team)
	c.Check()
}
