/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/csiabb/donation-service/common/log"
	"github.com/csiabb/donation-service/common/metadata"
	"github.com/csiabb/donation-service/config"
	"github.com/csiabb/donation-service/router"

	"github.com/gin-gonic/gin"
)

// gracefulTimeout controls how long we wait before forcefully terminating
const gracefulTimeout = 5 * time.Second

var (
	// DonationServer donation service instance
	DonationServer *ServerImpl
	logger         = log.MustGetLogger("service")
)

// ServerImpl is the unite did server
type ServerImpl struct {
	config     *config.SrvcCfg
	version    *metadata.Version
	httpSrv    *gin.Engine
	ShutdownCh <-chan struct{}
	myName     string
	serviceID  string
}

// Start ...
func (s *ServerImpl) Start() (err error) {
	logger.Infof("Starting %s server ...", s.version.ProgramName)

	// start to serve http connections
	address := fmt.Sprintf("%s:%d", s.config.ServerGeneral.Host, s.config.ServerGeneral.Port)
	logger.Infof("starting server on %s", address)
	s.httpSrv.Run(address)

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
	if err := s.httpControllerInit(); err != nil {
		return fmt.Errorf("failed to init http handler: %s", err.Error())
	}

	s.httpSrv = router.SetupRouter()

	return nil
}

func (s *ServerImpl) httpControllerInit() (err error) {

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
