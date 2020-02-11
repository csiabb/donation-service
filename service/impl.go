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

package service

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"donation-service/common/metadata"
	"donation-service/config"

	logging "github.com/op/go-logging"
)

// gracefulTimeout controls how long we wait before forcefully terminating
const gracefulTimeout = 5 * time.Second

var (
	// DonationServer donation service instance
	DonationServer *ServerImpl
	logger         = logging.MustGetLogger("service")
)

// ServerImpl is the unite did server
type ServerImpl struct {
	config     *config.SrvcCfg
	version    *metadata.Version
	httpSrv    *http.Server
	ShutdownCh <-chan struct{}
	myName     string
	serviceID  string
}

// Start ...
func (s *ServerImpl) Start() (err error) {
	logger.Infof("Starting %s server ...", s.version.ProgramName)
	// start to serve http connections
	if err := s.httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("failed to start http server: %s\n", err)
	}
	logger.Debugf("%s server started successfully.", s.version.ProgramName)

	return
}

// Shutdown ...
func (s *ServerImpl) Shutdown() {
	// shutting down http server
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger.Info("service shutted down successfully.")
}

func (s *ServerImpl) init() (err error) {
	logger.Debugf("Initializing %s server ...", s.version.ProgramName)
	//Init the rest http service
	if err = s.httpSrvInit(); err != nil {
		logger.Errorf("Failed to initialize %s restful API: %s", s.version.ProgramName, err)
		return err
	}
	return nil
}

func (s *ServerImpl) httpSrvInit() (err error) {
	// init api router
	if err := s.httpHandlerInit(); err != nil {
		return fmt.Errorf("failed to init http handler: %s", err.Error())
	}

	// TODO: handle errors here
	logger.Infof("starting server on %s", "0.0.0.0:8888")
	s.httpSrv = &http.Server{
		Addr: "0.0.0.0:8888",
		// Handler: r,
	}

	return nil
}

func (s *ServerImpl) httpHandlerInit() (err error) {

	return nil
}

// GetID ...
func (s *ServerImpl) GetID() string {
	return s.serviceID
}

// GetName ...
func (s *ServerImpl) GetName() string {
	return s.version.ProgramName
}

// GetTags ...
func (s *ServerImpl) GetTags() []string {
	return nil
}

// handleSignals blocks until we get an exit-causing signal
func (s *ServerImpl) handleSignals() int {
	signalCh := make(chan os.Signal, 4)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGPIPE)

	// Wait for a signal
WAIT:
	var sig os.Signal
	select {
	case s := <-signalCh:
		sig = s
	case <-s.ShutdownCh:
		sig = os.Interrupt
	}
	logger.Infof("Caught signal: %v", sig)

	// Skip any SIGPIPE signal (See issue #1798)
	if sig == syscall.SIGPIPE {
		goto WAIT
	}

	// Check if we should do a graceful leave
	graceful := false
	if sig == os.Interrupt && config.LeaveOnInt {
		graceful = true
	} else if sig == syscall.SIGTERM && config.LeaveOnTerm {
		graceful = true
	}

	// Bail fast if not doing a graceful leave
	if !graceful {
		return 1
	}

	// Attempt a graceful leave
	gracefulCh := make(chan struct{})
	logger.Infof("Gracefully shutting down %s server ...", s.version.ProgramName)
	go func() {
		s.Shutdown()
		close(gracefulCh)
	}()

	// Wait for leave or another signal
	select {
	case <-signalCh:
		return 1
	case <-time.After(gracefulTimeout):
		return 1
	case <-gracefulCh:
		return 0
	}
}
