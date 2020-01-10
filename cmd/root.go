package cmd

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/guigolab/rnaget-client/pkg/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// DefaultLocation is the default location
	DefaultLocation = "crg"
)

var (
	// Client is a shared GA4GHRnaget instance
	Client      *api.ClientWithResponses
	Ctx, Cancel = context.WithCancel(context.Background())

	rootCmd = &cobra.Command{
		Use:              "rnaget-client",
		Short:            "A demo client for the GA4GH RNAget API",
		Long:             `A demo client for the GA4GH RNAget API`,
		PersistentPreRun: getConfig,
		SilenceUsage:     true,
		SilenceErrors:    true,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

func init() {
	cobra.OnInitialize(initViper)
	rootCmd.PersistentFlags().StringP("location", "l", "", "Server location")
	err := viper.BindPFlag("location", rootCmd.PersistentFlags().Lookup("location"))
	if err != nil {
		log.Fatal(err)
	}
	viper.SetDefault("location", DefaultLocation)
}

func initViper() {
	viper.SetConfigName("client-config")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("rnaget")
	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatal(err)
	}
}

func setRequestEditor(c *api.Client) error {
	l := viper.GetString("location")
	server := viper.Sub(fmt.Sprintf("servers.%s", l))
	c.RequestEditor = func(req *http.Request, ctx context.Context) error {
		token := server.GetString("token")
		if len(token) > 0 {
			v := fmt.Sprintf("Bearer %s", token)
			req.Header.Add("Authorization", v)
		}
		return nil
	}
	return nil
}

func getConfig(*cobra.Command, []string) {
	l := viper.GetString("location")
	server := viper.Sub(fmt.Sprintf("servers.%s", l))
	if server == nil {
		log.Fatalf("Server location not found: %s", l)
	}
	var err error
	Client, err = api.NewClientWithResponses(server.GetString("baseUrl"), setRequestEditor)
	if err != nil {
		log.Fatal(err)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	http.DefaultClient.Transport = tr
}

func print(obj interface{}, l int, resp *http.Response) error {
	r, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println("   Host :", resp.Request.URL.Host)
	fmt.Println(" Status :", resp.Status)
	fmt.Println("Payload :")
	fmt.Printf("%s\n", r)
	re := regexp.MustCompile(`(\[\])?\*models.`)
	t := strings.ToLower(re.ReplaceAllString(fmt.Sprintf("%T", obj), ""))
	log.Debugf("Got %d %s(s) \n", l, t)
	return nil
}

// Execute is the main function of the root command
func Execute() error {
	return rootCmd.Execute()
}
