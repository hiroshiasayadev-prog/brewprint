package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

// Generate reads every .puppy.yaml declaration and .manifest.yaml binding
// beneath root and atomically replaces root/go_native/puppygen.
func Generate(root string) error {
	loaded, err := loadProgram(root)
	if err != nil {
		return err
	}

	goNativeRoot := filepath.Join(loaded.Root, "go_native")
	if err := os.MkdirAll(goNativeRoot, 0o755); err != nil {
		return fmt.Errorf("create Go native root: %w", err)
	}

	temporaryDirectory, err := os.MkdirTemp(goNativeRoot, ".puppygen-build-")
	if err != nil {
		return fmt.Errorf("create temporary generated directory: %w", err)
	}
	removeTemporary := true
	defer func() {
		if removeTemporary {
			_ = os.RemoveAll(temporaryDirectory)
		}
	}()

	if err := renderProgram(loaded, temporaryDirectory); err != nil {
		return err
	}

	outputDirectory := filepath.Join(goNativeRoot, "puppygen")
	backupDirectory := ""
	if _, statErr := os.Stat(outputDirectory); statErr == nil {
		backupDirectory, err = reserveBackupPath(goNativeRoot)
		if err != nil {
			return err
		}
		if err := os.Rename(outputDirectory, backupDirectory); err != nil {
			return fmt.Errorf("move existing generated directory to backup: %w", err)
		}
	} else if !os.IsNotExist(statErr) {
		return fmt.Errorf("stat generated directory: %w", statErr)
	}

	if err := os.Rename(temporaryDirectory, outputDirectory); err != nil {
		if backupDirectory != "" {
			if rollbackErr := os.Rename(backupDirectory, outputDirectory); rollbackErr != nil {
				return fmt.Errorf(
					"activate generated directory: %w; restore previous output: %v",
					err,
					rollbackErr,
				)
			}
		}
		return fmt.Errorf("activate generated directory: %w", err)
	}
	removeTemporary = false

	if backupDirectory != "" {
		if err := os.RemoveAll(backupDirectory); err != nil {
			return fmt.Errorf("remove previous generated directory backup %s: %w", backupDirectory, err)
		}
	}
	return nil
}

func reserveBackupPath(root string) (string, error) {
	path, err := os.MkdirTemp(root, ".puppygen-backup-")
	if err != nil {
		return "", fmt.Errorf("reserve generated directory backup: %w", err)
	}
	if err := os.Remove(path); err != nil {
		return "", fmt.Errorf("prepare generated directory backup path: %w", err)
	}
	return path, nil
}
