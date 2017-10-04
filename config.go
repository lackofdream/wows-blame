package wows_blame

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/lackofdream/wows-blame/model"
)

const (
	WOWS_AISA_API_URL string = "https://api.worldofwarships.asia/wows"
)

var (
	SETUP_FLAG     bool
	APPLICATION_ID string
	GAME_PATH      string
)

func init() {
	SETUP_FLAG = false
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
