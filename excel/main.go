package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	fmt.Println("len", len(os.Args))
	pFileName := "tmpFilename"
	pSheetName := "tmpSheetname"
	flag.StringVar(&pFileName, "fileName", "usx", "Filename")
	flag.StringVar(&pSheetName, "sheetName", "Sheet1", "SheetName")
	//This should also be flag
	//labels := []string{}
	//labelCellNumber := [5]string{"A1", "B1"}

	delimiter := ","
	//count := 0

	flag.Parse()
	fmt.Println(pFileName)

	/*for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}*/

	f, err := excelize.OpenFile(pFileName + ".xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// This will store values from column A based on column B
	valuesMap := make(map[string][]string)

	rows, err := f.GetRows(pSheetName)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range rows {
		if len(row) >= 2 { // Ensure that we have at least two columns in the current row
			colAValue := row[0]
			colBValue := row[1]

			// Add to map based on the value in column B
			valuesMap[colBValue] = append(valuesMap[colBValue], colAValue)
		}
	}

	// Open a text file for writing
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Print the map
	for key, values := range valuesMap {
		_, err := writer.WriteString(fmt.Sprintf("Key: %s\n%s\n\n", key, strings.Join(values, delimiter)))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
	writer.Flush()
	//fmt.Printf("\ncount: %d\n", count)
}

// Get all the rows in the sheet.
/*cols, err := f.GetCols("Sheet1")
if err != nil {
	fmt.Println(err)
	return
}
for _, col := range cols {
	for _, rowCell := range col {
		fmt.Print(rowCell, "\t")
	}
	fmt.Println()
}


// Get value from cell by given sheet name and axis.
cellValue, err := f.GetCellValue("Sheet1", labelCellNumber[0])
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(cellValue)

rows, err := f.GetRows("Sheet1")
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println("PRINTING OUT CELLS...")

// Iterate over the rows and print the cell values.
for _, row := range rows {
	for _, colCell := range row {
		fmt.Print(colCell, delimiter)
		count++
	}
}*/
