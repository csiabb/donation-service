/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"os"

	"github.com/csiabb/donation-service/common/metadata"
	"github.com/csiabb/donation-service/config"
	"github.com/csiabb/donation-service/service"

	"github.com/op/go-logging"
	"gopkg.in/alecthomas/kingpin.v2"
)

//command line flags
var (
	programName = "donation-service"
	logger      = logging.MustGetLogger("main")

	app = kingpin.New(metadata.ProgramName, "rest server for client business integration")

	startCmd   = app.Command("start", fmt.Sprintf("Start the %s server", metadata.ProgramName)).Default()
	versionCmd = app.Command("version", "Show version information")
)

func cleanup() {
}

func main() {
	defer cleanup()

	kingpin.Version("0.0.1")
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	// "start" command
	case startCmd.FullCommand():
		logger.Infof("Starting %s", metadata.ProgramVersion.FullVersion())
		logger.Infof("Beginning to serve requests")
		conf := config.GetServiceCfg(metadata.ProgramName)
		logger.Debugf("initialize configure %+v", conf)
		server, err := service.NewServer(conf, metadata.ProgramVersion)
		if err != nil {
			logger.Panicf("Failed to create %s server, %+v", metadata.ProgramName, err)
			return
		}
		logger.Infof("Beginning to serve requests")

		server.Start()
	// "version" command
	case versionCmd.FullCommand():
		fmt.Println(metadata.ProgramVersion.FullVersion())
	}
}
