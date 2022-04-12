package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
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
