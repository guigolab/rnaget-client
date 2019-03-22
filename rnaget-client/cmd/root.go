package cmd

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/guigolab/rnaget-client/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// DefaultLocation is the deafatul location
	DefaultLocation = "crg"
)

var (
	// Client is a shared GA4GHRnaget instance
	Client *client.GA4GHRnaget
	// AuthInfo stores authentication info
	AuthInfo runtime.ClientAuthInfoWriter

	rootCmd = &cobra.Command{
		Use:              "rnaget-client",
		Short:            "A demo client for the GA4gh rnaget API",
		Long:             `A demo client for the GA4gh rnaget API`,
		PersistentPreRun: getConfig,
		SilenceUsage:     true,
		SilenceErrors:    true,
		Run:              nil,
	}
)

func init() {
	cobra.OnInitialize(initViper)
	rootCmd.PersistentFlags().StringP("location", "l", "", "Server location")
	viper.BindPFlag("location", rootCmd.PersistentFlags().Lookup("location"))
	viper.SetDefault("location", DefaultLocation)
}

func initViper() {
	viper.SetConfigName("rnaget-client")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal(fmt.Errorf("Fatal error config file: %s", err))
		}
	}
}

func getConfig(*cobra.Command, []string) {
	switch strings.ToLower(viper.GetString("location")) {
	case "crg":
		AuthInfo = runtime.ClientAuthInfoWriterFunc(func(req runtime.ClientRequest, reg strfmt.Registry) error {
			return req.SetHeaderParam("Authorization", "Bearer abcdefuvwxyz")
		})
		tc := client.DefaultTransportConfig().WithHost("genome.crg.cat")
		Client = client.NewHTTPClientWithConfig(nil, tc)
	case "caltech":
		AuthInfo = nil
		tc := client.DefaultTransportConfig().
			WithHost("felcat.caltech.edu").
			WithSchemes([]string{"http"})
		Client = client.NewHTTPClientWithConfig(nil, tc)
	}
	return
}

func print(obj interface{}, l int) error {
	r, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", r)
	re := regexp.MustCompile(`(\[\])?\*models.`)
	t := strings.ToLower(re.ReplaceAllString(fmt.Sprintf("%T", obj), ""))
	log.Infof("Got %d %s(s) \n", l, t)
	return nil
}

// Execute is the main function of the root command
func Execute() error {
	return rootCmd.Execute()
}
