package cmd

import (
	"bufio"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/guigolab/rnaget-client/pkg/api"
	"github.com/guigolab/rnaget-client/pkg/version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	DefaultConfigURL = "https://raw.githubusercontent.com/guigolab/rnaget-client/master/.rnaget-client.yml"
)

var (
	// Client is a shared GA4GHRnaget instance
	Client      *api.ClientWithResponses
	Ctx, Cancel = context.WithCancel(context.Background())

	terms = map[string]string{
		"study":   "studies",
		"project": "projects",
		"matrix":  "matrices",
	}

	rootCmd = &cobra.Command{
		Use:              "rnaget-client",
		Short:            "A demo client for the GA4GH RNAget API",
		Long:             `A demo client for the GA4GH RNAget API`,
		Version:          version.Get(),
		PersistentPreRun: getConfig,
		SilenceUsage:     true,
		SilenceErrors:    true,
	}
)

func init() {
	initViper()
	rootCmd.PersistentFlags().StringP("location", "l", viper.GetString("location"), "Server location")
	rootCmd.PersistentFlags().CountP("verbose", "V", "Verbosity")
	err := viper.BindPFlag("location", rootCmd.PersistentFlags().Lookup("location"))
	if err != nil {
		log.Fatal(err)
	}
}

func initViper() {
	viper.SetConfigName(".rnaget-client")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("rnaget")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err = getDefaultConfig(); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}
}

func getDefaultConfig() error {
	viper.SetConfigType("yaml")

	// Get the default config data
	resp, err := http.Get(DefaultConfigURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Debug("Read remote config from ", DefaultConfigURL)

	return viper.ReadConfig(resp.Body)
}

func getConfig(cmd *cobra.Command, args []string) {
	var err error
	v, err := cmd.Flags().GetCount("verbose")
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(log.WarnLevel + log.Level(v))

	l := viper.GetString("location")
	if l == "" {
		log.Fatalf("Please specify a server location")
	}
	server := viper.Sub(fmt.Sprintf("servers.%s", l))
	if server == nil {
		log.Fatalf("Server location not found: %s", string(l))
	}
	reqEditor := func(ctx context.Context, req *http.Request) error {
		token := server.GetString("token")
		if len(token) > 0 {
			v := fmt.Sprintf("Bearer %s", token)
			req.Header.Add("Authorization", v)
		}
		return nil
	}
	Client, err = api.NewClientWithResponses(server.GetString("baseUrl"), api.WithRequestEditorFn(reqEditor))
	if err != nil {
		log.Fatal(err)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	http.DefaultClient.Transport = tr
}

func printBytes(payload []byte, l int, f *os.File, resp *http.Response) error {
	defer f.Close()
	contentType := strings.Split(resp.Header.Get("content-type"), ";")[0]

	fmt.Fprintln(os.Stderr, "         Host :", resp.Request.URL.Host)
	fmt.Fprintln(os.Stderr, "       Status :", resp.Status)
	fmt.Fprintln(os.Stderr, " Content-Type :", contentType)

	fo, _ := f.Stat()
	namedPipe := (fo.Mode() & os.ModeNamedPipe) == 0
	pipe := (fo.Mode() & os.ModeCharDevice) == 0
	redir := namedPipe && pipe
	var w io.Writer = f
	if redir {
		log.Debug("Using buffered writer")
		w = bufio.NewWriter(f)
	}
	switch contentType {
	case "text/tab-separated-values":
		if !redir {
			fmt.Fprintln(os.Stderr, "      Payload : ")
		}
		fmt.Fprintf(w, "%s\n", payload)
	case "application/vnd.loom":
		if !redir {
			log.Warn(`Writing loom file content to stdout might mess up your terminal screen. Please redirect stdout to a file or specify an output file by using the '--output|-o' flag instead.`)
		} else {
			fmt.Fprintf(w, "%s\n", payload)
		}
	default:
		fmt.Fprintln(os.Stderr, "      Payload : ")
		fmt.Fprintf(w, "%s\n", payload)
	}
	t := "matrix"
	if l != 1 {
		t = terms[t]
	}
	log.Infof("Got %d %s", l, t)
	return nil
}

func printJSON(obj interface{}, l int, resp *http.Response) error {
	fmt.Fprintln(os.Stderr, "   Host :", resp.Request.URL.Host)
	fmt.Fprintln(os.Stderr, " Status :", resp.Status)
	if obj != nil {
		payload, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stderr, "Payload :")
		fmt.Printf("%s\n", payload)
	}
	re := regexp.MustCompile(`(\[\])?\*(\[\])?api.`)
	t := strings.ToLower(re.ReplaceAllString(fmt.Sprintf("%T", obj), ""))
	if l != 1 {
		t = terms[t]
	}
	log.Infof("Got %d %s\n", l, t)
	return nil
}

func printError(body []byte, resp *http.Response) error {
	e := api.Error{}
	err := json.Unmarshal(body, &e)
	if err != nil {
		return printBytes(body, 0, os.Stderr, resp)
	}
	return printJSON(e, 0, resp)
}

// Execute is the main function of the root command
func Execute() error {
	return rootCmd.Execute()
}
