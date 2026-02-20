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

var restV1ConsentSettingsCreate = cli.Command{
	Name:            "create",
	Usage:           "Create a new consent setting. Requires scope: consentSettings:create",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRestV1ConsentSettingsCreate,
	HideHelpCommand: true,
}

var restV1ConsentSettingsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single consent setting by ID. Requires scope: consentSettings:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleRestV1ConsentSettingsRetrieve,
	HideHelpCommand: true,
}

var restV1ConsentSettingsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a consent setting. Requires scope: consentSettings:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleRestV1ConsentSettingsUpdate,
	HideHelpCommand: true,
}

var restV1ConsentSettingsList = cli.Command{
	Name:            "list",
	Usage:           "List all consent settings. Requires scope: consentSettings:list",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRestV1ConsentSettingsList,
	HideHelpCommand: true,
}

var restV1ConsentSettingsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a consent setting. Requires scope: consentSettings:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleRestV1ConsentSettingsDelete,
	HideHelpCommand: true,
}

func handleRestV1ConsentSettingsCreate(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacyplatform.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := oursprivacyplatform.RestV1ConsentSettingNewParams{}

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
	_, err = client.Rest.V1.ConsentSettings.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:consent-settings create", obj, format, transform)
}

func handleRestV1ConsentSettingsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Rest.V1.ConsentSettings.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:consent-settings retrieve", obj, format, transform)
}

func handleRestV1ConsentSettingsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacyplatform.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := oursprivacyplatform.RestV1ConsentSettingUpdateParams{}

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
	_, err = client.Rest.V1.ConsentSettings.Update(
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
	return ShowJSON(os.Stdout, "rest:v1:consent-settings update", obj, format, transform)
}

func handleRestV1ConsentSettingsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Rest.V1.ConsentSettings.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:consent-settings list", obj, format, transform)
}

func handleRestV1ConsentSettingsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Rest.V1.ConsentSettings.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:consent-settings delete", obj, format, transform)
}
