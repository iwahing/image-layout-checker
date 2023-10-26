package main


import (
	"flag"
	"github.com/iwahing/image-layout-checker/src/checker"
)

var templateFile string
var folderPath string

func main() {

	flag.StringVar(&templateFile, "template file", "./sizing.csv", "File of sizing to be used as a template")
	flag.StringVar(&folderPath, "path", "../MEDALLA NAPO MILAGROSA/Medalla Napo", "Folder path to scan")
	flag.Parse()

	c := checker.Checker{}
	c.Init(templateFile, folderPath)
	c.Check()
}
