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

var shortLinksList = cli.Command{
	Name:    "list",
	Usage:   "List all short links (QR codes / redirects) for this account, newest first.\nSupports cursor pagination and optional `status` and `nameContains` filters.\nEach entity bundles the destination URL, the composed public `shortUrl`, and the\nQR/campaign design. Requires scope: source:list",
	Suggest: true,
	Flags: []cli.Flag{
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
			Name:      "name-contains",
			Usage:     "Case-insensitive substring filter on the short link name.",
			QueryPath: "nameContains",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by short link status.",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleShortLinksList,
	HideHelpCommand: true,
}

var shortLinksCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a short link (QR code / redirect) with its destination, campaign tags,\nand QR styling in a single call. The short code is generated automatically; the\nresponse `shortUrl` is the public URL the QR encodes. All body fields are\noptional — send `{}` to create an unconfigured link and fill it in later with\nPATCH. A newly created short link only resolves at the edge once a version is\npublished. Requires scope: source:create",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "Human-readable name. Also sent as the tracked event name on every click/scan.",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "qr",
			Usage:    "QR code visual styling.",
			BodyPath: "qr",
		},
		&requestflag.Flag[*string]{
			Name:     "redirect-url",
			Usage:    "Destination URL the short link redirects to. Must be a valid URL.",
			BodyPath: "redirectUrl",
		},
		&requestflag.Flag[any]{
			Name:     "utm",
			Usage:    "Campaign / UTM tags appended to the tracked short-link URL.",
			BodyPath: "utm",
		},
	},
	Action:          handleShortLinksCreate,
	HideHelpCommand: true,
}

var shortLinksRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch a single short link by id, including its destination, composed `shortUrl`,\nand QR/campaign design. Returns 404 when no short link matches the id or it\nbelongs to a different account. Requires scope: source:view",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleShortLinksRetrieve,
	HideHelpCommand: true,
}

var shortLinksUpdate = cli.Command{
	Name:    "update",
	Usage:   "Partially update a short link. Only the fields you send are changed; omitted\nfields are unchanged. Send explicit `null` to clear `redirectUrl`. The `utm` and\n`qr` objects are replaced wholesale when sent. Returns the full short link\nentity after the update. Requires scope: source:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "qr",
			BodyPath: "qr",
		},
		&requestflag.Flag[*string]{
			Name:     "redirect-url",
			Usage:    "Destination URL the short link redirects to. Must be a valid URL. Send `null` to clear it.",
			BodyPath: "redirectUrl",
		},
		&requestflag.Flag[*string]{
			Name:     "status",
			Usage:    "Whether the short link resolves at the edge. Send `Enabled` or `Disabled`; `null` is rejected since storage cannot represent it.",
			BodyPath: "status",
		},
		&requestflag.Flag[any]{
			Name:     "utm",
			BodyPath: "utm",
		},
	},
	Action:          handleShortLinksUpdate,
	HideHelpCommand: true,
}

var shortLinksDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a short link and its QR/campaign design. After deletion the short URL\nstops resolving on the next publish. Requires scope: source:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleShortLinksDelete,
	HideHelpCommand: true,
}

var shortLinksResults = cli.Command{
	Name:    "results",
	Usage:   "Aggregate click analytics for a short link over a date window: total and unique\nclicks, a time series (daily or hourly), and breakdowns by country, city, and\ndevice. QR scans are counted as clicks. Pass `from`/`to` as UTC calendar days\n(`YYYY-MM-DD`); set `granularity=HOURLY` for hourly buckets and\n`excludeBots=false` to include bot traffic. Requires the `shortlink:reporting`\nscope, which is gated separately because analytics data is PHI-bearing. Requires\nscope: shortlink:reporting",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "from",
			Usage:     "Inclusive lower bound of the report window, a UTC calendar day in `YYYY-MM-DD`.",
			Required:  true,
			QueryPath: "from",
		},
		&requestflag.Flag[string]{
			Name:      "to",
			Usage:     "Inclusive upper bound of the report window, a UTC calendar day in `YYYY-MM-DD`.",
			Required:  true,
			QueryPath: "to",
		},
		&requestflag.Flag[bool]{
			Name:      "exclude-bots",
			Usage:     "Exclude bot traffic from the counts. Defaults to `true`.",
			QueryPath: "excludeBots",
		},
		&requestflag.Flag[string]{
			Name:      "granularity",
			Usage:     "Time-series bucket size. Defaults to `DAILY`.",
			QueryPath: "granularity",
		},
	},
	Action:          handleShortLinksResults,
	HideHelpCommand: true,
}

func handleShortLinksList(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.ShortLinkListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.ShortLinks.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "short-links list",
			Transform:      transform,
		})
	} else {
		iter := client.ShortLinks.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "short-links list",
			Transform:      transform,
		})
	}
}

func handleShortLinksCreate(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacy.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := oursprivacy.ShortLinkNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ShortLinks.New(ctx, params, options...)
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
		Title:          "short-links create",
		Transform:      transform,
	})
}

func handleShortLinksRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ShortLinks.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "short-links retrieve",
		Transform:      transform,
	})
}

func handleShortLinksUpdate(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := oursprivacy.ShortLinkUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ShortLinks.Update(
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
		Title:          "short-links update",
		Transform:      transform,
	})
}

func handleShortLinksDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ShortLinks.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "short-links delete",
		Transform:      transform,
	})
}

func handleShortLinksResults(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.ShortLinkResultsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ShortLinks.Results(
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
		Title:          "short-links results",
		Transform:      transform,
	})
}
