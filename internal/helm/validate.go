package helm

import (
	"errors"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

type BuildInfo struct {
	Version      string `json:"version"`
	GitCommit    string `json:"git_commit"`
	GitTreeState string `json:"git_tree_state"`
	GoVersion    string `json:"go_version"`
}

type Helm struct {
	BuildInfo BuildInfo
}

func CheckIntall() error {

	out, err := exec.Command("helm", "version", "--short").Output()

	if errors.Is(err, exec.ErrDot) {
		log.Errorln(`
		-> Unable to check the installation of Helm binary.
		-> Please check the installation.
		-> Site: https://helm.sh/`)
		log.Fatal("CheckIntall::Command", err, out)
		return err
	}
	if err != nil {
		log.Errorln(`
		-> Unable to check the installation of Helm binary.
		-> Please check the installation.
		-> Site: https://helm.sh/`)
		log.Fatal("CheckIntall::Command", err, out)
		return err
	}

	var helm = &Helm{
		BuildInfo{
			Version: strings.Trim(string(out), "\n"),
		},
	}

	log.Infof("|-> Helm Version %s <-|", helm.BuildInfo.Version)

	return nil

}
