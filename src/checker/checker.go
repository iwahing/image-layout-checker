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
	builder    strings.Builder
}

func (c *Checker) Init(templateFile string, folderPath string) {
	c.builder = strings.Builder{}
	c.template = Controller{}
	c.builder.WriteString("Reading Template File'" + templateFile + "'\n")
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
		c.builder.WriteString(fmt.Sprintf("%v\n", err))
		return 0, 0, err
	}

	return config.Width, config.Height, nil
}

func (c *Checker) FileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func (c *Checker) ScanFolder(folderPath string, sizeType string) []string {
	files, err := os.ReadDir(folderPath)

	if err != nil {
		c.builder.WriteString(fmt.Sprintf("%v\n", err))
		return nil
	}

	var incorrectSizes []string
	sizing, ok := c.template.Sizing[sizeType]
	if !ok {
		c.builder.WriteString("No Size Template for" + sizeType + "Skipping folder: " + folderPath + "\n")
		return incorrectSizes
	}

	for _, file := range files {
		fileName := file.Name()

		if filepath.Ext(fileName) == ".jpg" {
			absFilePath := folderPath + "/" + fileName
			size := strings.Split(c.FileNameWithoutExtSliceNotation(fileName), "_")

			if len(size) != 3 {
				c.builder.WriteString("	-" + fileName + " Incorrect File Name Format!\n")
				incorrectSizes = append(incorrectSizes, fileName)
				continue
			}

			size[2] = strings.ToLower(size[2])
			name := size[0] + "_" + size[1] + "_" + size[2]

			width, height, err := c.GetDimension(absFilePath)
			if err != nil {
				c.builder.WriteString(fmt.Sprintf("	-%s Error: %v\n", name, err))
				continue
			}

			actualSize := Size{width, height}

			correctSize, ok := sizing[size[2]]
			if !ok {
				c.builder.WriteString("	-" + name + " size " + size[2] + " not found in " + sizeType + "\n")
				continue
			}

			if correctSize == actualSize {
				c.builder.WriteString("	-" + name + " Correct!\n")
			} else {
				c.builder.WriteString("	-" + name + " Incorrect!\n")
				incorrectSizes = append(incorrectSizes, name)
			}

		}
	}

	return incorrectSizes
}

func (c *Checker) formatString(data []string) string {
	builder := strings.Builder{}

	for _, str := range data {
		builder.WriteString("\t" + str + "\n")
	}

	return builder.String()
}

func (c *Checker) Check() string {
	c.builder = strings.Builder{}

	c.builder.WriteString(fmt.Sprintf("Scanning folder '%s'\n", c.folderPath))

	files, err := os.ReadDir(c.folderPath)
	if err != nil {
		c.builder.WriteString(fmt.Sprintf("%v\n", err))
		return c.builder.String()
	}

	var incorrectSizes []string
	for _, file := range files {
		filename := file.Name()
		absPath := c.folderPath + "/" + filename

		info, err := os.Stat(absPath)
		if err != nil {
			c.builder.WriteString(fmt.Sprintf("%v\n", err))
			continue
		}

		size := strings.ToLower(filename)
		_, ok := c.template.Sizing[size]

		if info.IsDir() && ok {
			c.builder.WriteString(filename + "\n")
			result := c.ScanFolder(c.folderPath+"/"+filename, size)
			if result != nil {
				incorrectSizes = append(incorrectSizes, result...)
			}
		} else if filename == "Banner.jpg" {
			width, height, err := c.GetDimension(absPath)
			if err != nil {
				c.builder.WriteString(fmt.Sprintf("%v", err))
				continue
			}

			actualSize := Size{width, height}

			if bannerSize == actualSize {
				c.builder.WriteString("Banner.jpg Correct!\n")
			} else {
				c.builder.WriteString("Banner.jpg Incorrect!\n")
				incorrectSizes = append(incorrectSizes, "Banner")
			}

		} else {
			c.builder.WriteString("Skipping " + filename + "\n")
		}
	}

	c.builder.WriteString(fmt.Sprintf("# Failed resize:" + c.formatString(incorrectSizes)))

	return c.builder.String()
}
