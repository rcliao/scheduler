package main

import (
	"net/http"
	"time"

	"github.com/rcliao/scheduler"
	"github.com/rcliao/scheduler/web"
	"github.com/rcliao/scheduler/worker"
)

func main() {
	schedulerInstance := worker.NewNativeScheduler()
	duration, _ := time.ParseDuration("10s")
	schedulerInstance.DelayJob(duration, "test", scheduler.DebugFunc)
	http.HandleFunc("/hello", web.HelloHandler())
	http.ListenAndServe(":9000", nil)
}
