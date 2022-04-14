/*
Copyright © 2022 Anish De contact@anishde.dev

*/
package cmd

import (
	// "context"
	// "errors"
	"fmt"
	"os"
	// "time"

	// "cloud.google.com/go/storage"
	"github.com/AnishDe12020/starli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	// "golang.org/x/sync/errgroup"
	// "google.golang.org/api/option"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "starli",
	Short: "A CLI to generate boilerplace code for your project",
	Long:  `Starli lets you generate boilerplace code for your project via interactive prompts. You are able to select different frameworks, add libraries and other tools like linters.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		specsExists, err := utils.CheckIfSpecsExists()
		if err != nil {
			return err
		}

		if !specsExists {
			utils.DownloadSpecsDir()
		} else {
			utils.UpdateSpecs()
		}

		// errs := new(errgroup.Group)

		// errs.Go(func() error {
		// 	cacheDir, err := os.UserCacheDir()
		// 	if err != nil {
		// 		return err
		// 	}

		// 	starliDirPath := cacheDir + "/starli"

		// 	if _, err := os.Stat(starliDirPath); errors.Is(err, os.ErrNotExist) {
		// 		err := os.Mkdir(starliDirPath, os.ModePerm)
		// 		if err != nil {
		// 			return err
		// 		}
		// 	}

		// 	ctx := context.Background()

		// 	client, err := storage.NewClient(ctx, option.WithoutAuthentication())
		// 	if err != nil {
		// 		return err
		// 	}
		// 	defer client.Close()

		// 	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
		// 	defer cancel()

		// 	rc, err := client.Bucket("starli-cli.appspot.com").Object("specs.tar").NewReader(ctx)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	defer rc.Close()

		// 	err = utils.Untar(starliDirPath, rc)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	return nil

		// })

		// if err := errs.Wait(); err != nil {
		// 	return err
		// }

		return nil
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:,
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
