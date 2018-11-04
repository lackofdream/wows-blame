package main

import (
	"net/http"

	"github.com/skratchdot/open-golang/open"
	. "gitlab.com/lackofdream/wows-blame"
	"time"
)

func main() {
	go func() {
		time.Sleep(1 * time.Second)
		open.Start("http://localhost:8080")
	}()
	http.Handle("/", Router)
	http.ListenAndServe(":8080", nil)
}
