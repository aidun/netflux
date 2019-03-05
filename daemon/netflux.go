package daemon

import (
	netatmo "github.com/exzz/netatmo-api-go"
	client "github.com/influxdata/influxdb1-client"
	"log"
	"net/url"
	"time"
)

type netfluxDaemon struct {
	netatmo           *netatmo.Client
	influxdb          *client.Client
	influxdb_database string
}

func NewNetfluxDaemon(
	netatmo_clientid string,
	netatmo_clientsecret string,
	netatmo_user string,
	netatmo_password string,

	influxdb_url string,
	influxdb_user string,
	influxdb_password string,
	influxdb_database string,
) *netfluxDaemon {

	// Netatmo
	n, err := netatmo.NewClient(netatmo.Config{
		ClientID:     netatmo_clientid,
		ClientSecret: netatmo_clientsecret,
		Username:     netatmo_user,
		Password:     netatmo_password,
	})

	if err != nil {
		log.Fatalf("Could not create Netatmo client: %s\n", err)
	}

	// Influxdb
	host, err := url.Parse(influxdb_url)
	if err != nil {
		log.Fatalf("Could not create Influxdb-Config %s", err)
	}

	conf := client.Config{
		URL:      *host,
		Username: influxdb_user,
		Password: influxdb_password,
	}

	con, err := client.NewClient(conf)
	if err != nil {
		log.Fatalf("Could not create Influxdb connection %s", err)
	}
	log.Printf("Connection to influx db estabilished: %s", con)

	return &netfluxDaemon{
		netatmo:           n,
		influxdb:          con,
		influxdb_database: influxdb_database,
	}
}

func (nd *netfluxDaemon) Start() {
	for {
		dc, err := nd.netatmo.Read()

		if err != nil {
			log.Println(err)
			continue
		}

		var pts = make([]client.Point, 1)

		for _, station := range dc.Stations() {

			for _, module := range station.Modules() {

				if module.DashboardData.Temperature != nil {
					pts = append(pts, client.Point{
						Measurement: "temperature",
						Tags: map[string]string{
							"station_name": station.StationName,
							"module_name":  module.ModuleName,
						},
						Fields: map[string]interface{}{
							"now": *module.DashboardData.Temperature,
						},
						Time: time.Now(),
					})
				}

				if module.DashboardData.Humidity != nil {
					pts = append(pts, client.Point{
						Measurement: "humidity",
						Tags: map[string]string{
							"station_name": station.StationName,
							"module_name":  module.ModuleName,
						},
						Fields: map[string]interface{}{
							"now": *module.DashboardData.Humidity,
						},
						Time: time.Now(),
					})
				}

				bps := client.BatchPoints{
					Points:          pts,
					Database:        nd.influxdb_database,
					RetentionPolicy: "autogen",
				}

				_, err := nd.influxdb.Write(bps)

				if err != nil {
					log.Printf("Could not write data to influxdb %s", err)
				}
			}

		}
		time.Sleep(10 * time.Second)
	}
}
