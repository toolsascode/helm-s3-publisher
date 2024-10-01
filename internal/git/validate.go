package git

import (
	"errors"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type BuildInfo struct {
	Version string `json:"version"`
}

type Git struct {
	BuildInfo BuildInfo
}

func CheckIntall() error {

	out, err := exec.Command("git", "version").Output()

	if errors.Is(err, exec.ErrDot) {
		log.Errorln(`
		-> Unable to check the installation of Git binary.
		-> Please check the installation.
		-> Site: https://git-scm.com/`)
		log.Fatalln("CheckIntall::Command", err, out)
		return err
	}
	if err != nil {
		log.Errorln(`
		-> Unable to check the installation of Git binary.
		-> Please check the installation.
		-> Site: https://git-scm.com/`)
		log.Fatalln("CheckIntall::Command", err, out)
		return err
	}

	var (
		git = &Git{
			BuildInfo{
				Version: cases.Title(language.English, cases.Compact).String(strings.Trim(string(out), "\n")),
			},
		}
	)

	log.Infof("|-> %s <-|", git.BuildInfo.Version)

	return nil

}
