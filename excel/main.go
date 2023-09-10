package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {
	fmt.Println("len", len(os.Args))
	pfileName := "tmp"
	flag.StringVar(&pfileName, "fileName", "usx", "Filename")
	//This should also be flag
	delimiter := ","

	flag.Parse()
	fmt.Println(pfileName)

	/*for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}*/

	f, err := excelize.OpenFile(pfileName + ".xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get value from cell by given sheet name and axis.
	cellValue, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cellValue)

	// Get all the rows in the sheet.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Iterate over the rows and print the cell values.
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, delimiter)
		}
	}
}
