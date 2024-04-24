package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
)

// gobrew [-l|--lang] go
func main() {
	lang := "go" // get count of packags written in "go" language

	if len(os.Args) > 2 {
		allowedArgs := []string{"-l", "--lang"}

		if !slices.Contains(allowedArgs, os.Args[1]) {
			fmt.Printf("The argument (%v) is not supported.\n\nHere is how to use gobrew.\n  gobrew\n  gobrew -l rust\n  gobrew --lang cmake\n", os.Args[1])
			return
		}

		if len(os.Args[2]) < 20 {
			lang = os.Args[2]
		} else {
			fmt.Printf("The language is more than 20 characters long! which is weird! : language=%v\n", os.Args[2])
			return
		}
	}

	coreFormulasFile := "core_formulas.json"

	if fileDoNotExist(coreFormulasFile) {
		getCoreFormulas(coreFormulasFile)
	}

	formulas_list, err := getFormulasFromFile(coreFormulasFile, lang)
	if err != nil {
		fmt.Println("Error getting formulas list: ", err)
	}

	pkgCount := len(formulas_list)

	fmt.Println(pkgCount)

	// for _, f := range formulas_list {
	// 	if fileDoNotExist("./formulas/" + f + ".json") {
	// 		getFormulaInfo(f)
	// 	}
	// }
}

type Formula struct {
	Name                    string                 `json:"name"`
	FullName                string                 `json:"full_name"`
	Tap                     string                 `json:"tap"`
	Oldname                 interface{}            `json:"oldname"`
	Oldnames                []string               `json:"oldnames"`
	Aliases                 []string               `json:"aliases"`
	VersionedFormulae       []string               `json:"versioned_formulae"`
	Desc                    string                 `json:"desc"`
	License                 string                 `json:"license"`
	Homepage                string                 `json:"homepage"`
	Versions                map[string]interface{} `json:"versions"`
	Urls                    map[string]interface{} `json:"urls"`
	Revision                int                    `json:"revision"`
	VersionScheme           int                    `json:"version_scheme"`
	Bottle                  map[string]interface{} `json:"bottle"`
	PourBottleOnlyIf        interface{}            `json:"pour_bottle_only_if"`
	KeyOnly                 bool                   `json:"keg_only"`
	KegOnlyReason           interface{}            `json:"keg_only_reason"`
	Options                 []string               `json:"options"`
	BuildDependencies       []string               `json:"build_dependencies"`
	Dependencies            []string               `json:"dependencies"`
	TestDependencies        []string               `json:"test_dependencies"`
	RecommendedDependencies []string               `json:"recommended_dependencies"`
	OptionalDependencies    []string               `json:"optional_dependencies"`
	UsesFromMacos           interface{}            `json:"uses_from_macos"`
	UsesFromMacosBounds     interface{}            `json:"uses_from_macos_bounds"`
	Requirements            interface{}            `json:"requirements"`
	ConflictsWith           []string               `json:"conflicts_with"`
	ConflictsWithReasons    []string               `json:"conflicts_with_reasons"`
	LinkOverwrite           []string               `json:"link_overwrite"`
	Caveats                 interface{}            `json:"caveats"`
	Installed               interface{}            `json:"installed"`
	LinkedKeg               interface{}            `json:"linked_keg"`
	Pinned                  bool                   `json:"pinned"`
	Outdated                bool                   `json:"outdated"`
	Deprecated              bool                   `json:"deprecated"`
	DeprecationDate         interface{}            `json:"deprecation_date"`
	DeprecationReason       interface{}            `json:"deprecation_reason"`
	Disabled                bool                   `json:"disabled"`
	DisableDate             interface{}            `json:"disable_date"`
	DisableReason           interface{}            `json:"disable_reason"`
	PostInstallDefined      bool                   `json:"post_install_defined"`
	Service                 interface{}            `json:"service"`
	TapGitHead              string                 `json:"tap_git_head"`
	RubySourcePath          string                 `json:"ruby_source_path"`
	RubySourceChecksum      interface{}            `json:"ruby_source_checksum"`
	Variations              interface{}            `json:"variations"`
}

func getFormulasFromFile(fileName, langName string) ([]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	var formulas []Formula
	err = json.Unmarshal(data, &formulas)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
		return nil, err
	}

	var allFormulas []string
	for _, formula := range formulas {
		// BuildDependencies
		if len(formula.BuildDependencies) > 0 {
			for _, dep := range formula.BuildDependencies {
				if dep == langName {
					allFormulas = append(allFormulas, formula.Name)
				}
			}
		}
		// Dependencies
		if len(formula.Dependencies) > 0 {
			for _, dep := range formula.Dependencies {
				if dep == langName {
					allFormulas = append(allFormulas, formula.Name)
				}
			}
		}
		// TestDependencies
		if len(formula.TestDependencies) > 0 {
			for _, dep := range formula.TestDependencies {
				if dep == langName {
					allFormulas = append(allFormulas, formula.Name)
				}
			}
		}
		// RecommendedDependencies
		if len(formula.RecommendedDependencies) > 0 {
			for _, dep := range formula.RecommendedDependencies {
				if dep == langName {
					allFormulas = append(allFormulas, formula.Name)
				}
			}
		}
		// OptionalDependencies
		if len(formula.OptionalDependencies) > 0 {
			for _, dep := range formula.OptionalDependencies {
				if dep == langName {
					allFormulas = append(allFormulas, formula.Name)
				}
			}
		}
		// TODO: Requirements
	}

	return allFormulas, nil
}

// func getFormulaInfo(f string) {
// 	filePath := "./formulas/" + f + ".json"
// 	url := "https://formulae.brew.sh/api/formula/" + f + ".json"

// 	resp, err := http.Get(url)

// 	if err != nil {
// 		fmt.Println("Error: can not reach API endpoint", err.Error())
// 		return
// 	}

// 	defer resp.Body.Close()

// 	outFile, err := os.Create(filePath)

// 	if err != nil {
// 		fmt.Println("Error creating file: ", err.Error())
// 		return
// 	}

// 	defer outFile.Close()

// 	_, err = io.Copy(outFile, resp.Body)

// 	if err != nil {
// 		fmt.Println("Error writing to a file: ", err.Error())
// 		return
// 	}

// 	fmt.Println("successfully written JSON data into ", filePath)
// }

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
