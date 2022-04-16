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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Example: `
starli specs update
`,
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
