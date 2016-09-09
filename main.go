package main

import (
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/shad7/gochlog/commands"
)

func init() {
	loglevel, err := log.ParseLevel(strings.ToLower(os.Getenv("GOCHLOG_LOG_LEVEL")))
	if err != nil {
		loglevel = log.WarnLevel
	}
	log.SetLevel(loglevel)
}

// Main executes the CLI
func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}

}
