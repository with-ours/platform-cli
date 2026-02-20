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

var restV1SourcesCreate = cli.Command{
	Name:            "create",
	Usage:           "Create a new source. Requires scope: source:create",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRestV1SourcesCreate,
	HideHelpCommand: true,
}

var restV1SourcesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single source by ID. Requires scope: source:view",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleRestV1SourcesRetrieve,
	HideHelpCommand: true,
}

var restV1SourcesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a source. Requires scope: source:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleRestV1SourcesUpdate,
	HideHelpCommand: true,
}

var restV1SourcesList = cli.Command{
	Name:            "list",
	Usage:           "List all sources. Requires scope: source:list",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRestV1SourcesList,
	HideHelpCommand: true,
}

var restV1SourcesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a source. Requires scope: source:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleRestV1SourcesDelete,
	HideHelpCommand: true,
}

func handleRestV1SourcesCreate(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacyplatform.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := oursprivacyplatform.RestV1SourceNewParams{}

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
	_, err = client.Rest.V1.Sources.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:sources create", obj, format, transform)
}

func handleRestV1SourcesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Rest.V1.Sources.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:sources retrieve", obj, format, transform)
}

func handleRestV1SourcesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacyplatform.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := oursprivacyplatform.RestV1SourceUpdateParams{}

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
	_, err = client.Rest.V1.Sources.Update(
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
	return ShowJSON(os.Stdout, "rest:v1:sources update", obj, format, transform)
}

func handleRestV1SourcesList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Rest.V1.Sources.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:sources list", obj, format, transform)
}

func handleRestV1SourcesDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Rest.V1.Sources.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:sources delete", obj, format, transform)
}
