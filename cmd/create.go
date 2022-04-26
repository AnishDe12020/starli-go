/*
Copyright © 2022 Anish De contact@anishde.dev

*/
package cmd

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

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
		currentDirName, err := utils.GetCurrentDirName()

		if err != nil {
			return utils.Error("Failed to get current directory name")
		}

		prompt := &survey.Input{
			Message: "Name of the project",
			Default: currentDirName,
		}

		err = survey.AskOne(prompt, &opts.Name, survey.WithValidator(survey.Required))

		if err != nil {
			return utils.Error("Failed to get name")
		}
	}

	if opts.Path == "" {
		currentDirName, err := utils.GetCurrentDirName()

		if err != nil {
			return utils.Error("Failed to get current directory name")
		}

		var currentDirPath string

		if currentDirName == opts.Name {
			currentDirPath = "./" + opts.Name
		} else {
			currentDirPath, err = utils.GetCurrentDirPath()
		}

		if err != nil {
			return utils.Error("Failed to get current directory name")
		}

		prompt := &survey.Input{
			Message: "Path to create the project",
			Default: currentDirPath,
		}

		err = survey.AskOne(prompt, &opts.Path, survey.WithValidator(survey.Required))

		if err != nil {
			return utils.Error("Failed to get path")
		}
	}

	if opts.Template == "" {
		prompt := &survey.Select{
			Message: "Template to use",
			Options: templates,
		}

		err := survey.AskOne(prompt, &opts.Template, survey.WithValidator(survey.Required))

		if err != nil {
			return utils.Error("Failed to get template")
		}
	}

	starliSpecsDir := utils.GetStarliSpecsCacheDir()

	filepath.WalkDir(starliSpecsDir+"/templates/"+strings.ToLower(opts.Template), func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, ".tmpl") {
			fmt.Print(path)
			fmt.Print(" Is template file\n")
		} else if strings.HasSuffix(path, "starli.json") {
			fmt.Print(path)
			fmt.Print(" Is starli config file")
		} else {
			fmt.Print(path)
			fmt.Print(" Is file\n")
		}

		return nil
	})

	templateMeta, err := utils.GetTemplate(opts.Template)

	for _, question := range templateMeta.Questions {
		var answer string

		prompt := &survey.Input{
			Message: question.Message,
			Default: question.Default,
		}

		err = survey.AskOne(prompt, &answer, survey.WithValidator(survey.Required))

		if err != nil {
			return utils.Error("Failed to get answer")
		}

	}

	if err != nil {
		return utils.Error("Failed to get template")
	}

	// templatesParsed, err := utils.GetTemplate(opts.Template)

	// if err != nil {
	// 	return utils.Error("Failed to get template")
	// }

	// type ExecuteOptions struct {
	// 	Test string
	// }

	// execOpts := ExecuteOptions{
	// 	Test: "value",
	// }

	// execOpts := map[string]string{
	// 	"Test": "value",
	// }

	// for _, template := range templatesParsed.Templates() {
	// 	// fmt.Println("Name: ", template.)
	// 	fmt.Println("Value: ")
	// 	template.Execute(os.Stdout, execOpts)
	// }

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
