package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/azarc-io/action-module-scrape/temp/module_v1"
	"github.com/azarc-io/action-module-scrape/util"
	"github.com/sethvargo/go-githubactions"
	"net/http"
)

const (
	readme = "readme.md"
)

func main() {
	gitAction := githubactions.New()
	action := &module_v1.Action{}
	path := gitAction.Getenv("GITHUB_WORKSPACE")

	str := gitAction.Getenv("INPUT_TOKEN")

	gitAction.Infof("token %s", str)

	action.Module = loadModule(gitAction, path)
	action.Sparks = loadSparks(gitAction, path)
	action.Connectors = loadConnectors(gitAction, path)

	submitAction(gitAction, action)
}

//********************************************************************************************
// MODULE
//********************************************************************************************

func loadModule(gitAction *githubactions.Action, path string) *module_v1.Module {
	module := &module_v1.Module{}
	util.ParseYaml(gitAction, fmt.Sprintf("%s/module.yaml", path), &module)
	if module.Version = gitAction.Getenv("INPUT_VERSION"); module.Version == "" {
		if gitAction.Getenv("GITHUB_REF_TYPE") != "tag" {
			gitAction.Fatalf("you must either set the version or run on tag push")
		}
		if _, err := fmt.Sscanf(gitAction.Getenv("GITHUB_REF"), "refs/tags/%s", &module.Version); err != nil {
			gitAction.Fatalf("getting tag push version: %s", err.Error())
		}
	}
	module.Repo = gitAction.Getenv("GITHUB_REPOSITORY")
	module.Icon = util.LoadImage(gitAction, fmt.Sprintf("%s/icon", path), false)
	module.Readme = util.LoadFile(gitAction, readme, false)
	module.Licence = util.LoadFile(gitAction, fmt.Sprintf("%s/licence.txt", path), false)
	return module
}

//********************************************************************************************
// SPARK
//********************************************************************************************

func loadSpark(gitAction *githubactions.Action, sparks []*module_v1.Spark) func(string, string) {
	return func(path, name string) {
		spark := module_v1.Spark{}
		util.ParseYaml(gitAction, fmt.Sprintf("%s/%s", path, "spark.yaml"), &spark)
		spark.Name = name
		spark.Readme = util.LoadFile(gitAction, fmt.Sprintf("%s/%s", path, readme), false)
		spark.InputSchema = util.LoadSchema(gitAction, fmt.Sprintf("%s/%s", path, "input_schema.json"))
		spark.OutputSchema = util.LoadSchema(gitAction, fmt.Sprintf("%s/%s", path, "output_schema.json"))
		spark.Icon = util.LoadImage(gitAction, fmt.Sprintf("%s/icon", path), false)
		sparks = append(sparks, &spark)
	}
}

func loadSparks(gitAction *githubactions.Action, path string) []*module_v1.Spark {
	var sparks []*module_v1.Spark
	util.LoadDirs(gitAction, fmt.Sprintf("%s/sparks", path), loadSpark(gitAction, sparks))
	return sparks
}

//********************************************************************************************
// CONNECTOR
//********************************************************************************************

func loadConnector(gitAction *githubactions.Action, connectors []*module_v1.Connector) func(string, string) {
	return func(path, name string) {
		connector := module_v1.Connector{}
		util.ParseYaml(gitAction, fmt.Sprintf("%s/%s", path, "connector.yaml"), &connector)
		connector.Name = name
		connector.Readme = util.LoadFile(gitAction, fmt.Sprintf("%s/%s", path, readme), false)
		connector.Schema = util.LoadSchema(gitAction, fmt.Sprintf("%s/%s", path, "schema.json"))
		connector.Icon = util.LoadImage(gitAction, fmt.Sprintf("%s/icon", path), false)
		connectors = append(connectors, &connector)
	}
}

func loadConnectors(gitAction *githubactions.Action, path string) []*module_v1.Connector {
	var connectors []*module_v1.Connector
	util.LoadDirs(gitAction, fmt.Sprintf("%s/connectors", path), loadConnector(gitAction, connectors))
	return connectors
}

//********************************************************************************************
// SUBMISSION
//********************************************************************************************

func submitAction(gitAction *githubactions.Action, action *module_v1.Action) {
	if err := action.Validate(); err != nil {
		gitAction.Fatalf("action validation failed: %s", err.Error())
	}

	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(action); err != nil {
		gitAction.Fatalf("could not encode module: %s", err.Error())
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/module", gitAction.Getenv("INPUT_SUBMISSION_HOST")),
		buf,
	)
	req.Header.Set("Authorization", gitAction.Getenv("INPUT_TOKEN"))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		gitAction.Fatalf("could not add module: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		gitAction.Fatalf("add module response [code]: %d, [status]: '%s'", resp.StatusCode, resp.Status)
	}

	gitAction.Infof("scraped and submitted for module [package]: %s, [version]: %s, [sparks]: %d",
		action.Module.Package, action.Module.Version, len(action.Sparks))
}
