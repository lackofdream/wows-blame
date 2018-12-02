package wows_blame

import (
	"net/http"
	"strings"

	"encoding/json"

	"github.com/gorilla/mux"
)

type JSONRouter struct {
	H http.Handler
}

func (j *JSONRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api/") {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/api/setup" && !SetupFlag {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"ok":      false,
				"message": "not setup yet",
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
