package plugins

import (
	"errors"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

type BuildInfo struct {
	Version      string `json:"version"`
	GitCommit    string `json:"git_commit"`
	GitTreeState string `json:"git_tree_state"`
	GoVersion    string `json:"go_version"`
}

func S3CheckIntall() error {

	out, err := exec.Command("helm", "s3", "version").Output()

	if errors.Is(err, exec.ErrDot) {
		err = nil
		log.Fatal(err, out)
		return err
	}
	if err != nil {
		log.Errorln(`
		-> Unable to check the installation of Helm S3 plugin.
		-> Please check the installation.
		-> Site: https://helm-s3.hypnoglow.io/
		-> GitHub: https://github.com/hypnoglow/helm-s3
		=> Command: helm plugin install https://github.com/hypnoglow/helm-s3.git`)
		log.Fatal(err, out)
		return err
	}

	// if err := json.Unmarshal(out, &helm); err != nil {
	// 	log.Fatal(err)
	// 	return err
	// }

	// log.Infof("%s", buildInfo)

	// log.Infof("Helm s3 plugin version %v", helm)

	return nil

}
