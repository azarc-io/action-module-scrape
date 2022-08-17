package test_test

import (
	"github.com/azarc-io/action-module-scrape/util"
	"testing"
)

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
	s.t.Logf(msg, args...)
	s.failed = true
}

func TestLoadSmallPNG(t *testing.T) {
	util.LoadImage(t, "image/small.png")
}

func TestLoadLargePNG(t *testing.T) {
	util.LoadImage(newShouldFail(t), "image/large.png")
}

func TestLoadSmallJPG(t *testing.T) {
	util.LoadImage(t, "image/small.jpg")
}

func TestLoadUnknown(t *testing.T) {
	util.LoadImage(newShouldFail(t), "image/small.webp")
}
