package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

func PrintHelp() {
	fmt.Println("aab [-n worksheet name] filename.xlsx")
}

func main() {
	var sheetname, filename string
	flag.StringVar(&sheetname, "n", "", "Worksheet name")
	flag.Parse()

	if filename = flag.Arg(0); filename == "" {
		PrintHelp()
		os.Exit(1)
	}

	f, e := excelize.OpenFile(filename)
	if e != nil {
		log.Fatal(e)
	}
	defer func() {
		if e = f.Close(); e != nil {
			log.Fatal(e)
		}
	}()

	var rows [][]string
	if sheetname == "" {
		sheetname = f.GetSheetList()[0]
	}
	rows, e = f.GetRows(sheetname)
	if e != nil {
		log.Fatal(e)
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
