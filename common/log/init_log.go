/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package log

// Config config with rolling backend
// MaxSize is the maximum size in megabytes
// MaxBackups is the maximum number of old log files to retain
// MaxAge is the maximum number of days to retain old log files
type Config struct {
	LogFile    string
	LogLevel   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

// InitLogConfig Set the logging level with common ServerGeneral configurations
func InitLogConfig(conf *Config) {
	// Init backend before set log level
	InitRollingBackend(conf.LogFile, conf.MaxSize, conf.MaxBackups, conf.MaxAge)
	InitFromSpec(conf.LogLevel)
}
