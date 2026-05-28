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

var tagManagerTagsList = cli.Command{
	Name:    "list",
	Usage:   "List tags inside a single tag manager. Requires the `tagManagerId` query\nparameter — tags are always scoped to one parent container. Supports cursor\npagination via `limit` and `cursor`; the limit clamp is 1000 so a single request\ncan return the full set (the web-app workspace renders all tags in one shot).\nRequires scope: tagManagers:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "tag-manager-id",
			Usage:     "Parent tag manager whose tags should be returned.",
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
			Usage:     "Maximum number of tags to return. Defaults to 25; values below 1 are clamped to 1 and values above 1000 are clamped to 1000. The web-app passes 1000 to render the full workspace in one request.",
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleTagManagerTagsList,
	HideHelpCommand: true,
}

var tagManagerTagsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new tag inside a tag manager. `tagManagerId` is required in the body.\nNewly created tags are not assigned to any folder — assign after creation via\nPATCH with `folderId`. Requires scope: tagManagers:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "fire-trigger-id",
			Usage:    "Trigger ids that cause this tag to fire. Use `[]` only for placeholder tags.",
			Required: true,
			BodyPath: "fireTriggerIds",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable tag name.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "parameters",
			Usage:    "Type-specific JSON configuration. Send `{}` for a placeholder.",
			Required: true,
			BodyPath: "parameters",
		},
		&requestflag.Flag[string]{
			Name:     "tag-manager-id",
			Usage:    "Parent tag manager that will own the new tag.",
			Required: true,
			BodyPath: "tagManagerId",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Tag type discriminator. Pick from `GET /tag-manager-tags/types` for the canonical set (e.g. `OursTrackTag`, `OursInitTag`, `CustomHtmlTag`). Names like `GA4Event` are not valid ids.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[any]{
			Name:     "block-trigger-id",
			Usage:    "Optional trigger ids that block this tag when they match.",
			BodyPath: "blockTriggerIds",
		},
		&requestflag.Flag[*bool]{
			Name:     "enabled",
			Usage:    "Defaults to `true`.",
			BodyPath: "enabled",
		},
		&requestflag.Flag[*float64]{
			Name:     "priority",
			Usage:    "Defaults to 0.",
			BodyPath: "priority",
		},
	},
	Action:          handleTagManagerTagsCreate,
	HideHelpCommand: true,
}

var tagManagerTagsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch a single tag by id, including its `folderId` (read-only on this endpoint).\nRequires scope: tagManagers:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleTagManagerTagsRetrieve,
	HideHelpCommand: true,
}

var tagManagerTagsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Partially update a tag. Only the fields you send are changed. Tags cannot be\nmoved between tag managers. To assign a tag to a folder, use\n`POST /rest/v1/tag-manager-asset-folders`. Requires scope: tagManagers:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[any]{
			Name:     "block-trigger-id",
			Usage:    "Replaces the block trigger list wholesale. Send `null` to clear.",
			BodyPath: "blockTriggerIds",
		},
		&requestflag.Flag[*bool]{
			Name:     "enabled",
			Usage:    "Pause/resume the tag without changing other fields.",
			BodyPath: "enabled",
		},
		&requestflag.Flag[[]string]{
			Name:     "fire-trigger-id",
			Usage:    "Replaces the fire trigger list wholesale.",
			BodyPath: "fireTriggerIds",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Updated tag name.",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "parameters",
			Usage:    "Updated type-specific JSON configuration.",
			BodyPath: "parameters",
		},
		&requestflag.Flag[*float64]{
			Name:     "priority",
			Usage:    "Updated priority.",
			BodyPath: "priority",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Updated tag type. Pick from `GET /tag-manager-tags/types`.",
			BodyPath: "type",
		},
	},
	Action:          handleTagManagerTagsUpdate,
	HideHelpCommand: true,
}

var tagManagerTagsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a tag manager tag. Requires scope: tagManagers:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleTagManagerTagsDelete,
	HideHelpCommand: true,
}

var tagManagerTagsTypes = cli.Command{
	Name:            "types",
	Usage:           "Lists every tag template the platform supports — what `type` discriminator to\nsend on create/patch, and the shape of the type-specific `parameters` payload\n(fields, validators, required flags, available values for selects).\nAccount-agnostic: the response is the same for every API key. The same registry\npowers server-side validation on `POST` / `PATCH` so what this endpoint\nadvertises matches what the server enforces. Requires scope: tagManagers:find",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleTagManagerTagsTypes,
	HideHelpCommand: true,
}

func handleTagManagerTagsList(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.TagManagerTagListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.TagManagerTags.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "tag-manager-tags list",
			Transform:      transform,
		})
	} else {
		iter := client.TagManagerTags.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "tag-manager-tags list",
			Transform:      transform,
		})
	}
}

func handleTagManagerTagsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.TagManagerTagNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TagManagerTags.New(ctx, params, options...)
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
		Title:          "tag-manager-tags create",
		Transform:      transform,
	})
}

func handleTagManagerTagsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.TagManagerTags.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "tag-manager-tags retrieve",
		Transform:      transform,
	})
}

func handleTagManagerTagsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.TagManagerTagUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TagManagerTags.Update(
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
		Title:          "tag-manager-tags update",
		Transform:      transform,
	})
}

func handleTagManagerTagsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.TagManagerTags.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "tag-manager-tags delete",
		Transform:      transform,
	})
}

func handleTagManagerTagsTypes(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.TagManagerTags.Types(ctx, options...)
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
		Title:          "tag-manager-tags types",
		Transform:      transform,
	})
}
