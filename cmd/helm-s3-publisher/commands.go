package main

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/toolsascode/helm-s3-publisher/pkg/publisher"

	mango "github.com/muesli/mango-cobra"
	"github.com/muesli/roff"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "helm-s3-publiher REPO [PATHS]",
	Version: getVersion(),
	Short:   "Helm S3 Publisher is a project to help with the process of publishing new releases for charts.",
	Long: `
The Helm S3 Publisher project makes it easy to publish new releases to the AWS S3 Bucket.

Ideal for automating pipelines and follows these steps:
	1. Check the minimum requirements to start the process;
	2. The Git LS Tree feature is built into the CLI and helps you automatically check which charts have changed and will be updated.
	3. Validates whether the changed chart already has a published version. 
	   It is possible to force and override the version that exists in the repository. 
	   We do not recommend using this functionality in production, only in necessary cases.
	4. Then, package the chart.
	5. Finally, publish the chart to the indicated AWS S3 Bucket using the helm s3 plugin.
	   See: https://github.com/hypnoglow/helm-s3

$ helm s3-publisher REPO [CHART PATHS] [flags]

- REPO 	=> (Required)
	Repository for searching and publishing the new version of the chart.

- CHART PATHS => (Optional and Default: . ) 
	List of charts directories separated by commas.
	If the Git LS Tree feature is enabled, the CLI will attempt to identify all changed chart directories indicated in the PATHS parameter.
	Example: "dir-chart-1,dir-chart-2"

Complete documentation is available at https://github.com/toolsascode/helm-s3-publisher
	
	`,
	Args: func(_ *cobra.Command, args []string) error {

		if len(args) < 1 {
			log.Errorln("Repository name is required!!!")
			return errors.New("it is necessary to provide at least the name of the helm chart repository")
		}
		viper.Set("chart.repo", args[0])

		if len(args) == 2 {
			viper.Set("chart.paths", args[1])
		} else {
			viper.Set("chart.paths", ".")
		}

		log.Infof("%#v", args)
		log.Infof("Repo: %s and Paths: %s", viper.Get("chart.repo"), viper.Get("chart.paths"))

		return nil

	},
	Run: func(_ *cobra.Command, _ []string) {

		publisher.Publisher()

	},
}

var versionCmd = &cobra.Command{
	Use:                "version",
	Short:              "Print the version number of Helm S3 Publisher",
	Long:               `All software has versions. This is Helm S3 Publisher's`,
	DisableFlagParsing: true,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Version\t:", version)
		fmt.Println("Date\t:", date)
		fmt.Println("Commit\t: ", commit)
		fmt.Println("Built by: ", builtBy)
	},
}

var docsCmd = &cobra.Command{
	Use:                   "docs",
	Short:                 "Generating Helm S3 Publisher CLI markdown documentation.",
	Long:                  `Allow generating documentation in markdown format for Helm S3 Publisher CLI internal commands`,
	Hidden:                true,
	DisableFlagParsing:    true,
	SilenceUsage:          true,
	DisableFlagsInUseLine: true,
	Args:                  cobra.NoArgs,
	ValidArgsFunction:     cobra.NoFileCompletions,
	Run: func(_ *cobra.Command, _ []string) {

		var path = "./docs"

		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(path, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}

		log.Info("Generating markdown documentation")
		err := doc.GenMarkdownTree(rootCmd, path)
		if err != nil {
			log.Fatal(err)
		}

		log.Infof("Documentation successfully generated in %s", path)
	},
}

var manCmd = &cobra.Command{
	Use:                   "man",
	Short:                 "Generates manpages",
	DisableFlagParsing:    true,
	SilenceUsage:          true,
	DisableFlagsInUseLine: true,
	Hidden:                true,
	Args:                  cobra.NoArgs,
	RunE: func(_ *cobra.Command, _ []string) error {
		manPage, err := mango.NewManPage(1, rootCmd.Root())
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, manPage.Build(roff.NewDocument()))
		return err
	},
}
