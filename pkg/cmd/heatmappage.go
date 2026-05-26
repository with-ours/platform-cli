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

var heatmapPagesList = cli.Command{
	Name:    "list",
	Usage:   "List pages with heatmap coverage in a date window, ranked for triage. Each\nentity is identified by `pageKey` (origin + pathname, query string stripped);\nuse that value to drill into `GET /rest/v1/heatmap-pages/summary`. Supports\ncursor pagination, with cursor depth capped at roughly 10,000 entries; if you\nneed pages beyond that, narrow `from`/`to` or add filters rather than paginating\nfurther. `from`/`to` are UTC calendar days in `YYYY-MM-DD`; the window must be\n60 days or fewer. Requires scope: web-analytics:view",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "from",
			Usage:     "Inclusive lower bound of the heatmap window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or fewer.",
			Required:  true,
			QueryPath: "from",
		},
		&requestflag.Flag[string]{
			Name:      "to",
			Usage:     "Inclusive upper bound of the heatmap window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or fewer.",
			Required:  true,
			QueryPath: "to",
		},
		&requestflag.Flag[string]{
			Name:      "browser-name",
			Usage:     "Filter by browser name as captured on events.",
			QueryPath: "browserName",
		},
		&requestflag.Flag[string]{
			Name:      "country",
			Usage:     "Filter by visitor country (ISO country name or code as stored on events).",
			QueryPath: "country",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor from pagination.nextCursor in the previous response. Do not decode or modify it. Malformed cursors return 400 Bad Request.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[*int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return. Defaults to 25; values below 1 are clamped to 1 and values above 100 are clamped to 100.",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "region",
			Usage:     "Filter by visitor region (state/province as stored on events).",
			QueryPath: "region",
		},
		&requestflag.Flag[string]{
			Name:      "search",
			Usage:     "Case-insensitive substring match against `pageKey`.",
			QueryPath: "search",
		},
		&requestflag.Flag[string]{
			Name:      "sort-by",
			Usage:     "Sort key. Defaults to `CLICKS` (descending).",
			QueryPath: "sortBy",
		},
		&requestflag.Flag[string]{
			Name:      "sort-dir",
			Usage:     "Sort direction. Defaults to `DESC`.",
			QueryPath: "sortDir",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleHeatmapPagesList,
	HideHelpCommand: true,
}

var heatmapPagesSummary = cli.Command{
	Name:    "summary",
	Usage:   "Bundled per-page heatmap rollup: click bins, dead clicks, rage clicks, scroll\ndepth, and up to 5 curated in-app replay URLs in a single payload. Designed as a\none-call diagnosis surface for an AI assistant or marketer — answers \"what is\nhappening on this page?\" without fanning out across five endpoints. `pageKey`\ncomes from `GET /rest/v1/heatmap-pages`. The endpoint is identified by query\nparams rather than a path id because heatmap pages are not stored entities; this\nis a documented derived-read exception. `breakpoint` scopes the bin/scroll\naggregations; replays are returned across all breakpoints regardless of\n`breakpoint` (weighted to cover multiple viewports) so callers can compare\ndevices. Requires scope: web-analytics:view",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "breakpoint",
			Usage:     "Viewport bucket the click, dead-click, rage, and scroll-depth aggregations are computed for. Replays are returned for all breakpoints regardless of this value so callers can compare across devices.",
			Required:  true,
			QueryPath: "breakpoint",
		},
		&requestflag.Flag[string]{
			Name:      "from",
			Usage:     "Inclusive lower bound of the heatmap window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or fewer.",
			Required:  true,
			QueryPath: "from",
		},
		&requestflag.Flag[string]{
			Name:      "page-key",
			Usage:     "Page identifier returned by `GET /rest/v1/heatmap-pages`. Origin + pathname with the query string stripped (e.g. `https://example.com/pricing`).",
			Required:  true,
			QueryPath: "pageKey",
		},
		&requestflag.Flag[string]{
			Name:      "to",
			Usage:     "Inclusive upper bound of the heatmap window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or fewer.",
			Required:  true,
			QueryPath: "to",
		},
		&requestflag.Flag[string]{
			Name:      "browser-name",
			Usage:     "Filter by browser name as captured on events.",
			QueryPath: "browserName",
		},
		&requestflag.Flag[string]{
			Name:      "country",
			Usage:     "Filter by visitor country (ISO country name or code as stored on events).",
			QueryPath: "country",
		},
		&requestflag.Flag[string]{
			Name:      "region",
			Usage:     "Filter by visitor region (state/province as stored on events).",
			QueryPath: "region",
		},
	},
	Action:          handleHeatmapPagesSummary,
	HideHelpCommand: true,
}

func handleHeatmapPagesList(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.HeatmapPageListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.HeatmapPages.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "heatmap-pages list",
			Transform:      transform,
		})
	} else {
		iter := client.HeatmapPages.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "heatmap-pages list",
			Transform:      transform,
		})
	}
}

func handleHeatmapPagesSummary(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.HeatmapPageSummaryParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.HeatmapPages.Summary(ctx, params, options...)
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
		Title:          "heatmap-pages summary",
		Transform:      transform,
	})
}
