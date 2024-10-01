package helpers

import (
	"errors"
	"log"
	"os"
)

func CreateDir(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Fatalf("CreateDir: %q", err)
			return err
		}
	}

	return nil
}
