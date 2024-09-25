package jfs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Insulince/jlib/pkg/jmust"
)

// Tmp creates a temporary file in the OS temp directory.
// It returns the absolute path to the temp file.
func Tmp() (string, error) {
	program := filepath.Base(os.Args[0])

	// Create a temp file with a prefix
	file, err := os.CreateTemp("", fmt.Sprintf("%s-*", program))
	if err != nil {
		return "", err
	}
	jmust.MustClose(file)

	// Get the absolute path of the temp file
	absPath, err := filepath.Abs(file.Name())
	if err != nil {
		return "", err
	}

	return absPath, nil
}

func MustTmp() string {
	return jmust.Must[string](Tmp)[0]
}

// DeleteAllFilesInDir deletes all files in the specified directory.
func DeleteAllFilesInDir(dirPath string) error {
	// Open the directory
	dir, err := os.Open(dirPath)
	if err != nil {
		return fmt.Errorf("failed to open directory: %w", err)
	}
	defer jmust.MustClose(dir)

	// Get all file names in the directory
	files, err := dir.Readdirnames(-1)
	if err != nil {
		return fmt.Errorf("failed to read directory names: %w", err)
	}

	// Iterate through the files and delete each one
	for _, fileName := range files {
		filePath := filepath.Join(dirPath, fileName)
		err = os.Remove(filePath)
		if err != nil {
			return fmt.Errorf("failed to delete file %s: %w", filePath, err)
		}
	}

	return nil
}

func MustDeleteAllFilesInDir(dirPath string) {
	jmust.Must[any](DeleteAllFilesInDir, dirPath)
}
