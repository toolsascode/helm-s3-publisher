package helm

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

func CheckIntall() error {

	out, err := exec.Command("helm", "version").Output()

	if errors.Is(err, exec.ErrDot) {
		log.Fatal(err, out)
		return err
	}
	if err != nil {
		log.Fatal(err, out)
		return err
	}

	// var (
	// 	helm      = new(BuildInfo)
	// 	buildInfo = strings.Replace(string(out), "version.BuildInfo", "", 1)
	// )

	// buildInfoJson, err := json.Marshal(buildInfo)
	// if err != nil {
	// 	log.Fatal(err, out)
	// 	return err
	// }

	// if err := json.Unmarshal(buildInfoJson, &helm); err != nil {
	// 	log.Fatal(err, " ", buildInfo)
	// 	return err
	// }

	// log.Infof("%s", buildInfo)

	// log.Infof("Helm Version %v", helm)

	return nil

}
