package cmd

import (
	"encoding/json"

	"github.com/guigolab/rnaget-client/pkg/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	projectCmd = &cobra.Command{
		Use:   "projects [id]",
		Short: "Query projects",
		Long:  `Query projects.`,
		Args:  cobra.MaximumNArgs(1),
		RunE:  getProjects,
	}
)

func init() {
	setupProjectCommand()
	rootCmd.AddCommand(projectCmd)
}

func getProjectByID(id string) error {
	resp, err := Client.GetProjectByIdWithResponse(Ctx, id)
	if err != nil {
		return err
	}
	switch resp.StatusCode() {
	case 200:
		payload := resp.JSON200
		return printJSON(payload, 1, resp.HTTPResponse)
	case 400:
		payload := resp.JSON400
		return printJSON(payload, 1, resp.HTTPResponse)
	case 404:
		payload := resp.JSON404
		return printJSON(payload, 1, resp.HTTPResponse)
	case 406:
		payload := resp.JSON406
		return printJSON(payload, 1, resp.HTTPResponse)
	case 501:
		payload := resp.JSON501
		return printJSON(payload, 1, resp.HTTPResponse)
	default:
		return printError(resp.Body, resp.HTTPResponse)
	}
}

func getProjects(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return getProjectByID(args[0])
	}
	params := api.GetProjectsParams{}
	ver := cmd.Flag("version")
	if ver.Changed {
		verString, err := cmd.Flags().GetString("version")
		if err != nil {
			return err
		}
		verParam := api.VersionParam(verString)
		params.Version = &verParam
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		p, err := json.MarshalIndent(params, "", "  ")
		log.Debugf("\n%s", p)
		if err != nil {
			return err
		}
	}
	resp, err := Client.GetProjectsWithResponse(Ctx, &params)
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case 200:
		payload := resp.JSON200
		return printJSON(payload, len(*payload), resp.HTTPResponse)
	case 400:
		payload := resp.JSON400
		return printJSON(payload, 1, resp.HTTPResponse)
	case 406:
		payload := resp.JSON406
		return printJSON(payload, 1, resp.HTTPResponse)
	case 501:
		payload := resp.JSON501
		return printJSON(payload, 1, resp.HTTPResponse)
	default:
		return printError(resp.Body, resp.HTTPResponse)
	}
}

func getProjectFilters(cmd *cobra.Command, args []string) error {
	resp, err := Client.GetProjectFiltersWithResponse(Ctx)
	if err != nil {
		return err
	}
	switch resp.StatusCode() {
	case 200:
		payload := resp.JSON200
		return printJSON(payload, 1, resp.HTTPResponse)
	case 400:
		payload := resp.JSON400
		return printJSON(payload, 1, resp.HTTPResponse)
	default:
		return printError(resp.Body, resp.HTTPResponse)
	}
}

func setupProjectCommand() {
	var filtersCmd = &cobra.Command{
		Use:   "filters",
		Short: "Get filters for project searches",
		Long:  `Get filters for project searches`,
		RunE:  getProjectFilters,
	}
	projectCmd.Flags().StringP("version", "v", "", "Search for a specific version (ignored when [id] is specified)")
	projectCmd.AddCommand(filtersCmd)
}
