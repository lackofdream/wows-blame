package wows_blame

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"gitlab.com/lackofdream/wows-blame/model"
)

const (
	WowsAsiaApiUrl string = "https://api.worldofwarships.asia/wows"
)

var (
	SetupFlag     bool
	ApplicationId string
	GamePath      string
)

func init() {
	SetupFlag = false
	info, err := os.Stat(".config")
	if err != nil {
		return
	}
	if info.IsDir() {
		log.Println("please delete the .config directory then restart me")
		os.Exit(1)
	}
	content, err := ioutil.ReadFile(".config")
	if err != nil {
		return
	}

	var param model.WowsBlameSetupParam
	if err := json.Unmarshal(content, &param); err != nil {
		return
	}

	if err := setupFromParam(param, false); err != nil {
		return
	}
}
