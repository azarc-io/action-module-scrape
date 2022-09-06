package util_test

import (
	"fmt"
	"github.com/azarc-io/action-module-scrape/util"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

//********************************************************************************************
// CONFIG
//********************************************************************************************

func TestLoadSpark(t *testing.T) {
	sparks := util.LoadSparks(&testWrap{t: t}, &util.Config{Path: "test_files"})
	pb, err := structpb.NewStruct(map[string]interface{}{"foo": "bar"})
	assert.Nil(t, err)
	assert.Equal(t, sparks[0].Config, pb)
}

//********************************************************************************************
// FILE
//********************************************************************************************

func TestLoadFile(t *testing.T) {
	util.LoadFile(&testWrap{t: t}, "!exist", false)
	util.LoadFile(newShouldFail(t), "!exist", true)
}

//********************************************************************************************
// IMAGE
//********************************************************************************************

func TestLoadSmallPNG(t *testing.T) {
	util.LoadImage(&testWrap{t: t}, "test_files/small_png", true)
}

func TestLoadLargePNG(t *testing.T) {
	util.LoadImage(newShouldFail(t), "test_files/large_png", true)
}

func TestLoadSmallJPG(t *testing.T) {
	util.LoadImage(&testWrap{t: t}, "test_files/jpg", true)
}

func TestLoadSVG(t *testing.T) {
	util.LoadImage(&testWrap{t: t}, "test_files/svg", true)
}

func TestNotExistFile(t *testing.T) {
	util.LoadImage(&testWrap{t: t}, "!exists", false)
	util.LoadImage(newShouldFail(t), "!exists", true)
}

func TestLoadUnknown(t *testing.T) {
	util.LoadImage(newShouldFail(t), "test_files/webp", true)
}

//********************************************************************************************
// HELPER
//********************************************************************************************

type testWrap struct {
	t *testing.T
}

func (t *testWrap) Fatalf(fmt string, args ...interface{}) {
	t.t.Fatalf(fmt, args...)
}

func (t *testWrap) Infof(fmt string, args ...interface{}) {
	t.t.Logf(fmt, args...)
}

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

func (s *shouldFail) Infof(fmt string, args ...interface{}) {
	s.t.Logf(fmt, args...)
}
