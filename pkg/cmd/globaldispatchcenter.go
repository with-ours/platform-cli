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

var globalDispatchCentersList = cli.Command{
	Name:    "list",
	Usage:   "List global dispatch centers for this account. Supports cursor pagination via\n`limit` and `cursor`. Requires scope: globalDispatch:list",
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
	},
	Action:          handleGlobalDispatchCentersList,
	HideHelpCommand: true,
}

var globalDispatchCentersCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new global dispatch center. Requires scope: globalDispatch:create",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*bool]{
			Name:     "is-enabled",
			Usage:    "Whether the center starts enabled. Defaults to false — opt in by setting true here or via PATCH later.",
			BodyPath: "isEnabled",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    `Display name for the new center. Defaults to "Consent Dispatch Center".`,
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "notes",
			Usage:    "Free-form notes shown in the dashboard. Not used for routing.",
			BodyPath: "notes",
		},
	},
	Action:          handleGlobalDispatchCentersCreate,
	HideHelpCommand: true,
}

var globalDispatchCentersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single global dispatch center by ID. Requires scope: globalDispatch:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleGlobalDispatchCentersRetrieve,
	HideHelpCommand: true,
}

var globalDispatchCentersUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Partially update a global dispatch center. Only the fields you send are changed.\nRequires scope: globalDispatch:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[any]{
			Name:     "category",
			Usage:    "Full replacement of the categories list. Categories are sorted by priority on write and re-stamped 1..N — see the priority field. Omit to leave existing categories untouched.",
			BodyPath: "categories",
		},
		&requestflag.Flag[*bool]{
			Name:     "is-enabled",
			Usage:    "Toggle the dispatch center on/off without changing any other config.",
			BodyPath: "isEnabled",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "New display name for the center.",
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "notes",
			Usage:    "Replace the free-form notes.",
			BodyPath: "notes",
		},
	},
	Action:          handleGlobalDispatchCentersUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"category": {
		&requestflag.InnerFlag[*string]{
			Name:                  "category.description",
			Usage:                 "Optional human-readable description shown in the dashboard.",
			InnerField:            "description",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[any]{
			Name:                  "category.destination-ids",
			Usage:                 "Destinations that receive events matching this category. Stale IDs (deleted destinations or ones from another account) are silently filtered out at write time.",
			InnerField:            "destinationIds",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[any]{
			Name:                  "category.logic",
			Usage:                 "Optional condition tree. When set, only matching events route to this category.",
			InnerField:            "logic",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "category.name",
			Usage:                 "Display name for the category. Auto-generated if omitted.",
			InnerField:            "name",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*float64]{
			Name:                  "category.priority",
			Usage:                 "Used as a sort key on write. The server sorts categories by this value (ascending), then re-stamps priority as (sorted index + 1) on persist. Send any positive number — gaps are ironed out, duplicate values keep input order via stable sort. Omit to fall to the end.",
			InnerField:            "priority",
			OuterIsArrayOfObjects: true,
		},
	},
})

var globalDispatchCentersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a global dispatch center. Requires scope: globalDispatch:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleGlobalDispatchCentersDelete,
	HideHelpCommand: true,
}

func handleGlobalDispatchCentersList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.GlobalDispatchCenterListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.GlobalDispatchCenters.List(ctx, params, options...)
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
		Title:          "global-dispatch-centers list",
		Transform:      transform,
	})
}

func handleGlobalDispatchCentersCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.GlobalDispatchCenterNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.GlobalDispatchCenters.New(ctx, params, options...)
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
		Title:          "global-dispatch-centers create",
		Transform:      transform,
	})
}

func handleGlobalDispatchCentersRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.GlobalDispatchCenters.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "global-dispatch-centers retrieve",
		Transform:      transform,
	})
}

func handleGlobalDispatchCentersUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.GlobalDispatchCenters.Update(
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
		Title:          "global-dispatch-centers update",
		Transform:      transform,
	})
}

func handleGlobalDispatchCentersDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.GlobalDispatchCenters.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "global-dispatch-centers delete",
		Transform:      transform,
	})
}
