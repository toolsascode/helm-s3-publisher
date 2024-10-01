package plugins

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

type HelmS3 struct {
	BuildInfo BuildInfo
}

func S3CheckIntall() error {

	out, err := exec.Command("helm", "s3", "version").Output()

	if errors.Is(err, exec.ErrDot) {
		log.Errorln(`
		-> Unable to check the installation of Helm S3 plugin.
		-> Please check the installation.
		-> Site: https://helm-s3.hypnoglow.io/
		-> GitHub: https://github.com/hypnoglow/helm-s3
		=> Command: helm plugin install https://github.com/hypnoglow/helm-s3.git`)
		log.Fatal("S3CheckIntall", err, out)
		return err
	}
	if err != nil {
		log.Errorln(`
		-> Unable to check the installation of Helm S3 plugin.
		-> Please check the installation.
		-> Site: https://helm-s3.hypnoglow.io/
		-> GitHub: https://github.com/hypnoglow/helm-s3
		=> Command: helm plugin install https://github.com/hypnoglow/helm-s3.git`)
		log.Fatal("S3CheckIntall", err, out)
		return err
	}

	var helm = &HelmS3{
		BuildInfo{
			Version: strings.Trim(string(out), "\n"),
		},
	}

	log.Infof("|-> Helm S3 Plugin Version: %s <-|", helm.BuildInfo.Version)

	return nil

}
