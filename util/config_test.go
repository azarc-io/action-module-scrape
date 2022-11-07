package util_test

import (
	"github.com/azarc-io/action-module-scrape/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

type loader struct {
	m map[string]string
}

func (l *loader) Getenv(k string) string {
	return l.m[k]
}

func TestConfigLoad(t *testing.T) {
	const v = "blap"
	config := util.LoadConfig(&testWrap{t: t}, &loader{
		m: map[string]string{
			"INPUT_TOKEN":   v,
			"INPUT_VERSION": v,
			"INPUT_RESOURCES": `
				docker: foo/bar:v1
				linux-amd64: bar-linux-amd64-v1
			`,
		}})
	assert.Equal(t, config.Version, v)
	resources := config.ResourcesAsMap()
	if resources != nil {
		assert.Equal(t, resources, map[string]string{
			"docker":      "foo/bar:v1",
			"linux-amd64": "bar-linux-amd64-v1",
		})
	}
}
