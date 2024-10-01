package publish

type Commands struct{}

type Report struct {
	ChartName    string `json:"chart_name"`
	ChartVersion string `json:"chart_version"`
	ChartPath    string `json:"chart_path"`
	ChartURL     string `json:"chart_url"`
	RepoName     string `json:"repo_name"`
	GitLsTree    bool   `json:"git_ls_tree"`
	Published    bool   `json:"published"`
	Force        bool   `json:"force"`
}
