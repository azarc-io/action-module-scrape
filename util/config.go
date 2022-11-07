package util

import (
	"fmt"
	"reflect"
	"strings"
)

type (
	Config struct {
		Path           string `env:"GITHUB_WORKSPACE"`
		RefType        string `env:"GITHUB_REF_TYPE"`
		Ref            string `env:"GITHUB_REF"`
		Repo           string `env:"GITHUB_REPOSITORY"`
		Token          string `env:"INPUT_TOKEN"`
		Version        string `env:"INPUT_VERSION"`
		SubmissionHost string `env:"INPUT_SUBMISSION_HOST"`
		InputResources string `env:"INPUT_RESOURCES"`
	}
	ConfigLoader interface {
		Getenv(key string) string
	}
)

func (c *Config) ResourcesAsMap() map[string]string {
	result := map[string]string{}
	if c.InputResources != "" {
		// split on new line
		nls := strings.Split(strings.TrimSpace(c.InputResources), "\n")
		for _, nl := range nls {
			// split by first colon
			parts := strings.SplitN(strings.TrimSpace(nl), ":", 2)
			if len(parts) > 2 {
				panic("invalid resource format, should be <resource>: <value> when a new line separating each resource")
			}
			result[parts[0]] = strings.ReplaceAll(parts[1], " ", "")
		}
	}
	return result
}

func LoadConfig(log Logger, l ConfigLoader) *Config {
	config := &Config{}
	t := reflect.TypeOf(config).Elem()
	v := reflect.ValueOf(config).Elem()

	for i := 0; i < v.NumField(); i++ {
		v.Field(i).SetString(l.Getenv(t.Field(i).Tag.Get("env")))
	}

	if config.Token == "" {
		log.Fatalf("token is not set")
	}
	if config.Version == "" {
		if config.RefType != "tag" {
			log.Fatalf("you must either set the version or run on tag push")
		}
		if _, err := fmt.Sscanf(config.Ref, "refs/tags/%s", &config.Version); err != nil {
			log.Fatalf("getting tag push version: %s", err.Error())
		}
	}
	return config
}
