/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package middleware

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/csiabb/donation-service/common/log"
	"github.com/gin-gonic/gin"
)

var (
	logger = log.MustGetLogger("middleware")
)

// responseLogger save response content
type responseLogger struct {
	gin.ResponseWriter
	content *bytes.Buffer
}

// Write ...
func (rbl *responseLogger) Write(b []byte) (int, error) {
	rbl.content.Write(b)
	return rbl.ResponseWriter.Write(b)
}

// RequestResponseLogger we use this middleware to log api request/response and check api latency
func RequestResponseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// get request body and log it
		buf, _ := ioutil.ReadAll(c.Request.Body)
		logger.Debugf("got request '%s' with params '%s' and header '%v'", c.Request.URL.Path, string(buf), c.Request.Header)

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))

		// prepare for the response body logger, our response content will be saved to responseLogger.content
		rlw := &responseLogger{content: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = rlw

		c.Next()

		// check the response latency for this request
		latency := time.Since(t)
		outString := rlw.content.String()
		outLen := len(outString)

		if outLen < 1000 {
			logger.Debugf("handle request '%s' with latency of %s, response body: %s", c.Request.URL.Path, latency, outString)
		} else {
			logger.Debugf("handle request '%s' with latency of %s, response body length %v", c.Request.URL.Path, latency, outLen)
		}

		return
	}
}
