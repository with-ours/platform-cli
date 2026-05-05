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

var replaySettingsList = cli.Command{
	Name:    "list",
	Usage:   "List the replay configurations on this account. Supports cursor pagination via\n`limit` and `cursor`. Replay settings control which domains may capture session\nreplays and where the capture script is hosted. Requires scope:\nreplaySettings:list",
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
	},
	Action:          handleReplaySettingsList,
	HideHelpCommand: true,
}

var replaySettingsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create the replay configuration for this account. Each account is limited to one\nreplay configuration — calls made when one already exists return HTTP 409 with\nthe reason in the response `error` field. Requires scope: replaySettings:create",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:     "custom-domain",
			Usage:    "Optional custom domain (CNAME) for hosting the replay capture script. Leave null to use the default Ours Privacy domain.",
			BodyPath: "customDomain",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "Human-readable label for this replay configuration. Shown in the dashboard. May be empty.",
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "status",
			Usage:    `Whether session replay capture is currently active. Set to "Enabled" to start capturing replays from whitelisted domains, or "Disabled" to pause capture without losing the configuration.`,
			BodyPath: "status",
		},
		&requestflag.Flag[any]{
			Name:     "whitelist-domain",
			Usage:    "Hostnames where session replay capture is permitted. Replays initiated from any host not in this list are dropped. PATCH replaces the list — partial updates are not merged.",
			BodyPath: "whitelistDomains",
		},
	},
	Action:          handleReplaySettingsCreate,
	HideHelpCommand: true,
}

var replaySettingsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch a single replay configuration by ID, including its whitelisted domains and\ncustom domain. Requires scope: replaySettings:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleReplaySettingsRetrieve,
	HideHelpCommand: true,
}

var replaySettingsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update one or more fields on an existing replay configuration. Only the fields\nyou send are changed; omitted fields keep their current value. Note that\n`whitelistDomains` is replaced wholesale (not merged with the existing list).\nRequires scope: replaySettings:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[*string]{
			Name:     "custom-domain",
			Usage:    "Optional custom domain (CNAME) for hosting the replay capture script. Leave null to use the default Ours Privacy domain.",
			BodyPath: "customDomain",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "Human-readable label for this replay configuration. Shown in the dashboard. May be empty.",
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "status",
			Usage:    `Whether session replay capture is currently active. Set to "Enabled" to start capturing replays from whitelisted domains, or "Disabled" to pause capture without losing the configuration.`,
			BodyPath: "status",
		},
		&requestflag.Flag[any]{
			Name:     "whitelist-domain",
			Usage:    "Hostnames where session replay capture is permitted. Replays initiated from any host not in this list are dropped. PATCH replaces the list — partial updates are not merged.",
			BodyPath: "whitelistDomains",
		},
	},
	Action:          handleReplaySettingsUpdate,
	HideHelpCommand: true,
}

var replaySettingsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete the replay configuration. Capture stops immediately for all whitelisted\ndomains. Requires scope: replaySettings:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleReplaySettingsDelete,
	HideHelpCommand: true,
}

func handleReplaySettingsList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.ReplaySettingListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ReplaySettings.List(ctx, params, options...)
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
		Title:          "replay-settings list",
		Transform:      transform,
	})
}

func handleReplaySettingsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.ReplaySettingNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ReplaySettings.New(ctx, params, options...)
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
		Title:          "replay-settings create",
		Transform:      transform,
	})
}

func handleReplaySettingsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ReplaySettings.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "replay-settings retrieve",
		Transform:      transform,
	})
}

func handleReplaySettingsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.ReplaySettingUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ReplaySettings.Update(
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
		Title:          "replay-settings update",
		Transform:      transform,
	})
}

func handleReplaySettingsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ReplaySettings.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "replay-settings delete",
		Transform:      transform,
	})
}
