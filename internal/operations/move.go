package operations

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"filemanager/internal/utils"
)

func MoveFiles(source, dest, pattern string, recursive, dryRun bool) error {
	if !dryRun {
		err := os.MkdirAll(dest, 0755)
		if err != nil {
			return fmt.Errorf("error creating destination directory: %w", err)
		}
	}

	var re *regexp.Regexp
	var err error
	if pattern != "" {
		re, err = regexp.Compile(pattern)
		if err != nil {
			return fmt.Errorf("invalid regular expression: %w", err)
		}
	}

	files, err := utils.GetFilePaths(source, recursive)
	if err != nil {
		return fmt.Errorf("error scanning source: %w", err)
	}

	for _, file := range files {
		_, filename := filepath.Split(file)

		if re != nil && !re.MatchString(filename) {
			continue
		}

		destPath := filepath.Join(dest, filename)

		if dryRun {
			fmt.Printf("[DRY RUN] Would move: %s -> %s\n", file, destPath)
		} else {
			fmt.Printf("Moving: %s -> %s\n", file, destPath)
			err := os.Rename(file, destPath)
			if err != nil {
				fmt.Printf("Error moving %s: %v\n", file, err)
			}
		}
	}

	return nil
}
