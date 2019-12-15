package netatmo

type Client interface {
	GetStations() []string
	GetModules(station string) []string
	GetMetrics(station string, module string) []Metric
	UpdateCache()
}
