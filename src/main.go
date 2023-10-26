package main


import (
	"flag"
	// "fmt"
	// "os"
	// "strings"
	// "path/filepath"
	// "bufio"
	// "image"
    // _ "image/jpeg"
	"github.com/iwahing/image-layout-checker/src/template"
	"github.com/iwahing/image-layout-checker/src/checker"
)

// var meshShort = map[string]Size{
// 	"xs": Size{9056, 2500}, 
// 	"small": Size{9559, 2850},
// 	"medium": Size{10062, 3000}, 
// 	"large":Size{10565, 3150},
// 	"XL": Size{11068, 3300},
// 	"2XL": Size{11571, 3450},
// 	"3XL": Size{12074, 3600}, 
// 	"4XL": Size{12577, 3750}, 
// 	"5XL": Size{13080, 3900},
// }

// var longsleeve = map[string]Size{
// 	"xs": Size{0, 0}, 
// 	"small": Size{0, 0},
// 	"medium": Size{0, 0}, 
// 	"large":Size{0, 0},
// 	"XL": Size{0, 0},
// 	"2XL": Size{0, 0},
// 	"3XL": Size{0, 0}, 
// 	"4XL": Size{0, 0}, 
// 	"5XL": Size{0, 0},
// }

// var tshirt = map[string]Size{
// 	"xs": Size{0, 0}, 
// 	"small": Size{0, 0},
// 	"medium": Size{0, 0}, 
// 	"large":Size{0, 0},
// 	"XL": Size{0, 0},
// 	"2XL": Size{0, 0},
// 	"3XL": Size{0, 0}, 
// 	"4XL": Size{0, 0}, 
// 	"5XL": Size{0, 0},
// }

// var warmer = map[string]Size{
// 	"xs": Size{5775, 4200}, 
// 	"small": Size{6075, 4350},
// 	"medium": Size{6375, 4500}, 
// 	"large": Size{6675, 4650}, 
// 	"XL": Size{6975, 4800},
// 	"2XL": Size{7275, 4950},
// 	"3XL": Size{7575, 5100}, 
// 	"4XL": Size{7875, 5250}, 
// 	"5XL": Size{8175, 5400},
// }

// var warmerSleeves = map[string]Size{
// 	"xs": Size{5074, 3825}, 
// 	"s": Size{3900, 5173},
// 	"m": Size{5272, 3975}, 
// 	"l": Size{5371, 4050},
// 	"xl": Size{5470, 4125},
// 	"2xl": Size{5569, 4200},
// 	"3xl": Size{5668, 4275}, 
// 	"4xl": Size{5767, 4350}, 
// 	"5xl": Size{5866, 4425},
// }

// var hood = Size{4069, 2850}

// var jacket = map[string]Size{
// 	"xs": Size{0, 0}, 
// 	"small": Size{0, 0},
// 	"medium": Size{0, 0}, 
// 	"large":Size{0, 0},
// 	"XL": Size{0, 0},
// 	"2XL": Size{0, 0},
// 	"3XL": Size{0, 0}, 
// 	"4XL": Size{0, 0}, 
// 	"5XL": Size{0, 0},
// }

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

var templateFile string
var folderPath string

func main() {

	flag.StringVar(&templateFile, "template file", "./sizing.csv", "File of sizing to be used as a template")
	flag.StringVar(&folderPath, "path", "./", "Folder path to scan")
	flag.Parse()

	c := template.Controller{}
	c.Init(templateFile)

	checker.Check(c.Sizing, folderPath)
	// fmt.Println("Scanning folder '", folderPath, "'")


	// files, err := os.ReadDir(folderPath)
	// if err != nil {
	// 	fmt.Println(err)
	// }


	// var incorrectSizes []string 
	// for _, file := range files {
	// 	filename := file.Name()
	// 	absPath := folderPath + "/" + filename

	// 	info, err := os.Stat(absPath)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}

	// 	if info.IsDir() {
	// 		if filename == "Jersey" || filename == "Short" {
	// 			fmt.Println(filename)
	// 			result := ScanFolder(folderPath+"/"+filename, filename)
	// 			if result != nil {
	// 				incorrectSizes = append(incorrectSizes, result...)
	// 			}
	// 		} else {
	// 			fmt.Println("Skipping ", filename)
	// 		}
	// 	} else if filename == "Banner.jpg" {
	// 		width, height, err := GetDimension(absPath)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			continue
	// 		}

	// 		actualSize := Size{width, height}

	// 		if bannerSize == actualSize {
	// 			fmt.Println("Banner.jpg Correct!")
	// 		} else {
	// 			fmt.Println("Banner.jpg Incorrect")
	// 			incorrectSizes = append(incorrectSizes, "Banner") 
	// 		}
		
	// 	} else {
	// 		fmt.Println("Skipping ", filename)
	// 	}
	// }

	// fmt.Printf("# Failed resize: %v\n", incorrectSizes)
}
