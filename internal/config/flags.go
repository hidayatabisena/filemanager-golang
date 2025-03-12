package config

import (
	"flag"
	"fmt"
)

type Config struct {
	Operation   string
	Source      string
	Destination string
	Pattern     string
	Replacement string
	Recursive   bool
	DryRun      bool
}

const (
	RENAME   = "rename"
	MOVE     = "move"
	ORGANIZE = "organize"
)

func ParseFlags() (*Config, error) {
	cfg := &Config{}

	flag.StringVar(&cfg.Operation, "op", "", "Operation type: rename, move, or organize")
	flag.StringVar(&cfg.Source, "src", "", "Source directory or file pattern")
	flag.StringVar(&cfg.Destination, "dest", "", "Destination directory (for move operation)")
	flag.StringVar(&cfg.Pattern, "pattern", "", "Regular expression pattern to match files")
	flag.StringVar(&cfg.Replacement, "replace", "", "Replacement pattern for rename operation")
	flag.BoolVar(&cfg.Recursive, "recursive", false, "Process directories recursively")
	flag.BoolVar(&cfg.DryRun, "dry-run", false, "Show what would be done without actually doing it")

	flag.Parse()

	// Validate inputs
	if cfg.Operation == "" {
		return nil, fmt.Errorf("operation type (-op) is required")
	}

	if cfg.Source == "" {
		return nil, fmt.Errorf("source (-src) is required")
	}

	// Validate operation-specific requirements
	switch cfg.Operation {
	case RENAME:
		if cfg.Pattern == "" || cfg.Replacement == "" {
			return nil, fmt.Errorf("both pattern and replacement are required for rename operation")
		}
	case MOVE:
		if cfg.Destination == "" {
			return nil, fmt.Errorf("destination directory is required for move operation")
		}
	case ORGANIZE:
	default:
		return nil, fmt.Errorf("unsupported operation type: %s", cfg.Operation)
	}

	return cfg, nil
}

func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("  filemanager -op rename -src <source_dir> -pattern <regex> -replace <replacement> [-recursive] [-dry-run]")
	fmt.Println("  filemanager -op move -src <source_dir> -dest <dest_dir> [-pattern <regex>] [-recursive] [-dry-run]")
	fmt.Println("  filemanager -op organize -src <source_dir> [-recursive] [-dry-run]")
	fmt.Println("\nExamples:")
	fmt.Println("  filemanager -op rename -src ./photos -pattern 'IMG_(\\d+)' -replace 'Photo_$1' -recursive")
	fmt.Println("  filemanager -op move -src ./downloads -dest ./sorted -pattern '\\.pdf$' -recursive")
	fmt.Println("  filemanager -op organize -src ./documents -recursive")
}