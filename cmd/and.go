package cmd

import (
	"github.com/Clement-Jean/obs-tools/pkg"

	"github.com/spf13/cobra"
)

var andCmd = &cobra.Command{
	Use:   "and -- COMMAND1 <OPTS> - COMMAND2 <OPTS> ...",
	Short: "run multiple commands separated by -",
	Run: func(cobraCmd *cobra.Command, args []string) {
		cmdList := pkg.ParseCommands(args, "-")

		for _, cmdParts := range cmdList {
			rootCmd.SetArgs(cmdParts)
			rootCmd.Execute()
		}
	},
}

func init() {
	andCmd.Flags().SetInterspersed(false)
	rootCmd.AddCommand(andCmd)
}
