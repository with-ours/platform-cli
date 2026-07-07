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

var attributionInitial = cli.Command{
	Name:    "initial",
	Usage:   "Returns the top-15 values for each UTM dimension (source, medium, campaign,\ncontent, term, name) and referring domain attributed to the conversion event on\na first-touch basis for the given date window. Use `from`/`to` to set the\nanalysis window (max 60 days). Optionally filter to a specific UTM combo with\n`utmSource`, `utmMedium`, etc. The counts represent unique visitors who\nperformed the specified `eventName` and were attributed to each UTM value.\nRequires scope: web-analytics:view",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "event-name",
			Usage:     "Conversion event to count. Must be a selectable conversion event.",
			Required:  true,
			QueryPath: "eventName",
		},
		&requestflag.Flag[string]{
			Name:      "from",
			Usage:     "Inclusive lower bound of the analytics window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or fewer.",
			Required:  true,
			QueryPath: "from",
		},
		&requestflag.Flag[string]{
			Name:      "to",
			Usage:     "Inclusive upper bound of the analytics window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or fewer.",
			Required:  true,
			QueryPath: "to",
		},
		&requestflag.Flag[string]{
			Name:      "attribution-type",
			Usage:     "Attribution type for UTM filter matching. Defaults to `INITIAL`.",
			QueryPath: "attributionType",
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
	Action:          handleAttributionInitial,
	HideHelpCommand: true,
}

var attributionLastTouch = cli.Command{
	Name:    "last-touch",
	Usage:   "Returns the top-15 values for each UTM dimension (source, medium, campaign,\ncontent, term, name) and referring domain attributed to the conversion event on\na last-touch basis for the given date window. Use `from`/`to` to set the\nanalysis window (max 60 days). The counts represent unique visitors who\nperformed the specified `eventName` and were attributed to each UTM value on\ntheir most recent session. Requires scope: web-analytics:view",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "event-name",
			Usage:     "Conversion event to count. Must be a selectable conversion event.",
			Required:  true,
			QueryPath: "eventName",
		},
		&requestflag.Flag[string]{
			Name:      "from",
			Usage:     "Inclusive lower bound of the analytics window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or fewer.",
			Required:  true,
			QueryPath: "from",
		},
		&requestflag.Flag[string]{
			Name:      "to",
			Usage:     "Inclusive upper bound of the analytics window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or fewer.",
			Required:  true,
			QueryPath: "to",
		},
		&requestflag.Flag[string]{
			Name:      "attribution-type",
			Usage:     "Attribution type for UTM filter matching. Defaults to `LAST_TOUCH`.",
			QueryPath: "attributionType",
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
	Action:          handleAttributionLastTouch,
	HideHelpCommand: true,
}

var attributionConversion = cli.Command{
	Name:    "conversion",
	Usage:   "Multi-touch conversion attribution: returns a source → medium → campaign\nhierarchy with attributed converter credits distributed according to the\nselected `attributionModel`. Scoped to all web sources by default; optionally\nnarrow to a single web source via `webSourceId`. Date range is capped at 31\ndays; lookback window is capped at 60 days. Requires scope: web-analytics:view",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "attribution-model",
			Usage:     "Attribution model to apply to multi-touch conversion paths.",
			Required:  true,
			QueryPath: "attributionModel",
		},
		&requestflag.Flag[string]{
			Name:      "event-name",
			Usage:     "Conversion event to attribute. Must be a selectable conversion event.",
			Required:  true,
			QueryPath: "eventName",
		},
		&requestflag.Flag[string]{
			Name:      "from",
			Usage:     "Inclusive lower bound of the analytics window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 31 days or fewer.",
			Required:  true,
			QueryPath: "from",
		},
		&requestflag.Flag[string]{
			Name:      "to",
			Usage:     "Inclusive upper bound of the analytics window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 31 days or fewer.",
			Required:  true,
			QueryPath: "to",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of leaf-level attribution rows to return. Defaults to 1000.",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "lookback-window",
			Usage:     "How far back before each conversion to consider touchpoints. Capped at 60 days for this report. Defaults to `THIRTY_DAYS`.",
			QueryPath: "lookbackWindow",
		},
		&requestflag.Flag[string]{
			Name:      "web-source-id",
			Usage:     "Scope to a single web source by id, or omit for all sources (account-wide).",
			QueryPath: "webSourceId",
		},
	},
	Action:          handleAttributionConversion,
	HideHelpCommand: true,
}

var attributionAudienceConversion = cli.Command{
	Name:    "audience-conversion",
	Usage:   "Audience performance conversion report: returns a summary of converters and\nconversion rate for the selected event and date window, a per-day timeseries,\nand a UTM source/medium/campaign breakdown. Optionally compare against the\npreceding period of equal length when `attributionWindow` is `IN_RANGE`. Date\nrange is capped at 60 days. Requires scope: web-analytics:view",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "event-name",
			Usage:     "Conversion event to analyze.",
			Required:  true,
			QueryPath: "eventName",
		},
		&requestflag.Flag[string]{
			Name:      "from",
			Usage:     "Inclusive lower bound of the analytics window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or fewer.",
			Required:  true,
			QueryPath: "from",
		},
		&requestflag.Flag[string]{
			Name:      "to",
			Usage:     "Inclusive upper bound of the analytics window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or fewer.",
			Required:  true,
			QueryPath: "to",
		},
		&requestflag.Flag[string]{
			Name:      "attribution-window",
			Usage:     "Attribution window: `IN_RANGE` or a number of lookback days (e.g. `7`, `30`). Defaults to `IN_RANGE`.",
			QueryPath: "attributionWindow",
		},
		&requestflag.Flag[string]{
			Name:      "value-property",
			Usage:     "Event property to sum as conversion value.",
			QueryPath: "valueProperty",
		},
	},
	Action:          handleAttributionAudienceConversion,
	HideHelpCommand: true,
}

var attributionUtmComparison = cli.Command{
	Name:    "utm-comparison",
	Usage:   "Compare up to 5 UTM dimension combinations side-by-side for a single conversion\nevent. Each combo returns the unique visitors, sessions, total events, and\nderived conversion rate for that UTM filter within the window. Requires both\n`web-analytics:view` and `report:event-count-by-day` API-key scopes. Date range\nis capped at 31 days. Pass `combos` as a single JSON-encoded array:\n`combos=[{\"utmSource\":\"google\",\"utmMedium\":\"cpc\"},{\"utmSource\":\"meta\"}]`.\nRequires scope: web-analytics:view",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "combos",
			Usage:     "JSON-encoded array of UTM dimension combos to compare side-by-side (min 1, max 5). Each combo is an object with optional `utmSource`, `utmMedium`, `utmCampaign`, `utmContent`, `utmTerm` fields.",
			Required:  true,
			QueryPath: "combos",
		},
		&requestflag.Flag[string]{
			Name:      "event-name",
			Usage:     "Conversion event to compare across UTM combos.",
			Required:  true,
			QueryPath: "eventName",
		},
		&requestflag.Flag[string]{
			Name:      "from",
			Usage:     "Inclusive lower bound of the analytics window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 31 days or fewer.",
			Required:  true,
			QueryPath: "from",
		},
		&requestflag.Flag[string]{
			Name:      "to",
			Usage:     "Inclusive upper bound of the analytics window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 31 days or fewer.",
			Required:  true,
			QueryPath: "to",
		},
	},
	Action:          handleAttributionUtmComparison,
	HideHelpCommand: true,
}

func handleAttributionInitial(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.AttributionInitialParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Attribution.Initial(ctx, params, options...)
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
		Title:          "attribution initial",
		Transform:      transform,
	})
}

func handleAttributionLastTouch(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.AttributionLastTouchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Attribution.LastTouch(ctx, params, options...)
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
		Title:          "attribution last-touch",
		Transform:      transform,
	})
}

func handleAttributionConversion(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.AttributionConversionParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Attribution.Conversion(ctx, params, options...)
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
		Title:          "attribution conversion",
		Transform:      transform,
	})
}

func handleAttributionAudienceConversion(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.AttributionAudienceConversionParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Attribution.AudienceConversion(ctx, params, options...)
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
		Title:          "attribution audience-conversion",
		Transform:      transform,
	})
}

func handleAttributionUtmComparison(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.AttributionUtmComparisonParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Attribution.UtmComparison(ctx, params, options...)
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
		Title:          "attribution utm-comparison",
		Transform:      transform,
	})
}
