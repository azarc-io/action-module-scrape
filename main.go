package main

import (
	"fmt"
	ga "github.com/sethvargo/go-githubactions"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	sparksRoot = "sparks"
	readme     = "readme.md"
)

func main() {
	action := ga.New()

	module := Module{Config: ModuleConfig{}}
	parseYaml(action, "module.yaml", &module.Config)
	module.Readme = readFile(action, readme)

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

		sparkRoot := fmt.Sprintf("%s/%s", sparksRoot, dir.Name())
		parseYaml(action, fmt.Sprintf("%s/%s", sparkRoot, "spark.yaml"), &spark.Config)
		loadSchema(action, fmt.Sprintf("%s/%s", sparkRoot, "input_schema.json"), &spark.InputSchema)
		spark.Readme = readFile(action, fmt.Sprintf("%s/%s", sparkRoot, readme))
	}
	fmt.Println(module)
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
