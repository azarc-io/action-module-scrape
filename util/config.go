package util

import (
	"fmt"
	"reflect"
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
	}
	ConfigLoader interface {
		Getenv(key string) string
	}
)

func LoadConfig(l ConfigLoader) *Config {
	config := &Config{}
	t := reflect.TypeOf(config).Elem()
	v := reflect.ValueOf(config).Elem()

	for i := 0; i < v.NumField(); i++ {
		v.Field(i).SetString(l.Getenv(t.Field(i).Tag.Get("env")))
	}
	return config
}

func (config *Config) Validate(log Logger) {
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
}
