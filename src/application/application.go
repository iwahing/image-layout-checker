package application


import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/iwahing/image-layout-checker/src/checker"
)

type Appication struct {
	app     		fyne.App
	mainWin 		fyne.Window

	templateFile	*widget.Label
	teamFolder		*widget.Label

	template		string
	team			string
}

func (this *Appication) Init(){
	this.app = app.New()

	this.mainWin = this.app.NewWindow("Image Layout Checker")

	this.templateFile = widget.NewLabel("template")
	this.teamFolder = widget.NewLabel("team")

	this.template = ""
	this.team = ""

	this.mainWin.SetContent(this.makeGUI())
	this.mainWin.Resize(fyne.NewSize(1200, 750))
	this.mainWin.ShowAndRun()
}

func (this *Appication) makeGUI() fyne.CanvasObject {
	main := container.NewVBox(
		this.templateFile,
		widget.NewButton("Select Template File", this.openTemplateFile),
		this.teamFolder,
		widget.NewButton("Select Team Folder", this.openTeamFolder),
		widget.NewButton("Check", this.checkFiles),
	)

	return main
}

func (this *Appication) openTemplateFile() {
	// this.content.SetText("Welcome :)")
	fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
        if err != nil {
            // Handle error
            return
        }
        // Read file contents from reader
		this.template = reader.URI().String()
		this.templateFile.SetText("Template File:" + reader.URI().String())
		}, this.mainWin)

	fileDialog.Resize(fyne.NewSize(900, 500))
	fileDialog.Show()
}

func (this *Appication) openTeamFolder() {
	// this.content.SetText("Welcome :)")
	fileDialog := dialog.NewFolderOpen(func(list  fyne.ListableURI, err error) {
        if err != nil {
            // Handle error
            return
        }
        // Read file contents from reader\
		this.team = list.String()
		this.teamFolder.SetText("Team Folder:" + list.String())
		}, this.mainWin)

	fileDialog.Resize(fyne.NewSize(900, 500))
	fileDialog.Show()
}

func (this *Appication) checkFiles() {
	c := checker.Checker{}
	c.Init(this.template, this.team)
	c.Check()
}
