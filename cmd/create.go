/*
Copyright © 2022 Anish De contact@anishde.dev

*/
package cmd

import (
	"fmt"

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
	Args: cobra.MaximumNArgs(2),
	RunE: handleCreate,
}

type CreateOptions struct {
	Name     string "json:name"
	Path     string "json:path"
	Template string "json:template"
}

var opts = CreateOptions{}

func handleCreate(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		opts.Name = args[0]
	}

	if len(args) > 1 {
		opts.Path = args[1]
	}

	templates, err := utils.GetTemplates()

	if err != nil {
		return utils.Error("Failed to get templates")
	}

	if opts.Template != "" {
		if utils.Contains(templates, opts.Template) {
			return utils.Error("Template not found")
		}
	}

	err = createProject(opts, templates)

	return err
}

func createProject(opts CreateOptions, templates []string) error {
	if opts.Name == "" {
		prompt := &survey.Input{
			Message: "Name of the project",
		}

		err := survey.AskOne(prompt, &opts.Name, nil)

		if err != nil {
			return utils.Error("Failed to get name")
		}
	}

	if opts.Path == "" {
		prompt := &survey.Input{
			Message: "Path to create the project",
		}

		err := survey.AskOne(prompt, &opts.Path, nil)

		if err != nil {
			return utils.Error("Failed to get path")
		}
	}

	if opts.Template == "" {
		prompt := &survey.Select{
			Message: "Template to use",
			Options: templates,
		}

		err := survey.AskOne(prompt, &opts.Template, nil)

		if err != nil {
			return utils.Error("Failed to get template")
		}
	}

	fmt.Println(opts)

	return nil
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&opts.Name, "name", "n", "", "Name of the project")
	createCmd.Flags().StringVarP(&opts.Path, "path", "p", "", "Path to create the project")
	createCmd.Flags().StringVarP(&opts.Template, "template", "t", "", "Template to use")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
