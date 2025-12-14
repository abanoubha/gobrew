package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func ensureFileExists(fileName string) error {
	if isFileOld(fileName) {
		return getCoreFormulas(fileName)
	}
	return nil
}

type Formula struct {
	Name                    string   `json:"name"`
	BuildDependencies       []string `json:"build_dependencies"`
	Dependencies            []string `json:"dependencies"`
	TestDependencies        []string `json:"test_dependencies"`
	RecommendedDependencies []string `json:"recommended_dependencies"`
	OptionalDependencies    []string `json:"optional_dependencies"`
	Desc                    string   `json:"desc"`
	// Requirements            interface{} `json:"requirements"`
	// FullName                string      `json:"full_name"`
	// UsesFromMacos           interface{} `json:"uses_from_macos"`
	// UsesFromMacosBounds     interface{} `json:"uses_from_macos_bounds"`
	// Tap                     string      `json:"tap"`
	// Oldname                 interface{}            `json:"oldname"`
	// Oldnames                []string               `json:"oldnames"`
	// Aliases                 []string               `json:"aliases"`
	// VersionedFormulae       []string               `json:"versioned_formulae"`
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

func getFormulasFromFile(fileName, langName string) (map[string]string, error) {
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

	allFormulas := map[string]string{}
	langAt := langName + "@"

	isMatch := func(dep string) bool {
		return dep == langName || strings.HasPrefix(dep, langAt)
	}

	for _, formula := range formulas {
		matched := false

		checkDeps := func(deps []string) {
			for _, dep := range deps {
				if isMatch(dep) {
					matched = true
					return
				}
			}
		}

		checkDeps(formula.BuildDependencies)
		if !matched {
			checkDeps(formula.Dependencies)
		}
		if !matched {
			checkDeps(formula.TestDependencies)
		}
		if !matched {
			checkDeps(formula.RecommendedDependencies)
		}
		if !matched {
			checkDeps(formula.OptionalDependencies)
		}
		// TODO: Requirements

		if matched {
			allFormulas[formula.Name] = formula.Desc
		}
	}

	return allFormulas, nil
}
