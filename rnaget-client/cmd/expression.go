package cmd

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/guigolab/rnaget-client/client/expressions"
	"github.com/spf13/cobra"
)

var (
	expressionCmd = &cobra.Command{
		Use:   "expression",
		Short: "Expression queries",
		Long:  `Expression queries.`,
		Run:   nil,
	}
)

func init() {
	setupExpressionCommands()
	rootCmd.AddCommand(expressionCmd)
}

func expressionGet(cmd *cobra.Command, args []string) error {
	params := expressions.NewGetExpressionByIDParams().WithExpressionID(args[0])
	result, err := Client.Expressions.GetExpressionByID(params, AuthInfo)
	if err != nil {
		return err
	}

	return print(result.Payload, 1)
}

func expressionSearch(cmd *cobra.Command, args []string) error {
	params := expressions.NewSearchExpressionsParams()
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
	if cmd.Flag("study-id").Changed {
		studyID, err := cmd.Flags().GetString("study-id")
		if err != nil {
			return err
		}
		params.SetStudyID(&studyID)
	}
	if cmd.Flag("project-id").Changed {
		projectID, err := cmd.Flags().GetString("project-id")
		if err != nil {
			return err
		}
		params.SetProjectID(&projectID)
	}
	if cmd.Flag("sample-id").Changed {
		sampleSlice, err := cmd.Flags().GetStringSlice("sample-id")
		if err != nil {
			return err
		}
		params.SetSampleIDList(sampleSlice)
	}
	if cmd.Flag("feature-id").Changed {
		featureSlice, err := cmd.Flags().GetStringSlice("feature-id")
		if err != nil {
			return err
		}
		params.SetFeatureIDList(featureSlice)
	}
	if cmd.Flag("feature-name").Changed {
		featureSlice, err := cmd.Flags().GetStringSlice("feature-name")
		if err != nil {
			return err
		}
		params.SetFeatureNameList(featureSlice)
	}
	if cmd.Flag("feature-accession").Changed {
		featureSlice, err := cmd.Flags().GetStringSlice("feature-accession")
		if err != nil {
			return err
		}
		params.SetFeatureAccessionList(featureSlice)
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		p, _ := json.MarshalIndent(params, "", "  ")
		log.Debugf("\n%s", p)
	}
	result, err := Client.Expressions.SearchExpressions(params, AuthInfo)
	if err != nil {
		return err
	}

	return print(result.Payload, len(result.Payload))
}

func setupExpressionCommands() {
	var get = &cobra.Command{
		Use:   "get",
		Short: "Get expression by ID",
		Long:  `Get expression by ID.`,
		Args:  cobra.ExactArgs(1),
		RunE:  expressionGet,
	}
	var search = &cobra.Command{
		Use:   "search",
		Short: "Search expressions",
		Long:  `Search expressions.`,
		// Args:  cobra.MinimumNArgs(1),
		RunE: expressionSearch,
	}
	search.Flags().StringP("version", "v", "", "Search for a specific version")
	search.Flags().StringP("study-id", "s", "", "Search for a specific study id")
	search.Flags().StringP("project-id", "p", "", "Search for a specific project id")
	search.Flags().StringSliceP("tags", "t", nil, "Search for specific tags")
	search.Flags().StringSliceP("sample-id", "i", nil, "Slice by sample id")
	search.Flags().StringSliceP("feature-id", "f", nil, "Slice by feature id")
	search.Flags().StringSliceP("feature-name", "n", nil, "Slice by feature name")
	search.Flags().StringSliceP("feature-accession", "a", nil, "Slice by feature accession")
	expressionCmd.AddCommand(get, search)
}
