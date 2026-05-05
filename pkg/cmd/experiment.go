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

var experimentsList = cli.Command{
	Name:    "list",
	Usage:   "List experiments for this account. Supports cursor pagination and filtering by\n`status`, `type`, and free-text `search` matched against experiment id, name,\nand description. Combine filters with AND semantics. Requires scope:\nexperiment:list",
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
			Name:      "search",
			Usage:     "Optional case-insensitive text search matched against experiment ID, name, and description.",
			QueryPath: "search",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Optional lifecycle-state filter.",
			QueryPath: "status",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "Optional experiment-type filter.",
			QueryPath: "type",
		},
	},
	Action:          handleExperimentsList,
	HideHelpCommand: true,
}

var experimentsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new experiment. Requires scope: experiment:create",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "experiment-settings-id",
			Usage:    "ID of the experiment settings record that owns this experiment. Call `GET /rest/v1/experiment-settings` first if you need to discover the correct ID for the current account.",
			Required: true,
			BodyPath: "experimentSettingsId",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Short experiment name.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			Usage:    "Optional hypothesis or operator note.",
			BodyPath: "description",
		},
		&requestflag.Flag[*bool]{
			Name:     "include-query-string",
			Usage:    "Whether redirect variants in this experiment should preserve the original request query string.",
			BodyPath: "includeQueryString",
		},
		&requestflag.Flag[*string]{
			Name:     "key",
			Usage:    "Optional stable code-facing key. When omitted, the API slugifies the name automatically.",
			BodyPath: "key",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metrics",
			Usage:    "Goal events. If you send `metrics.primary`, `metrics.primary.eventName` must be a non-blank string. A primary event name is required before the experiment can be started.",
			BodyPath: "metrics",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "targeting-rules",
			Usage:    "Eligibility rules — URL patterns, audience, visitor status, query-param conditions. Omit to inherit defaults.",
			BodyPath: "targetingRules",
		},
		&requestflag.Flag[*float64]{
			Name:     "traffic-allocation",
			Usage:    "Initial traffic allocation percentage from 0 to 100.",
			BodyPath: "trafficAllocation",
		},
		&requestflag.Flag[*string]{
			Name:     "type",
			Usage:    "Experiment mode to create. `ab` and `multivariate` use traffic allocation and results; `personalization` is always-on targeting. Omit to create a standard `ab` experiment.",
			BodyPath: "type",
		},
	},
	Action:          handleExperimentsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"metrics": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "metrics.primary",
			Usage:      "Primary success metric. When provided, `eventName` must be a non-blank string.",
			InnerField: "primary",
		},
		&requestflag.InnerFlag[any]{
			Name:       "metrics.secondary",
			Usage:      "Optional secondary metrics tracked alongside the primary goal.",
			InnerField: "secondary",
		},
	},
	"targeting-rules": {
		&requestflag.InnerFlag[[]string]{
			Name:       "targeting-rules.url-patterns",
			Usage:      "Glob-style URL patterns that must match for the experiment to be eligible. Up to 200 patterns; each pattern up to 2000 characters. An empty array (or omitting the field) matches all URLs — equivalent to `[\"**\"]`.",
			InnerField: "urlPatterns",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "targeting-rules.audience-id",
			Usage:      "Optional audience identifier used for server-side eligibility filtering.",
			InnerField: "audienceId",
		},
		&requestflag.InnerFlag[any]{
			Name:       "targeting-rules.query-params",
			Usage:      "Additional query-string conditions that must all match for the visitor to qualify.",
			InnerField: "queryParams",
		},
		&requestflag.InnerFlag[any]{
			Name:       "targeting-rules.visitor-properties",
			Usage:      "Optional visitor-property matching rules. These are passed through as JSON for experimentation targeting.",
			InnerField: "visitorProperties",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "targeting-rules.visitor-status",
			Usage:      "Whether the experiment should target new visitors, returning visitors, or any visitor.",
			InnerField: "visitorStatus",
		},
	},
})

var experimentsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single experiment by ID. Requires scope: experiment:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleExperimentsRetrieve,
	HideHelpCommand: true,
}

var experimentsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Partially update an experiment. Only the fields you send are changed. Edits are\nonly accepted while the experiment is in `draft` status — running, paused, and\ncompleted experiments return 409 with\n`Experiment can only be edited in draft status`. Use the lifecycle endpoints\n(`/start`, `/pause`, `/resume`, `/stop`) to change status. Requires scope:\nexperiment:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			Usage:    "Updated experiment hypothesis or operator note.",
			BodyPath: "description",
		},
		&requestflag.Flag[*bool]{
			Name:     "include-query-string",
			Usage:    "Updated redirect query-string forwarding behavior for the experiment.",
			BodyPath: "includeQueryString",
		},
		&requestflag.Flag[*string]{
			Name:     "key",
			Usage:    "Updated stable code-facing key. When blank, the API falls back to a slugified key derived from the current experiment name.",
			BodyPath: "key",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metrics",
			Usage:    "Updated goal events. Send the full nested object — replaces the previous value, not merged. If you send `metrics.primary`, `metrics.primary.eventName` must be a non-blank string.",
			BodyPath: "metrics",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "Updated experiment name.",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "targeting-rules",
			Usage:    "Updated eligibility rules. Send the full nested object — replaces the previous value, not merged.",
			BodyPath: "targetingRules",
		},
		&requestflag.Flag[*float64]{
			Name:     "traffic-allocation",
			Usage:    "Updated traffic allocation percentage from 0 to 100.",
			BodyPath: "trafficAllocation",
		},
	},
	Action:          handleExperimentsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"metrics": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "metrics.primary",
			Usage:      "Primary success metric. When provided, `eventName` must be a non-blank string.",
			InnerField: "primary",
		},
		&requestflag.InnerFlag[any]{
			Name:       "metrics.secondary",
			Usage:      "Optional secondary metrics tracked alongside the primary goal.",
			InnerField: "secondary",
		},
	},
	"targeting-rules": {
		&requestflag.InnerFlag[[]string]{
			Name:       "targeting-rules.url-patterns",
			Usage:      "Glob-style URL patterns that must match for the experiment to be eligible. Up to 200 patterns; each pattern up to 2000 characters. An empty array (or omitting the field) matches all URLs — equivalent to `[\"**\"]`.",
			InnerField: "urlPatterns",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "targeting-rules.audience-id",
			Usage:      "Optional audience identifier used for server-side eligibility filtering.",
			InnerField: "audienceId",
		},
		&requestflag.InnerFlag[any]{
			Name:       "targeting-rules.query-params",
			Usage:      "Additional query-string conditions that must all match for the visitor to qualify.",
			InnerField: "queryParams",
		},
		&requestflag.InnerFlag[any]{
			Name:       "targeting-rules.visitor-properties",
			Usage:      "Optional visitor-property matching rules. These are passed through as JSON for experimentation targeting.",
			InnerField: "visitorProperties",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "targeting-rules.visitor-status",
			Usage:      "Whether the experiment should target new visitors, returning visitors, or any visitor.",
			InnerField: "visitorStatus",
		},
	},
})

var experimentsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an experiment. Requires scope: experiment:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleExperimentsDelete,
	HideHelpCommand: true,
}

var experimentsStart = cli.Command{
	Name:    "start",
	Usage:   "Start an experiment. The request body is optional — send `{}` to use defaults.\nRequires scope: experiment:start",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[bool]{
			Name:     "publish-after-start",
			Usage:    "When true (default on the REST surface), publish the current draft version immediately after starting the experiment. Any other unpublished changes in the same account version are included in that publish. Pass `false` explicitly to stage the change without publishing; the response will report `pending_publish`.",
			BodyPath: "publishAfterStart",
		},
	},
	Action:          handleExperimentsStart,
	HideHelpCommand: true,
}

var experimentsStop = cli.Command{
	Name:    "stop",
	Usage:   "Stop an experiment. The request body is optional — send `{}` to stop without\nrecording a winner. Requires scope: experiment:stop",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "winner-variant-id",
			Usage:    "Optional winning variant ID to persist when completing the experiment.",
			BodyPath: "winnerVariantId",
		},
	},
	Action:          handleExperimentsStop,
	HideHelpCommand: true,
}

var experimentsPause = cli.Command{
	Name:    "pause",
	Usage:   "Pause a running experiment. Stops new variant assignments while preserving\nexisting ones; the experiment can later be resumed. The request body is\noptional. Requires scope: experiment:stop",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[bool]{
			Name:     "publish-after-pause",
			Usage:    "When true (default on the REST surface), publish the current draft version immediately after pausing the experiment. Any other unpublished changes in the same account version are included in that publish. Pass `false` explicitly to stage the change without publishing; the response will report `pending_publish`.",
			BodyPath: "publishAfterPause",
		},
	},
	Action:          handleExperimentsPause,
	HideHelpCommand: true,
}

var experimentsResume = cli.Command{
	Name:    "resume",
	Usage:   "Resume a previously-paused experiment so new visitors can be assigned again. The\nrequest body is optional. Requires scope: experiment:start",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[bool]{
			Name:     "publish-after-resume",
			Usage:    "When true (default on the REST surface), publish the current draft version immediately after resuming the experiment. Any other unpublished changes in the same account version are included in that publish. Pass `false` explicitly to stage the change without publishing; the response will report `pending_publish`.",
			BodyPath: "publishAfterResume",
		},
	},
	Action:          handleExperimentsResume,
	HideHelpCommand: true,
}

var experimentsResults = cli.Command{
	Name:    "results",
	Usage:   "Aggregate per-variant impressions, conversions, conversion rate, and Bayesian\nprobability-to-be-best across the experiment runtime window. Requires scope:\nexperiment:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "event-name",
			Usage:     "Optional override for the conversion event name. When omitted, the experiment primary metric event is used.",
			QueryPath: "eventName",
		},
	},
	Action:          handleExperimentsResults,
	HideHelpCommand: true,
}

var experimentsResultsTimeSeries = cli.Command{
	Name:    "results-time-series",
	Usage:   "Per-day per-variant impressions, conversions, and conversion rate, sliced to a\ndate range. Use this to chart trends, compare windows, or zoom in on a specific\nperiod. Pass `startDate` / `endDate` (`YYYY-MM-DD`, UTC, both inclusive) to set\nthe window; both default to the full experiment runtime when omitted, so the\nno-arg call returns every day from start to today (or to `stoppedAt` for\ncompleted experiments). The response orders days oldest-first and omits days\nwith no impressions, so an empty `days` array means there was no measured\ntraffic in the window. Requires scope: experiment:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "end-date",
			Usage:     "Inclusive upper bound of the response window, as a UTC calendar day in `YYYY-MM-DD` format. Defaults to the experiment stop date for completed experiments, or today for running experiments. Values after that are silently clamped. The window between `startDate` and `endDate` must be 366 days or fewer.",
			QueryPath: "endDate",
		},
		&requestflag.Flag[string]{
			Name:      "event-name",
			Usage:     "Optional override for the conversion event name. When omitted, the experiment primary metric event is used.",
			QueryPath: "eventName",
		},
		&requestflag.Flag[string]{
			Name:      "start-date",
			Usage:     "Inclusive lower bound of the response window, as a UTC calendar day in `YYYY-MM-DD` format. Defaults to the experiment start date when omitted. Values before the experiment started are silently clamped to the experiment start.",
			QueryPath: "startDate",
		},
	},
	Action:          handleExperimentsResultsTimeSeries,
	HideHelpCommand: true,
}

func handleExperimentsList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.ExperimentListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.List(ctx, params, options...)
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
		Title:          "experiments list",
		Transform:      transform,
	})
}

func handleExperimentsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.ExperimentNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.New(ctx, params, options...)
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
		Title:          "experiments create",
		Transform:      transform,
	})
}

func handleExperimentsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Experiments.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "experiments retrieve",
		Transform:      transform,
	})
}

func handleExperimentsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.ExperimentUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.Update(
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
		Title:          "experiments update",
		Transform:      transform,
	})
}

func handleExperimentsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Experiments.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "experiments delete",
		Transform:      transform,
	})
}

func handleExperimentsStart(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.ExperimentStartParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.Start(
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
		Title:          "experiments start",
		Transform:      transform,
	})
}

func handleExperimentsStop(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.ExperimentStopParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.Stop(
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
		Title:          "experiments stop",
		Transform:      transform,
	})
}

func handleExperimentsPause(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.ExperimentPauseParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.Pause(
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
		Title:          "experiments pause",
		Transform:      transform,
	})
}

func handleExperimentsResume(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.ExperimentResumeParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.Resume(
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
		Title:          "experiments resume",
		Transform:      transform,
	})
}

func handleExperimentsResults(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.ExperimentResultsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.Results(
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
		Title:          "experiments results",
		Transform:      transform,
	})
}

func handleExperimentsResultsTimeSeries(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.ExperimentResultsTimeSeriesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.ResultsTimeSeries(
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
		Title:          "experiments results-time-series",
		Transform:      transform,
	})
}
