package netatmo

type NetatmoApiGoClient struct {
	Client
}

func (c *NetatmoApiGoClient) GetStations() []string {
	return nil
}

func (c *NetatmoApiGoClient) GetModules(station string) []string {
	return nil
}

func (c *NetatmoApiGoClient) GetMetrics(station string, module string) []Metric {
	return nil
}
