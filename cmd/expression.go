package cmd

import (
	"encoding/json"
	"os"

	"github.com/guigolab/rnaget-client/pkg/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	expressionCmd = &cobra.Command{
		Use:   "expressions",
		Short: "Query expression data",
		Long:  `Query expression data`,
	}
)

type slicingParams struct {
	sampleIDList    *api.SampleIDListParam
	featureIDList   *api.FeatureIDListParam
	featureNameList *api.FeatureNameListParam
}

type filterParams struct {
	projectID *api.ProjectIDParam
	studyID   *api.StudyIDParam
	version   *api.VersionParam
}

func init() {
	setupExpressionCommands()
	rootCmd.AddCommand(expressionCmd)
}

func getExpressionBytesById(id string, params *api.GetExpressionFileByIdParams, w *os.File) error {
	resp, err := Client.GetExpressionFileByIdWithResponse(Ctx, id, params)
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case 200:
		payload := resp.Body
		printBytes(payload, 1, w, resp.HTTPResponse)
	case 400:
		payload := resp.JSON400
		printJSON(payload, 1, resp.HTTPResponse)
	case 404:
		payload := resp.JSON404
		printJSON(payload, 1, resp.HTTPResponse)
	case 406:
		payload := resp.JSON406
		printJSON(payload, 1, resp.HTTPResponse)
	case 501:
		payload := resp.JSON501
		printJSON(payload, 1, resp.HTTPResponse)
	}

	return nil
}

func getExpressionTicketById(id string, params *api.GetExpressionTicketByIdParams) error {
	resp, err := Client.GetExpressionTicketByIdWithResponse(Ctx, id, params)
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case 200:
		payload := resp.JSON200
		printJSON(payload, 1, resp.HTTPResponse)
	case 400:
		payload := resp.JSON400
		printJSON(payload, 1, resp.HTTPResponse)
	case 404:
		payload := resp.JSON404
		printJSON(payload, 1, resp.HTTPResponse)
	case 406:
		payload := resp.JSON406
		printJSON(payload, 1, resp.HTTPResponse)
	case 501:
		payload := resp.JSON501
		printJSON(payload, 1, resp.HTTPResponse)
	}

	return nil
}

func getExpressionBytes(cmd *cobra.Command, args []string) error {
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}
	w := getPayloadWriter(output)
	if len(args) > 0 {
		params, err := getExpressionFileByIdParamsFromFlags(cmd)
		if err != nil {
			return err
		}
		return getExpressionBytesById(args[0], params, w)
	}
	params, err := getExpressionFileParamsFromFlags(cmd)
	if err != nil {
		return err
	}
	resp, err := Client.GetExpressionFileWithResponse(Ctx, params)
	if err != nil {
		return err
	}
	switch resp.StatusCode() {
	case 200:
		payload := resp.Body
		printBytes(payload, 1, w, resp.HTTPResponse)
	case 400:
		payload := resp.JSON400
		printJSON(payload, 1, resp.HTTPResponse)
	case 406:
		payload := resp.JSON406
		printJSON(payload, 1, resp.HTTPResponse)
	case 501:
		payload := resp.JSON501
		printJSON(payload, 1, resp.HTTPResponse)
	}

	return nil
}

func getExpressionTicket(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		params, err := getExpressionTicketByIdParamsFromFlags(cmd)
		if err != nil {
			return err
		}
		return getExpressionTicketById(args[0], params)
	}
	params, err := getExpressionTicketParamsFromFlags(cmd)
	if err != nil {
		return err
	}
	resp, err := Client.GetExpressionTicketWithResponse(Ctx, params)
	if err != nil {
		return err
	}
	switch resp.StatusCode() {
	case 200:
		payload := resp.JSON200
		printJSON(payload, 1, resp.HTTPResponse)
	case 400:
		payload := resp.JSON400
		printJSON(payload, 1, resp.HTTPResponse)
	case 406:
		payload := resp.JSON406
		printJSON(payload, 1, resp.HTTPResponse)
	case 501:
		payload := resp.JSON501
		printJSON(payload, 1, resp.HTTPResponse)
	}

	return nil
}

func getExpressionFormats(cmd *cobra.Command, args []string) error {
	resp, err := Client.GetExpressionFormatsWithResponse(Ctx)
	if err != nil {
		return err
	}
	switch resp.StatusCode() {
	case 200:
		payload := resp.JSON200
		printJSON(payload, 1, resp.HTTPResponse)
	case 404:
		payload := resp.JSON404
		printJSON(payload, 1, resp.HTTPResponse)
	case 406:
		payload := resp.JSON406
		printJSON(payload, 1, resp.HTTPResponse)
	case 501:
		payload := resp.JSON501
		printJSON(payload, 1, resp.HTTPResponse)
	}

	return nil
}

func getExpressionFilters(cmd *cobra.Command, args []string) error {
	params, err := getExpressionFiltersParamsFromFlags(cmd)
	if err != nil {
		return err
	}
	resp, err := Client.GetExpressionFiltersWithResponse(Ctx, params)
	if err != nil {
		return err
	}
	switch resp.StatusCode() {
	case 200:
		payload := resp.JSON200
		printJSON(payload, 1, resp.HTTPResponse)
	case 400:
		payload := resp.JSON400
		printJSON(payload, 1, resp.HTTPResponse)
	case 406:
		payload := resp.JSON406
		printJSON(payload, 1, resp.HTTPResponse)
	case 501:
		payload := resp.JSON501
		printJSON(payload, 1, resp.HTTPResponse)
	}

	return nil
}

func getPayloadWriter(output string) *os.File {
	switch output {
	case "-", "stdout":
		return os.Stdout
	default:
		f, err := os.Create(output)
		if err != nil {
			log.Fatal(err)
		}
		return f
	}
}

func getFirstSupportedFormat() (string, error) {
	var format string
	resp, err := Client.GetExpressionFormatsWithResponse(Ctx)
	if err != nil {
		return "", err
	}
	if resp.JSON200 != nil {
		format = (*resp.JSON200)[0]
	}
	return format, nil
}

func getFormatFromCommand(cmd *cobra.Command) (string, error) {
	if cmd.Flag("format").Changed {
		formatString, err := cmd.Flags().GetString("format")
		if err != nil {
			return "", err
		}
		return formatString, nil
	} else {
		fmt, err := getFirstSupportedFormat()
		if err != nil {
			return "", err
		}
		return fmt, nil
	}
}

func getFilterParamsFromCommand(cmd *cobra.Command) (*filterParams, error) {
	params := filterParams{}
	if cmd.Flag("project-id").Changed {
		projetctID, err := cmd.Flags().GetString("project-id")
		if err != nil {
			return nil, err
		}
		params.projectID = (*api.ProjectIDParam)(&projetctID)
	}
	if cmd.Flag("study-id").Changed {
		studyID, err := cmd.Flags().GetString("study-id")
		if err != nil {
			return nil, err
		}
		params.studyID = (*api.StudyIDParam)(&studyID)
	}
	if cmd.Flag("version").Changed {
		version, err := cmd.Flags().GetString("version")
		if err != nil {
			return nil, err
		}
		params.version = (*api.VersionParam)(&version)
	}

	return &params, nil
}

func getSlicingParamsFromCommand(cmd *cobra.Command) (*slicingParams, error) {
	params := slicingParams{}
	if cmd.Flag("sample-id").Changed {
		sampleIDList, err := cmd.Flags().GetStringSlice("sample-id")
		if err != nil {
			return nil, err
		}
		params.sampleIDList = (*api.SampleIDListParam)(&sampleIDList)
	}
	if cmd.Flag("feature-id").Changed {
		featureIdList, err := cmd.Flags().GetStringSlice("feature-id")
		if err != nil {
			return nil, err
		}
		params.featureIDList = (*api.FeatureIDListParam)(&featureIdList)
	}
	if cmd.Flag("feature-name").Changed {
		featureNameList, err := cmd.Flags().GetStringSlice("feature-name")
		if err != nil {
			return nil, err
		}
		params.featureNameList = (*api.FeatureNameListParam)(&featureNameList)
	}
	return &params, nil
}

func getExpressionFileByIdParamsFromFlags(cmd *cobra.Command) (*api.GetExpressionFileByIdParams, error) {
	slicingParams, err := getSlicingParamsFromCommand(cmd)
	if err != nil {
		return nil, err
	}
	params := api.GetExpressionFileByIdParams{
		SampleIDList:    slicingParams.sampleIDList,
		FeatureIDList:   slicingParams.featureIDList,
		FeatureNameList: slicingParams.featureNameList,
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		p, _ := json.MarshalIndent(params, "", "  ")
		log.Debugf("\n%s", p)
	}

	return &params, nil
}

func getExpressionFileParamsFromFlags(cmd *cobra.Command) (*api.GetExpressionFileParams, error) {
	fmt, err := getFormatFromCommand(cmd)
	if err != nil {
		return nil, err
	}
	filterParams, err := getFilterParamsFromCommand(cmd)
	if err != nil {
		return nil, err
	}
	slicingParams, err := getSlicingParamsFromCommand(cmd)
	if err != nil {
		return nil, err
	}
	params := api.GetExpressionFileParams{
		Format:          fmt,
		ProjectID:       filterParams.projectID,
		StudyID:         filterParams.studyID,
		Version:         filterParams.version,
		SampleIDList:    slicingParams.sampleIDList,
		FeatureIDList:   slicingParams.featureIDList,
		FeatureNameList: slicingParams.featureNameList,
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		p, _ := json.MarshalIndent(params, "", "  ")
		log.Debugf("\n%s", p)
	}

	return &params, nil
}

func getExpressionTicketByIdParamsFromFlags(cmd *cobra.Command) (*api.GetExpressionTicketByIdParams, error) {
	slicingParams, err := getSlicingParamsFromCommand(cmd)
	if err != nil {
		return nil, err
	}
	params := api.GetExpressionTicketByIdParams{
		SampleIDList:    slicingParams.sampleIDList,
		FeatureIDList:   slicingParams.featureIDList,
		FeatureNameList: slicingParams.featureNameList,
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		p, _ := json.MarshalIndent(params, "", "  ")
		log.Debugf("\n%s", p)
	}

	return &params, nil
}

func getExpressionTicketParamsFromFlags(cmd *cobra.Command) (*api.GetExpressionTicketParams, error) {
	fmt, err := getFormatFromCommand(cmd)
	if err != nil {
		return nil, err
	}
	filterParams, err := getFilterParamsFromCommand(cmd)
	if err != nil {
		return nil, err
	}
	slicingParams, err := getSlicingParamsFromCommand(cmd)
	if err != nil {
		return nil, err
	}
	params := api.GetExpressionTicketParams{
		Format:          fmt,
		ProjectID:       filterParams.projectID,
		StudyID:         filterParams.studyID,
		Version:         filterParams.version,
		SampleIDList:    slicingParams.sampleIDList,
		FeatureIDList:   slicingParams.featureIDList,
		FeatureNameList: slicingParams.featureNameList,
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		p, _ := json.MarshalIndent(params, "", "  ")
		log.Debugf("\n%s", p)
	}

	return &params, nil
}

func getExpressionFiltersParamsFromFlags(cmd *cobra.Command) (*api.GetExpressionFiltersParams, error) {
	params := api.GetExpressionFiltersParams{}
	if cmd.Flag("type").Changed {
		typeString, err := cmd.Flags().GetString("type")
		if err != nil {
			return nil, err
		}
		params.Type = &typeString
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		p, _ := json.MarshalIndent(params, "", "  ")
		log.Debugf("\n%s", p)
	}
	return &params, nil
}

func addCommonFlags(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		cmd.Flags().StringP("format", "f", "", "Data format to return")
		cmd.Flags().StringP("version", "v", "", "Search for a specific version (ignored when [id] is specified)")
		cmd.Flags().StringP("study-id", "s", "", "Search for a specific study id (ignored when [id] is specified)")
		cmd.Flags().StringP("project-id", "p", "", "Search for a specific project id (ignored when [id] is specified)")
		cmd.Flags().StringSliceP("sample-id", "i", nil, "Slice by sample id")
		cmd.Flags().StringSliceP("feature-id", "t", nil, "Slice by feature id")
		cmd.Flags().StringSliceP("feature-name", "n", nil, "Slice by feature name")
	}
}

func setupExpressionCommands() {
	var bytesCmd = &cobra.Command{
		Use:   "bytes [id]",
		Short: "Get raw expression data",
		Long:  `Get raw expression data.`,
		Args:  cobra.MaximumNArgs(1),
		RunE:  getExpressionBytes,
	}
	var ticketCmd = &cobra.Command{
		Use:   "ticket [id]",
		Short: "Get expression ticket",
		Long:  `Get expression ticket.`,
		Args:  cobra.MaximumNArgs(1),
		RunE:  getExpressionTicket,
	}
	var formatsCmd = &cobra.Command{
		Use:   "formats",
		Short: "Get expression formats",
		Long:  `Get expression formats.`,
		RunE:  getExpressionFormats,
	}
	var filtersCmd = &cobra.Command{
		Use:   "filters",
		Short: "Get expression filters",
		Long:  `Get expression filters.`,
		RunE:  getExpressionFilters,
	}
	addCommonFlags(bytesCmd, ticketCmd)
	bytesCmd.Flags().StringP("output", "o", "stdout", "Output file")
	filtersCmd.Flags().StringP("type", "t", "", "Filter type")
	expressionCmd.AddCommand(bytesCmd, ticketCmd, formatsCmd, filtersCmd)
}
