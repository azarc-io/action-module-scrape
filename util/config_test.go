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
	config := util.LoadConfig(&testWrap{t: t}, &loader{m: map[string]string{"INPUT_TOKEN": v, "INPUT_VERSION": v}})
	assert.Equal(t, config.Version, v)
}
