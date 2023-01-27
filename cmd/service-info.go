package cmd

import (
	"github.com/spf13/cobra"
)

var (
	serviceInfoCmd = &cobra.Command{
		Use:   "service-info",
		Short: "Get service info",
		Long:  `Get service info.`,
		Args:  cobra.MaximumNArgs(0),
		RunE:  getServiceInfo,
	}
)

func init() {
	rootCmd.AddCommand(serviceInfoCmd)
}

func getServiceInfo(cmd *cobra.Command, args []string) error {
	resp, err := Client.GetServiceInfoWithResponse(Ctx)
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case 200:
		payload := resp.JSON200
		return printJSON(payload, 1, resp.HTTPResponse)
	default:
		return printError(resp.Body, resp.HTTPResponse)
	}
}
