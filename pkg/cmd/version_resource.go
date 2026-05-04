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
	Usage:   "Publish the current draft (i.e. all unpublished entity changes) as a new\nversion. Returns the full Version on success. Returns HTTP 409 with the reason\nin the response `error` field when there are no draft changes to publish, when\nanother publish is already in flight, or when the action otherwise conflicts\nwith current state. To re-publish an existing version, use POST\n/rest/v1/versions/{id}/publish instead. Requires scope: version:publish",
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
		&requestflag.Flag[*string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
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
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleVersionsRetrieve,
	HideHelpCommand: true,
}

var versionsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Partially update a version. Only the fields you send are changed. Requires\nscope: version:update",
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
		&requestflag.Flag[*string]{
			Name:     "notes",
			BodyPath: "notes",
		},
	},
	Action:          handleVersionsUpdate,
	HideHelpCommand: true,
}

var versionsList = cli.Command{
	Name:    "list",
	Usage:   "List versions for this account, newest first. Supports cursor pagination and\nfiltering by `isPublished`, `nameContains`, and `notesContains`. Combine filters\nwith AND semantics. Requires scope: version:list",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor from pagination.nextCursor in the previous response. Do not decode or modify it. Malformed cursors return 400 Bad Request.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "is-published",
			Usage:     "Filter to only published or unpublished versions.",
			QueryPath: "isPublished",
		},
		&requestflag.Flag[*int64]{
			Name:      "limit",
			Usage:     "Maximum number of versions to return. Defaults to 25; values below 1 are clamped to 1 and values above 100 are clamped to 100.",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "name-contains",
			Usage:     "Case-insensitive substring match on the version name.",
			QueryPath: "nameContains",
		},
		&requestflag.Flag[string]{
			Name:      "notes-contains",
			Usage:     "Case-insensitive substring match on the version notes.",
			QueryPath: "notesContains",
		},
	},
	Action:          handleVersionsList,
	HideHelpCommand: true,
}

func handleVersionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomwithoursplatformsdkgo.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := githubcomwithoursplatformsdkgo.VersionNewParams{}

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

	params := githubcomwithoursplatformsdkgo.VersionUpdateParams{}

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

	params := githubcomwithoursplatformsdkgo.VersionListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Versions.List(ctx, params, options...)
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
