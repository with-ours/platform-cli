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

var mappingsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a mapping. Two body shapes are accepted:",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "allowed-event-id",
			Usage:    "Quick-create variant: allowed event to bind the new mapping to. Required together with `destinationId`. Mutually exclusive with `templateId`/`entityId`.",
			BodyPath: "allowedEventId",
		},
		&requestflag.Flag[string]{
			Name:     "destination-id",
			Usage:    "Quick-create variant: destination that should receive events matched by this mapping. Required together with `allowedEventId`.",
			BodyPath: "destinationId",
		},
		&requestflag.Flag[string]{
			Name:     "entity-id",
			Usage:    "Template fat-create variant: destination or source id this mapping belongs to. Required together with `templateId`.",
			BodyPath: "entityId",
		},
		&requestflag.Flag[int64]{
			Name:     "insert-after-idx",
			Usage:    "Template fat-create only. Zero-based position in the destination/source priority order to insert this mapping after. Omit to append at the end.",
			BodyPath: "insertAfterIdx",
		},
		&requestflag.Flag[bool]{
			Name:     "is-enabled",
			Usage:    "Template fat-create only. Initial enabled state. Defaults to `true`.",
			BodyPath: "isEnabled",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "logic",
			Usage:    "Condition tree gating when this mapping fires. A node is either a leaf `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are themselves `MappingLogic` nodes, so trees nest arbitrarily. Example leaf: `{ \"condition\": { \"property\": \"$event.event\", \"operator\": \"Is\", \"value\": \"page_view\" } }`. Example combinator: `{ \"AND\": [{ \"condition\": ... }, { \"OR\": [...] }] }`.",
			BodyPath: "logic",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "mapping",
			Usage:    "Template fat-create only. Optional initial property mappings. When omitted the mapping is seeded with template defaults for sources and non-default destination templates, and with `[]` for default destination templates.",
			BodyPath: "mappings",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "Template fat-create only. Override the auto-generated mapping name.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "template-id",
			Usage:    "Template fat-create variant: template id from `GET /rest/v1/mapping-templates`. Picks the property descriptor set used to validate `mappings[].property`. Required together with `entityId`.",
			BodyPath: "templateId",
		},
	},
	Action:          handleMappingsCreate,
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
	Usage:   "Partially update a mapping. Only the fields you send are changed. Send\n`isEnabled: false` to pause the mapping without changing other fields (mirrors\n`status` on destinations). `mappings[]` is replaced wholesale when sent.\nRequires scope: mapping:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[*bool]{
			Name:     "is-enabled",
			Usage:    "Flip the mapping on/off without changing other fields. `null` is treated as omitted.",
			BodyPath: "isEnabled",
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

var mappingsReorder = cli.Command{
	Name:    "reorder",
	Usage:   "Reassign `priority` for a set of mappings. Pass `{ uuids: [...] }` with the\nmapping ids in their new order — index 0 becomes the highest-priority mapping.\nAll ids must belong to the same parent entity (source or destination); mixing\nentities returns 400. Requires scope: mapping:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "uuid",
			Usage:    "Mapping ids in their new priority order, low priority index first. All ids must belong to the same parent entity (source or destination).",
			Required: true,
			BodyPath: "uuids",
		},
	},
	Action:          handleMappingsReorder,
	HideHelpCommand: true,
}

var mappingsTemplates = cli.Command{
	Name:    "templates",
	Usage:   "Discover every mapping template available for a destination or source, with full\nproperty descriptors inlined. Use the returned `id` as `templateId` when calling\n`POST /rest/v1/mappings` (template fat-create variant), and use each entry under\n`mappings[]` to learn the valid `property`, `kind`, `modificationOptions`, and\nany enforced `options`. The `isDefault: true` entry is the destination's\nbuilt-in default template (the one stored at `MAPPER#{destinationId}` when\nconfigured via `PUT /rest/v1/default-mappings/{destinationId}`). Requires scope:\nmapping:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "entity-id",
			Usage:     "Destination or source id. Required.",
			Required:  true,
			QueryPath: "entityId",
		},
	},
	Action:          handleMappingsTemplates,
	HideHelpCommand: true,
}

var mappingsDefaultVariables = cli.Command{
	Name:            "default-variables",
	Usage:           "Lists the platform-provided variables that any mapping `value` can reference\n(e.g. `event.email`, `event.request_context.ip`, `visitor.id`). Account-agnostic\ndiscovery — use these paths as the right-hand side of a mapping field. Requires\nscope: variables:find-default",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleMappingsDefaultVariables,
	HideHelpCommand: true,
}

var mappingsCustomVariables = cli.Command{
	Name:            "custom-variables",
	Usage:           "Lists the custom variables observed in this account’s recent event stream (last\n14 days). These are dot-paths under `event.event_properties.*` that callers can\ntarget in mapping `value` fields. The result is cached for 10 minutes; an empty\nlist means no custom properties have been seen yet for this account. Requires\nscope: variables:find-custom",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleMappingsCustomVariables,
	HideHelpCommand: true,
}

var mappingsModifications = cli.Command{
	Name:            "modifications",
	Usage:           "Lists every value accepted on a mapping field’s `modification` property, with a\nhuman-readable label and one-sentence description. Account-agnostic. Use this\nalongside `GET /rest/v1/mapping-templates` to render a labelled modification\npicker without hardcoding the enum. Requires scope: variables:find-default",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleMappingsModifications,
	HideHelpCommand: true,
}

func handleMappingsList(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.MappingListParams{}

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

	params := oursprivacy.MappingNewParams{}

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

	params := oursprivacy.MappingUpdateParams{}

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

func handleMappingsReorder(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.MappingReorderParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Mappings.Reorder(ctx, params, options...)
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
		Title:          "mappings reorder",
		Transform:      transform,
	})
}

func handleMappingsTemplates(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.MappingTemplatesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Mappings.Templates(ctx, params, options...)
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
		Title:          "mappings templates",
		Transform:      transform,
	})
}

func handleMappingsDefaultVariables(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Mappings.DefaultVariables(ctx, options...)
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
		Title:          "mappings default-variables",
		Transform:      transform,
	})
}

func handleMappingsCustomVariables(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Mappings.CustomVariables(ctx, options...)
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
		Title:          "mappings custom-variables",
		Transform:      transform,
	})
}

func handleMappingsModifications(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Mappings.Modifications(ctx, options...)
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
		Title:          "mappings modifications",
		Transform:      transform,
	})
}
