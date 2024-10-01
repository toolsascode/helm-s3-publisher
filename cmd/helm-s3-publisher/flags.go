package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile          string
	logLevel         string
	setFormatter     string
	s3Acl            string
	s3ContentType    string
	reportType       string
	reportName       string
	reportOutput     string
	outputPath       string
	dryRun           bool
	logQuiet         bool
	gitLsTree        bool
	helmPackageForce bool
	excludePaths     []string
)

func initFlag() {

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&logQuiet, "quiet", "q", false, "Make some output more `quiet`.")
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Executes the entire process without performing any publishing operations.")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file (default is $HOME/.helm-s3-publisher.yaml or ./.configs/helm-s3-publisher.yaml or .helm-s3-publisher.yaml)")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "info", "Log level [debug, info, warn, error, fatal, panic].")
	rootCmd.PersistentFlags().StringVarP(&setFormatter, "output", "o", "text", "Sets the standard logger formatter. [text, json].")
	rootCmd.PersistentFlags().BoolVarP(&gitLsTree, "git-ls-tree", "g", false, "Enable the Git LS Tree feature and automatically disables the CHART PATHS parameter if it was specified.")
	rootCmd.PersistentFlags().StringSliceVarP(&excludePaths, "exclude-paths", "x", []string{}, "`List` of directories to ignore separated by commas.")
	rootCmd.PersistentFlags().StringVar(&outputPath, "output-path", ".", "`Location` where generated chart files will be saved.")

	rootCmd.PersistentFlags().StringVarP(&reportType, "report", "r", "", "Generate report on helm charts published or not. [json, text and txt].")
	rootCmd.PersistentFlags().StringVarP(&reportOutput, "report-path", "p", ".", "Path where the report should be saved.")
	rootCmd.PersistentFlags().StringVarP(&reportName, "report-name", "n", "helm-s3-publisher", "File name of the generated report without the extension.")

	rootCmd.PersistentFlags().StringVarP(&s3Acl, "s3-acl", "a", "", "S3 Object ACL to use for charts and indexes. Can be sourced from S3_ACL environment variable.")
	rootCmd.PersistentFlags().StringVarP(&s3ContentType, "s3-content-type", "t", "application/gzip", "Set the content-type for the chart file. Can be sourced from S3_CHART_CONTENT_TYPE environment variable.")
	rootCmd.PersistentFlags().BoolVarP(&helmPackageForce, "s3-force", "f", false, "Replace the chart if it already exists. This can cause the repository to lose existing chart; use it with care.")

	_ = viper.BindPFlag("command.dry-run", rootCmd.PersistentFlags().Lookup("dry-run"))
	_ = viper.BindPFlag("helm.s3.force", rootCmd.PersistentFlags().Lookup("s3-force"))
	_ = viper.BindPFlag("helm.s3.acl", rootCmd.PersistentFlags().Lookup("s3-acl"))
	_ = viper.BindPFlag("helm.s3.content-type", rootCmd.PersistentFlags().Lookup("s3-content-type"))
	_ = viper.BindPFlag("output.path", rootCmd.PersistentFlags().Lookup("output-path"))
	_ = viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))
	_ = viper.BindPFlag("log.output.format", rootCmd.PersistentFlags().Lookup("output"))
	_ = viper.BindPFlag("report.type", rootCmd.PersistentFlags().Lookup("report"))
	_ = viper.BindPFlag("report.name", rootCmd.PersistentFlags().Lookup("report-name"))
	_ = viper.BindPFlag("report.path", rootCmd.PersistentFlags().Lookup("report-path"))
	_ = viper.BindPFlag("git.lsTree", rootCmd.PersistentFlags().Lookup("git-ls-tree"))
	_ = viper.BindPFlag("git.exclude.paths", rootCmd.PersistentFlags().Lookup("exclude-paths"))

	viper.SetDefault("author", "Carlos Freitas <carlosrfjunior@gmail.com>")
	viper.SetDefault("license", "MIT")

	viper.AutomaticEnv()

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(docsCmd)
	rootCmd.AddCommand(manCmd)

}
