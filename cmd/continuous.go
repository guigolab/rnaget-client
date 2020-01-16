package cmd

import (
	"encoding/json"
	"os"

	"github.com/guigolab/rnaget-client/pkg/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	continuousCmd = &cobra.Command{
		Use:   "continuous",
		Short: "Query continuous data",
		Long:  `Query continuous data`,
	}
)

type genomicParams struct {
	chr   *api.ChrParam
	start *api.StartParam
	end   *api.EndParam
}

func init() {
	setupcontinuousCommands()
	rootCmd.AddCommand(continuousCmd)
}

func getContinuousBytesById(id string, params *api.GetContinuousFileByIdParams, w *os.File) error {
	resp, err := Client.GetContinuousFileByIdWithResponse(Ctx, id, params)
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

func getContinuousTicketById(id string, params *api.GetContinuousTicketByIdParams) error {
	resp, err := Client.GetContinuousTicketByIdWithResponse(Ctx, id, params)
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

func getContinuousBytes(cmd *cobra.Command, args []string) error {
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}
	w := getPayloadWriter(output)
	if len(args) > 0 {
		params, err := getContinuousFileByIdParamsFromFlags(cmd)
		if err != nil {
			return err
		}
		return getContinuousBytesById(args[0], params, w)
	}
	params, err := getContinuousFileParamsFromFlags(cmd)
	if err != nil {
		return err
	}
	resp, err := Client.GetContinuousFileWithResponse(Ctx, params)
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case 200:
		payload := resp.Body
		return printBytes(payload, 1, w, resp.HTTPResponse)
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
		e := api.Error{}
		err := json.Unmarshal(resp.Body, &e)
		if err != nil {
			return printBytes(resp.Body, 0, w, resp.HTTPResponse)

		}
		return printJSON(e, 0, resp.HTTPResponse)
	}
}

func getContinuousTicket(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		params, err := getContinuousTicketByIdParamsFromFlags(cmd)
		if err != nil {
			return err
		}
		return getContinuousTicketById(args[0], params)
	}
	params, err := getContinuousTicketParamsFromFlags(cmd)
	if err != nil {
		return err
	}
	resp, err := Client.GetContinuousTicketWithResponse(Ctx, params)
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

func getContinuousFormats(cmd *cobra.Command, args []string) error {
	resp, err := Client.GetContinuousFormatsWithResponse(Ctx)
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

func getContinuousFilters(cmd *cobra.Command, args []string) error {
	resp, err := Client.GetContinuousFiltersWithResponse(Ctx)
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

func getContinuousFirstSupportedFormat() (string, error) {
	var format string
	var err error
	resp, err := Client.GetContinuousFormatsWithResponse(Ctx)
	if err != nil {
		return "", err
	}
	switch resp.StatusCode() {
	case 200:
		format = (*resp.JSON200)[0]
	}
	return format, err
}

func getContinuousFormatFromCommand(cmd *cobra.Command) (string, error) {
	if cmd.Flag("format").Changed {
		formatString, err := cmd.Flags().GetString("format")
		if err != nil {
			return "", err
		}
		return formatString, nil
	} else {
		fmt, err := getContinuousFirstSupportedFormat()
		if err != nil {
			return "", err
		}
		return fmt, nil
	}
}

func getSampleIDListFromCommand(cmd *cobra.Command) (*api.SampleIDListParam, error) {
	if cmd.Flag("sample-id").Changed {
		sampleIDList, err := cmd.Flags().GetStringSlice("sample-id")
		if err != nil {
			return nil, err
		}
		return (*api.SampleIDListParam)(&sampleIDList), nil
	}
	return nil, nil
}

func getGenomicParamsFromCommand(cmd *cobra.Command) (*genomicParams, error) {
	params := genomicParams{}
	if cmd.Flag("chr").Changed {
		chr, err := cmd.Flags().GetString("chr")
		if err != nil {
			return nil, err
		}
		params.chr = (*api.ChrParam)(&chr)
	}
	if cmd.Flag("start").Changed {
		start, err := cmd.Flags().GetInt32("start")
		if err != nil {
			return nil, err
		}
		params.start = (*api.StartParam)(&start)
	}
	if cmd.Flag("end").Changed {
		end, err := cmd.Flags().GetInt32("end")
		if err != nil {
			return nil, err
		}
		params.end = (*api.EndParam)(&end)
	}
	return &params, nil
}

func getContinuousFileByIdParamsFromFlags(cmd *cobra.Command) (*api.GetContinuousFileByIdParams, error) {
	genomicParams, err := getGenomicParamsFromCommand(cmd)
	if err != nil {
		return nil, err
	}
	params := api.GetContinuousFileByIdParams{
		Chr:   genomicParams.chr,
		Start: genomicParams.start,
		End:   genomicParams.end,
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		p, _ := json.MarshalIndent(params, "", "  ")
		log.Debugf("\n%s", p)
	}

	return &params, nil
}

func getContinuousFileParamsFromFlags(cmd *cobra.Command) (*api.GetContinuousFileParams, error) {
	fmt, err := getContinuousFormatFromCommand(cmd)
	if err != nil {
		return nil, err
	}
	sampleIDList, err := getSampleIDListFromCommand(cmd)
	if err != nil {
		return nil, err
	}
	filterParams, err := getFilterParamsFromCommand(cmd)
	if err != nil {
		return nil, err
	}
	genomicParams, err := getGenomicParamsFromCommand(cmd)
	if err != nil {
		return nil, err
	}
	params := api.GetContinuousFileParams{
		Format:       fmt,
		ProjectID:    filterParams.projectID,
		StudyID:      filterParams.studyID,
		Version:      filterParams.version,
		SampleIDList: sampleIDList,
		Chr:          genomicParams.chr,
		Start:        genomicParams.start,
		End:          genomicParams.end,
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		p, _ := json.MarshalIndent(params, "", "  ")
		log.Debugf("\n%s", p)
	}

	return &params, nil
}

func getContinuousTicketByIdParamsFromFlags(cmd *cobra.Command) (*api.GetContinuousTicketByIdParams, error) {
	genomicParams, err := getGenomicParamsFromCommand(cmd)
	if err != nil {
		return nil, err
	}
	params := api.GetContinuousTicketByIdParams{
		Chr:   genomicParams.chr,
		Start: genomicParams.start,
		End:   genomicParams.end,
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		p, _ := json.MarshalIndent(params, "", "  ")
		log.Debugf("\n%s", p)
	}

	return &params, nil
}

func getContinuousTicketParamsFromFlags(cmd *cobra.Command) (*api.GetContinuousTicketParams, error) {
	fmt, err := getContinuousFormatFromCommand(cmd)
	if err != nil {
		return nil, err
	}
	sampleIDList, err := getSampleIDListFromCommand(cmd)
	if err != nil {
		return nil, err
	}
	filterParams, err := getFilterParamsFromCommand(cmd)
	if err != nil {
		return nil, err
	}
	genomicParams, err := getGenomicParamsFromCommand(cmd)
	if err != nil {
		return nil, err
	}
	params := api.GetContinuousTicketParams{
		Format:       fmt,
		ProjectID:    filterParams.projectID,
		StudyID:      filterParams.studyID,
		Version:      filterParams.version,
		SampleIDList: sampleIDList,
		Chr:          genomicParams.chr,
		Start:        genomicParams.start,
		End:          genomicParams.end,
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		p, _ := json.MarshalIndent(params, "", "  ")
		log.Debugf("\n%s", p)
	}

	return &params, nil
}

func addContinuousFlags(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		cmd.Flags().StringP("format", "f", "", "Data format to return")
		cmd.Flags().StringP("version", "v", "", "Search for a specific version (ignored when [id] is specified)")
		cmd.Flags().StringP("study-id", "s", "", "Search for a specific study id (ignored when [id] is specified)")
		cmd.Flags().StringP("project-id", "p", "", "Search for a specific project id (ignored when [id] is specified)")
		cmd.Flags().StringSliceP("sample-id", "i", nil, "Slice by sample id")
		cmd.Flags().StringP("chr", "c", "", "The refererence to which start and end apply")
		cmd.Flags().Int32("start", -1, "The start position of the range on the sequence, 0-based, inclusive.")
		cmd.Flags().Int32("end", -1, "The end position of the range on the sequence, 0-based, exclusive.")
	}
}

func setupcontinuousCommands() {
	var bytesCmd = &cobra.Command{
		Use:   "bytes [id]",
		Short: "Get raw continuous data",
		Long:  `Get raw continuous data.`,
		Args:  cobra.MaximumNArgs(1),
		RunE:  getContinuousBytes,
	}
	var ticketCmd = &cobra.Command{
		Use:   "ticket [id]",
		Short: "Get continuous ticket",
		Long:  `Get continuous ticket.`,
		Args:  cobra.MaximumNArgs(1),
		RunE:  getContinuousTicket,
	}
	var formatsCmd = &cobra.Command{
		Use:   "formats",
		Short: "Get continuous formats",
		Long:  `Get continuous formats.`,
		RunE:  getContinuousFormats,
	}
	var filtersCmd = &cobra.Command{
		Use:   "filters",
		Short: "Get continuous filters",
		Long:  `Get continuous filters.`,
		RunE:  getContinuousFilters,
	}
	addContinuousFlags(bytesCmd, ticketCmd)
	bytesCmd.Flags().StringP("output", "o", "stdout", "Output file")
	continuousCmd.AddCommand(bytesCmd, ticketCmd, formatsCmd, filtersCmd)
}
