package main

import (
	"os"

	"github.com/guigolab/rnaget-client/rnaget-client/cmd"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
