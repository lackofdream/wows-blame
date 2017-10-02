package wows_blame

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type JSONRouter struct {
	H http.Handler
}

func (j *JSONRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api/") {
		w.Header().Set("Content-Type", "application/json")
		if !setup_flag {
			http.Error(w, "not set up yet", http.StatusUpgradeRequired)
			return
		}
	}
	j.H.ServeHTTP(w, r)
}

var Router *JSONRouter

func init() {
	r := mux.NewRouter()

	registerSubRouter(r.PathPrefix("/api").Subrouter())

	r.Handle("/", http.FileServer(http.Dir("./webapp/dist")))

	Router = &JSONRouter{r}
}
