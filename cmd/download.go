/*
Copyright © 2022 Anish De contact@anishde.dev

*/
package cmd

import (
	"github.com/AnishDe12020/starli/utils"
	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Manually download Starli Specs",
	Long:  `This command downloads the Starli specs to the Starli cache directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.DownloadSpecsDir()
	},
}

func init() {
	specsCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
