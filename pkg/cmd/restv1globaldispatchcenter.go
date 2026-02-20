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

var restV1GlobalDispatchCentersCreate = cli.Command{
	Name:            "create",
	Usage:           "Create a new global dispatch center. Requires scope: globalDispatch:create",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRestV1GlobalDispatchCentersCreate,
	HideHelpCommand: true,
}

var restV1GlobalDispatchCentersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single global dispatch center by ID. Requires scope: globalDispatch:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleRestV1GlobalDispatchCentersRetrieve,
	HideHelpCommand: true,
}

var restV1GlobalDispatchCentersUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a global dispatch center. Requires scope: globalDispatch:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleRestV1GlobalDispatchCentersUpdate,
	HideHelpCommand: true,
}

var restV1GlobalDispatchCentersList = cli.Command{
	Name:            "list",
	Usage:           "List all global dispatch centers. Requires scope: globalDispatch:list",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRestV1GlobalDispatchCentersList,
	HideHelpCommand: true,
}

var restV1GlobalDispatchCentersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a global dispatch center. Requires scope: globalDispatch:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleRestV1GlobalDispatchCentersDelete,
	HideHelpCommand: true,
}

func handleRestV1GlobalDispatchCentersCreate(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacyplatform.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := oursprivacyplatform.RestV1GlobalDispatchCenterNewParams{}

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
	_, err = client.Rest.V1.GlobalDispatchCenters.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:global-dispatch-centers create", obj, format, transform)
}

func handleRestV1GlobalDispatchCentersRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Rest.V1.GlobalDispatchCenters.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:global-dispatch-centers retrieve", obj, format, transform)
}

func handleRestV1GlobalDispatchCentersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacyplatform.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := oursprivacyplatform.RestV1GlobalDispatchCenterUpdateParams{}

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
	_, err = client.Rest.V1.GlobalDispatchCenters.Update(
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
	return ShowJSON(os.Stdout, "rest:v1:global-dispatch-centers update", obj, format, transform)
}

func handleRestV1GlobalDispatchCentersList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Rest.V1.GlobalDispatchCenters.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:global-dispatch-centers list", obj, format, transform)
}

func handleRestV1GlobalDispatchCentersDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Rest.V1.GlobalDispatchCenters.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rest:v1:global-dispatch-centers delete", obj, format, transform)
}
