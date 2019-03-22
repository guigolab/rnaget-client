package cmd

import (
	"encoding/json"

	"github.com/guigolab/rnaget-client/client/projects"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	projectCmd = &cobra.Command{
		Use:   "project",
		Short: "Project queries",
		Long:  `Project queries.`,
		// Args:  cobra.MinimumNArgs(1),
		Run: nil,
	}
)

func init() {
	setupProjectCommands()
	rootCmd.AddCommand(projectCmd)
}

func projectGet(cmd *cobra.Command, args []string) error {
	params := projects.NewGetProjectByIDParams().WithProjectID(args[0])
	result, err := Client.Projects.GetProjectByID(params, AuthInfo)
	if err != nil {
		return err
	}

	return print(result.Payload, 1)
}

func projectSearch(cmd *cobra.Command, args []string) error {
	params := projects.NewSearchProjectsParams()
	ver := cmd.Flag("version")
	tags := cmd.Flag("tags")
	if ver.Changed {
		verString, err := cmd.Flags().GetString("version")
		if err != nil {
			return err
		}
		params.SetVersion(&verString)
	}
	if tags.Changed {
		tagSlice, err := cmd.Flags().GetStringSlice("tags")
		if err != nil {
			return err
		}
		params.SetTags(tagSlice)
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		p, _ := json.MarshalIndent(params, "", "  ")
		log.Debugf("\n%s", p)
	}
	result, err := Client.Projects.SearchProjects(params, AuthInfo)
	if err != nil {
		return err
	}

	return print(result.Payload, len(result.Payload))
}

func setupProjectCommands() {
	var get = &cobra.Command{
		Use:   "get",
		Short: "Get projects by ID",
		Long:  `Get projects by ID.`,
		Args:  cobra.ExactArgs(1),
		RunE:  projectGet,
	}
	var search = &cobra.Command{
		Use:   "search",
		Short: "Search projects",
		Long:  `Search projects.`,
		// Args:  cobra.MinimumNArgs(1),
		RunE: projectSearch,
	}
	search.Flags().StringP("version", "v", "", "Search for a specific version")
	search.Flags().StringSliceP("tags", "t", nil, "Search for specific tags")
	projectCmd.AddCommand(get, search)
}
