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

var tagManagerTriggersList = cli.Command{
	Name:    "list",
	Usage:   "List triggers inside a single tag manager. Requires the `tagManagerId` query\nparameter — triggers are always scoped to one parent container. Supports cursor\npagination via `limit` and `cursor`; the limit clamp is 1000 so a single request\ncan return the full set (the web-app workspace renders all triggers in one\nshot). Requires scope: tagManagers:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "tag-manager-id",
			Usage:     "Parent tag manager whose triggers should be returned.",
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
			Usage:     "Maximum number of triggers to return. Defaults to 25; values below 1 are clamped to 1 and values above 1000 are clamped to 1000. The web-app passes 1000 to render the full workspace in one request.",
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleTagManagerTriggersList,
	HideHelpCommand: true,
}

var tagManagerTriggersCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new trigger inside a tag manager. `tagManagerId` is required in the\nbody. Send `conditions: []` for an unconditional trigger; otherwise supply\ntype-specific condition objects. Requires scope: tagManagers:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "condition",
			Usage:    "Match conditions; use `[]` for an unconditional trigger.",
			Required: true,
			BodyPath: "conditions",
		},
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
			Usage:    "Parent tag manager that will own the new trigger.",
			Required: true,
			BodyPath: "tagManagerId",
		},
		&requestflag.Flag[string]{
			Name:     "trigger",
			Usage:    "Trigger implementation identifier (typically equals `type`).",
			Required: true,
			BodyPath: "Trigger",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Trigger type discriminator.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[*bool]{
			Name:     "enabled",
			BodyPath: "enabled",
		},
	},
	Action:          handleTagManagerTriggersCreate,
	HideHelpCommand: true,
}

var tagManagerTriggersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single tag manager trigger by ID. Requires scope: tagManagers:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleTagManagerTriggersRetrieve,
	HideHelpCommand: true,
}

var tagManagerTriggersUpdate = cli.Command{
	Name:    "update",
	Usage:   "Partially update a trigger. Only the fields you send are changed. `conditions`\nis replaced wholesale when sent. Requires scope: tagManagers:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "condition",
			Usage:    "Replaces conditions wholesale when sent. Use `[]` for an unconditional trigger.",
			BodyPath: "conditions",
		},
		&requestflag.Flag[*bool]{
			Name:     "enabled",
			Usage:    "Pause/resume the trigger without changing other fields.",
			BodyPath: "enabled",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Updated trigger name.",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "parameters",
			Usage:    "Updated type-specific JSON configuration.",
			BodyPath: "parameters",
		},
		&requestflag.Flag[string]{
			Name:     "trigger",
			Usage:    "Updated trigger implementation identifier.",
			BodyPath: "Trigger",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Updated trigger type.",
			BodyPath: "type",
		},
	},
	Action:          handleTagManagerTriggersUpdate,
	HideHelpCommand: true,
}

var tagManagerTriggersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a tag manager trigger. Requires scope: tagManagers:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleTagManagerTriggersDelete,
	HideHelpCommand: true,
}

var tagManagerTriggersTypes = cli.Command{
	Name:            "types",
	Usage:           "Lists every trigger template the platform supports — what `type` discriminator\nto send on create/patch, and the shape of the type-specific `parameters`\npayload. Trigger `conditions` are evaluated at runtime (per-trigger, see the\nresource docs) and are not part of this descriptor. Account-agnostic: the\nresponse is the same for every API key. Requires scope: tagManagers:find",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleTagManagerTriggersTypes,
	HideHelpCommand: true,
}

func handleTagManagerTriggersList(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.TagManagerTriggerListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.TagManagerTriggers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "tag-manager-triggers list",
			Transform:      transform,
		})
	} else {
		iter := client.TagManagerTriggers.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "tag-manager-triggers list",
			Transform:      transform,
		})
	}
}

func handleTagManagerTriggersCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.TagManagerTriggerNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TagManagerTriggers.New(ctx, params, options...)
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
		Title:          "tag-manager-triggers create",
		Transform:      transform,
	})
}

func handleTagManagerTriggersRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.TagManagerTriggers.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "tag-manager-triggers retrieve",
		Transform:      transform,
	})
}

func handleTagManagerTriggersUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.TagManagerTriggerUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TagManagerTriggers.Update(
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
		Title:          "tag-manager-triggers update",
		Transform:      transform,
	})
}

func handleTagManagerTriggersDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.TagManagerTriggers.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "tag-manager-triggers delete",
		Transform:      transform,
	})
}

func handleTagManagerTriggersTypes(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.TagManagerTriggers.Types(ctx, options...)
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
		Title:          "tag-manager-triggers types",
		Transform:      transform,
	})
}
