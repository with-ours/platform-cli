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

var sourcesList = cli.Command{
	Name:    "list",
	Usage:   "List all sources for this account. Supports cursor pagination and optional\nfilters for `type`, `status`, and `nameContains`. Results are sorted by creation\ndate descending. Requires scope: source:list",
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
			Usage:     "Case-insensitive substring filter on the source name.",
			QueryPath: "nameContains",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by source status.",
			QueryPath: "status",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "Filter by source type.",
			QueryPath: "type",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleSourcesList,
	HideHelpCommand: true,
}

var sourcesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new source. Returns the full source entity (same shape as GET\n/sources/{id}) so callers can read all server-assigned fields without a\nfollow-up GET. Requires scope: source:create",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "AlchemerWebhook", "AndroidNativeApi", "CSharpApi", "CalComWebhooks", "CalendlyWebhook", "CallRail", "CallTrackingMetrics", "DotNetApi", "FacebookLeadAds", "FormsortWebhooks", "Formstack", "GoLangApi", "HTTPApiSource", "Healthie", "HubspotAppActions", "HubspotFormWebhook", "JotFormWebhooks", "KotlinApi", "NodejsApi", "PHPApi", "PixelImage", "PythonApi", "ReactNativeApi", "RedirectSource", "RubyApi", "SegmentWebPlugin", "TypeformWebhooks", "WebSource", "Webhook", "WhatConverts", "iOSNativeApi".`,
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			BodyPath: "name",
		},
	},
	Action:          handleSourcesCreate,
	HideHelpCommand: true,
}

var sourcesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single source by ID. Requires scope: source:view",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleSourcesRetrieve,
	HideHelpCommand: true,
}

var sourcesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Partially update a source. Only the fields you send are changed; omitted fields\nare unchanged. Send explicit `null` to clear a nullable field. Returns the full\nsource entity after the update. Requires scope: source:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[*string]{
			Name:     "bot-control-mode",
			BodyPath: "botControlMode",
		},
		&requestflag.Flag[*float64]{
			Name:     "bot-score-threshold",
			BodyPath: "botScoreThreshold",
		},
		&requestflag.Flag[*bool]{
			Name:     "exclude-request-context",
			BodyPath: "excludeRequestContext",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "probabilistic-identity",
			BodyPath: "probabilisticIdentity",
		},
		&requestflag.Flag[*string]{
			Name:     "project-api-key",
			BodyPath: "projectAPIKey",
		},
		&requestflag.Flag[*string]{
			Name:     "redirect-url",
			BodyPath: "redirectUrl",
		},
		&requestflag.Flag[*string]{
			Name:     "selected-account-id",
			BodyPath: "selectedAccountId",
		},
		&requestflag.Flag[*string]{
			Name:     "status",
			BodyPath: "status",
		},
		&requestflag.Flag[any]{
			Name:     "whitelist-domain",
			BodyPath: "whitelistDomains",
		},
		&requestflag.Flag[any]{
			Name:     "whitelist-ip",
			BodyPath: "whitelistIps",
		},
	},
	Action:          handleSourcesUpdate,
	HideHelpCommand: true,
}

var sourcesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a source. Requires scope: source:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleSourcesDelete,
	HideHelpCommand: true,
}

var sourcesTokens = cli.Command{
	Name:    "tokens",
	Usage:   "Returns the install or ingest tokens for a source. The response is a\ndiscriminated union on `sourceType`: pixel sources return\n`{ sourceType: \"pixel\", token, testToken, installScript, testInstallScript }`,\nand webhook sources return\n`{ sourceType: \"webhook\", token, testToken, ingestUrl, testIngestUrl, sampleCurl }`.\nInspect the source's `type` field (`GET /rest/v1/sources/{id}`) to know which\nvariant to expect. Requires scope: source:view",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleSourcesTokens,
	HideHelpCommand: true,
}

func handleSourcesList(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.SourceListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Sources.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "sources list",
			Transform:      transform,
		})
	} else {
		iter := client.Sources.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "sources list",
			Transform:      transform,
		})
	}
}

func handleSourcesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.SourceNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sources.New(ctx, params, options...)
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
		Title:          "sources create",
		Transform:      transform,
	})
}

func handleSourcesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Sources.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "sources retrieve",
		Transform:      transform,
	})
}

func handleSourcesUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.SourceUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sources.Update(
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
		Title:          "sources update",
		Transform:      transform,
	})
}

func handleSourcesDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Sources.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "sources delete",
		Transform:      transform,
	})
}

func handleSourcesTokens(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Sources.Tokens(ctx, cmd.Value("id").(string), options...)
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
		Title:          "sources tokens",
		Transform:      transform,
	})
}
