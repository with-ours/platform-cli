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

var restV1VersionsCreate = cli.Command{
	Name:            "create",
	Usage:           "Create a new version. Requires scope: version:publish",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRestV1VersionsCreate,
	HideHelpCommand: true,
}

var restV1VersionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single version by ID. Requires scope: version:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleRestV1VersionsRetrieve,
	HideHelpCommand: true,
}

var restV1VersionsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a version. Requires scope: version:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleRestV1VersionsUpdate,
	HideHelpCommand: true,
}

var restV1VersionsList = cli.Command{
	Name:            "list",
	Usage:           "List all versions. Requires scope: version:list",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRestV1VersionsList,
	HideHelpCommand: true,
}

func handleRestV1VersionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacyplatform.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := oursprivacyplatform.RestV1VersionNewParams{}

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
	_, err = client.Rest.V1.Versions.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:versions create", obj, format, transform)
}

func handleRestV1VersionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Rest.V1.Versions.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:versions retrieve", obj, format, transform)
}

func handleRestV1VersionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacyplatform.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := oursprivacyplatform.RestV1VersionUpdateParams{}

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
	_, err = client.Rest.V1.Versions.Update(
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
	return ShowJSON(os.Stdout, "rest:v1:versions update", obj, format, transform)
}

func handleRestV1VersionsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Rest.V1.Versions.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:versions list", obj, format, transform)
}
