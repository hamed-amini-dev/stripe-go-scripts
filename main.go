/*
Banking Transferring money between account Go programs.

loading config file with viper package and also using cli package for getting argument from console
initialize all module after loading configuration and then run service

Usage:

main [arg]

The args are:

s
Run service and loading mock database
Ready to transfer balance between account
*/
package main

import (
	"github.com/hamed-amini-dev/stripe-go-scripts/cmd"
	"github.com/hamed-amini-dev/stripe-go-scripts/pkg/config"
)

// Load Configuration File
// Create New Commander package for getting args console
// Run Commander package with args inputted
func main() {

	//load configuration with viper and init it
	err := config.InitConfig(".")
	if err != nil {
		panic(err)
	}

	//commander initialize for getting args from console
	commander, err := cmd.New()
	if err != nil {
		panic(err)
	}

	//run commander
	if err := commander.Run(); err != nil {
		panic(err)
	}

}
