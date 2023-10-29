package checker

import (
	"bufio"
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
	"path/filepath"
	"strings"
)

type Checker struct {
	template   Controller
	folderPath string
}

func (c *Checker) Init(templateFile string, folderPath string) {
	c.template = Controller{}
	c.template.Init(templateFile)
	// c.PrintTemplate()
	c.folderPath = folderPath
}

var bannerSize = Size{3395, 2396}

func (c *Checker) GetDimension(file string) (int, int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, 0, err
	}

	reader := bufio.NewReader(f)
	config, _, err := image.DecodeConfig(reader)
	if err != nil {
		fmt.Println(err)
	}

	return config.Width, config.Height, nil
}

func (c *Checker) FileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func (c *Checker) ScanFolder(folderPath string, sizeType string) []string {
	files, err := os.ReadDir(folderPath)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var incorrectSizes []string
	for _, file := range files {
		fileName := file.Name()

		if filepath.Ext(fileName) == ".jpg" {
			absFilePath := folderPath + "/" + fileName
			size := strings.Split(c.FileNameWithoutExtSliceNotation(fileName), "_")
			size[1] = strings.ToLower(size[1])

			name := size[0] + "_" + size[1]
			if len(size) > 3 {
				name = size[0] + "_" + size[2] + "_" + size[1]
			}

			width, height, err := c.GetDimension(absFilePath)
			if err != nil {
				fmt.Println("	-", name, " Error: ", err)
				continue
			}

			actualSize := Size{width, height}
			// fmt.Println(fmt.Sprintf("Item: %s | Size: %s", sizeType, size[1]))
			if c.template.Sizing[sizeType][size[1]] == actualSize {
				fmt.Println("	-", name, " Correct!")
			} else {
				fmt.Println("	-", name, " Incorrect")
				incorrectSizes = append(incorrectSizes, name)
			}

		}
	}

	return incorrectSizes
}

func (c *Checker) Check() {
	fmt.Println("Scanning folder '", c.folderPath, "'")

	files, err := os.ReadDir(c.folderPath)
	if err != nil {
		fmt.Println(err)
	}

	var incorrectSizes []string
	for _, file := range files {
		filename := file.Name()
		absPath := c.folderPath + "/" + filename

		info, err := os.Stat(absPath)
		if err != nil {
			fmt.Println(err)
			continue
		}

		size := strings.ToLower(filename)
		_, ok := c.template.Sizing[size]

		if info.IsDir() && ok {
			fmt.Println(filename)
			result := c.ScanFolder(c.folderPath+"/"+filename, size)
			if result != nil {
				incorrectSizes = append(incorrectSizes, result...)
			}
		} else if filename == "Banner.jpg" {
			width, height, err := c.GetDimension(absPath)
			if err != nil {
				fmt.Println(err)
				continue
			}

			actualSize := Size{width, height}

			if bannerSize == actualSize {
				fmt.Println("Banner.jpg Correct!")
			} else {
				fmt.Println("Banner.jpg Incorrect")
				incorrectSizes = append(incorrectSizes, "Banner")
			}

		} else {
			fmt.Println("Skipping ", filename)
		}
	}

	fmt.Printf("# Failed resize: %v\n", incorrectSizes)
	return
}
