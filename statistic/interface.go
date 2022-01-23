package statistic

type EventStatistic interface {
	IncEvent(name string)
	GetEventStatisticByName(name string) float64
	GetAllEventStatistic() map[string]float64
	PrintStatistic()
}
