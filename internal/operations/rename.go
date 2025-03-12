package operations

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"filemanager/internal/utils"
)

func RenameFiles(source, pattern, replacement string, recursive, dryRun bool) error {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("invalid regular expression: %w", err)
	}

	files, err := utils.GetFilePaths(source, recursive)
	if err != nil {
		return fmt.Errorf("error scanning source: %w", err)
	}

	for _, file := range files {
		dir, filename := filepath.Split(file)

		newFilename := re.ReplaceAllString(filename, replacement)

		if newFilename == filename {
			continue
		}

		newPath := filepath.Join(dir, newFilename)

		if dryRun {
			fmt.Printf("[DRY RUN] Would rename: %s -> %s\n", file, newPath)
		} else {
			fmt.Printf("Renaming: %s -> %s\n", file, newPath)
			err := os.Rename(file, newPath)
			if err != nil {
				fmt.Printf("Error renaming %s: %v\n", file, err)
			}
		}
	}

	return nil
}
