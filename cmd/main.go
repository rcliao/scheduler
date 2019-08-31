package main

import (
	"net/http"

	"github.com/rcliao/scheduler/web"
)

func main() {
	http.HandleFunc("/hello", web.HelloHandler())
	http.ListenAndServe(":9000", nil)
}
