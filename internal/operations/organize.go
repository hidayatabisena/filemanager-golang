package operations

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"filemanager/internal/utils"
	"filemanager/pkg/fileutils"
)

func OrganizeFiles(source string, recursive, dryRun bool) error {
	files, err := utils.GetFilePaths(source, recursive)
	if err != nil {
		return fmt.Errorf("error scanning source: %w", err)
	}

	for _, file := range files {
		fileInfo, err := os.Stat(file)
		if err != nil || fileInfo.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file))
		category := fileutils.GetFileCategory(ext)

		modTime := fileInfo.ModTime()
		yearMonth := fmt.Sprintf("%d-%02d", modTime.Year(), modTime.Month())

		categoryDir := filepath.Join(source, category)
		dateDir := filepath.Join(categoryDir, yearMonth)

		_, filename := filepath.Split(file)

		destPath := filepath.Join(dateDir, filename)

		if dryRun {
			fmt.Printf("[DRY RUN] Would organize: %s -> %s\n", file, destPath)
		} else {
			err := os.MkdirAll(dateDir, 0755)
			if err != nil {
				fmt.Printf("Error creating directory %s: %v\n", dateDir, err)
				continue
			}

			fmt.Printf("Organizing: %s -> %s\n", file, destPath)
			err = os.Rename(file, destPath)
			if err != nil {
				fmt.Printf("Error organizing %s: %v\n", file, err)
			}
		}
	}

	return nil
}
