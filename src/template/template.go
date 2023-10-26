package template

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
	"strings"
	"strconv"
)

type Size struct {
	Width int
	Height int
}

type Controller struct {
	Sizing map[string]map[string]Size
}

func (cr *Controller) Init(filepath string) {
	if filepath == "" {
		filepath = "sizing.csv"
	}

	fmt.Println("Reading Template File'", filepath, "'")

	file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open file", err)
		return
	}

	defer file.Close()
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error reading records")
    }

	numCols := len(records[0])
	numRows := len(records)
	totalSizes := numRows - 1

	items := records[0][1:]
	totalItems := len(items)

	sizes := make([]string, totalSizes)

	for i := 1; i < numRows; i++ {
		sizes[i-1] = records[i][0]
	}

	// fmt.Println(sizes)
	// fmt.Println(numRows)
	// fmt.Println(numCols)
	// fmt.Println(totalSizes)
	// fmt.Println(totalItems)

	cr.Sizing = make(map[string]map[string]Size, totalItems)
	for i := 0; i < totalItems; i++ {
		cr.Sizing[fmt.Sprintf("%s", items[i])] = make(map[string]Size, totalSizes-1)
	}

	for col := 1; col < numCols; col++ {
		for row := 1; row < numRows; row++ {

			if records[row][col] == "" {
				cr.Sizing[fmt.Sprintf("%s", items[col-1])][fmt.Sprintf("%s", sizes[row-1])] = Size{0, 0}
				continue
			}

			sizeSlice := strings.Split(records[row][col], "x")
			width, err := strconv.Atoi(sizeSlice[0])
			if err != nil {
				fmt.Println("Error converting string to integer")
			}

			height, err := strconv.Atoi(sizeSlice[1])
			if err != nil {
				fmt.Println("Error converting string to integer")
			}

			cr.Sizing[fmt.Sprintf("%s", items[col-1])][fmt.Sprintf("%s", sizes[row-1])] = Size{width, height}
		}
	}
}

func (cr *Controller) PrintTemplate() {
    for item, sizes := range cr.Sizing {
        fmt.Println(item)
        for size, value := range sizes {
            fmt.Printf("\t%s: %dx%d\n", size, value.Width, value.Height)
        }
    }
}
