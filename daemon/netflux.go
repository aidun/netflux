package daemon

import (
	"log"
	"net/url"
	"time"

	netatmo "github.com/exzz/netatmo-api-go"
	client "github.com/influxdata/influxdb1-client"
)

type NetfluxDaemon struct {
	netatmo              *netatmo.Client
	influxdb             *client.Client
	influxdbDatabaseName string
}

func NewNetfluxDaemon(
	netatmoClientid string,
	netatmoClientsecret string,
	netatmoUser string,
	netatmoPassword string,

	influxdbURL string,
	influxdbUser string,
	influxdbPassword string,
	influxdbDatabaseName string,
) *NetfluxDaemon {

	// Netatmo
	n, err := netatmo.NewClient(netatmo.Config{
		ClientID:     netatmoClientid,
		ClientSecret: netatmoClientsecret,
		Username:     netatmoUser,
		Password:     netatmoPassword,
	})

	if err != nil {
		log.Fatalf("Could not create Netatmo client: %s\n", err)
	}

	// Influxdb
	host, err := url.Parse(influxdbURL)
	if err != nil {
		log.Fatalf("Could not create Influxdb-Config %s", err)
	}

	conf := client.Config{
		URL:      *host,
		Username: influxdbUser,
		Password: influxdbPassword,
	}

	con, err := client.NewClient(conf)
	if err != nil {
		log.Fatalf("Could not create Influxdb connection %s", err)
	}
	log.Printf("Connection to influx db estabilished: %v", con)

	return &NetfluxDaemon{
		netatmo:              n,
		influxdb:             con,
		influxdbDatabaseName: influxdbDatabaseName,
	}
}

func (nd *NetfluxDaemon) Start() {
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
					Database:        nd.influxdbDatabaseName,
					RetentionPolicy: "autogen",
				}

				_, err := nd.influxdb.Write(bps)

				if err != nil {
					log.Printf("Could not write data to influxdb: %s", err)
				}
			}

		}
		time.Sleep(10 * time.Second)
	}
}
