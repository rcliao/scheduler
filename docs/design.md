# Design Thoughts

This doc is to store ideas I have with scheduler design before settling with API.

## Why

In the past systems I've worked with, we eventually came across a need of
scheduling a one off event somewhere in the future. For example, I want to
schedule a SMS campaign at 8am tomorrow or at Tuesday. Furthermore, in the
micro-services world, it's hard to keep the implementation of scheduling jobs
consistent across services. Like, one service may decide to use CRON job to
handle its delayed jobs while other uses a delayed job functionality in their
message queue. Inconsistency causes the debugging and reasoning difficult. Plus,
relying on message queue to provide delay job functionality limits the choices
you have for message queue.

Therefore, I want to take my attempt to abstract out the scheduling needs
regardless of the system developers use. To do so, the most basic needs is simple
- to be able to schedule a event that should be triggered in the future. Once
the event is popped from the scheduled queue, we should put it to the actual
worker queue if necessary. In other word, the scheduler does not and should not
handle specific business needs; instead, it should sent off the work to worker
queue so the downstream worker can import or be more business specific.

## Why Choose GoLang?

Because of `time.AfterFunc` built-in function, we can easily schedule a event
in the runtime scheduler without writing a `while true` loop and check ourselves.
It's simple!

However, this does not persist the scheduling events. It implies that the scheduled
events will be lost if the process is ever restarted. To avoid data loss, we need
to persist the event data somewhere so we can put the events in memory whenever
the process is restarted.
