/*
Copyright © 2022 Anish De contact@anishde.dev

*/
package cmd

import (
	"github.com/AnishDe12020/starli/utils"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Manually update Starli Specs",
	Long:  `This command updates the Starli specs present in the Starli cache directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.UpdateSpecs(true)
	},
}

func init() {
	specsCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
