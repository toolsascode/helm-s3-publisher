package helpers

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

func CreateFileJSON(data any, filename string) {

	fileJSON, err := json.Marshal(data)
	if err != nil {
		log.Errorln("helpers::CreateFileJson::Marshal", err)
		return
	}

	if err = os.WriteFile(filename, fileJSON, 0755); err != nil {
		log.Errorln("helpers::CreateFileJson", err)
	}
}

func CreateFilePrettyJSON(data any, filename string) {

	fileJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Errorln("helpers::CreateFilePrettyJson::MarshalIndent", err)
		return
	}

	if err = os.WriteFile(filename, fileJSON, 0755); err != nil {
		log.Errorln("helpers::CreateFilePrettyJson", err)
	}
}
