package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	flagVersion    bool
	flagBuildDep   bool
	flagStatistics bool
	flagLang       string
	flagDependants string
	flagChart      string
	flagReport     string
)

var rootCmd = &cobra.Command{
	Use:   "gobrew",
	Short: "Count all programs written/built in X language or Y build system or Z library distributed via Homebrew.",
	Long:  `Count all programs written/built in X language or Y build system or Z library distributed via Homebrew. Get all build dependencies of all packages in Homebrew Core formulae`,
	Run: func(cmd *cobra.Command, args []string) {

		if flagVersion {

			fmt.Printf(`
gobrew v%v

Software Developer  : Abanoub Hanna
Source code         : https://github.com/abanoubha/gobrew
Twitter             : https://x.com/@AbanoubHA
Developer's Website : https://AbanoubHanna.com`, VERSION)
			return
		}

		if flagBuildDep {
			if err := getAllBuildDeps(coreFormulaeFilePath); err != nil {
				fmt.Println("error: ", err)
			}
			return
		}

		if flagChart != "" {
			if err := generateSVGChart(coreFormulaeFilePath, flagChart); err != nil {
				fmt.Println("error:", err)
			}
			return
		}

		if flagStatistics {
			if err := getAllStatistics(coreFormulaeFilePath); err != nil {
				fmt.Println("error: ", err)
			}
			return
		}

		if flagLang != "" {
			count, err := getPackageCount(coreFormulaeFilePath, flagLang)
			if err != nil {
				fmt.Println("error: ", err)
				return
			}
			fmt.Println(count)
			return
		}

		if flagDependants != "" {
			printDependants(coreFormulaeFilePath, flagDependants)
			return
		}

		// default report
		targetLang := flagReport
		if targetLang == "" {
			targetLang = "go"
		}

		count, err := getPackageCount(coreFormulaeFilePath, targetLang)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		countI, err := strconv.Atoi(count)
		if err != nil {
			fmt.Println("error converting string to int, err:", err)
			return
		}

		if countI < 1 {
			fmt.Printf("There is NO language/library named %s", targetLang)
			return
		}

		fmt.Printf(
			"## Statistics of %s language/library\n\nNumber of CLI apps depend on %s and distributed via Homebrew Core Formulae is %s apps.\n\n",
			targetLang,
			targetLang,
			count,
		)

		formulas, err := getFormulasFromFile(coreFormulaeFilePath, targetLang)
		if err != nil {
			fmt.Println("Error getting formulas list: ", err)
		}

		fmt.Printf("\n### Apps that depend on %s and distributed via Homebrew Core Formulae\n\n", targetLang)

		for k, v := range formulas {
			fmt.Println("-", k, ":", v)
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&flagVersion, "version", "v", false, "show release version")
	rootCmd.Flags().BoolVarP(&flagBuildDep, "build-dep", "b", false, "show building dependencies for all packages in Homebrew Core")
	rootCmd.Flags().BoolVarP(&flagStatistics, "statistics", "s", false, "show all languages and the count of packages which depends on each one of them")

	rootCmd.Flags().StringVarP(&flagLang, "lang", "l", "", "get count for specific language/build-system/library")
	rootCmd.Flags().StringVarP(&flagDependants, "dependants", "d", "", "show all dependants of a certain language/build-system/library")
	rootCmd.Flags().StringVarP(&flagChart, "chart", "c", "", "create an SVG chart for statistics of specified languages (comma separated langs)")
	rootCmd.Flags().StringVarP(&flagReport, "report", "r", "go", "show detailed report of usage statistics of the specified language")
}
