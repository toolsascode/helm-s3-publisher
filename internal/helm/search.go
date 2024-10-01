package helm

import (
	"encoding/json"
	"errors"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

type searchResult struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	AppVersion  string `json:"app_version"`
	Description string `json:"description"`
}

func Search(keyword, version string) (bool, error) {
	out, err := exec.Command("helm", "search", "repo", keyword, "--version", version, "--output", "json").Output()
	if errors.Is(err, exec.ErrDot) {
		log.Fatalf("Search: %q <-> %s", err, out)
		return false, err
	}

	var s []searchResult

	err = json.Unmarshal(out, &s)
	if err != nil {
		log.Fatalf("Search::Unmarshal: %q <-> %s", err, out)
		return false, err
	}

	if len(s) > 0 {
		log.Warnf("|-> [NOk] <-| Helm Chart %s and %s version found!", keyword, version)
		log.Warnf("|-> [NOk] <-| The %s version of the %s chart already exists and it was not possible to continue publishing it!", version, keyword)
		log.Warnln("|-> [NOk] <-| If you want to continue publishing the chart anyway, re-execute the command adding the `--s3-force` flag.")
		return true, nil
	}

	log.Infof("|-> [Ok] <-| Helm Chart %s and %s version, not found!", keyword, version)
	log.Infof("|-> [Ok] <-| We will continue with the publication of the new version of the %s chart.", keyword)

	return false, nil
}
