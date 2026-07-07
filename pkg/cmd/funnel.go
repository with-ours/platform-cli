// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
	"github.com/with-ours/platform-cli/internal/apiquery"
	"github.com/with-ours/platform-cli/internal/requestflag"
	"github.com/with-ours/platform-sdk-go"
	"github.com/with-ours/platform-sdk-go/option"
)

var funnelsList = cli.Command{
	Name:            "list",
	Usage:           "List every funnel configured on this account. Each funnel includes its step\nconfiguration, funnel type, conversion window, and current processing status.\nThe available report date range (if any pre-computed reports exist) is returned\nin `reportDateRange`. Requires scope: web-analytics:view",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleFunnelsList,
	HideHelpCommand: true,
}

var funnelsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch a single funnel configuration by its id. Returns `404` when the funnel\ndoes not exist or belongs to a different account. Requires scope:\nweb-analytics:view",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleFunnelsRetrieve,
	HideHelpCommand: true,
}

var funnelsResults = cli.Command{
	Name:    "results",
	Usage:   "Compute funnel step analytics for a funnel over a date window. Returns per-step\nvisitor counts, conversion rates, drop-off rates, average time to next step, and\nsample session IDs for replay. Funnel results are pre-computed daily from S3;\nreports outside the `reportDateRange` shown on the funnel config will return\nempty steps. Requires scope: web-analytics:view",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "from",
			Usage:     "Inclusive lower bound of the analysis window, as a UTC calendar day in `YYYY-MM-DD` format.",
			Required:  true,
			QueryPath: "from",
		},
		&requestflag.Flag[string]{
			Name:      "to",
			Usage:     "Inclusive upper bound of the analysis window, as a UTC calendar day in `YYYY-MM-DD` format.",
			Required:  true,
			QueryPath: "to",
		},
		&requestflag.Flag[string]{
			Name:      "attribution-type",
			Usage:     "Attribution type for UTM filter matching in funnel steps.",
			QueryPath: "attributionType",
		},
		&requestflag.Flag[string]{
			Name:      "device-type",
			Usage:     "Filter funnel analytics to a specific device type. Defaults to `ALL`.",
			Default:   "ALL",
			QueryPath: "deviceType",
		},
		&requestflag.Flag[string]{
			Name:      "utm-campaign",
			Usage:     "Filter by UTM campaign.",
			QueryPath: "utmCampaign",
		},
		&requestflag.Flag[string]{
			Name:      "utm-content",
			Usage:     "Filter by UTM content.",
			QueryPath: "utmContent",
		},
		&requestflag.Flag[string]{
			Name:      "utm-medium",
			Usage:     "Filter by UTM medium.",
			QueryPath: "utmMedium",
		},
		&requestflag.Flag[string]{
			Name:      "utm-name",
			Usage:     "Filter by UTM name.",
			QueryPath: "utmName",
		},
		&requestflag.Flag[string]{
			Name:      "utm-source",
			Usage:     "Filter by UTM source.",
			QueryPath: "utmSource",
		},
		&requestflag.Flag[string]{
			Name:      "utm-term",
			Usage:     "Filter by UTM term.",
			QueryPath: "utmTerm",
		},
	},
	Action:          handleFunnelsResults,
	HideHelpCommand: true,
}

func handleFunnelsList(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacy.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Funnels.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "funnels list",
		Transform:      transform,
	})
}

func handleFunnelsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacy.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Funnels.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "funnels retrieve",
		Transform:      transform,
	})
}

func handleFunnelsResults(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacy.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := oursprivacy.FunnelResultsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Funnels.Results(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "funnels results",
		Transform:      transform,
	})
}
