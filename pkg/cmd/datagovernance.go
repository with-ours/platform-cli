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

var dataGovernanceList = cli.Command{
	Name:    "list",
	Usage:   "List the data-governance record(s) on this account. Each account has at most one\nrecord, so this list returns either an empty array or a single entity. Cursor\npagination is exposed for consistency with other list endpoints but is rarely\nmeaningful here. Data governance is the second stage of the dispatch flow\n(Source → Allowed Events → Data Governance → Mappings → Destination) — it\nevaluates each event against the configured category logic and stops dispatch to\nthe destinations on any matching category. Requires scope: globalDispatch:list",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleDataGovernanceList,
	HideHelpCommand: true,
}

var dataGovernanceCreate = cli.Command{
	Name:    "create",
	Usage:   "Create the data-governance record for this account. Each account may have at\nmost one — a second POST returns 409. Body is optional; defaults are\n`isEnabled: false` and no categories. Categories are added later via PATCH.\nRequires scope: globalDispatch:create",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*bool]{
			Name:     "is-enabled",
			Usage:    "Whether the record starts enabled. Defaults to false — opt in by setting true here or via PATCH later. When disabled, every category is bypassed and inbound events flow through to destinations regardless of consent state.",
			BodyPath: "isEnabled",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    `Display name for the new record. Defaults to "Data Governance".`,
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "notes",
			Usage:    "Free-form notes shown in the dashboard. Not used for routing.",
			BodyPath: "notes",
		},
	},
	Action:          handleDataGovernanceCreate,
	HideHelpCommand: true,
}

var dataGovernanceRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch the data-governance record by id, including its categories (logic,\ndestinations, priority). Returns 404 when no record matches. Requires scope:\nglobalDispatch:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleDataGovernanceRetrieve,
	HideHelpCommand: true,
}

var dataGovernanceUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Partially update the data-governance record. Top-level fields (`name`, `notes`,\n`isEnabled`) follow the standard PATCH semantic — only the fields you send are\nchanged.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[any]{
			Name:     "category",
			Usage:    "Full replacement of the categories list. The sent array becomes the new state — there is no partial-merge semantic for categories. Categories are sorted by priority on write and re-stamped 1..N. Omit to leave existing categories untouched.",
			BodyPath: "categories",
		},
		&requestflag.Flag[*bool]{
			Name:     "is-enabled",
			Usage:    "Toggle data governance on/off without changing any other config.",
			BodyPath: "isEnabled",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "New display name for the record.",
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "notes",
			Usage:    "Replace the free-form notes.",
			BodyPath: "notes",
		},
	},
	Action:          handleDataGovernanceUpdate,
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
			Usage:                 "Destinations gated by this category. When the category logic evaluates to TRUE for an inbound event, dispatch to every destination on this list is stopped. Stale IDs (deleted destinations or destinations on another account) are silently filtered out at write time.",
			InnerField:            "destinationIds",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:                  "category.logic",
			Usage:                 "Condition tree evaluated against each inbound event. Write conditions that evaluate **TRUE for events you want to STOP**. A node is either a leaf `condition` or a combinator (`AND`, `OR`, `NOT`); combinator children are themselves logic nodes, so trees nest arbitrarily.\n\nDiscovery: `GET /rest/v1/mappings/default-variables` lists the canonical platform-provided `property` paths (visitor consent arrays, event fields, request context, identity fields). Custom `event.event_properties.*` paths are caller-defined.\n\nExample leaf (stop dispatch when the visitor rejected the `advertising` consent category): `{ \"condition\": { \"property\": \"visitor.consent.rejected_categories\", \"operator\": \"Contains\", \"value\": \"advertising\" } }`.\n\nExample combinator: `{ \"AND\": [{ \"condition\": { \"property\": \"visitor.consent.rejected_categories\", \"operator\": \"Contains\", \"value\": \"advertising\" } }, { \"OR\": [/* nested logic nodes */] }] }`.",
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

var dataGovernanceDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete the data-governance record. After deletion, inbound events flow through\nto destinations without category-level gating. Create a new record with POST to\nreinstate governance. Requires scope: globalDispatch:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleDataGovernanceDelete,
	HideHelpCommand: true,
}

func handleDataGovernanceList(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.DataGovernanceListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.DataGovernance.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "data-governance list",
			Transform:      transform,
		})
	} else {
		iter := client.DataGovernance.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "data-governance list",
			Transform:      transform,
		})
	}
}

func handleDataGovernanceCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.DataGovernanceNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.DataGovernance.New(ctx, params, options...)
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
		Title:          "data-governance create",
		Transform:      transform,
	})
}

func handleDataGovernanceRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.DataGovernance.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "data-governance retrieve",
		Transform:      transform,
	})
}

func handleDataGovernanceUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.DataGovernanceUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.DataGovernance.Update(
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
		Title:          "data-governance update",
		Transform:      transform,
	})
}

func handleDataGovernanceDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.DataGovernance.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "data-governance delete",
		Transform:      transform,
	})
}
