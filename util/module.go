package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/azarc-io/action-module-scrape/temp/module_v1"
	"google.golang.org/protobuf/types/known/structpb"
	"net/http"
)

const (
	readme = "readme.md"
)

type config struct {
	Config map[string]interface{} `yaml:"config"`
}

//********************************************************************************************
// MODULE
//********************************************************************************************

func LoadModule(log Logger, config *Config) *module_v1.Module {
	module := &module_v1.Module{}
	ParseYaml(log, fmt.Sprintf("%s/module.yaml", config.Path), &module)
	module.Version = config.Version
	module.Repo = config.Repo
	module.Icon = LoadImage(log, fmt.Sprintf("%s/icon", config.Path), false)
	module.Readme = LoadFileString(log, fmt.Sprintf("%s/%s", config.Path, readme), false)
	module.Licence = LoadFileString(log, fmt.Sprintf("%s/licence.txt", config.Path), false)
	return module
}

//********************************************************************************************
// SPARK
//********************************************************************************************

func loadSpark(log Logger, sparks *[]*module_v1.Spark) func(string, string) {
	return func(path, name string) {
		spark := module_v1.Spark{}
		cfg := config{}
		var err error

		ParseYaml(log, fmt.Sprintf("%s/%s", path, "spark.yaml"), &spark, &cfg)
		if spark.Config, err = structpb.NewStruct(cfg.Config); err != nil {
			log.Fatalf("unmarshalling config [spark]: %s, [err]: %s", name, err.Error())
		}
		spark.Name = name
		spark.Readme = LoadFileString(log, fmt.Sprintf("%s/%s", path, readme), false)
		spark.Icon = LoadImage(log, fmt.Sprintf("%s/icon", path), false)
		for _, input := range spark.Inputs {
			input.Schema = LoadSchema(log, fmt.Sprintf("%s/%s", path, input.Schema))
		}
		for _, output := range spark.Outputs {
			if len(output.Schema) != 0 {
				if output.MimeType != "application/json" {
					log.Fatalf("can only set schema on application/json outputs")
				}
				LoadSchema(log, fmt.Sprintf("%s/%s", path, output.Schema))
			}
		}
		*sparks = append(*sparks, &spark)
	}
}

func LoadSparks(log Logger, config *Config) []*module_v1.Spark {
	var sparks []*module_v1.Spark
	LoadDirs(log, fmt.Sprintf("%s/sparks", config.Path), loadSpark(log, &sparks))
	return sparks
}

//********************************************************************************************
// CONNECTOR
//********************************************************************************************

func loadConnector(log Logger, connectors *[]*module_v1.Connector) func(string, string) {
	return func(path, name string) {
		connector := module_v1.Connector{}
		cfg := config{}
		var err error

		ParseYaml(log, fmt.Sprintf("%s/%s", path, "connector.yaml"), &connector, &cfg)
		if connector.Config, err = structpb.NewStruct(cfg.Config); err != nil {
			log.Fatalf("unmarshalling config [connector]: %s, [err]: %s", name, err.Error())
		}
		connector.Name = name
		connector.Readme = LoadFileString(log, fmt.Sprintf("%s/%s", path, readme), false)
		connector.Schema = LoadSchema(log, fmt.Sprintf("%s/%s", path, "schema.json"))
		connector.Icon = LoadImage(log, fmt.Sprintf("%s/icon", path), false)
		*connectors = append(*connectors, &connector)
	}
}

func LoadConnectors(log Logger, config *Config) []*module_v1.Connector {
	var connectors []*module_v1.Connector
	LoadDirs(log, fmt.Sprintf("%s/connectors", config.Path), loadConnector(log, &connectors))
	return connectors
}

//********************************************************************************************
// SUBMISSION
//********************************************************************************************

func SubmitAction(log Logger, config *Config, action *module_v1.Action) {
	if err := action.Validate(); err != nil {
		log.Fatalf("action validation failed: %s", err.Error())
	}

	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(action); err != nil {
		log.Fatalf("could not encode module: %s", err.Error())
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/module", config.SubmissionHost),
		buf,
	)
	req.Header.Set("Authorization", config.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("could not add module: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("add module response [code]: %d, [status]: '%s'", resp.StatusCode, resp.Status)
	}

	log.Infof("scraped and submitted for module [package]: %s, [version]: %s, [sparks]: %d, connectors: %d",
		action.Module.Package, action.Module.Version, len(action.Sparks), len(action.Connectors))
}

//********************************************************************************************
// HELPER
//********************************************************************************************
