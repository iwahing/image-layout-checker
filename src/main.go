package main


import (
	"flag"
	"github.com/iwahing/image-layout-checker/src/template"
	"github.com/iwahing/image-layout-checker/src/checker"
)

var templateFile string
var folderPath string

func main() {

	flag.StringVar(&templateFile, "template file", "./sizing.csv", "File of sizing to be used as a template")
	flag.StringVar(&folderPath, "path", "./", "Folder path to scan")
	flag.Parse()

	t := template.Controller{}
	t.Init(templateFile)
	t.PrintTemplate()

	checker.Check(t.Sizing, folderPath)
}
