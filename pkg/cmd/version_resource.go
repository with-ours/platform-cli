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

var versionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new version. Requires scope: version:publish",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "include-allowed-event",
			BodyPath: "includeAllowedEvents",
		},
		&requestflag.Flag[any]{
			Name:     "include-consent-setting",
			BodyPath: "includeConsentSettings",
		},
		&requestflag.Flag[any]{
			Name:     "include-destination",
			BodyPath: "includeDestinations",
		},
		&requestflag.Flag[any]{
			Name:     "include-external-allowed-event-data",
			BodyPath: "includeExternalAllowedEventData",
		},
		&requestflag.Flag[any]{
			Name:     "include-global-dispatch-center",
			BodyPath: "includeGlobalDispatchCenters",
		},
		&requestflag.Flag[any]{
			Name:     "include-mapping",
			BodyPath: "includeMappings",
		},
		&requestflag.Flag[any]{
			Name:     "include-replay-setting",
			BodyPath: "includeReplaySettings",
		},
		&requestflag.Flag[any]{
			Name:     "include-source",
			BodyPath: "includeSources",
		},
		&requestflag.Flag[any]{
			Name:     "include-tag-manager-tag",
			BodyPath: "includeTagManagerTags",
		},
		&requestflag.Flag[any]{
			Name:     "include-tag-manager-trigger",
			BodyPath: "includeTagManagerTriggers",
		},
		&requestflag.Flag[any]{
			Name:     "include-tag-manager-variable",
			BodyPath: "includeTagManagerVariables",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "notes",
			BodyPath: "notes",
		},
	},
	Action:          handleVersionsCreate,
	HideHelpCommand: true,
}

var versionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single version by ID. Requires scope: version:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleVersionsRetrieve,
	HideHelpCommand: true,
}

var versionsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a version. Requires scope: version:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "notes",
			BodyPath: "notes",
		},
	},
	Action:          handleVersionsUpdate,
	HideHelpCommand: true,
}

var versionsList = cli.Command{
	Name:            "list",
	Usage:           "List all versions. Requires scope: version:list",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleVersionsList,
	HideHelpCommand: true,
}

func handleVersionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomwithoursplatformsdkgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomwithoursplatformsdkgo.VersionNewParams{}

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
	_, err = client.Versions.New(ctx, params, options...)
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
		Title:          "versions create",
		Transform:      transform,
	})
}

func handleVersionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Versions.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "versions retrieve",
		Transform:      transform,
	})
}

func handleVersionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomwithoursplatformsdkgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomwithoursplatformsdkgo.VersionUpdateParams{}

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
	_, err = client.Versions.Update(
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
		Title:          "versions update",
		Transform:      transform,
	})
}

func handleVersionsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Versions.List(ctx, options...)
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
		Title:          "versions list",
		Transform:      transform,
	})
}
