package web

import (
	"fmt"
	"net/http"
)

// HelloHandler renders the hello world!
func HelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	}
}
