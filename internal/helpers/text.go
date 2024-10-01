package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
)

func CreateFileTextPlain(report any, filename string) {

	fileJSON, err := json.Marshal(report)
	if err != nil {
		log.Errorln("GenerateReport::Marshal", err)
		return
	}

	var v interface{}

	err = json.Unmarshal(fileJSON, &v)
	if err != nil {
		log.Errorln("CreateFileTextPlain::Unmarshal", err)
		return
	}

	data := v.([]interface{})

	file, err := CreateOrOpenFile(filename)
	if err != nil {
		log.Errorln("CreateFileTextPlain::CreateOrOpenFile", err)
		return
	}

	defer file.Close()

	for k, v := range data {
		switch v := v.(type) {
		case string:
			writeFile(file, v)
		case float64:
			writeFile(file, fmt.Sprintf("%.2f", v))
		case []interface{}:
			var list []string
			for _, value := range v {
				list = append(list, fmt.Sprintf("%v", value))
			}
			sort.Strings(list)
			writeFile(file, strings.Join(list, " "))
		case map[string]interface{}:
			var list []string
			for key, value := range v {
				log.Tracef("%v: %v", key, value)
				list = append(list, fmt.Sprintf("%v", value))
			}
			log.Tracef("%s", strings.Repeat("=", 90))
			sort.Strings(list)
			writeFile(file, strings.Join(list, ";"))

		default:
			log.Infoln("Text::CreateFileTextPlain::unknown", k, v)
		}
	}
}

func writeFile(file *os.File, line string) {

	_, err := fmt.Fprintln(file, line)
	if err != nil {
		log.Fatalln("Text::writeFile::WriteString", err)
	}

	// file.Sync()
}
