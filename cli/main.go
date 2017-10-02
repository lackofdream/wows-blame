package main

import (
	"net/http"

	. "github.com/lackofdream/wows-blame"
)

func main() {
	http.Handle("/", Router)
	http.ListenAndServe(":8080", nil)
}
