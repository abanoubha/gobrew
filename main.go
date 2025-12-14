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

	if err := ensureFileExists(fileName); err != nil {
		return err
	}

	for lang := range langs {
		lang = strings.TrimSpace(lang)
		if len(lang) > 30 {
			fmt.Printf("The language is more than 30 characters long! which is weird! : language=%v\n", lang)
			break
		}

		formulasList, err := getFormulasFromFile(fileName, lang)
		if err != nil {
			fmt.Println("Error getting formulas list: ", err)
			return err
		}

		langStats[lang] = len(formulasList)
	}

	// Prepare data for the chart
	type bar struct {
		Language string
		Count    int
	}
	var bars []bar
	maxCount := 0

	for lang, count := range langStats {
		bars = append(bars, bar{lang, count})
		if count > maxCount {
			maxCount = count
		}
	}

	// Sort languages by count (descending) for better visualization
	sort.Slice(bars, func(i, j int) bool {
		return bars[i].Count > bars[j].Count
	})

	// Generate terminal bar chart
	const maxBarWidth = 60

	for _, bar := range bars {
		barLength := int(float64(bar.Count) / float64(maxCount) * float64(maxBarWidth))
		fmt.Printf("%-10s %s %d\n\n", bar.Language, strings.Repeat("░", barLength), bar.Count)
	}

	// Generate SVG chart
	barHeight := 30
	barPadding := 10
	graphWidth := 600
	labelWidth := 100
	graphHeight := len(bars)*(barHeight+barPadding) + 50 // Add some padding at the top and bottom

	var svg strings.Builder
	fmt.Fprintf(&svg, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" style="background-color: #f0f0f0;">`, graphWidth+labelWidth, graphHeight)

	// Draw bars and labels
	for i, bar := range bars {
		y := i*(barHeight+barPadding) + barPadding + 20 // Adjusted y position for top padding
		barWidth := 0
		if maxCount > 0 {
			barWidth = int(float64(bar.Count) / float64(maxCount) * float64(graphWidth))
		}

		// Draw bar
		fmt.Fprintf(&svg, `<rect x="%d" y="%d" width="%d" height="%d" fill="#4CAF50"/>`, labelWidth-40, y, barWidth, barHeight)

		// Draw language label
		fmt.Fprintf(&svg, `<text x="%d" y="%d" dy="%d" font-family="Arial" font-size="14">%s</text>`, labelWidth-90, y, barHeight/2+5, bar.Language)

		// Draw count label
		fmt.Fprintf(&svg, `<text x="%d" y="%d" dy="%d" font-family="Arial" font-size="14">%d</text>`, labelWidth+barWidth-38, y, barHeight/2+5, bar.Count)
	}

	// Add X-axis label
	fmt.Fprintf(&svg, `<text x="%d" y="%d" text-anchor="middle" font-family="Arial" font-size="16">Package Count</text>`, labelWidth+graphWidth/2, graphHeight-10)

	languages_vs := strings.ReplaceAll(chart, ",", " vs ")

	// Add title
	fmt.Fprintf(&svg, `<text x="%d" y="%d" text-anchor="middle" font-family="Arial" font-size="20">%s Statistics</text></svg>`, (graphWidth+labelWidth)/2, 20, languages_vs)

	filename2save := strings.ReplaceAll(chart, ",", "-") + time.Now().Format("_2006-01-02_15-04-05.svg")
	if err := saveToFile(filename2save, svg.String()); err != nil {
		return fmt.Errorf("error saving SVG file: %w", err)
	}

	return nil
}

func getPackageCount(fileName, lang string) (string, error) {
	if len(lang) > 30 {
		return "", fmt.Errorf("error: the language is more than 30 characters long! which is weird! : language=%v", lang)
	}

	if err := ensureFileExists(fileName); err != nil {
		return "", err
	}

	langCountCache := filepath.Join(cachePath, lang)

	if _, err := os.Stat(langCountCache); err == nil {
		data, err := os.ReadFile(langCountCache)
		if err == nil {
			return string(data), nil
		}
	}

	// calculate if cache missing
	formulasList, err := getFormulasFromFile(fileName, lang)
	if err != nil {
		return "", fmt.Errorf("error getting homebrew formulas list: %w", err)
	}

	pkgCount := len(formulasList)
	pkgCountStr := strconv.Itoa(pkgCount)

	if err := saveToFile(langCountCache, pkgCountStr); err != nil {
		fmt.Printf("error caching: %w\n", err)
	}

	return pkgCountStr, nil
}

func printDependants(fileName, lang string) {
	if len(lang) > 30 {
		fmt.Printf("The language is more than 30 characters long! which is weird! : language=%v\n", lang)
		return
	}

	if err := ensureFileExists(fileName); err != nil {
		fmt.Println(err)
		return
	}

	formulasList, err := getFormulasFromFile(fileName, lang)
	if err != nil {
		fmt.Println("Error getting formulas list: ", err)
		return
	}

	for k, v := range formulasList {
		fmt.Printf("\n%s:\n %s\n", k, v)
	}
}

func getAllBuildDeps(fileName string) error {
	if err := ensureFileExists(fileName); err != nil {
		return err
	}

	cacheFile := filepath.Join(cachePath, "allBuildDepsCache")
	if _, err := os.Stat(cacheFile); err == nil {
		data, err := os.ReadFile(cacheFile)
		if err == nil {
			fmt.Println(string(data))
			return nil
		}
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	var formulas []Formula
	if err := json.Unmarshal(data, &formulas); err != nil {
		return err
	}

	buildDeps := map[string]struct{}{}
	for _, f := range formulas {
		for _, dep := range f.BuildDependencies {
			buildDeps[dep] = struct{}{}
		}
	}

	var keys []string
	for k := range buildDeps {
		keys = append(keys, k)
	}

	output := fmt.Sprintf("All Build Dependencies Count: %v\n%v", len(keys), keys)
	fmt.Println(output)

	if err := saveToFile(cacheFile, output); err != nil {
		fmt.Println("error caching:", err)
	}

	return nil
}

func getAllStatistics(fileName string) error {
	if err := ensureFileExists(fileName); err != nil {
		return err
	}

	statCache := filepath.Join(cachePath, "allStatisticsCache")
	if _, err := os.Stat(statCache); err == nil {
		data, err := os.ReadFile(statCache)
		if err == nil {
			fmt.Println(string(data))
			return nil
		}
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	var formulas []Formula
	if err = json.Unmarshal(data, &formulas); err != nil {
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

	type KV struct {
		Key string
		Val int
	}

	// sort all languages by the count of their packages
	kvPairs := make([]KV, 0, len(data))
	for k, v := range deps {
		kvPairs = append(kvPairs, KV{k, v})
	}
	sort.Slice(kvPairs, func(i, j int) bool {
		return kvPairs[i].Val > kvPairs[j].Val
	})

	var result string
	for _, pair := range kvPairs {
		result += fmt.Sprintf("%v: %v\n", pair.Key, pair.Val)
	}

	fmt.Println(result)

	if err := saveToFile(statCache, result); err != nil {
		fmt.Println("error caching: ", err)
	}

	return nil
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

func getCoreFormulas(fileName string) error {
	resp, err := http.Get("https://formulae.brew.sh/api/formula.json")
	if err != nil {
		return fmt.Errorf("error: can not reach API endpoint: %w", err)
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
			return fmt.Errorf("error creating '%s' directory: %w", cachePath, err)
		}
	}

	outFile, err := os.Create(fileName) //os.CreateTemp("", fileName)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("error writing to a file: %w", err)
	}

	fmt.Println("successfully written JSON data into ", fileName)
	return nil
}
