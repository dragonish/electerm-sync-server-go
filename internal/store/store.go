package store

import (
	"electerm/internal/logger"
	"electerm/internal/model"
	"encoding/json"
	"errors"
	"io/fs"
	"os"
)

// Writer writes content into the file.
func Writer(id string, content interface{}) error {
	if !isPathExists(model.STORE_DIR + "/") {
		logger.Info("create upload directory", "dir", model.STORE_DIR)
		if err := os.MkdirAll(model.STORE_DIR, 0777); err != nil {
			return logger.Err("create upload directory error", err, "dir", model.STORE_DIR)
		}
	}

	filepath := model.STORE_DIR + "/" + id + ".json"

	file, err := os.Create(filepath)
	if err != nil {
		return logger.Err("unable to open file", err, "path", filepath)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(content); err != nil {
		return logger.Err("unable to write file", err, "path", filepath)
	}

	return nil
}

// Reader reads content from the file.
func Reader(id string) string {
	filepath := model.STORE_DIR + "/" + id + ".json"
	if isPathExists(filepath) {
		return filepath
	}
	return ""
}

// isPathExists returns true when the path exists.
func isPathExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, fs.ErrNotExist)
}
