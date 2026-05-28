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

var tagManagersList = cli.Command{
	Name:            "list",
	Usage:           "List every tag manager on this account. Each tag manager is a pixel-scoped\ncontainer of tags, triggers, variables, and folders. Not paginated — accounts\nare capped at a small number of tag managers in practice, so the response fits\nin a single page. Requires scope: tagManagers:list",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleTagManagersList,
	HideHelpCommand: true,
}

var tagManagersCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new tag manager. The server seeds three default triggers\n(`Initialization`, `PageView`, `DomReady`) and one `OursInitTag` so the\ncontainer is immediately usable — call\n`GET /tag-manager-triggers?tagManagerId={id}` right after create to grab their\nserver-assigned ids so you can reuse them in `fireTriggerIds` instead of\nredundantly creating a second `PageView`/`DomReady`/`Initialization`. Returns\nthe bare entity. Requires scope: tagManagers:create",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable name for the new tag manager.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "data-layer-name",
			Usage:    "Optional global data-layer name (e.g. `dataLayer`). Omit to use the SDK default.",
			BodyPath: "dataLayerName",
		},
	},
	Action:          handleTagManagersCreate,
	HideHelpCommand: true,
}

var tagManagersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch a single tag manager by id, including its server-assigned `pixel` token\nused by the install snippet. Requires scope: tagManagers:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleTagManagersRetrieve,
	HideHelpCommand: true,
}

var tagManagersUpdate = cli.Command{
	Name:    "update",
	Usage:   "Partially update a tag manager. Only the fields you send are changed; omitted\nfields keep their current value. Send `dataLayerName: null` to clear the\noverride and fall back to the SDK default. Requires scope: tagManagers:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[*string]{
			Name:     "data-layer-name",
			Usage:    "New data-layer name. Send `null` to clear and fall back to the SDK default.",
			BodyPath: "dataLayerName",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "New display name.",
			BodyPath: "name",
		},
	},
	Action:          handleTagManagersUpdate,
	HideHelpCommand: true,
}

var tagManagersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a tag manager. Child tags, triggers, variables, and folders are\ncascade-deleted with the container. Requires scope: tagManagers:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleTagManagersDelete,
	HideHelpCommand: true,
}

func handleTagManagersList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.TagManagers.List(ctx, options...)
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
		Title:          "tag-managers list",
		Transform:      transform,
	})
}

func handleTagManagersCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.TagManagerNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TagManagers.New(ctx, params, options...)
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
		Title:          "tag-managers create",
		Transform:      transform,
	})
}

func handleTagManagersRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.TagManagers.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "tag-managers retrieve",
		Transform:      transform,
	})
}

func handleTagManagersUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.TagManagerUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TagManagers.Update(
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
		Title:          "tag-managers update",
		Transform:      transform,
	})
}

func handleTagManagersDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.TagManagers.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "tag-managers delete",
		Transform:      transform,
	})
}
