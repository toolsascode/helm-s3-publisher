package helpers

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

func GetListPath(list string) []string {

	var paths = []string{}

	var (
		commas     = strings.Split(list, ",")
		semicolons = strings.Split(list, ";")
		spaces     = strings.Split(list, " ")
	)

	paths = append(paths, commas...)
	paths = append(paths, semicolons...)
	paths = append(paths, spaces...)

	return paths
}

func GetExcludePaths(list []string) map[string]bool {

	items := make(map[string]bool)

	for _, l := range list {
		paths := GetListPath(l)
		for _, v := range paths {
			if v != "" {
				t := strings.Trim(v, " ")
				s, _ := strings.CutPrefix(t, "/")
				s, _ = strings.CutPrefix(s, "./")
				items[s] = true
			}
		}
	}

	log.Tracef("GetExcludePaths: %#v", items)

	return items
}
