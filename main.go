package main

import (
	"github.com/azarc-io/action-module-scrape/temp/module_v1"
	"github.com/azarc-io/action-module-scrape/util"
	"github.com/sethvargo/go-githubactions"
)

func main() {
	gitAction := githubactions.New()
	config := util.LoadConfig(gitAction, gitAction)

	action := &module_v1.Action{}
	action.Module = util.LoadModule(gitAction, config)
	action.Sparks = util.LoadSparks(gitAction, config)
	action.Connectors = util.LoadConnectors(gitAction, config)

	util.SubmitAction(gitAction, config, action)
}
