// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/ours-privacy-platform-cli/internal/apiquery"
	"github.com/stainless-sdks/ours-privacy-platform-cli/internal/requestflag"
	"github.com/stainless-sdks/ours-privacy-platform-go"
	"github.com/stainless-sdks/ours-privacy-platform-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var restV1ReplaySettingsCreate = cli.Command{
	Name:            "create",
	Usage:           "Create a new replay setting. Requires scope: replaySettings:create",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRestV1ReplaySettingsCreate,
	HideHelpCommand: true,
}

var restV1ReplaySettingsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single replay setting by ID. Requires scope: replaySettings:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleRestV1ReplaySettingsRetrieve,
	HideHelpCommand: true,
}

var restV1ReplaySettingsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a replay setting. Requires scope: replaySettings:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleRestV1ReplaySettingsUpdate,
	HideHelpCommand: true,
}

var restV1ReplaySettingsList = cli.Command{
	Name:            "list",
	Usage:           "List all replay settings. Requires scope: replaySettings:list",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRestV1ReplaySettingsList,
	HideHelpCommand: true,
}

var restV1ReplaySettingsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a replay setting. Requires scope: replaySettings:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleRestV1ReplaySettingsDelete,
	HideHelpCommand: true,
}

func handleRestV1ReplaySettingsCreate(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacyplatform.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := oursprivacyplatform.RestV1ReplaySettingNewParams{}

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
	_, err = client.Rest.V1.ReplaySettings.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:replay-settings create", obj, format, transform)
}

func handleRestV1ReplaySettingsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacyplatform.NewClient(getDefaultRequestOptions(cmd)...)
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
	_, err = client.Rest.V1.ReplaySettings.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:replay-settings retrieve", obj, format, transform)
}

func handleRestV1ReplaySettingsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacyplatform.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := oursprivacyplatform.RestV1ReplaySettingUpdateParams{}

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
	_, err = client.Rest.V1.ReplaySettings.Update(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:replay-settings update", obj, format, transform)
}

func handleRestV1ReplaySettingsList(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacyplatform.NewClient(getDefaultRequestOptions(cmd)...)
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
	_, err = client.Rest.V1.ReplaySettings.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:replay-settings list", obj, format, transform)
}

func handleRestV1ReplaySettingsDelete(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacyplatform.NewClient(getDefaultRequestOptions(cmd)...)
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
	_, err = client.Rest.V1.ReplaySettings.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:replay-settings delete", obj, format, transform)
}
