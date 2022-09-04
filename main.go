package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dlopes7/bmv-xbrl-parser/xbrl"
	"io"
	"os"
	"strings"
)

const DATA_FOLDER = "D:/projects/python/invictus/bmv-scrapper/data"

func getDataFromZipFile(zipFileData []byte) ([]byte, error) {
	r, err := zip.NewReader(bytes.NewReader(zipFileData), int64(len(zipFileData)))
	if err != nil {
		return nil, err
	}
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		defer rc.Close()
		return io.ReadAll(rc)
	}
	return nil, fmt.Errorf("no file found in zip file")
}

func parseToXBRL(data []byte) (xbrl.XBRL, error) {
	var processed xbrl.XBRL
	if err := json.Unmarshal(data, &processed); err != nil {
		return processed, err
	}
	return processed, nil
}

func main() {

	interestingConcepts := []string{
		"Revenue",
		"CostOfSales",
		"GrossProfit",
		"DistributionCosts",
		"AdministrativeExpense",
		"OtherIncome",
		"OtherExpenseByFunction",
		"ProfitLossFromOperatingActivities",
		"ProfitLossBeforeTax",
		"IncomeTaxExpenseContinuingOperations",
		"ProfitLoss",
		"FinanceIncome",
		"InterestIncome",
	}

	// Loop through all files in the data folder
	files, err := os.ReadDir(DATA_FOLDER)
	if err != nil {
		panic(err)
	}

	uniqueUnits := make(map[string]bool)
	for i, file := range files {
		// Open the file
		filePath := DATA_FOLDER + "/" + file.Name()
		fmt.Printf("%d/%d Processing file %s\n", i, len(files), filePath)
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Error opening file %s: %v\n", filePath, err)
			continue
		}
		defer file.Close()
		// Read the file
		fileData, err := io.ReadAll(file)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", filePath, err)
			continue
		}

		// If the file name ends in .zip, get the data from inside the zip file
		if strings.HasSuffix(file.Name(), ".zip") {
			fileData, err = getDataFromZipFile(fileData)
			if err != nil {
				fmt.Printf("Error getting data from zip file %s: %v\n", filePath, err)
				continue
			}
		}

		// Parse the file
		data, err := xbrl.ParseXBRLData(fileData)
		if err != nil {
			panic(err)
		}

		for _, context := range data.QuarterContexts() {
			// fmt.Printf("Context: %s, facts: %d\n", context.ID, len(context.Facts))

			for _, fact := range context.Facts {
				if contains(interestingConcepts, fact.ConceptName) {
					uniqueUnits[fact.UnityID] = true
				}
			}
		}
		if i > 1000 {
			break
		}

	}
	fmt.Printf("Unique units: %+v\n", uniqueUnits)

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
