package utils

import (
	"github.com/fenixvlabs/meshkit/pkg/meshlog"
	log "github.com/sirupsen/logrus"
	"os"
)

var verbose bool

func SetupLogrusFormatter() {
	log.SetFormatter(&log.TextFormatter{})
}

func SetupMeshkitLogger(debugLevel bool) {
	_, err := meshlog.New(meshlog.AdapterMeshkit, meshlog.HandlerOptions{
		Format: meshlog.TerminalLogFormat,
		Level:  debugLevel,
	})

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	return
}

func setupLogger() {
	SetupMeshkitLogger(verbose)
}
