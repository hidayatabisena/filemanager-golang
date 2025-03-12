package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetFilePaths(source string, recursive bool) ([]string, error) {
	var files []string

	fileInfo, err := os.Stat(source)
	if err != nil {
		matches, globErr := filepath.Glob(source)
		if globErr != nil {
			return nil, fmt.Errorf("invalid source path or pattern: %w", err)
		}
		return matches, nil
	}

	if fileInfo.IsDir() {
		walkFunc := func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() && path != source {
				if !recursive {
					return filepath.SkipDir
				}
				return nil
			}

			if !info.IsDir() {
				files = append(files, path)
			}

			return nil
		}

		err = filepath.Walk(source, walkFunc)
		if err != nil {
			return nil, err
		}
	} else {
		files = append(files, source)
	}

	return files, nil
}
