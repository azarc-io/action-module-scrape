package test_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/azarc-io/action-module-scrape/temp/module_v1"
	"github.com/azarc-io/action-module-scrape/util"
	"net/http"
	"testing"
)

func TestSubmit(t *testing.T) {
	action := &module_v1.Action{
		Module: &module_v1.Module{
			Package:     "vth.azarc.hello-world",
			Version:     "v0.0.0",
			Repo:        "0",
			Label:       "0",
			Description: "0",
			Tags:        []string{"0"},
		},
	}

	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(action); err != nil {
		t.Fatalf(err.Error())
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"https://auth-events.cloud.azarc.dev/api/v1/module",
		buf,
	)
	req.Header.Set("Authorization", "fkdfjslkff-fd")
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf(resp.Status)
}

//********************************************************************************************
// FILE
//********************************************************************************************

func TestLoadFile(t *testing.T) {
	util.LoadFile(t, "!exist", false)
	util.LoadFile(newShouldFail(t), "!exist", true)
}

//********************************************************************************************
// IMAGE
//********************************************************************************************

func TestLoadSmallPNG(t *testing.T) {
	util.LoadImage(t, "image/small_png", true)
}

func TestLoadLargePNG(t *testing.T) {
	util.LoadImage(newShouldFail(t), "image/large_png", true)
}

func TestLoadSmallJPG(t *testing.T) {
	util.LoadImage(t, "image/jpg", true)
}

func TestLoadSVG(t *testing.T) {
	util.LoadImage(t, "image/svg", true)
}

func TestNotExistFile(t *testing.T) {
	util.LoadImage(t, "!exists", false)
	util.LoadImage(newShouldFail(t), "!exists", true)
}

func TestLoadUnknown(t *testing.T) {
	util.LoadImage(newShouldFail(t), "image/webp", true)
}

//********************************************************************************************
// HELPER
//********************************************************************************************

type shouldFail struct {
	t      *testing.T
	failed bool
}

func newShouldFail(t *testing.T) util.Logger {
	s := &shouldFail{t: t}
	t.Cleanup(func() {
		if !s.failed {
			s.t.Fatalf("should have failed")
		}
	})
	return s
}

func (s *shouldFail) Fatalf(msg string, args ...any) {
	if !s.failed {
		s.t.Logf(fmt.Sprintf("failed successfully: %s", msg), args...)
		s.failed = true
	}
}
