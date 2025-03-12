package operations

import (
	"fmt"

	"filemanager/internal/config"
)

func Execute(cfg *config.Config) error {
	switch cfg.Operation {
	case config.RENAME:
		return RenameFiles(cfg.Source, cfg.Pattern, cfg.Replacement, cfg.Recursive, cfg.DryRun)
	case config.MOVE:
		return MoveFiles(cfg.Source, cfg.Destination, cfg.Pattern, cfg.Recursive, cfg.DryRun)
	case config.ORGANIZE:
		return OrganizeFiles(cfg.Source, cfg.Recursive, cfg.DryRun)
	default:
		return fmt.Errorf("unsupported operation: %s", cfg.Operation)
	}
}