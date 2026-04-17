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

var sourcesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new source. Requires scope: source:create",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "AlchemerWebhook", "AndroidNativeApi", "CSharpApi", "CalComWebhooks", "CalendlyWebhook", "CallRail", "CallTrackingMetrics", "DotNetApi", "FacebookLeadAds", "FormsortWebhooks", "Formstack", "GoLangApi", "HTTPApiSource", "Healthie", "HubspotAppActions", "HubspotFormWebhook", "JotFormWebhooks", "KotlinApi", "NodejsApi", "PHPApi", "PixelImage", "PythonApi", "ReactNativeApi", "RedirectSource", "RubyApi", "SegmentWebPlugin", "TypeformWebhooks", "WebSource", "Webhook", "WhatConverts", "iOSNativeApi".`,
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[any]{
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
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleSourcesRetrieve,
	HideHelpCommand: true,
}

var sourcesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a source. Requires scope: source:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    `Allowed values: "Disabled", "Enabled".`,
			Required: true,
			BodyPath: "status",
		},
		&requestflag.Flag[any]{
			Name:     "bot-control-mode",
			BodyPath: "botControlMode",
		},
		&requestflag.Flag[any]{
			Name:     "bot-score-threshold",
			BodyPath: "botScoreThreshold",
		},
		&requestflag.Flag[any]{
			Name:     "exclude-request-context",
			BodyPath: "excludeRequestContext",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "probabilistic-identity",
			BodyPath: "probabilisticIdentity",
		},
		&requestflag.Flag[any]{
			Name:     "project-api-key",
			BodyPath: "projectAPIKey",
		},
		&requestflag.Flag[any]{
			Name:     "redirect-url",
			BodyPath: "redirectUrl",
		},
		&requestflag.Flag[any]{
			Name:     "selected-account-id",
			BodyPath: "selectedAccountId",
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

var sourcesList = cli.Command{
	Name:            "list",
	Usage:           "List all sources. Requires scope: source:list",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleSourcesList,
	HideHelpCommand: true,
}

var sourcesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a source. Requires scope: source:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleSourcesDelete,
	HideHelpCommand: true,
}

func handleSourcesCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomwithoursplatformsdkgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomwithoursplatformsdkgo.SourceNewParams{}

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
	client := githubcomwithoursplatformsdkgo.NewClient(getDefaultRequestOptions(cmd)...)
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
	client := githubcomwithoursplatformsdkgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomwithoursplatformsdkgo.SourceUpdateParams{}

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

func handleSourcesList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomwithoursplatformsdkgo.NewClient(getDefaultRequestOptions(cmd)...)
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
	_, err = client.Sources.List(ctx, options...)
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
		Title:          "sources list",
		Transform:      transform,
	})
}

func handleSourcesDelete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomwithoursplatformsdkgo.NewClient(getDefaultRequestOptions(cmd)...)
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
