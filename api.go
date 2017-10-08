package wows_blame

import (
	"encoding/json"
	"net/http"

	"fmt"
	"io/ioutil"

	"os"

	"strings"

	"errors"

	"github.com/gorilla/mux"
	"github.com/lackofdream/wows-blame/model"
)

func registerSubRouter(r *mux.Router) {
	r.HandleFunc("/version", version)
	r.HandleFunc("/player", playerWithShip)
	r.HandleFunc("/setup", setup)
	r.HandleFunc("/match", match)
}

func version(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.WowsBlameVersion{
		Ok: true,
		Data: struct {
			Version string `json:"version"`
		}{Version: "0.1.0"},
		Message: "",
	})
}

func getPlayerID(name string) (string, error) {
	if strings.HasPrefix(name, "Bot_") {
		return "", fmt.Errorf("cannot find player")
	}
	req, err := http.NewRequest("GET", WOWS_AISA_API_URL+"/account/list/", nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("application_id", APPLICATION_ID)
	q.Add("search", name)
	req.URL.RawQuery = q.Encode()

	rsp, err := http.Get(req.URL.String())
	if err != nil {
		return "", err
	}

	var accountList *model.AccountListResponse

	json.NewDecoder(rsp.Body).Decode(&accountList)

	for _, account := range accountList.Data {
		if strings.ToLower(account.Nickname) == strings.ToLower(name) {
			return fmt.Sprintf("%d", account.AccountID), nil
		}
	}
	return "", fmt.Errorf("cannot find player")
}

func getPlayerInfo(accountID string) ([]byte, error) {
	req, err := http.NewRequest("GET", WOWS_AISA_API_URL+"/account/info/", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("application_id", APPLICATION_ID)
	q.Add("account_id", accountID)
	req.URL.RawQuery = q.Encode()

	rsp, err := http.Get(req.URL.String())
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func handleWGPlayerInfoResponse(body []byte, accountID string, payload *model.WowsBlamePlayerPayload) error {

	payload.AccountID = accountID

	var playerInfoObj map[string]*json.RawMessage
	if err := json.Unmarshal(body, &playerInfoObj); err != nil {
		return err
	}

	var playerInfoData map[string]*json.RawMessage
	if err := json.Unmarshal(*playerInfoObj["data"], &playerInfoData); err != nil {
		return err
	}

	var playerInfo map[string]*json.RawMessage
	if err := json.Unmarshal(*playerInfoData[accountID], &playerInfo); err != nil {
		return err
	}

	var hiddenProfile bool
	if err := json.Unmarshal(*playerInfo["hidden_profile"], &hiddenProfile); err != nil {
		return err
	}

	if hiddenProfile {
		return errors.New("player's data not public")
	}

	if err := json.Unmarshal(*playerInfo["nickname"], &payload.AccountName); err != nil {
		return err
	}

	var statistics map[string]*json.RawMessage
	if err := json.Unmarshal(*playerInfo["statistics"], &statistics); err != nil {
		return err
	}

	var pvpStatistics map[string]*json.RawMessage
	if err := json.Unmarshal(*statistics["pvp"], &pvpStatistics); err != nil {
		return err
	}

	var battles int
	var wins int
	if err := json.Unmarshal(*pvpStatistics["battles"], &battles); err != nil {
		return err
	}
	if err := json.Unmarshal(*pvpStatistics["wins"], &wins); err != nil {
		return err
	}
	payload.TotalBattleCount = battles
	payload.WinRate = float64(wins) / float64(battles)

	return nil
}

func getShipWiki(shipID string) ([]byte, error) {
	req, err := http.NewRequest("GET", WOWS_AISA_API_URL+"/encyclopedia/ships/", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("application_id", APPLICATION_ID)
	q.Add("ship_id", shipID)
	q.Add("language", "en")
	req.URL.RawQuery = q.Encode()

	rsp, err := http.Get(req.URL.String())
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func handleWGShipWiki(body []byte, shipID string, payload *model.WowsBlamePlayerPayload) error {
	payload.ShipID = shipID

	var shipInfoObj map[string]*json.RawMessage
	if err := json.Unmarshal(body, &shipInfoObj); err != nil {
		return err
	}

	var shipInfoData map[string]*json.RawMessage
	if err := json.Unmarshal(*shipInfoObj["data"], &shipInfoData); err != nil {
		return err
	}

	var shipInfo map[string]*json.RawMessage
	if err := json.Unmarshal(*shipInfoData[shipID], &shipInfo); err != nil {
		return err
	}

	if err := json.Unmarshal(*shipInfo["name"], &payload.ShipName); err != nil {
		return err
	}

	if err := json.Unmarshal(*shipInfo["type"], &payload.ShipType); err != nil {
		return err
	}
	return nil
}

func getShipStat(accountID string, shipID string) ([]byte, error) {
	req, err := http.NewRequest("GET", WOWS_AISA_API_URL+"/ships/stats/", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("application_id", APPLICATION_ID)
	q.Add("ship_id", shipID)
	q.Add("account_id", accountID)
	req.URL.RawQuery = q.Encode()

	rsp, err := http.Get(req.URL.String())
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
func handleWGShipStatWiki(body []byte, accountID string, payload *model.WowsBlamePlayerPayload) error {
	var shipStatInfoObj map[string]*json.RawMessage
	if err := json.Unmarshal(body, &shipStatInfoObj); err != nil {
		return err
	}

	var shipStatInfoData map[string][]*json.RawMessage
	if err := json.Unmarshal(*shipStatInfoObj["data"], &shipStatInfoData); err != nil {
		return err
	}

	var shipStatInfo map[string]*json.RawMessage
	if err := json.Unmarshal(*shipStatInfoData[accountID][0], &shipStatInfo); err != nil {
		return err
	}

	var pvpStat map[string]*json.RawMessage
	if err := json.Unmarshal(*shipStatInfo["pvp"], &pvpStat); err != nil {
		return err
	}

	var battles int
	var wins int
	if err := json.Unmarshal(*pvpStat["battles"], &battles); err != nil {
		return err
	}
	if err := json.Unmarshal(*pvpStat["wins"], &wins); err != nil {
		return err
	}
	payload.ShipBattleCount = battles
	payload.ShipWinRate = float64(wins) / float64(battles)

	return nil
}

func playerWithShip(w http.ResponseWriter, r *http.Request) {
	apiResponseData := model.WowsBlamePlayerPayload{}

	playerName := r.URL.Query().Get("name")
	shipID := r.URL.Query().Get("ship_id")

	accountID, err := getPlayerID(playerName)
	if err != nil {
		if err.Error() == "cannot find player" {
			json.NewEncoder(w).Encode(model.WowsBlamePlayer{
				Message: err.Error(),
				Ok:      false,
			})
			return
		} else {
			json.NewEncoder(w).Encode(model.WowsBlamePlayer{
				Message: err.Error(),
				Ok:      false,
			})
		}
		return
	}

	playerInfo, err := getPlayerInfo(accountID)
	if err != nil {
		json.NewEncoder(w).Encode(model.WowsBlamePlayer{
			Message: err.Error(),
			Ok:      false,
		})
		return
	}

	if err = handleWGPlayerInfoResponse(playerInfo, accountID, &apiResponseData); err != nil {
		json.NewEncoder(w).Encode(model.WowsBlamePlayer{
			Message: err.Error(),
			Ok:      false,
		})
		return
	}

	shipWiki, err := getShipWiki(shipID)
	if err != nil {
		json.NewEncoder(w).Encode(model.WowsBlamePlayer{
			Message: err.Error(),
			Ok:      false,
		})
		return
	}

	if err = handleWGShipWiki(shipWiki, shipID, &apiResponseData); err != nil {
		json.NewEncoder(w).Encode(model.WowsBlamePlayer{
			Message: err.Error(),
			Ok:      false,
		})
		return
	}

	shipStat, err := getShipStat(accountID, shipID)
	if err != nil {
		json.NewEncoder(w).Encode(model.WowsBlamePlayer{
			Message: err.Error(),
			Ok:      false,
		})
		return
	}

	if err = handleWGShipStatWiki(shipStat, accountID, &apiResponseData); err != nil {
		json.NewEncoder(w).Encode(model.WowsBlamePlayer{
			Message: err.Error(),
			Ok:      false,
		})
		return
	}

	json.NewEncoder(w).Encode(model.WowsBlamePlayer{
		Ok:   true,
		Data: apiResponseData,
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

func match(w http.ResponseWriter, r *http.Request) {
	var rsp model.WowsBlameMatchResponse
	rsp.Ok = false
	body, err := ioutil.ReadFile(GAME_PATH + string(os.PathSeparator) + "replays" + string(os.PathSeparator) + "tempArenaInfo.json")
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

	SETUP_FLAG = true
	APPLICATION_ID = param.ApplicationID
	GAME_PATH = param.GamePath

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

	req, err := http.NewRequest("GET", WOWS_AISA_API_URL+"/encyclopedia/info/", nil)
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
