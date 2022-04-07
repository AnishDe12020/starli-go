package utils

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type SpecTemplate struct {
	Name        string                    "json:name"
	StaticFiles []SpecTemplateStaticFiles "json:staticFiles"
}

type SpecTemplateStaticFiles struct {
	Name    string "json:name"
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
