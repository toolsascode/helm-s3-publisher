// Helm S3 Publisher is a small CLI that brings the powerful features of the golang template into a simplified form.
package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// Start application
func init() {

	initFlag()

	log.SetReportCaller(false)
	log.SetOutput(os.Stdout)

}

func Execute() {

	err := rootCmd.Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func main() {
	Execute()
}
