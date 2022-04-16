/*
Copyright © 2022 Anish De contact@anishde.dev

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// specsCmd represents the specs command
var specsCmd = &cobra.Command{
	Use:   "specs",
	Short: "Command for operations on Starli specs.",
	Long:  `This is a parent command with subcommands to perform operations on Starli specs.`,
}

func init() {
	rootCmd.AddCommand(specsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// specsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// specsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
