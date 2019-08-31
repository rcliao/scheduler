# Scheduler

A generic scheduler for scheduling delay jobs

## Installation

`go get -u github.com/rcliao/scheduler`

## Usages (TBD)

```
import "github.com/rcliao/scheduler"

client = scheduler.NewClient()
queueClient = rabbitmqClient()

client.delayJob(2 * scheduler.DAYS, func() {
    queueClient.send(job)
})
```
