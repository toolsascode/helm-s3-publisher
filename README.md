# helm-s3-publisher 
[![Go Reference](https://pkg.go.dev/badge/github.com/toolsascode/helm-s3-publisher.svg)](https://pkg.go.dev/github.com/toolsascode/helm-s3-publisher) [![Testing](https://github.com/toolsascode/helm-s3-publisher/actions/workflows/go.yml/badge.svg)](https://github.com/toolsascode/helm-s3-publisher/actions/workflows/go.yml) [![Published Versions](https://github.com/toolsascode/helm-s3-publisher/actions/workflows/releaser.yml/badge.svg)](https://github.com/toolsascode/helm-s3-publisher/actions/workflows/releaser.yml)

Helm S3 Publisher is a small project with the purpose of helping in the process of publishing new helm charts using the helm s3 plugin already known by the community.

## Requirements
- [helm]()
- [helm-s3](https://github.com/hypnoglow/helm-s3)

## How to install?

### Helm Plugin

```shell
helm plugin install https://github.com/toolsascode/helm-s3-publisher.git
```

or

```shell
helm plugin install https://github.com/toolsascode/helm-s3-publisher.git --version 1.1.0-beta
```

### Binary via Github

```shell
curl -fLSs https://raw.githubusercontent.com/toolsascode/helm-s3-publisher/main/scripts/install.sh | bash
```

Or

```shell
curl -fLSs https://raw.githubusercontent.com/toolsascode/helm-s3-publisher/main/scripts/install.sh | sudo bash
```

### Go Install

```shell
go install github.com/toolsascode/helm-s3-publisher/cmd/helm-s3-publisher@latest
```

### Homebrew

```shell
brew install toolsascode/tap/helm-s3-publisher
```

### Scoop

1. Run **PowerShell as an Administrator** and:
2. To add this bucket, run `scoop bucket add helm-s3-publisher-scoop https://github.com/toolsascode/scoop-bucket`.
3. To install, do `scoop install helm-s3-publisher`.

## Usage
- Ideal for automating pipelines and follows these steps:
    - Check the minimum requirements to start the process;
    - The Git LS Tree feature is built into the CLI and helps you automatically check which charts have changed and will be updated.
    - Validates whether the changed chart already has a published version. 
    - It is possible to force and override the version that exists in the repository. 
    - We do not recommend using this functionality in production, only in necessary cases.
    - Then, package the chart.
    - Finally, publish the chart to the indicated AWS S3 Bucket using the helm s3 plugin.

```shell
helm s3-publisher REPO [CHART PATHS] [flags]
```
### Examples:

1. In this first example, the CLI will search for directories of the changed charts and publish them to the repository automatically. Since in this case we chose to use the Git Ls Tree feature, it is good practice to exclude paths that should not be processed and avoid failures. At the end, a JSON file will be generated with all interactions performed, published or not.

```shell
helm s3-publisher myrepo /path/to/helm-charts \
--git-ls-tree --exclude-paths ".git, .github" \
--log-level debug --report json
```

2. In the following example, we inform exactly which chart(s) were manually changed.

```shell
helm s3-publisher myrepo \
/path/to/helm-charts/chart1,/path/to/helm-charts/chart2 \
--log-level debug --report json
```

| Inputs | Required | Description |
|---     | :---:       |---          |
**REPO** | Yes | _(Required)_ Repository for searching and publishing the new version of the chart. |
**CHART PATHS** | No | _(Optional and Default: . )_ List of charts directories separated by commas. If the **Git LS Tree** feature is enabled, the CLI will attempt identify all changed chart directories indicated in the `PATHS` parameter. **Example:** _"dir-chart-1,dir-chart-2"_ |

## SEE ALSO

- [helm-s3-publisher](docs/helm-s3-publisher.md) - All available CLI options
