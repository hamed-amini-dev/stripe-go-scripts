package config

/*
initialize viper module for loading file and load default config
loading default most important config if user don't set on config file
*/

import (
	"github.com/hamed-amini-dev/stripe-go-scripts/pkg/constants"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// InitConfig initializes the viper and loads the configuration.
func InitConfig(path string) error {
	var err error

	// ─── CONFIG THE PATH AND FILE NAME ──────────────────────────────────────────────
	viper.SetConfigName(constants.ConfigFileName) // name of config file (without extension)
	viper.AddConfigPath(constants.PathWorkingDirecotry)
	viper.AddConfigPath(path) // call multiple times to add many search paths
	viper.AutomaticEnv()

	// ─── INIT THE DEFAULT VALUES ────────────────────────────────────────────────────
	err = initDefault()
	if err != nil {
		return err
	}

	// ─── READ THE CONFIG ────────────────────────────────────────────────────────────
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	// ─── BINDING FLAGS ──────────────────────────────────────────────────────────────
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	return nil
}

// initDefault is where you can write the default value
// Priority
// default < config.yaml < osEnvironment < flags
func initDefault() error {

	//
	// ───────────────────────────────────────────── ADD YOUR DEFAULT VALUES HERE ─────
	//

	// ─── SERVER CONFIGURATION ───────────────────────────────────────────────────────

	pflag.Int(constants.Port, constants.PortDefault, "http port of server")
	viper.SetDefault(constants.Port, constants.PortDefault)

	return nil
}
