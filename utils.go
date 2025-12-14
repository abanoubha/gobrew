package main

import (
	"os"
	"time"
)

func saveToFile(filename, content string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return file.Sync()
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
