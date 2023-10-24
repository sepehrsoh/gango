package utils

import (
	"os"
	"path/filepath"
)

func WriteFile(dir, path string, data string) error {
	err := os.WriteFile(filepath.Join(dir, path), []byte(data), 0777)
	if err != nil {
		return err
	}
	return nil
}
