package git

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/toolsascode/helm-s3-publisher/internal/helpers"
)

func LsTree(path string) []string {

	out, err := exec.Command("git", "-C", path, "ls-tree", "-r", "--name-only", "HEAD..").Output()
	if errors.Is(err, exec.ErrDot) {
		log.Fatal(err, out)
		return nil
	}

	if err != nil {
		log.Fatal(err, out)
		return nil
	}

	var listPaths []string
	var items = make(map[string]bool)

	lines := string(out)
	list := regexp.MustCompile("\r?\n").Split(lines, -1)
	excludePaths := helpers.GetExcludePaths(viper.GetStringSlice("git.exclude.paths"))

	for _, v := range list {

		setPath := ""

		if strings.Contains(v, "/") {
			rootPath := strings.Split(filepath.Dir(v), "/")[0]
			if strings.HasSuffix(path, "/") {
				setPath = fmt.Sprintf("%s%s", path, rootPath)
			} else {
				setPath = fmt.Sprintf("%s/%s", path, rootPath)
			}
			if excludePaths[rootPath] {
				setPath = ""
			}
		}

		if !items[setPath] {
			if _, err := os.Stat(setPath); err == nil {
				items[setPath] = true
				listPaths = append(listPaths, setPath)
			}
		}
	}

	log.Tracef("LsTree: %#v", listPaths)

	return listPaths
}

func MergeLsTree(paths []string) []string {
	var listPaths []string
	var items = make(map[string]bool)
	for _, v := range paths {
		if !items[v] {
			items[v] = true
			listPaths = append(listPaths, LsTree(v)...)
		}
	}
	log.Tracef("MergeLsTree: %#v", listPaths)
	return listPaths
}
