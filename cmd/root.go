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
			netatmoClientid,
			netatmoClientsecret,
			netatmoUser,
			netatmoPassword,
			influxdbURL,
			influxdbUser,
			influxdbPassword,
			influxdbDatabase,
		)

		nd.Start()
	},
}

var netatmoUser string
var netatmoPassword string
var netatmoClientid string
var netatmoClientsecret string
var influxdbURL string
var influxdbUser string
var influxdbPassword string
var influxdbDatabase string

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().StringVar(&netatmoUser, "netatmo_user", "", "User to Netatmo Weather API")
	rootCmd.PersistentFlags().StringVar(&netatmoPassword, "netatmo_password", "", "Password to Netatmo Weather API")
	rootCmd.PersistentFlags().StringVar(&netatmoClientid, "netatmo_clientid", "", "Client-ID to Netatmo Weather API")
	rootCmd.PersistentFlags().StringVar(&netatmoClientsecret, "netatmo_clientsecret", "", "Client-Secret to Netatmo Weather API")

	rootCmd.PersistentFlags().StringVar(&influxdbURL, "influxdb_url", "", "API of the influxdb instance")
	rootCmd.PersistentFlags().StringVar(&influxdbUser, "influxdb_user", "", "User of the influxdb with write access")
	rootCmd.PersistentFlags().StringVar(&influxdbPassword, "influxdb_password", "", "Password of the influxdb user")
	rootCmd.PersistentFlags().StringVar(&influxdbDatabase, "influxdb_database", "", "Database")

	rootCmd.MarkPersistentFlagRequired("netatmo_user")
	rootCmd.MarkPersistentFlagRequired("netatmo_password")
	rootCmd.MarkPersistentFlagRequired("netatmo_clientid")
	rootCmd.MarkPersistentFlagRequired("netatmo_clientsecret")

	rootCmd.MarkPersistentFlagRequired("influxdb_url")
	rootCmd.MarkPersistentFlagRequired("influxdb_user")
	rootCmd.MarkPersistentFlagRequired("influxdb_password")
	rootCmd.MarkPersistentFlagRequired("influxdb_database")
}
