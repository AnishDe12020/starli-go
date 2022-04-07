/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AnishDe12020/starli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "starli",
	Short: "A CLI to generate boilerplace code for your project",
	Long:  `Starli lets you generate boilerplace code for your project via interactive prompts. You are able to select different frameworks, add libraries and other tools like linters.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		templates, err := utils.GetTemplates()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var questions = []*survey.Question{
			{
				Name:     "name",
				Prompt:   &survey.Input{Message: "What is the name of the project?"},
				Validate: survey.Required,
			},
			{
				Name: "path",
				Prompt: &survey.Input{
					Message: "Where do you want to create the project?",
					Default: ".",
				},
			},
			{
				Name: "template",
				Prompt: &survey.Select{
					Message: "What template do you want to use?",
					Options: templates,
				},
			},
		}

		answers := struct {
			Name string
			Path string
		}{}

		survey.Ask(questions, &answers)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the irootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.starli.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".starli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".starli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}