// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
	"github.com/with-ours/platform-cli/internal/apiquery"
	"github.com/with-ours/platform-cli/internal/requestflag"
	"github.com/with-ours/platform-sdk-go"
	"github.com/with-ours/platform-sdk-go/option"
)

var consentSettingsCreate = cli.Command{
	Name:            "create",
	Usage:           "Create a new consent setting. Requires scope: consentSettings:create",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleConsentSettingsCreate,
	HideHelpCommand: true,
}

var consentSettingsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single consent setting by ID. Requires scope: consentSettings:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleConsentSettingsRetrieve,
	HideHelpCommand: true,
}

var consentSettingsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a consent setting. Requires scope: consentSettings:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[[]any]{
			Name:     "category",
			Required: true,
			BodyPath: "categories",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]any]{
			Name:     "region",
			Required: true,
			BodyPath: "regions",
		},
		&requestflag.Flag[[]any]{
			Name:     "service",
			Required: true,
			BodyPath: "services",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Required: true,
			BodyPath: "status",
		},
		&requestflag.Flag[any]{
			Name:     "consent-cookie-name",
			BodyPath: "consentCookieName",
		},
		&requestflag.Flag[any]{
			Name:     "custom-domain",
			BodyPath: "customDomain",
		},
		&requestflag.Flag[any]{
			Name:     "default",
			BodyPath: "default",
		},
		&requestflag.Flag[any]{
			Name:     "revision",
			BodyPath: "revision",
		},
		&requestflag.Flag[any]{
			Name:     "skip-blocking-class-name",
			BodyPath: "skipBlockingClassNames",
		},
		&requestflag.Flag[any]{
			Name:     "web-sdk-token",
			BodyPath: "webSDKToken",
		},
		&requestflag.Flag[any]{
			Name:     "whitelist-domain",
			BodyPath: "whitelistDomains",
		},
	},
	Action:          handleConsentSettingsUpdate,
	HideHelpCommand: true,
}

var consentSettingsList = cli.Command{
	Name:            "list",
	Usage:           "List all consent settings. Requires scope: consentSettings:list",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleConsentSettingsList,
	HideHelpCommand: true,
}

var consentSettingsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a consent setting. Requires scope: consentSettings:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleConsentSettingsDelete,
	HideHelpCommand: true,
}

func handleConsentSettingsCreate(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ConsentSettings.New(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "consent-settings create", obj, format, transform)
}

func handleConsentSettingsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ConsentSettings.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "consent-settings retrieve", obj, format, transform)
}

func handleConsentSettingsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomwithoursplatformsdkgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomwithoursplatformsdkgo.ConsentSettingUpdateParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ConsentSettings.Update(
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
	return ShowJSON(os.Stdout, "consent-settings update", obj, format, transform)
}

func handleConsentSettingsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ConsentSettings.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "consent-settings list", obj, format, transform)
}

func handleConsentSettingsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ConsentSettings.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "consent-settings delete", obj, format, transform)
}
