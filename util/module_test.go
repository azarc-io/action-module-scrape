package util_test

import (
	"github.com/azarc-io/action-module-scrape/temp/module_v1"
	"github.com/azarc-io/action-module-scrape/util"
	"testing"
)

//********************************************************************************************
// SUBMIT
//********************************************************************************************

func TestSubmit(t *testing.T) {
	util.SubmitAction(
		&testWrap{t: t},
		&util.Config{
			Token:          "54545454545",
			SubmissionHost: "https://auth-events.cloud.azarc.dev",
		},
		&module_v1.Action{
			Module: &module_v1.Module{
				Package:     "vth.azarc.hello-world",
				Version:     "v0.0.0",
				Repo:        "0",
				Label:       "0",
				Description: "0",
			},
		},
	)
}
