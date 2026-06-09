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
			Usage:     "Maximum number of items to return. Defaults to 25; values below 1 are clamped to 1 and values above 100 are clamped to 100.",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleVersionsList,
	HideHelpCommand: true,
}

var versionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Publish the current draft of non-experiment entities (destinations, mappings,\nexperiment settings, governance rules, etc.) as a new version. Newly created or\nmodified DRAFT experiments and experiment variants are NOT shipped here — call\n`POST /rest/v1/experiments/{id}/start` instead, which atomically publishes the\nexperiment and its variants by default (`publishAfterStart: true`). Returns the\nfull Version on success. Returns HTTP 409 with the reason in the response\n`error` field when there are no draft changes to publish, when another publish\nis already in flight, or when the action otherwise conflicts with current state.\nTo re-publish an existing version, use POST /rest/v1/versions/{id}/publish\ninstead.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "include-allowed-event",
			Usage:    "Cherry-pick: allowed event ids (slug-like strings) to include from the draft. Omit or send `[]` to include all draft changes in this collection.",
			BodyPath: "includeAllowedEvents",
		},
		&requestflag.Flag[any]{
			Name:     "include-consent-setting",
			Usage:    "Cherry-pick: consent settings ids to include from the draft. Omit or send `[]` to include all draft changes in this collection.",
			BodyPath: "includeConsentSettings",
		},
		&requestflag.Flag[any]{
			Name:     "include-data-governance-event",
			Usage:    "Cherry-pick: data governance event UUIDs to include from the draft. Omit or send `[]` to include all draft changes in this collection.",
			BodyPath: "includeDataGovernanceEvents",
		},
		&requestflag.Flag[any]{
			Name:     "include-data-governance-rule",
			Usage:    "Cherry-pick: data governance rule UUIDs to include from the draft. Omit or send `[]` to include all draft changes in this collection.",
			BodyPath: "includeDataGovernanceRules",
		},
		&requestflag.Flag[any]{
			Name:     "include-destination",
			Usage:    "Cherry-pick: destination UUIDs to include from the draft. Omit or send `[]` to include all draft changes in this collection.",
			BodyPath: "includeDestinations",
		},
		&requestflag.Flag[any]{
			Name:     "include-experiment",
			Usage:    "Cherry-pick: experiment UUIDs to include from the draft. Omit or send `[]` to include all draft changes in this collection.",
			BodyPath: "includeExperiments",
		},
		&requestflag.Flag[any]{
			Name:     "include-experiment-setting",
			Usage:    "Cherry-pick: experiment settings UUIDs to include from the draft. Omit or send `[]` to include all draft changes in this collection.",
			BodyPath: "includeExperimentSettings",
		},
		&requestflag.Flag[any]{
			Name:     "include-experiment-variant",
			Usage:    "Cherry-pick: experiment variant UUIDs to include from the draft. Omit or send `[]` to include all draft changes in this collection.",
			BodyPath: "includeExperimentVariants",
		},
		&requestflag.Flag[any]{
			Name:     "include-external-allowed-event-data",
			Usage:    "Cherry-pick: external allowed event data ids to include from the draft. Omit or send `[]` to include all draft changes in this collection.",
			BodyPath: "includeExternalAllowedEventData",
		},
		&requestflag.Flag[any]{
			Name:     "include-global-dispatch-center",
			Usage:    "Cherry-pick: global dispatch center ids to include from the draft. Omit or send `[]` to include all draft changes in this collection.",
			BodyPath: "includeGlobalDispatchCenters",
		},
		&requestflag.Flag[any]{
			Name:     "include-mapping",
			Usage:    "Cherry-pick: mapping ids to include from the draft. Omit or send `[]` to include all draft changes in this collection.",
			BodyPath: "includeMappings",
		},
		&requestflag.Flag[any]{
			Name:     "include-replay-setting",
			Usage:    "Cherry-pick: replay settings ids to include from the draft. Omit or send `[]` to include all draft changes in this collection.",
			BodyPath: "includeReplaySettings",
		},
		&requestflag.Flag[any]{
			Name:     "include-source",
			Usage:    "Cherry-pick: source UUIDs to include from the draft. Omit or send `[]` to include all draft changes in this collection.",
			BodyPath: "includeSources",
		},
		&requestflag.Flag[any]{
			Name:     "include-tag-manager-tag",
			Usage:    "Cherry-pick: tag manager tag UUIDs to include from the draft. Omit or send `[]` to include all draft changes in this collection.",
			BodyPath: "includeTagManagerTags",
		},
		&requestflag.Flag[any]{
			Name:     "include-tag-manager-trigger",
			Usage:    "Cherry-pick: tag manager trigger UUIDs to include from the draft. Omit or send `[]` to include all draft changes in this collection.",
			BodyPath: "includeTagManagerTriggers",
		},
		&requestflag.Flag[any]{
			Name:     "include-tag-manager-variable",
			Usage:    "Cherry-pick: tag manager variable UUIDs to include from the draft. Omit or send `[]` to include all draft changes in this collection.",
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

var versionsPublish = cli.Command{
	Name:    "publish",
	Usage:   "Re-publish an existing (previously created) version by ID. Use this to roll back\nto an older snapshot. Returns 409 if the version is already published, was\ncreated more than 45 days ago, or another publish is already in flight. To\ncreate-and-publish from current draft state, use POST /rest/v1/versions instead.\nRequires scope: version:publish",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleVersionsPublish,
	HideHelpCommand: true,
}

var versionsSnapshot = cli.Command{
	Name:    "snapshot",
	Usage:   "Retrieve the full JSON snapshot captured by a version — every entity\n(destinations, sources, mappings, consent settings, etc.) as it existed when\nthis version was published. Sensitive fields (API keys, tokens, secrets) are\nredacted. Useful for IaC export, audit, and backup workflows. Requires scope:\nversion:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleVersionsSnapshot,
	HideHelpCommand: true,
}

var versionsDiff = cli.Command{
	Name:    "diff",
	Usage:   "Compare two versions of the account configuration. Returns\nadded/removed/modified entities grouped by collection, plus a total `count`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "against",
			Usage:     "Baseline version id to compare the path version against. Omit for the latest published version. Pass a version UUID to compute a version-vs-version diff.",
			QueryPath: "against",
		},
	},
	Action:          handleVersionsDiff,
	HideHelpCommand: true,
}

var versionsRevert = requestflag.WithInnerFlags(cli.Command{
	Name:    "revert",
	Usage:   "Revert one or more pending draft changes back to the latest published version\n(the \"abandon a change\" action). Send a body listing the entities to revert;\neach is handled independently — a modified entity is restored to its published\nvalues, a deleted entity is recreated, and a newly-added entity is discarded. An\nentity that already matches the published version is left untouched and is not\ncounted (a no-op). An entity present in neither the draft nor the published\nversion is reported in `notFoundCount`. `remainingChanges` is the total number\nof pending draft changes left after this call, which may stay above zero when\nother unreverted changes remain.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     `Allowed values: "draft".`,
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "entity",
			Usage:    "Draft entities to revert back to the latest published version. Each entry is processed independently: a modified entity is restored to its published values, a deleted entity is recreated, and a newly-added entity is discarded. Pass a single entry to revert one change.",
			Required: true,
			BodyPath: "entities",
		},
	},
	Action:          handleVersionsRevert,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"entity": {
		&requestflag.InnerFlag[string]{
			Name:       "entity.id",
			Usage:      "Id of the draft entity to revert.",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "entity.collection",
			Usage:      "Entity collection the id belongs to. `experiments` and `experimentVariants` are not supported — use the experiment lifecycle endpoints for those — but `experimentSettings` is.",
			InnerField: "collection",
		},
	},
})

var versionsAbandon = cli.Command{
	Name:    "abandon",
	Usage:   "Abandon ALL pending draft changes, resetting every supported entity collection\nback to the latest published version. Running experiments and experiment\nvariants are not affected (ship those via the experiment lifecycle endpoints),\nbut experiment settings are reset. Returns HTTP 409 when there is no published\nversion to revert to, or when another publish or abandon is already in flight.\nTo revert only specific changes, use `POST /rest/v1/versions/draft/revert`.\n(`draft` is a literal path segment identifying the live draft as the target.)\nRequires scope: version:publish",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     `Allowed values: "draft".`,
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleVersionsAbandon,
	HideHelpCommand: true,
}

func handleVersionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.VersionListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Versions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "versions list",
			Transform:      transform,
		})
	} else {
		iter := client.Versions.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "versions list",
			Transform:      transform,
		})
	}
}

func handleVersionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.VersionNewParams{}

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

	params := oursprivacy.VersionUpdateParams{}

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

func handleVersionsPublish(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Versions.Publish(ctx, cmd.Value("id").(string), options...)
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
		Title:          "versions publish",
		Transform:      transform,
	})
}

func handleVersionsSnapshot(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Versions.Snapshot(ctx, cmd.Value("id").(string), options...)
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
		Title:          "versions snapshot",
		Transform:      transform,
	})
}

func handleVersionsDiff(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.VersionDiffParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Versions.Diff(
		ctx,
		oursprivacy.VersionDiffParamsID(cmd.Value("id").(string)),
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
		Title:          "versions diff",
		Transform:      transform,
	})
}

func handleVersionsRevert(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.VersionRevertParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Versions.Revert(
		ctx,
		oursprivacy.VersionRevertParamsID(cmd.Value("id").(string)),
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
		Title:          "versions revert",
		Transform:      transform,
	})
}

func handleVersionsAbandon(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Versions.Abandon(ctx, oursprivacy.VersionAbandonParamsID(cmd.Value("id").(string)), options...)
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
		Title:          "versions abandon",
		Transform:      transform,
	})
}
