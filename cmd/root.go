package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "netflux",
	Short: "Pushing weather data to influxdb",
	Long:  `netflux is a tool to push netatmo weather data to influxdb`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exec")
	},
}

var netatmorurl string
var netatmor_user string
var netatmor_password string
var netatmor_clientid string
var netatmor_clientsecret string

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().StringVar(&netatmorurl, "netatmo_url", "r", "Url to Netatmo Weather API")
	rootCmd.PersistentFlags().StringVar(&netatmor_user, "netatmo_user", "", "User to Netatmo Weather API")
	rootCmd.PersistentFlags().StringVar(&netatmor_password, "netatmo_password", "", "Password to Netatmo Weather API")
	rootCmd.PersistentFlags().StringVar(&netatmor_clientid, "netatmo_clientid", "", "Client-ID to Netatmo Weather API")
	rootCmd.PersistentFlags().StringVar(&netatmor_clientsecret, "netatmo_clientsecret", "", "Client-Secret to Netatmo Weather API")

	rootCmd.MarkPersistentFlagRequired("netatmo_user")
	rootCmd.MarkPersistentFlagRequired("netatmo_password")
	rootCmd.MarkPersistentFlagRequired("netatmo_clientid")
	rootCmd.MarkPersistentFlagRequired("netatmo_clientsecret")
}
