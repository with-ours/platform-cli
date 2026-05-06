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

var experimentSettingsList = cli.Command{
	Name:            "list",
	Usage:           "List experiment settings records for the account. Use the returned `id` as\n`experimentSettingsId` when creating an experiment. Requires scope:\nexperimentSettings:list",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleExperimentSettingsList,
	HideHelpCommand: true,
}

var experimentSettingsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create the account-level experimentation bootstrap record. Most accounts should\nonly ever have one. Requires scope: experimentSettings:create",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:     "cookie-name",
			Usage:    "Cookie name used to persist sticky variant assignments in the browser. Defaults to `_cord_exp` when omitted on create.",
			BodyPath: "cookieName",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "Human-readable name for this experimentation configuration. Defaults to `Experiment Settings` when omitted on create.",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "whitelist-domain",
			Usage:    "Optional domain allowlist for experiment SDK delivery. When set, experiments using this settings record are only served on these domains. This is separate from source `whitelistDomains`, which gates CDP event ingestion.",
			BodyPath: "whitelistDomains",
		},
	},
	Action:          handleExperimentSettingsCreate,
	HideHelpCommand: true,
}

var experimentSettingsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single experiment settings record by ID. Returns 404 when no record\nmatches the supplied id. Requires scope: experimentSettings:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleExperimentSettingsRetrieve,
	HideHelpCommand: true,
}

var experimentSettingsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Partially update an experiment settings. Only the fields you send are changed.\nRequires scope: experimentSettings:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[*string]{
			Name:     "cookie-name",
			Usage:    "Cookie name used to persist sticky variant assignments in the browser. Defaults to `_cord_exp` when omitted on create.",
			BodyPath: "cookieName",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "Human-readable name for this experimentation configuration. Defaults to `Experiment Settings` when omitted on create.",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "whitelist-domain",
			Usage:    "Optional domain allowlist for experiment SDK delivery. When set, experiments using this settings record are only served on these domains. This is separate from source `whitelistDomains`, which gates CDP event ingestion.",
			BodyPath: "whitelistDomains",
		},
	},
	Action:          handleExperimentSettingsUpdate,
	HideHelpCommand: true,
}

var experimentSettingsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete the experimentation bootstrap record. This also deletes child\nexperiments, variants, and personalization properties owned by it. Requires\nscope: experimentSettings:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleExperimentSettingsDelete,
	HideHelpCommand: true,
}

func handleExperimentSettingsList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ExperimentSettings.List(ctx, options...)
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
		Title:          "experiment-settings list",
		Transform:      transform,
	})
}

func handleExperimentSettingsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.ExperimentSettingNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ExperimentSettings.New(ctx, params, options...)
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
		Title:          "experiment-settings create",
		Transform:      transform,
	})
}

func handleExperimentSettingsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ExperimentSettings.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "experiment-settings retrieve",
		Transform:      transform,
	})
}

func handleExperimentSettingsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.ExperimentSettingUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ExperimentSettings.Update(
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
		Title:          "experiment-settings update",
		Transform:      transform,
	})
}

func handleExperimentSettingsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ExperimentSettings.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "experiment-settings delete",
		Transform:      transform,
	})
}
