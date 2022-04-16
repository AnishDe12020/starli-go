/*
Copyright © 2022 Anish De contact@anishde.dev

*/
package cmd

import (
	"github.com/AnishDe12020/starli/utils"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Manually update Starli Specs",
	Long:  `This command deletes the Starli specs present in the Starli cache directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.DeleteSpecs()
	},
}

func init() {
	specsCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
