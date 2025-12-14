package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

const VERSION = "25.12.14"

const coreFormulasFile = "core_formulas.json"

// var coreFormulaeFilePath = filepath.Join(os.TempDir(), coreFormulasFile)
var cachePath = filepath.Join(os.Getenv("HOME"), ".gobrew")
var coreFormulaeFilePath = filepath.Join(cachePath, coreFormulasFile)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

	// for _, f := range formulas_list {
	// 	if fileDoNotExist("./formulas/" + f + ".json") {
	// 		getFormulaInfo(f)
	// 	}
	// }
}

func generateSVGChart(fileName, chart string) error {

	langStats := make(map[string]int)

	langs := strings.SplitSeq(chart, ",")

	for lang := range langs {
		if len(lang) > 30 {
			fmt.Printf("The language is more than 30 characters long! which is weird! : language=%v\n", lang)
			break
		}

		// if !isFileFound(fileName) || isFileOld(fileName) {
		if isFileOld(fileName) { // if true, either old or not found
			getCoreFormulas(fileName)
		}

		formulas_list, err := getFormulasFromFile(fileName, lang)

		if err != nil {
			fmt.Println("Error getting formulas list: ", err)
		}

		langStats[lang] = len(formulas_list)
	}

	// Prepare data for the chart
	var bars []struct {
		Language string
		Count    int
	}
	for lang, count := range langStats {
		bars = append(bars, struct {
			Language string
			Count    int
		}{lang, count})
	}

	// Sort languages by count (descending) for better visualization
	sort.Slice(bars, func(i, j int) bool {
		return bars[i].Count > bars[j].Count
	})

	// Generate terminal bar chart
	const maxBarWidth = 60

	// Find the maximum count for scaling
	maxCount := 0
	for _, bar := range bars {
		if bar.Count > maxCount {
			maxCount = bar.Count
		}

	}

	for _, bar := range bars {
		barLength := int(float64(bar.Count) / float64(maxCount) * float64(maxBarWidth))
		barStr := strings.Repeat("░", barLength)
		fmt.Printf("%-10s %s %d\n\n", bar.Language, barStr, bar.Count)
	}

	// Generate SVG chart
	var (
		barHeight   int = 30
		barPadding  int = 10
		graphWidth  int = 600
		graphHeight int = 500 // Will be adjusted based on the number of languages
		labelWidth  int = 100
	)

	// Adjust graph height based on the number of languages
	graphHeight = len(bars)*(barHeight+barPadding) + 50 // Add some padding at the top and bottom

	var svg strings.Builder
	svg.WriteString(fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" style="background-color: #f0f0f0;">`, graphWidth+labelWidth, graphHeight))

	// Find the maximum count for scaling
	maxCount = 0
	for _, bar := range bars {
		if bar.Count > maxCount {
			maxCount = bar.Count
		}
	}

	// Draw bars and labels
	for i, bar := range bars {
		y := i*(barHeight+barPadding) + barPadding + 20 // Adjusted y position for top padding
		barWidth := int(float64(bar.Count) / float64(maxCount) * float64(graphWidth))

		// Draw bar
		svg.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" fill="#4CAF50"/>`, labelWidth-40, y, barWidth, barHeight))

		// Draw language label
		svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" dy="%d" font-family="Arial" font-size="14">%s</text>`, labelWidth-90, y, barHeight/2+5, bar.Language))

		// Draw count label
		svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" dy="%d" font-family="Arial" font-size="14">%d</text>`, labelWidth+barWidth-38, y, barHeight/2+5, bar.Count))
	}

	// Add X-axis label
	svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" text-anchor="middle" font-family="Arial" font-size="16">Package Count</text>`, labelWidth+graphWidth/2, graphHeight-10))

	languages_vs := strings.ReplaceAll(chart, ",", " vs ")

	// Add title
	svg.WriteString(fmt.Sprintf(`<text x="%d" y="%d" text-anchor="middle" font-family="Arial" font-size="20">%s Statistics</text>`, (graphWidth+labelWidth)/2, 20, languages_vs))

	svg.WriteString(`</svg>`)

	languages := strings.ReplaceAll(chart, ",", "-")
	timestamp := time.Now().Format("_2006-01-02_15-04-05.svg")
	err := saveToFile(languages+timestamp, svg.String())

	if err != nil {
		return fmt.Errorf("error saving SVG file: %w", err)
	}

	return nil
}

func getPackageCount(fileName, lang string) (string, error) {
	if len(lang) > 30 {
		return "", fmt.Errorf("error: the language is more than 30 characters long! which is weird! : language=%v", lang)
	}

	var langCountCache = filepath.Join(cachePath, lang)
	if _, err := os.Stat(langCountCache); os.IsNotExist(err) {
		// if !isFileFound(fileName) || isFileOld(fileName) {
		if isFileOld(fileName) { // if true, either old or not found
			getCoreFormulas(fileName)
		}
		formulas_list, err := getFormulasFromFile(fileName, lang)
		if err != nil {
			return "", fmt.Errorf("error getting homebrew formulas list: %v", err)
		}
		pkgCount := len(formulas_list)

		outFile, err := os.Create(langCountCache)
		if err != nil {
			return "", fmt.Errorf("error creating langCountCache (lang is %v) file: %v", lang, err)
		}
		defer outFile.Close()

		_, err = fmt.Fprintf(outFile, "%v", pkgCount)
		if err != nil {
			fmt.Println("Error writing to a file: ", err)
			return "", fmt.Errorf("error writing to langCountCache file (%v): %v", langCountCache, err)
		}

		pkgCountStr := strconv.Itoa(pkgCount)

		return pkgCountStr, nil

	} else {
		data, err := os.ReadFile(langCountCache)
		if err != nil {
			return "", fmt.Errorf("error reading langCountCache file (%v): %v", langCountCache, err)
		}
		return string(data), nil
	}

}

func getDependants(fileName, lang string) {
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

	for k, v := range formulas_list {
		fmt.Println("\n", k, ":\n  ", v)
	}
}

func getAllBuildDeps(fileName string) error {
	var allBuildDepsCache = filepath.Join(cachePath, "allBuildDepsCache")
	if _, err := os.Stat(allBuildDepsCache); os.IsNotExist(err) {

		if isFileOld(fileName) { // if true, either old or not found
			getCoreFormulas(fileName)
		}

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

		buildDeps := map[any]struct{}{}

		for _, formula := range formulas {
			if len(formula.BuildDependencies) > 0 {
				for _, dep := range formula.BuildDependencies {
					buildDeps[dep] = struct{}{}
				}
			}
		}

		allBuildDeps := getKeysAsString(buildDeps)

		allbuildDepsStr := fmt.Sprintf("All Build Dependencies Count: %v\n%v", len(allBuildDeps), allBuildDeps)

		fmt.Println(allbuildDepsStr)

		outfile, err := os.Create(allBuildDepsCache)
		if err != nil {
			return err
		}
		defer outfile.Close()

		_, err = outfile.WriteString(allbuildDepsStr)
		if err != nil {
			return err
		}

	} else {
		allBuildDeps, err := os.ReadFile(allBuildDepsCache)
		if err != nil {
			return err
		}
		fmt.Println(string(allBuildDeps))
	}

	return nil
}

type KV struct {
	Key string
	Val int
}

func getAllStatistics(fileName string) error {
	var allStatisticsCache = filepath.Join(cachePath, "allStatisticsCache")
	if _, err := os.Stat(allStatisticsCache); os.IsNotExist(err) {

		if isFileOld(fileName) { // if true, either old or not found
			getCoreFormulas(fileName)
		}

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

		deps := map[string]int{}
		countDeps := func(list []string) {
			for _, dep := range list {
				deps[dep]++
			}
		}

		for _, formula := range formulas {
			countDeps(formula.BuildDependencies)
			countDeps(formula.Dependencies)
			countDeps(formula.TestDependencies)
			countDeps(formula.RecommendedDependencies)
			countDeps(formula.OptionalDependencies)
		}

		fmt.Println("# of all languages/libraries/frameworks: ", len(deps))

		// sort all languages by the count of their packages
		kvPairs := make([]KV, 0, len(data))
		for k, v := range deps {
			kvPairs = append(kvPairs, KV{k, v})
		}
		sort.Slice(kvPairs, func(i, j int) bool {
			return kvPairs[i].Val > kvPairs[j].Val
		})

		var allStatisticsStr string
		for _, pair := range kvPairs {
			allStatisticsStr += fmt.Sprintf("%v: %v\n", pair.Key, pair.Val)
		}

		fmt.Println(allStatisticsStr)

		outfile, err := os.Create(allStatisticsCache)
		if err != nil {
			return err
		}
		defer outfile.Close()

		_, err = outfile.WriteString(allStatisticsStr)
		if err != nil {
			return err
		}

	} else {
		allStatistics, err := os.ReadFile(allStatisticsCache)
		if err != nil {
			return err
		}
		fmt.Println(string(allStatistics))
	}

	return nil
}

func getKeysAsString(m map[any]struct{}) []string {
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

	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		err = os.Mkdir(cachePath, 0755)
		if err != nil {
			fmt.Println("Error creating '~/.gobrew' directory: ", err.Error())
			return
		}
	}

	outFile, err := os.Create(fileName) //os.CreateTemp("", fileName)

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
