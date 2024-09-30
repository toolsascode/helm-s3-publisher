package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile          string
	logLevel         string
	outputPath       string
	setFormatter     string
	logQuiet         bool
	gitLsTree        bool
	helmPackageForce bool
	excludePaths     []string
)

func initFlag() {

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&logQuiet, "quiet", "q", false, "Make some output more `quiet`.")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file (default is $HOME/.helm-s3-publisher.yaml or ./.configs/helm-s3-publisher.yaml or .helm-s3-publisher.yaml)")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "info", "Log level [debug, info, warn, error, fatal, panic]")
	rootCmd.PersistentFlags().StringVarP(&setFormatter, "output-format", "o", "text", "Sets the standard logger formatter. [text, json] ")
	rootCmd.PersistentFlags().BoolVarP(&gitLsTree, "git-ls-tree", "g", false, "Enable the Git LS Tree feature and automatically disables the CHART PATHS parameter if it was specified.")
	rootCmd.PersistentFlags().BoolVarP(&helmPackageForce, "force", "f", false, "Replace the chart if it already exists. This can cause the repository to lose existing chart; use it with care.")
	rootCmd.PersistentFlags().StringVar(&outputPath, "output-path", ".", "`Location` where rendered files are saved.")
	rootCmd.PersistentFlags().StringSliceVarP(&excludePaths, "exclude-paths", "x", []string{}, "`List` of directories to ignore separated by commas.")

	_ = viper.BindPFlag("helm.s3.force", rootCmd.PersistentFlags().Lookup("force"))
	_ = viper.BindPFlag("exclude.paths", rootCmd.PersistentFlags().Lookup("exclude-paths"))
	_ = viper.BindPFlag("output.path", rootCmd.PersistentFlags().Lookup("output-path"))
	_ = viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))
	_ = viper.BindPFlag("log.output.format", rootCmd.PersistentFlags().Lookup("output-format"))
	_ = viper.BindPFlag("git.lsTree", rootCmd.PersistentFlags().Lookup("git-ls-tree"))

	viper.SetDefault("author", "Carlos Freitas <carlosrfjunior@gmail.com>")
	viper.SetDefault("license", "MIT")

	viper.AutomaticEnv()

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(docsCmd)
	rootCmd.AddCommand(manCmd)

}
