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

var tagManagerVariablesList = cli.Command{
	Name:    "list",
	Usage:   "List variables inside a single tag manager. Requires the `tagManagerId` query\nparameter — variables are always scoped to one parent container. Supports cursor\npagination via `limit` and `cursor`; the limit clamp is 1000 so a single request\ncan return the full set (the web-app workspace renders all variables in one\nshot). Requires scope: tagManagers:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "tag-manager-id",
			Usage:     "Parent tag manager whose variables should be returned.",
			Required:  true,
			QueryPath: "tagManagerId",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor from pagination.nextCursor in the previous response. Do not decode or modify it. Malformed cursors return 400 Bad Request.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[*int64]{
			Name:      "limit",
			Usage:     "Maximum number of variables to return. Defaults to 25; values below 1 are clamped to 1 and values above 1000 are clamped to 1000. The web-app passes 1000 to render the full workspace in one request.",
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleTagManagerVariablesList,
	HideHelpCommand: true,
}

var tagManagerVariablesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new variable inside a tag manager. `tagManagerId` is required in the\nbody. Known input failures (e.g. duplicate variable name within the tag manager)\nare returned as HTTP 409 with the reason in the response `error` field. Requires\nscope: tagManagers:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "parameters",
			Usage:    "Type-specific JSON configuration.",
			Required: true,
			BodyPath: "parameters",
		},
		&requestflag.Flag[string]{
			Name:     "tag-manager-id",
			Usage:    "Parent tag manager that will own the new variable.",
			Required: true,
			BodyPath: "tagManagerId",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Variable type discriminator.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "variable",
			Usage:    "Variable implementation identifier (typically equals `type`).",
			Required: true,
			BodyPath: "Variable",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "default-value",
			Usage:    "Optional default value. JSON value of any type.",
			BodyPath: "defaultValue",
		},
		&requestflag.Flag[*bool]{
			Name:     "enabled",
			BodyPath: "enabled",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "look-up-table",
			Usage:    "Optional lookup table for `LookUpTable` variables.",
			BodyPath: "lookUpTable",
		},
	},
	Action:          handleTagManagerVariablesCreate,
	HideHelpCommand: true,
}

var tagManagerVariablesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single tag manager variable by ID. Requires scope: tagManagers:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleTagManagerVariablesRetrieve,
	HideHelpCommand: true,
}

var tagManagerVariablesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Partially update a variable. Only the fields you send are changed. Name\ncollisions with other variables in the same tag manager return 409 with the\nreason in the response `error` field. Requires scope: tagManagers:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "default-value",
			Usage:    "Updated default value. JSON value of any type.",
			BodyPath: "defaultValue",
		},
		&requestflag.Flag[*bool]{
			Name:     "enabled",
			Usage:    "Pause/resume the variable without changing other fields.",
			BodyPath: "enabled",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "look-up-table",
			Usage:    "Updated lookup table payload.",
			BodyPath: "lookUpTable",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Updated variable name.",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "parameters",
			Usage:    "Updated type-specific JSON configuration.",
			BodyPath: "parameters",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Updated variable type.",
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "variable",
			Usage:    "Updated variable implementation identifier.",
			BodyPath: "Variable",
		},
	},
	Action:          handleTagManagerVariablesUpdate,
	HideHelpCommand: true,
}

var tagManagerVariablesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a tag manager variable. Requires scope: tagManagers:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleTagManagerVariablesDelete,
	HideHelpCommand: true,
}

var tagManagerVariablesTypes = cli.Command{
	Name:            "types",
	Usage:           "Lists every variable template the platform supports — what `type` discriminator\nto send on create/patch, the shape of the type-specific `parameters` payload,\nand `supportsVariables` (whether the variable's own parameter fields may\nreference `{{OtherVariable}}` at runtime). Account-agnostic: the response is the\nsame for every API key. Requires scope: tagManagers:find",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleTagManagerVariablesTypes,
	HideHelpCommand: true,
}

func handleTagManagerVariablesList(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.TagManagerVariableListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.TagManagerVariables.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "tag-manager-variables list",
			Transform:      transform,
		})
	} else {
		iter := client.TagManagerVariables.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "tag-manager-variables list",
			Transform:      transform,
		})
	}
}

func handleTagManagerVariablesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.TagManagerVariableNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TagManagerVariables.New(ctx, params, options...)
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
		Title:          "tag-manager-variables create",
		Transform:      transform,
	})
}

func handleTagManagerVariablesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.TagManagerVariables.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "tag-manager-variables retrieve",
		Transform:      transform,
	})
}

func handleTagManagerVariablesUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.TagManagerVariableUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TagManagerVariables.Update(
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
		Title:          "tag-manager-variables update",
		Transform:      transform,
	})
}

func handleTagManagerVariablesDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.TagManagerVariables.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "tag-manager-variables delete",
		Transform:      transform,
	})
}

func handleTagManagerVariablesTypes(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.TagManagerVariables.Types(ctx, options...)
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
		Title:          "tag-manager-variables types",
		Transform:      transform,
	})
}
