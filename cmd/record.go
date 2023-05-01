package cmd

import (
	"log"

	"github.com/Clement-Jean/obs-tools/pkg"

	"github.com/spf13/cobra"
)

var fileName string
var timestamp string
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Launch recording",
	Run: func(cmd *cobra.Command, args []string) {
		if err := pkg.RecordWithName(obs, fileName, timestamp); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	recordCmd.Flags().StringVarP(
		&fileName, "name", "n",
		"",
		"Output's name",
	)

	recordCmd.Flags().StringVarP(
		&timestamp, "timestamp", "t",
		"",
		"Timestamp format to be added after name",
	)

	rootCmd.AddCommand(recordCmd)
}
