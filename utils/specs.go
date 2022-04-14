package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type SpecTemplate struct {
	Name        string                    "json:name"
	StaticFiles []SpecTemplateStaticFiles "json:staticFiles"
}

type SpecTemplateStaticFiles struct {
	Name    string "json:name"
	Path    string "json:path"
	Content string "json:content"
}

func GetTemplates() ([]string, error) {
	templates := []string{}
	matches, err := filepath.Glob("specs/templates/**/starli.json")

	if err != nil {
		return nil, err
	}

	for _, path := range matches {
		templateData, err := ioutil.ReadFile(path)

		if err != nil {
			return nil, err
		}

		var template SpecTemplate

		err = json.Unmarshal(templateData, &template)

		if err != nil {
			return nil, err
		}

		templates = append(templates, template.Name)
	}

	return templates, nil
}

func GetTemplate(name string) (SpecTemplate, error) {
	var template SpecTemplate

	templateData, err := ioutil.ReadFile("specs/templates/" + strings.ToLower(name) + "/starli.json")

	if err != nil {
		fmt.Println(err)
		return template, err
	}

	err = json.Unmarshal(templateData, &template)

	if err != nil {
		fmt.Println(err)
		return template, err
	}

	return template, nil
}

func CheckIfSpecsDirExists() (bool, error) {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		return false, err
	}
	starliSpecsDir := filepath.Join(userCacheDir, "starli", "specs")

	if _, err := os.Stat(starliSpecsDir); errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func DownloadSpecsDir() error {
	fmt.Println("Downloading specs...")
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return err
	}

	starliDirPath := cacheDir + "/starli"

	if _, err := os.Stat(starliDirPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(starliDirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithoutAuthentication())
	if err != nil {
		return err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	rc, err := client.Bucket("starli-cli.appspot.com").Object("specs.tar").NewReader(ctx)
	if err != nil {
		return err
	}
	defer rc.Close()

	err = Untar(starliDirPath, rc)
	if err != nil {
		return err
	}

	fmt.Println("Specs downloaded.")

	return nil

}
