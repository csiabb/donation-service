/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Globle var
var (
	myViper *viper.Viper
)

// defaultCfgPath default conigure file path
const defaultCfgPath = "/etc/donation-service/"

// get component's name from configuration file path
// like for /etc/arxanchain/baymax.yaml return baymax
// func getSrvNameFromCfgPath(cfgFile string) string {
// 	fName := filepath.Base(cfgFile)
// 	extName := filepath.Ext(cfgFile)
// 	componentName := fName[:len(fName)-len(extName)]
// 	return componentName
// }

// SetupConfig setup the config during test execution
// func SetupConfig(cfgFile string) {
// 	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
// 		panic(fmt.Errorf("Fatal error when reading %s config file: %s\n", cfgFile, err))
// 	}

// 	MyViper = viper.New()
// 	cName := getSrvNameFromCfgPath(cfgFile)

// 	MyViper.SetEnvPrefix(strings.ToUpper(cName))
// 	MyViper.AutomaticEnv()
// 	replacer := strings.NewReplacer(".", "_")
// 	MyViper.SetEnvKeyReplacer(replacer)

// 	// Now set the configuration file.
// 	MyViper.SetConfigName(cName)                 // Name of config file (without extension)
// 	MyViper.AddConfigPath(filepath.Dir(cfgFile)) // Path to look for the config file in

// 	err := MyViper.ReadInConfig() // Find and read the config file
// 	if err != nil {               // Handle errors reading the config file
// 		panic(fmt.Errorf("Fatal error when reading %s config file: %s\n", cfgFile, err))
// 	}
// 	// err = cachConfigure()
// 	// if err != nil { // Handle errors reading the config file
// 	// 	panic(fmt.Errorf("Fatal error when reading %s config file: %s\n", cmdRoot, err))
// 	// }
// 	runtime.GOMAXPROCS(MyViper.GetInt("common.gomaxprocs"))
// 	flag.Parse()

// }

// initConfig Init configuration and return the viper instance
func initConfig(configName string) *viper.Viper {
	config := viper.New()
	initViper(config, configName)

	//get the environment prefix name from configName
	prefix := strings.ToUpper(configName)

	// for environment variables
	config.SetEnvPrefix(prefix)
	config.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	config.SetEnvKeyReplacer(replacer)

	err := config.ReadInConfig()
	if err != nil {
		logger.Errorf("Error reading configuration: %s, make sure your config file exists or format is correct.", err.Error())
	} else {
		logger.Debugf("Using config file: %s", config.ConfigFileUsed())
	}
	return config
}

func dirExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func addConfigPath(v *viper.Viper, p string) {
	if v != nil {
		v.AddConfigPath(p)
	} else {
		viper.AddConfigPath(p)
	}
}

// initViper ...
// Performs basic initialization of our viper-based configuration layer.
// Primary thrust is to establish the paths that should be consulted to find
// the configuration we need.  If v == nil, we will initialize the global
// Viper instance
//----------------------------------------------------------------------------------
func initViper(v *viper.Viper, configName string) error {
	prefix := strings.ToUpper(configName)
	logger.Debugf("Init component configuration with prefix: %s", prefix)
	var altPath = os.Getenv(fmt.Sprintf("%s_CFG_PATH", prefix))
	if altPath != "" {
		// If the user has overridden the path with an envvar, its the only path
		// we will consider
		logger.Debugf("adding altPath [%s] for viper configuration", altPath)
		addConfigPath(v, altPath)
	} else {
		// If we get here, we should use the default paths in priority order:
		//
		// *) CWD
		// *) The $GOPATH based development tree
		// *) /etc/hyperledger/fabric
		//
		// CWD
		addConfigPath(v, "./")

		// And finally, the default path
		if dirExists(defaultCfgPath) {
			logger.Debugf("adding [%s] to config path", defaultCfgPath)
			addConfigPath(v, defaultCfgPath)
		}
	}

	// Now set the configuration file.
	if v != nil {
		v.SetConfigName(configName)
	} else {
		viper.SetConfigName(configName)
	}

	return nil
}
