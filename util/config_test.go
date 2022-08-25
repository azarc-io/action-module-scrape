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
	config := util.LoadConfig(&loader{m: map[string]string{"GITHUB_WORKSPACE": v}})
	assert.Equal(t, config.Path, v)
}
