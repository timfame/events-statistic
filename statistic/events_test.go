package statistic

import (
	"github.com/stretchr/testify/assert"
	"github.com/timfame/events-statistic.git/clock"
	"math"
	"testing"
	"time"
)

func equalFloats(x, y float64) bool {
	return math.Abs(x - y) < 0.000001
}

func TestEventsHandler_Simple(t *testing.T) {
	now := time.Now()
	fake := clock.NewFake()
	handler := NewEventsHandler(fake)

	fake.SetNext(now)
	handler.IncEvent("first")

	fake.SetNext(now.Add(time.Minute))
	handler.IncEvent("second")
	handler.IncEvent("first")

	stats := handler.GetAllEventStatistic()

	assert.Equal(t, 2, len(stats))
	assert.True(t, equalFloats(2. / 60., stats["first"]))
	assert.True(t, equalFloats(1. / 60., stats["second"]))
}

func TestEventsHandler_PartialExpired(t *testing.T) {
	now := time.Now()
	fake := clock.NewFake()
	handler := NewEventsHandler(fake)

	fake.SetNext(now)
	handler.IncEvent("first")

	fake.SetNext(now.Add(time.Minute))
	handler.IncEvent("second")

	fake.SetNext(now.Add(time.Hour + time.Second))
	stats := handler.GetAllEventStatistic()

	assert.Equal(t, 1, len(stats))
	assert.True(t, equalFloats(1. / 60., stats["second"]))
}

func TestEventsHandler_AllExpired(t *testing.T) {
	now := time.Now()
	fake := clock.NewFake()
	handler := NewEventsHandler(fake)

	fake.SetNext(now)
	handler.IncEvent("first")

	fake.SetNext(now.Add(time.Minute))
	handler.IncEvent("second")

	fake.SetNext(now.Add(time.Hour + time.Minute * 2))
	stats := handler.GetAllEventStatistic()

	assert.Equal(t, 0, len(stats))
}

func TestEventsHandler(t *testing.T) {
	now := time.Now()
	fake := clock.NewFake()
	handler := NewEventsHandler(fake)

	fake.SetNext(now)
	handler.IncEvent("first")
	handler.IncEvent("first")
	handler.IncEvent("second")
	handler.IncEvent("second")
	handler.IncEvent("second")
	handler.IncEvent("third")

	fake.SetNext(now.Add(time.Minute * 20))
	handler.IncEvent("second")
	handler.IncEvent("first")
	handler.IncEvent("first")
	handler.IncEvent("second")

	stats := handler.GetAllEventStatistic()

	assert.Equal(t, 3, len(stats))
	assert.True(t, equalFloats(4. / 60., stats["first"]), "Statistic for name \"first\"")
	assert.True(t, equalFloats(5. / 60., stats["second"]), "Statistic for name \"second\"")
	assert.True(t, equalFloats(1. / 60., stats["third"]), "Statistic for name \"third\"")

	fake.SetNext(now.Add(time.Minute * 65))

	stats = handler.GetAllEventStatistic()

	assert.Equal(t, 2, len(stats))
	assert.True(t, equalFloats(2. / 60., stats["first"]), "Statistic for name \"first\"")
	assert.True(t, equalFloats(2. / 60., stats["second"]), "Statistic for name \"second\"")
}
