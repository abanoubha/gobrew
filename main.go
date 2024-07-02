package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const coreFormulasFile = "core_formulas.json"

var (
	buildDep   bool
	statistics bool
	lang       string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "gobrew",
		Short: "Count all programs written/built in X language or Y build system or Z library distributed via Homebrew.",
		Long:  `Count all programs written/built in X language or Y build system or Z library distributed via Homebrew. Get all build dependencies of all packages in Homebrew Core formulae`,
		Example: `gobrew -l go          # count all packages that depend on Go programming language.
gobrew --lang rust    # count all packages that depend on Rust programming language.
gobrew -b             # show all build dependencies of all Homebrew Core formulae.
gobrew -s             # show all languages and the count of packages which depends on each one of them.`,
	}

	rootCmd.Flags().BoolVarP(&buildDep, "build-dep", "b", false, "show building dependencies for all packages in Homebrew Core")

	rootCmd.Flags().BoolVarP(&statistics, "statistics", "s", false, "show all languages and the count of packages which depends on each one of them")

	rootCmd.Flags().StringVarP(&lang, "lang", "l", "", "get count of all packages which have this language/build-system/library as a dependency (required)")

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		if buildDep {
			getAllBuildDeps(coreFormulasFile)
		} else if statistics {
			getAllStatistics(coreFormulasFile)
		} else if lang != "" {
			getPackageCount(coreFormulasFile, lang)
		} else {
			fmt.Println("No language nor build system nor library is specified. Counting packages built in Go (by default):")
			getPackageCount(coreFormulasFile, "go")
		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

	// for _, f := range formulas_list {
	// 	if fileDoNotExist("./formulas/" + f + ".json") {
	// 		getFormulaInfo(f)
	// 	}
	// }
}

func getPackageCount(fileName, lang string) {
	if len(lang) > 30 {
		fmt.Printf("The language is more than 30 characters long! which is weird! : language=%v\n", lang)
		return
	}

	// if !isFileFound(fileName) || isFileOld(fileName) {
	if isFileOld(fileName) { // if true, either old or not found
		getCoreFormulas(fileName)
	}
	formulas_list, err := getFormulasFromFile(fileName, lang)
	if err != nil {
		fmt.Println("Error getting formulas list: ", err)
	}
	pkgCount := len(formulas_list)

	fmt.Println(pkgCount)
}

type Formula struct {
	Name                    string   `json:"name"`
	BuildDependencies       []string `json:"build_dependencies"`
	Dependencies            []string `json:"dependencies"`
	TestDependencies        []string `json:"test_dependencies"`
	RecommendedDependencies []string `json:"recommended_dependencies"`
	OptionalDependencies    []string `json:"optional_dependencies"`
	// Requirements            interface{} `json:"requirements"`
	// FullName                string      `json:"full_name"`
	// UsesFromMacos           interface{} `json:"uses_from_macos"`
	// UsesFromMacosBounds     interface{} `json:"uses_from_macos_bounds"`
	// Tap                     string      `json:"tap"`
	// Oldname                 interface{}            `json:"oldname"`
	// Oldnames                []string               `json:"oldnames"`
	// Aliases                 []string               `json:"aliases"`
	// VersionedFormulae       []string               `json:"versioned_formulae"`
	// Desc                    string                 `json:"desc"`
	// License                 string                 `json:"license"`
	// Homepage                string                 `json:"homepage"`
	// Versions                map[string]interface{} `json:"versions"`
	// Urls                    map[string]interface{} `json:"urls"`
	// Revision                int                    `json:"revision"`
	// VersionScheme           int                    `json:"version_scheme"`
	// Bottle                  map[string]interface{} `json:"bottle"`
	// PourBottleOnlyIf        interface{}            `json:"pour_bottle_only_if"`
	// KeyOnly                 bool                   `json:"keg_only"`
	// KegOnlyReason           interface{}            `json:"keg_only_reason"`
	// Options                 []string               `json:"options"`
	// ConflictsWith           []string               `json:"conflicts_with"`
	// ConflictsWithReasons    []string               `json:"conflicts_with_reasons"`
	// LinkOverwrite           []string               `json:"link_overwrite"`
	// Caveats                 interface{}            `json:"caveats"`
	// Installed               interface{}            `json:"installed"`
	// LinkedKeg               interface{}            `json:"linked_keg"`
	// Pinned                  bool                   `json:"pinned"`
	// Outdated                bool                   `json:"outdated"`
	// Deprecated              bool                   `json:"deprecated"`
	// DeprecationDate         interface{}            `json:"deprecation_date"`
	// DeprecationReason       interface{}            `json:"deprecation_reason"`
	// Disabled                bool                   `json:"disabled"`
	// DisableDate             interface{}            `json:"disable_date"`
	// DisableReason           interface{}            `json:"disable_reason"`
	// PostInstallDefined      bool                   `json:"post_install_defined"`
	// Service                 interface{}            `json:"service"`
	// TapGitHead              string                 `json:"tap_git_head"`
	// RubySourcePath          string                 `json:"ruby_source_path"`
	// RubySourceChecksum      interface{}            `json:"ruby_source_checksum"`
	// Variations              interface{}            `json:"variations"`
}

func getFormulasFromFile(fileName, langName string) (map[interface{}]struct{}, error) {
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

	allFormulas := map[interface{}]struct{}{}

	for _, formula := range formulas {

		langAt := langName + "@"

		// BuildDependencies
		if len(formula.BuildDependencies) > 0 {
			for _, dep := range formula.BuildDependencies {
				if dep == langName || strings.HasPrefix(dep, langAt) {
					allFormulas[formula.Name] = struct{}{}
				}
			}
		}
		// Dependencies
		if len(formula.Dependencies) > 0 {
			for _, dep := range formula.Dependencies {
				if dep == langName || strings.HasPrefix(dep, langAt) {
					allFormulas[formula.Name] = struct{}{}
				}
			}
		}
		// TestDependencies
		if len(formula.TestDependencies) > 0 {
			for _, dep := range formula.TestDependencies {
				if dep == langName || strings.HasPrefix(dep, langAt) {
					allFormulas[formula.Name] = struct{}{}
				}
			}
		}
		// RecommendedDependencies
		if len(formula.RecommendedDependencies) > 0 {
			for _, dep := range formula.RecommendedDependencies {
				if dep == langName || strings.HasPrefix(dep, langAt) {
					allFormulas[formula.Name] = struct{}{}
				}
			}
		}
		// OptionalDependencies
		if len(formula.OptionalDependencies) > 0 {
			for _, dep := range formula.OptionalDependencies {
				if dep == langName || strings.HasPrefix(dep, langAt) {
					allFormulas[formula.Name] = struct{}{}
				}
			}
		}
		// TODO: Requirements
	}

	return allFormulas, nil
}

func getAllBuildDeps(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	var formulas []Formula
	err = json.Unmarshal(data, &formulas)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
		return err
	}

	buildDeps := map[interface{}]struct{}{}

	for _, formula := range formulas {
		if len(formula.BuildDependencies) > 0 {
			for _, dep := range formula.BuildDependencies {
				buildDeps[dep] = struct{}{}
			}
		}
	}

	allBuildDeps := getKeysAsString(buildDeps)

	fmt.Println("All Build Dependencies Count: ", len(allBuildDeps), "\n", allBuildDeps)
	return nil
}

func getAllStatistics(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	var formulas []Formula
	err = json.Unmarshal(data, &formulas)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
		return err
	}

	buildDeps := map[string]int{}

	for _, formula := range formulas {
		if len(formula.BuildDependencies) > 0 {
			for _, dep := range formula.BuildDependencies {
				buildDeps[dep] = buildDeps[dep] + 1
			}
		}
	}

	fmt.Println("All Build Dependencies Count: ", len(buildDeps))

	for k, v := range buildDeps {
		fmt.Println(k, ":", v)
	}

	return nil
}

func getKeysAsString(m map[interface{}]struct{}) []string {
	var keys []string
	for k := range m {
		keys = append(keys, fmt.Sprintf("%v", k))
	}
	return keys
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

// func isFileFound(fileName string) bool {
// 	_, err := os.Open(fileName)
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			fmt.Println("The file", fileName, "does not exist")
// 			return false
// 		} else {
// 			fmt.Println("Error opening the file ", fileName, ": ", err.Error())
// 			return false
// 		}
// 	}
// 	return true
// }

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

func isFileOld(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		// file not found
		return true // consider it old, so we'll re-download it
	}

	sevenDaysAgo := time.Now().AddDate(0, 0, -7)

	return fileInfo.ModTime().Before(sevenDaysAgo)
}
