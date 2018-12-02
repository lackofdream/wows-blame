package wows_blame

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gitlab.com/lackofdream/wows-blame/model"
)

func registerSubRouter(r *mux.Router) {
	r.HandleFunc("/token", token)
	r.HandleFunc("/setup", setup)
	r.HandleFunc("/match", match)
}

func token(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ok": true,
		"data": map[string]string{
			"application_id": ApplicationId,
		},
		"message": "",
	})
}

func setup(w http.ResponseWriter, r *http.Request) {
	var rsp model.WowsBlameSetupResponse
	rsp.Ok = false
	if r.Method != "POST" {
		rsp.Message = "GET method not allowed"
		json.NewEncoder(w).Encode(rsp)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rsp.Message = err.Error()
		json.NewEncoder(w).Encode(rsp)
		return
	}

	var param model.WowsBlameSetupParam
	if err := json.Unmarshal(body, &param); err != nil {
		rsp.Message = err.Error()
		json.NewEncoder(w).Encode(rsp)
		return
	}

	validateFromParam(param, &rsp)

	if !rsp.Ok {
		json.NewEncoder(w).Encode(rsp)
		return
	}

	setupFromParam(param, true)
	json.NewEncoder(w).Encode(rsp)
}

func match(w http.ResponseWriter, _ *http.Request) {
	var rsp model.WowsBlameMatchResponse
	rsp.Ok = false
	body, err := ioutil.ReadFile(GamePath + string(os.PathSeparator) + "replays" + string(os.PathSeparator) + "tempArenaInfo.json")
	if err != nil {
		rsp.Message = err.Error()
		json.NewEncoder(w).Encode(rsp)
		return
	}

	if err := json.Unmarshal(body, &rsp.Data); err != nil {
		rsp.Message = err.Error()
		json.NewEncoder(w).Encode(rsp)
		return
	}

	rsp.Ok = true
	json.NewEncoder(w).Encode(rsp)
}

func setupFromParam(param model.WowsBlameSetupParam, writeFile bool) error {

	if writeFile {
		body, err := json.Marshal(param)
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(".config", body, 0644); err != nil {
			return err
		}
	}

	SetupFlag = true
	ApplicationId = param.ApplicationID
	GamePath = param.GamePath

	return nil
}

func validateFromParam(param model.WowsBlameSetupParam, result *model.WowsBlameSetupResponse) {

	result.AppIDOk = false
	result.AppIDMessage = "please check your application id again"
	result.PathOk = false
	result.PathMessage = "WorldOfWarships.exe not found in this path"

	if _, err := os.Stat(param.GamePath + string(os.PathSeparator) + "WorldOfWarships.exe"); err == nil {
		result.PathOk = true
		result.PathMessage = ""
	}

	req, err := http.NewRequest("GET", WowsAsiaApiUrl+"/encyclopedia/info/", nil)
	if err != nil {
		return
	}
	q := req.URL.Query()
	q.Add("application_id", param.ApplicationID)
	req.URL.RawQuery = q.Encode()

	rsp, err := http.Get(req.URL.String())
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return
	}

	var infoObj map[string]*json.RawMessage
	if err := json.Unmarshal(data, &infoObj); err != nil {
		return
	}

	var status string
	if err := json.Unmarshal(*infoObj["status"], &status); err != nil {
		return
	}
	if status == "ok" {
		result.AppIDOk = true
		result.AppIDMessage = ""
	}

	if result.AppIDOk && result.PathOk {
		result.Ok = true
	}
}
