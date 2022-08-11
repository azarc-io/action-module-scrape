package main

import (
	"fmt"
	ga "github.com/sethvargo/go-githubactions"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
)

const (
	readme = "readme.md"
)

func main() {
	action := ga.New()
	workspace := action.Getenv("GITHUB_WORKSPACE")

	module := Module{Config: ModuleConfig{}}

	parseYaml(action, fmt.Sprintf("%s/module.yaml", workspace), &module.Config)
	module.Readme = readFile(action, readme)

	sparksRoot := fmt.Sprintf("%s/sparks", workspace)
	files, err := ioutil.ReadDir(sparksRoot)
	if err != nil {
		action.Fatalf("listing sparks", err.Error())
	}

	for _, dir := range files {
		if !dir.IsDir() {
			continue
		}

		spark := Spark{}
		module.Sparks = append(module.Sparks, &spark)

		path.Join()
		sparkRoot := fmt.Sprintf("%s/%s", sparksRoot, dir.Name())
		parseYaml(action, fmt.Sprintf("%s/%s", sparkRoot, "spark.yaml"), &spark.Config)
		loadSchema(action, fmt.Sprintf("%s/%s", sparkRoot, "input_schema.json"), &spark.InputSchema)
		spark.Readme = readFile(action, fmt.Sprintf("%s/%s", sparkRoot, readme))
	}
	action.Infof("scraped %d sparks", len(module.Sparks))
}

func readFile(action *ga.Action, file string) []byte {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		action.Fatalf("reading file %s", err.Error())
	}
	return data
}

func parseYaml(action *ga.Action, file string, v interface{}) {
	data := readFile(action, file)
	if err := yaml.Unmarshal(data, v); err != nil {
		action.Fatalf("unmarshal file %s", err.Error())
	}
}

func loadSchema(action *ga.Action, file string, v *[]byte) {
	*v = readFile(action, file)
	if _, err := gojsonschema.NewSchema(gojsonschema.NewBytesLoader(*v)); err != nil {
		action.Fatalf("loading schema %s", err.Error())
	}
}
