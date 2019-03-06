package cmd

import (
	"fmt"
	"os"

	"github.com/aidun/netflux/daemon"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "netflux",
	Short: "Pushing weather data to influxdb",
	Long:  `netflux is a tool to push netatmo weather data to influxdb`,
	Run: func(cmd *cobra.Command, args []string) {
		nd := daemon.NewNetfluxDaemon(
			netatmo_clientid,
			netatmo_clientsecret,
			netatmo_user,
			netatmo_password,
			influxdb_url,
			influxdb_user,
			influxdb_password,
			influxdb_database,
		)

		nd.Start()
	},
}

var netatmo_user string
var netatmo_password string
var netatmo_clientid string
var netatmo_clientsecret string
var influxdb_url string
var influxdb_user string
var influxdb_password string
var influxdb_database string

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().StringVar(&netatmo_user, "netatmo_user", "", "User to Netatmo Weather API")
	rootCmd.PersistentFlags().StringVar(&netatmo_password, "netatmo_password", "", "Password to Netatmo Weather API")
	rootCmd.PersistentFlags().StringVar(&netatmo_clientid, "netatmo_clientid", "", "Client-ID to Netatmo Weather API")
	rootCmd.PersistentFlags().StringVar(&netatmo_clientsecret, "netatmo_clientsecret", "", "Client-Secret to Netatmo Weather API")

	rootCmd.PersistentFlags().StringVar(&influxdb_url, "influxdb_url", "", "API of the influxdb instance")
	rootCmd.PersistentFlags().StringVar(&influxdb_user, "influxdb_user", "", "User of the influxdb with write access")
	rootCmd.PersistentFlags().StringVar(&influxdb_password, "influxdb_password", "", "Password of the influxdb user")
	rootCmd.PersistentFlags().StringVar(&influxdb_database, "influxdb_database", "", "Database")

	rootCmd.MarkPersistentFlagRequired("netatmo_user")
	rootCmd.MarkPersistentFlagRequired("netatmo_password")
	rootCmd.MarkPersistentFlagRequired("netatmo_clientid")
	rootCmd.MarkPersistentFlagRequired("netatmo_clientsecret")

	rootCmd.MarkPersistentFlagRequired("influxdb_url")
	rootCmd.MarkPersistentFlagRequired("influxdb_user")
	rootCmd.MarkPersistentFlagRequired("influxdb_password")
	rootCmd.MarkPersistentFlagRequired("influxdb_database")
}
