package statistic

import (
	"fmt"
	"github.com/timfame/events-statistic.git/clock"
	"time"
)

type EventsHandler struct {
	clock        clock.Clock
	eventsByName map[string][]time.Time
}

func NewEventsHandler(clock clock.Clock) *EventsHandler {
	return &EventsHandler{
		clock:        clock,
		eventsByName: make(map[string][]time.Time),
	}
}

func (h *EventsHandler) IncEvent(name string) {
	h.eventsByName[name] = append(h.eventsByName[name], h.clock.Now())
}

func (h *EventsHandler) getEventCountByName(name string) int {
	timestamps, ok := h.eventsByName[name]
	if !ok {
		return 0
	}
	now := h.clock.Now()
	count := 0
	for _, timestamp := range timestamps {
		if timestamp.Add(time.Hour).After(now) {
			count++
		}
	}
	return count
}

func (h *EventsHandler) GetEventStatisticByName(name string) float64 {
	return float64(h.getEventCountByName(name)) / 60.0
}

func (h *EventsHandler) GetAllEventStatistic() map[string]float64 {
	all := make(map[string]float64)
	for name := range h.eventsByName {
		if stat := h.GetEventStatisticByName(name); stat != 0.0 {
			all[name] = stat
		}
	}
	return all
}

func (h *EventsHandler) PrintStatistic() {
	for name, rpm := range h.GetAllEventStatistic() {
		fmt.Printf("Name: %s, RPM: %f;\n", name, rpm)
	}
}
