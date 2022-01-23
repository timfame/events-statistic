package main

import (
	"github.com/timfame/events-statistic.git/clock"
	"github.com/timfame/events-statistic.git/statistic"
	"time"
)

func main() {
	handler := statistic.NewEventsHandler(clock.NewReal())
	handler.IncEvent("before")
	time.Sleep(time.Second)
	handler.IncEvent("after")
	handler.IncEvent("before")
	handler.PrintStatistic()
}
