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

var destinationsList = cli.Command{
	Name:    "list",
	Usage:   "List all destinations. Requires scope: destination:list",
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
			Name:      "status",
			Usage:     "Filter destinations by status.",
			QueryPath: "status",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "Filter destinations by destination type.",
			QueryPath: "type",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleDestinationsList,
	HideHelpCommand: true,
}

var destinationsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new destination. Requires scope: destination:create",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Event-dispatch destination type to create. Warehouse and cloud-storage destination types may appear on read responses but are not creatable through POST.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "settings",
			Usage:    "Per-type configuration keys and values. Call GET /rest/v1/destination-types/{id} to get the valid keys for your destination type.",
			BodyPath: "settings",
		},
	},
	Action:          handleDestinationsCreate,
	HideHelpCommand: true,
}

var destinationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single destination by ID. Requires scope: destination:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleDestinationsRetrieve,
	HideHelpCommand: true,
}

var destinationsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Partially update a destination. Only the fields you send are changed; omitted\nfields are unchanged. The `settings` object is patch-only: omitted keys keep\ntheir current value, and send `null` to clear a specific setting. Requires\nscope: destination:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[*string]{
			Name:     "hashing-salt",
			BodyPath: "hashingSalt",
		},
		&requestflag.Flag[any]{
			Name:     "limited-to-source-id",
			BodyPath: "limitedToSourceIds",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "settings",
			Usage:    "Per-type configuration keys and values. Call GET /rest/v1/destination-types/{id} to get the valid keys for your destination type.",
			BodyPath: "settings",
		},
		&requestflag.Flag[*string]{
			Name:     "status",
			Usage:    `Allowed values: "Disabled", "Enabled".`,
			BodyPath: "status",
		},
	},
	Action:          handleDestinationsUpdate,
	HideHelpCommand: true,
}

var destinationsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a destination. Requires scope: destination:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleDestinationsDelete,
	HideHelpCommand: true,
}

func handleDestinationsList(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.DestinationListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Destinations.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "destinations list",
			Transform:      transform,
		})
	} else {
		iter := client.Destinations.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "destinations list",
			Transform:      transform,
		})
	}
}

func handleDestinationsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.DestinationNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Destinations.New(ctx, params, options...)
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
		Title:          "destinations create",
		Transform:      transform,
	})
}

func handleDestinationsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Destinations.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "destinations retrieve",
		Transform:      transform,
	})
}

func handleDestinationsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.DestinationUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Destinations.Update(
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
		Title:          "destinations update",
		Transform:      transform,
	})
}

func handleDestinationsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Destinations.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "destinations delete",
		Transform:      transform,
	})
}
