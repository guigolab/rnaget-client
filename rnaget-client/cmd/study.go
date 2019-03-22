package cmd

import (
	"github.com/guigolab/rnaget-client/client/studies"
	"github.com/spf13/cobra"
)

var (
	studyCmd = &cobra.Command{
		Use:   "study",
		Short: "Study queries",
		Long:  `Study queries.`,
		// Args:  cobra.MinimumNArgs(1),
		Run: nil,
	}
)

func init() {
	setupStudyCommands()
	rootCmd.AddCommand(studyCmd)
}

func studyGet(cmd *cobra.Command, args []string) error {
	params := studies.NewGetStudyByIDParams().WithStudyID(args[0])
	result, err := Client.Studies.GetStudyByID(params, AuthInfo)
	if err != nil {
		return err
	}

	return print(result.Payload, 1)
}

func studySearch(cmd *cobra.Command, args []string) error {
	params := studies.NewSearchStudiesParams()
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
	result, err := Client.Studies.SearchStudies(params, AuthInfo)
	if err != nil {
		return err
	}

	return print(result.Payload, len(result.Payload))
}

func setupStudyCommands() {
	var get = &cobra.Command{
		Use:   "get",
		Short: "Get studies by ID",
		Long:  `Get studies by ID.`,
		Args:  cobra.ExactArgs(1),
		RunE:  studyGet,
	}
	var search = &cobra.Command{
		Use:   "search",
		Short: "Search studies",
		Long:  `Search studies.`,
		RunE:  studySearch,
	}
	search.Flags().StringP("version", "v", "", "Search for a specific version")
	search.Flags().StringSliceP("tags", "t", nil, "Search for specific tags")
	studyCmd.AddCommand(get, search)
}
