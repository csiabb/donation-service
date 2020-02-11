/*
Copyright Arxan Chain Ltd. 2020 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os"

	"donation-service/common/metadata"
	"donation-service/config"
	"donation-service/service"

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
		conf := &config.SrvcCfg{}
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
