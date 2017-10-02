package wows_blame

import (
	"encoding/json"
	"net/http"

	"log"

	"fmt"
	"io/ioutil"

	"github.com/gorilla/mux"
	"github.com/lackofdream/wows-blame/model"
)

const (
	WOWS_AISA_API_URL string = "https://api.worldofwarships.asia/wows"
)

var (
	setup_flag     bool = false
	application_id string
	game_path      string
)

func registerSubRouter(r *mux.Router) {
	r.HandleFunc("/version", version)
	r.HandleFunc("/player", player)
}

func version(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.WowsBlameVersion{
		Ok:      true,
		Version: "0.1.0",
	})
}

func player(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", WOWS_AISA_API_URL+"/account/list/", nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("application_id", application_id)
	q.Add("search", r.URL.Query().Get("name"))
	req.URL.RawQuery = q.Encode()

	rsp, err := http.Get(req.URL.String())
	if err != nil {
		log.Fatal(err)
	}

	var accountList *model.AccountListResponse

	json.NewDecoder(rsp.Body).Decode(&accountList)

	if len(accountList.Data) == 0 {
		http.Error(w, "player not found", http.StatusNotFound)
		return
	}

	req, err = http.NewRequest("GET", WOWS_AISA_API_URL+"/account/info/", nil)
	if err != nil {
		log.Fatal(err)
	}

	q = req.URL.Query()
	q.Add("application_id", application_id)
	q.Add("account_id", fmt.Sprintf("%d", accountList.Data[0].AccountID))
	req.URL.RawQuery = q.Encode()

	rsp, err = http.Get(req.URL.String())
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(body)
}
