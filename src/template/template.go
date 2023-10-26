package template

import ( 
    "encoding/csv"
    "fmt"
    "log"
    "os"
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
	
	var items = records[0][1:]
	fmt.Println(items)
	fmt.Println(len(records))

	column := len(records)
	sizes := make([]string, column-1)

	for i := 1; i < column; i++ { 
		sizes[i-1] = records[i][0]
	}

	fmt.Println(sizes)
}
