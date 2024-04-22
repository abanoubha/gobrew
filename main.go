package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	coreFormulasFile := "core_formula.txt"

	if fileDoNotExist(coreFormulasFile) {
		getCoreFormulas()
	}
}

func fileDoNotExist(fileName string) bool {
	_, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("The file", fileName, "does not exist")
			return true
		} else {
			fmt.Println("Error opening the file ", fileName, ": ", err.Error())
			return true
		}
	}
	return false // the file exists
}

func getCoreFormulas() {
	resp, err := http.Get("https://formulae.brew.sh/api/formula.json")

	if err != nil {
		fmt.Println("Error: can not reach API endpoint", err.Error())
		return
	}

	defer resp.Body.Close()

	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("Error reading response body", err.Error())
	//	return
	//}

	outFile, err := os.Create("core_formulas.txt")

	if err != nil {
		fmt.Println("Error creating file: ", err.Error())
		return
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)

	if err != nil {
		fmt.Println("Error writing to a file: ", err.Error())
		return
	}

	fmt.Println("successfully written JSON data into core_formulas.txt")
}
