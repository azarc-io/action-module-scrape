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
	gitAction, action, path := loadAction()
	action.Module = loadModule(gitAction, path)
	action.Sparks = loadSparks(gitAction, path)
	action.Connectors = loadConnectors(gitAction, path)
	submitAction(gitAction, action)
}

//********************************************************************************************
// ACTION
//********************************************************************************************

func loadAction() (*githubactions.Action, *module_v1.Action, string) {
	gitAction := githubactions.New()
	if gitAction.Getenv("GITHUB_REF_TYPE") != "tag" {
		gitAction.Fatalf("action can only be used on push tag")
	}

	action := &module_v1.Action{Repo: gitAction.Getenv("GITHUB_REPOSITORY")}
	if _, err := fmt.Sscanf(gitAction.Getenv("GITHUB_REF"), "refs/tags/%s", &action.Version); err != nil {
		gitAction.Fatalf("getting version: %s", err.Error())
	}
	return gitAction, action, gitAction.Getenv("GITHUB_WORKSPACE")
}

func submitAction(gitAction *githubactions.Action, action *module_v1.Action) {
	if err := action.Validate(); err != nil {
		gitAction.Fatalf("action validation failed: %s", err.Error())
	}

	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(action); err != nil {
		gitAction.Fatalf("could not encode module: %s", err.Error())
	}
	resp, err := http.Post("https://auth-events.cloud.azarc.dev/api/v1/module", "application/json", buf)
	if err != nil {
		gitAction.Fatalf("could not add module: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		gitAction.Fatalf("received %d from add module request", resp.StatusCode)
	}
	gitAction.Infof("scraped and submitted for module [repo]: %s, [version]: %s, [sparks]: %d",
		action.Repo, action.Version, len(action.Sparks))
}

//********************************************************************************************
// MODULE
//********************************************************************************************

func loadModule(gitAction *githubactions.Action, path string) *module_v1.Module {
	module := &module_v1.Module{}
	util.ParseYaml(gitAction, fmt.Sprintf("%s/module.yaml", path), &module)
	module.Icon = util.LoadImage(gitAction, fmt.Sprintf("%s/icon", path))
	module.Readme = util.ReadFile(gitAction, readme)
	module.Licence = util.ReadFile(gitAction, fmt.Sprintf("%s/licence.txt", path))
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
		spark.Readme = util.ReadFile(gitAction, fmt.Sprintf("%s/%s", path, readme))
		spark.InputSchema = util.LoadSchema(gitAction, fmt.Sprintf("%s/%s", path, "input_schema.json"))
		spark.OutputSchema = util.LoadSchema(gitAction, fmt.Sprintf("%s/%s", path, "output_schema.json"))
		spark.Icon = util.LoadImage(gitAction, fmt.Sprintf("%s/icon", path))
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
		connector.Readme = util.ReadFile(gitAction, fmt.Sprintf("%s/%s", path, readme))
		connector.Schema = util.LoadSchema(gitAction, fmt.Sprintf("%s/%s", path, "schema.json"))
		connector.Icon = util.LoadImage(gitAction, fmt.Sprintf("%s/icon", path))
		connectors = append(connectors, &connector)
	}
}

func loadConnectors(gitAction *githubactions.Action, path string) []*module_v1.Connector {
	var connectors []*module_v1.Connector
	util.LoadDirs(gitAction, fmt.Sprintf("%s/connectors", path), loadConnector(gitAction, connectors))
	return connectors
}
