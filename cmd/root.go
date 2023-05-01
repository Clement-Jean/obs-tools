package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/andreykaipov/goobs"

	"github.com/spf13/cobra"
)

const version = "0.0.1"

var obs *goobs.Client
var addr string
var password string
var rootCmd = &cobra.Command{
	Use:     "obs-tools",
	Version: version,
	Short:   "Automate obs tasks",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if obs == nil {
			client, err := goobs.New(addr, goobs.WithPassword(password))

			if err != nil {
				log.Fatal(err)
			}

			obs = client
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&addr, "addr", "a",
		"localhost:4455",
		"Websocket server address",
	)

	rootCmd.PersistentFlags().StringVarP(
		&password, "pass", "p",
		"",
		"Websocket password",
	)
}

func Execute() {
	defer func() {
		if obs != nil {
			obs.Disconnect()
			obs = nil
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
