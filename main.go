package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	coreFormulasFile := "core_formulas.json"

	if fileDoNotExist(coreFormulasFile) {
		getCoreFormulas(coreFormulasFile)
	}

	formulas_list, err := getFormulasFromFile(coreFormulasFile)
	if err != nil {
		fmt.Println("Error getting formulas list: ", err)
	}

	for _, f := range formulas_list {
		if fileDoNotExist("./formulas/" + f + ".json") {
			getFormulaInfo(f)
		}
	}
}

type Formula struct {
	Name              string                 `json:"name"`
	FullName          string                 `json:"full_name"`
	Tap               string                 `json:"tap"`
	Oldname           interface{}            `json:"oldname"` // Can be null
	Oldnames          []string               `json:"oldnames"`
	Aliases           []string               `json:"aliases"`
	VersionedFormulae []string               `json:"versioned_formulae"`
	Desc              string                 `json:"desc"`
	License           string                 `json:"license"`
	Homepage          string                 `json:"homepage"`
	Versions          map[string]interface{} `json:"versions"`
	Urls              map[string]interface{} `json:"urls"`
	Revision          int                    `json:"revision"`
	VersionScheme     int                    `json:"version_scheme"`
	Bottle            map[string]interface{} `json:"bottle"`
	// ... other fields (pour_bottle_only_if, keg_only, etc.) are omitted for brevity
}

func getFormulasFromFile(fileName string) ([]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []string{}, err
	}

	var formulas []Formula
	err = json.Unmarshal(data, &formulas)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
		return []string{}, err
	}

	var allFormulas []string
	for _, formula := range formulas {
		allFormulas = append(allFormulas, formula.Name)
	}

	return allFormulas, nil
}

func getFormulaInfo(f string) {
	filePath := "./formulas/" + f + ".json"
	url := "https://formulae.brew.sh/api/formula/" + f + ".json"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error: can not reach API endpoint", err.Error())
		return
	}

	defer resp.Body.Close()

	outFile, err := os.Create(filePath)

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

	fmt.Println("successfully written JSON data into ", filePath)
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

func getCoreFormulas(fileName string) {
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

	outFile, err := os.Create(fileName)

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

	fmt.Println("successfully written JSON data into ", fileName)
}
