package helpers

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

/*
It is necessary to add `defer file.Close()`` to process the file
*/
// CreateOrOpenFile <filename>
// This function checks the need to create or just open an existing file before proceeding.
func CreateOrOpenFile(filename string) (*os.File, error) {
	var (
		file *os.File
		err  error
	)
	if FileExists(filename) {
		file, err = os.OpenFile(filename, os.O_RDWR, os.ModePerm)
		if err != nil {
			log.Fatalln("Text::CreateOrOpenFile::Open", err)
			return nil, err
		}
	} else {
		file, err = os.Create(filename)
		if err != nil {
			log.Fatalln("Text::CreateOrOpenFile::Create", err)
			return nil, err
		}
	}

	return file, nil
}
