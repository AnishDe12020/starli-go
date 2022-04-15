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
	"github.com/AnishDe12020/spintron"
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

func CheckIfSpecsExists() (bool, error) {
	starliSpecsDir := GetStarliSpecsCacheDir()
	specsEtagFile := GetStarliSpecsEtagFile()

	if _, err := os.Stat(specsEtagFile); errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	if _, err := os.Stat(starliSpecsDir); errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func DownloadSpecsDir() error {
	s := spintron.New(spintron.Options{
		Text: "Downloading Starli specs...",
	})
	s.Start()

	starliDirPath := GetStarliCacheDir()

	if _, err := os.Stat(starliDirPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(starliDirPath, os.ModePerm)
		if err != nil {
			s.Fail("Failed to create starli directory")
			return err
		}
	}

	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithoutAuthentication())
	if err != nil {
		s.Fail("Failed to initialize a Google Cloud Storage client")
		return err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	rc, err := client.Bucket("starli-cli.appspot.com").Object("specs.tar").NewReader(ctx)
	if err != nil {
		s.Fail("Failed to download Starli specs")
		return err
	}
	defer rc.Close()

	err = Untar(starliDirPath, rc)
	if err != nil {
		s.Fail("Failed to untar Starli specs")
		return err
	}

	attrs, err := client.Bucket("starli-cli.appspot.com").Object("specs.tar").Attrs(ctx)
	if err != nil {
		s.Fail("Failed to get Starli specs attributes")
		return err
	}

	starliSpecsEtagPath := GetStarliSpecsEtagFile()

	os.WriteFile(starliSpecsEtagPath, []byte(attrs.Etag), 0644)

	s.Succeed("Specs downloaded")

	return nil

}

func UpdateSpecs(isVerbose bool) error {
	s := spintron.New(spintron.Options{
		Text: "Updating Starli specs...",
	})

	if isVerbose {
		s.Start()
	}

	starliDirPath := GetStarliCacheDir()

	if _, err := os.Stat(starliDirPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(starliDirPath, os.ModePerm)
		if err != nil {
			if isVerbose {
				s.Fail("Failed to create starli directory")
			}
			return err
		}
	}

	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithoutAuthentication())
	if err != nil {
		if isVerbose {
			s.Fail("Failed to initialize a Google Cloud Storage client")
		}
		return err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	attrs, err := client.Bucket("starli-cli.appspot.com").Object("specs.tar").Attrs(ctx)
	if err != nil {
		if isVerbose {
			s.Fail("Failed to get Starli specs attributes")
		}
		return err
	}

	starliSpecsEtagPath := GetStarliSpecsEtagFile()

	existingEtag, err := os.ReadFile(starliSpecsEtagPath)
	if err != nil {
		if isVerbose {
			s.Fail("Failed to read Starli specs etag")
		}
		return err
	}

	if string(existingEtag) == attrs.Etag {
		if isVerbose {
			s.Succeed("Specs up to date")
		}
		return nil
	}

	rc, err := client.Bucket("starli-cli.appspot.com").Object("specs.tar").NewReader(ctx)
	if err != nil {
		if isVerbose {
			s.Fail("Failed to download Starli specs")
		}
		return err
	}
	defer rc.Close()

	err = Untar(starliDirPath, rc)

	if err != nil {
		if isVerbose {
			s.Fail("Failed to untar Starli specs")
		}
		return err
	}

	os.WriteFile(starliSpecsEtagPath, []byte(attrs.Etag), 0644)

	if isVerbose {
		s.Succeed("Specs updated")
	}

	return nil
}
