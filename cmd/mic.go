package cmd

import (
	"log"

	"github.com/Clement-Jean/obs-tools/pkg"

	"github.com/spf13/cobra"
)

var audioSource string
var deviceID string
var micCmd = &cobra.Command{
	Use:   "mic",
	Short: "Check microphone usage",
	Run: func(cmd *cobra.Command, args []string) {
		if err := pkg.AssertDeviceID(obs, audioSource, deviceID); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	micCmd.Flags().StringVarP(
		&audioSource, "source", "s",
		"",
		"Audio Source's name (required)",
	)
	micCmd.MarkFlagRequired("source")

	micCmd.Flags().StringVarP(
		&deviceID, "device", "d",
		"",
		"Device id (required)",
	)
	micCmd.MarkFlagRequired("device")

	rootCmd.AddCommand(micCmd)
}
