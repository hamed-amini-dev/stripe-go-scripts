package cmd

import (
	"github.com/hamed-amini-dev/stripe-go-scripts/app"
	"github.com/urfave/cli/v2"
)

// Run server when user input args serve from console
// Initialize app and module imported

func (c *Commander) RunServer(ctx *cli.Context) error {
	//init app and creates objects
	appServer, err := app.NewApp()
	if err != nil {
		return err
	}

	//after module are ready , listen to port for handling request
	err = appServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
