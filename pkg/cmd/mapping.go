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

var mappingsList = cli.Command{
	Name:    "list",
	Usage:   "List mappings for an entity (a source or destination). Requires the `entityId`\nquery parameter. Supports cursor pagination via `limit` and `cursor`. Sorted by\n`priority` ascending, then by `id` for deterministic pagination. Requires scope:\nmapping:list",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "entity-id",
			Usage:     "Filter mappings by their parent entity id. Must be a destination id or source id.",
			Required:  true,
			QueryPath: "entityId",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor from pagination.nextCursor in the previous response. Do not decode or modify it. Malformed cursors return 400 Bad Request.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[*int64]{
			Name:      "limit",
			Usage:     "Maximum number of mappings to return. Defaults to 1000; values below 1 are clamped to 1 and values above 1000 are clamped to 1000. Most accounts can fetch the full list in one request.",
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleMappingsList,
	HideHelpCommand: true,
}

var mappingsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new mapping. Requires scope: mapping:create",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "allowed-event-id",
			Required: true,
			BodyPath: "allowedEventId",
		},
		&requestflag.Flag[string]{
			Name:     "destination-id",
			Required: true,
			BodyPath: "destinationId",
		},
	},
	Action:          handleMappingsCreate,
	HideHelpCommand: true,
}

var mappingsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single mapping by ID. Requires scope: mapping:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleMappingsRetrieve,
	HideHelpCommand: true,
}

var mappingsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Partially update a mapping. Only the fields you send are changed. Requires\nscope: mapping:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "logic",
			Usage:    "Condition tree gating when this mapping fires. A node is either a leaf `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are themselves `MappingLogic` nodes, so trees nest arbitrarily. Example leaf: `{ \"condition\": { \"property\": \"$event.event\", \"operator\": \"Is\", \"value\": \"page_view\" } }`. Example combinator: `{ \"AND\": [{ \"condition\": ... }, { \"OR\": [...] }] }`.",
			BodyPath: "logic",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "mapping",
			BodyPath: "mappings",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			BodyPath: "name",
		},
	},
	Action:          handleMappingsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"logic": {
		&requestflag.InnerFlag[any]{
			Name:       "logic.and",
			Usage:      "All child nodes must match. Each child is a `MappingLogic` node.",
			InnerField: "AND",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "logic.condition",
			InnerField: "condition",
		},
		&requestflag.InnerFlag[any]{
			Name:       "logic.not",
			Usage:      "Negates a single child `MappingLogic` node.",
			InnerField: "NOT",
		},
		&requestflag.InnerFlag[any]{
			Name:       "logic.or",
			Usage:      "Any child node must match. Each child is a `MappingLogic` node.",
			InnerField: "OR",
		},
	},
	"mapping": {
		&requestflag.InnerFlag[string]{
			Name:       "mapping.map",
			InnerField: "map",
		},
		&requestflag.InnerFlag[string]{
			Name:       "mapping.property",
			InnerField: "property",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "mapping.modification",
			Usage:      `Allowed values: "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs", "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash", "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender", "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone", "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null", "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase".`,
			InnerField: "modification",
		},
	},
})

var mappingsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a mapping. Requires scope: mapping:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleMappingsDelete,
	HideHelpCommand: true,
}

func handleMappingsList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.MappingListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Mappings.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "mappings list",
			Transform:      transform,
		})
	} else {
		iter := client.Mappings.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "mappings list",
			Transform:      transform,
		})
	}
}

func handleMappingsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.MappingNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Mappings.New(ctx, params, options...)
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
		Title:          "mappings create",
		Transform:      transform,
	})
}

func handleMappingsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Mappings.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "mappings retrieve",
		Transform:      transform,
	})
}

func handleMappingsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.MappingUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Mappings.Update(
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
		Title:          "mappings update",
		Transform:      transform,
	})
}

func handleMappingsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Mappings.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "mappings delete",
		Transform:      transform,
	})
}
