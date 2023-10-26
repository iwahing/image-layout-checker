package checker

import (
	"fmt"
	// "os"
	// "strings"
	// "path/filepath"
	// "bufio"
	// "image"
    // _ "image/jpeg"
	"github.com/iwahing/image-layout-checker/src/template"
)


// var sizing = map[string]map[string]Size{
// 	"Assorted_Jersey": assortedJersey,
// 	"Jersey": teamJersey,
// 	"Short": short,
// 	"Mesh_Short": meshShort,
// 	"LONGSLEEVE": longsleeve,
// 	"Hoodie": longsleeve,
// 	"TSHIRT": tshirt,
// 	"WARMER": warmer,
// 	"JACKET": jacket,
// }

// var bannerSize = Size{3395, 2396}

// func GetDimension(filepath string) (int, int, error) {
// 	f, err := os.Open(filepath)
// 	if err != nil {
// 		return 0, 0, err
// 	}

// 	reader := bufio.NewReader(f)

// 	// Decode image.
// 	config, _, err:= image.DecodeConfig(reader)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return config.Width, config.Height, nil
// }

// func FileNameWithoutExtSliceNotation(fileName string) string {
// 	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
// }


// func ScanFolder(folderPath string, sizeType string) []string {
// 	files, err := os.ReadDir(folderPath)

// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	}

// 	var incorrectSizes []string
// 	for _, file := range files {
// 		fileName := file.Name()

// 		if filepath.Ext(fileName) == ".jpg" {
// 			absFilePath := folderPath + "/" + fileName
// 			size := strings.ToLower(strings.Split(FileNameWithoutExtSliceNotation(fileName), "_"))
// 			name := size[0] + ":" + size[1]

// 			width, height, err := GetDimension(absFilePath)
// 			if err != nil {
// 				fmt.Println("	-", name, " Error: ", err)
// 				continue
// 			}

// 			actualSize := Size{width, height}

// 			if sizing[sizeType][size[1]] == actualSize {
// 				fmt.Println("	-", name, " Correct!")
// 			} else {
// 				fmt.Println("	-", name, " Incorrect")
// 				incorrectSizes = append(incorrectSizes, name)
// 			}

// 		}
// 	}

// 	return incorrectSizes
// }

func Check(sizing map[string]map[string]template.Size, folderPath string) {
    fmt.Println("Scanning folder '", folderPath, "'")

    files, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println(err)
	}

	var incorrectSizes []string
	for _, file := range files {
		filename := file.Name()
		absPath := folderPath + "/" + filename

		info, err := os.Stat(absPath)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if info.IsDir() {
			if filename == "Jersey" || filename == "Short" {
				fmt.Println(filename)
				result := ScanFolder(folderPath+"/"+filename, filename)
				if result != nil {
					incorrectSizes = append(incorrectSizes, result...)
				}
			} else {
				fmt.Println("Skipping ", filename)
			}
		} else if filename == "Banner.jpg" {
			width, height, err := GetDimension(absPath)
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
}