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
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: createTemplate,
}

func createTemplate(cmd *cobra.Command, args []string) {

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
		Name     string
		Path     string
		Template string
	}{}

	survey.Ask(questions, &answers)
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
