package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// Parse File Args
// Read File
// Map Csv File
// Marshall Json
// Write Json to Disk

func main() {
	// Parse flags in order to get arguments -- since flags area parsed before arguments
	flag.Parse()
	// Get the file directory from the first arguments, that is the only thing that matters to us.
	fileName := flag.Arg(0)
	// read the file and get the content of the file
	fileResults, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	csvReader := csv.NewReader(strings.NewReader(string(fileResults)))
	results, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
	}
	type csvRow map[string]string

	parsedCsv := make([]csvRow, 0)
	csvHeader := make(map[int]string)
	for index, result := range results {
		if index == 0 {
			for key, resultHead := range result {
				csvHeader[key] = resultHead
			}
			continue
		}
		row := csvRow{}
		for key, resultHead := range result {
			row[csvHeader[key]] = resultHead
		}
		parsedCsv = append(parsedCsv, row)
	}

	jsonPrint, err := json.MarshalIndent(parsedCsv, "", "   ")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.Mkdir("results", 0777)
	if err != nil {
		if err == os.ErrPermission {
			fmt.Println("You do not have sufficient permission to create this folder.")
			os.Exit(1)
		}
	}
	err = ioutil.WriteFile("results/"+time.Now().String()+".json", jsonPrint, 0777)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("CSV to JSON conversion complete")
	}
}
