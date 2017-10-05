package wows_blame

import (
	"net/http"
	"strings"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/lackofdream/wows-blame/model"
)

type JSONRouter struct {
	H http.Handler
}

func (j *JSONRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api/") {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/api/setup" && !SETUP_FLAG {
			json.NewEncoder(w).Encode(model.WowsBlameVersion{
				Ok:      false,
				Message: "not setup yet",
			})
			return
		}
	}
	j.H.ServeHTTP(w, r)
}

var Router *JSONRouter

func init() {
	r := mux.NewRouter()

	registerSubRouter(r.PathPrefix("/api").Subrouter())

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./webapp/dist")))

	Router = &JSONRouter{r}
}
