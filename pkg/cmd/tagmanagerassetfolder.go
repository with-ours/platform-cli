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

var tagManagerAssetFoldersCreate = cli.Command{
	Name:    "create",
	Usage:   "Assign a tag, trigger, or variable to a folder within its tag manager, or send\n`folderId: null` to remove the asset from its current folder. The assignment is\na full replace — calling it again with a different `folderId` silently moves the\nasset. Requires scope: tagManagers:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "asset-id",
			Usage:    "UUID of the tag, trigger, or variable to assign.",
			Required: true,
			BodyPath: "assetId",
		},
		&requestflag.Flag[string]{
			Name:     "asset-type",
			Usage:    "Asset type to assign. Must be one of `tagManagerTag`, `tagManagerTrigger`, or `tagManagerVariable`.",
			Required: true,
			BodyPath: "assetType",
		},
		&requestflag.Flag[*string]{
			Name:     "folder-id",
			Usage:    "Folder UUID to assign to. Send `null` to remove the asset from its current folder.",
			Required: true,
			BodyPath: "folderId",
		},
	},
	Action:          handleTagManagerAssetFoldersCreate,
	HideHelpCommand: true,
}

func handleTagManagerAssetFoldersCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.TagManagerAssetFolderNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TagManagerAssetFolders.New(ctx, params, options...)
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
		Title:          "tag-manager-asset-folders create",
		Transform:      transform,
	})
}
