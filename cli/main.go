package main

import (
	"net/http"

	. "github.com/lackofdream/wows-blame"
	"time"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	go func() {
		time.Sleep(1 * time.Second)
		open.Start("http://localhost:8080")
	}()
	http.Handle("/", Router)
	http.ListenAndServe(":8080", nil)
}
