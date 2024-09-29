package main

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func initConfig() {

	log.New()

	logLevel, err := log.ParseLevel(viper.GetString("log.level"))
	if err != nil {
		fmt.Println("Invalid log level specified:", err)
		os.Exit(1)
	}

	log.SetLevel(logLevel)

	if logQuiet {
		log.SetOutput(io.Discard)
	}

	// Log as JSON instead of the default ASCII formatter.
	if setFormatter == "json" {
		log.SetFormatter(&log.JSONFormatter{
			DisableTimestamp: false,
		})

	} else {
		log.SetFormatter(&log.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
	}

	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalln(err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath("./.config/")
		viper.AddConfigPath(".")
		viper.SetConfigName(".protomagic")
	}

	_ = viper.ReadInConfig()
}
