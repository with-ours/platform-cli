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

var mappingTemplatesList = cli.Command{
	Name:    "list",
	Usage:   "Discover every mapping template available for a destination or source, with full\nproperty descriptors inlined. Use the returned `id` as `templateId` when calling\n`POST /rest/v1/mappings` (template fat-create variant), and use each entry under\n`mappings[]` to learn the valid `property`, `kind`, `modificationOptions`, and\nany enforced `options`. The `isDefault: true` entry is the destination's\nbuilt-in default template (the one stored at `MAPPER#{destinationId}` when\nconfigured via `PUT /rest/v1/default-mappings/{destinationId}`). Requires scope:\nmapping:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "entity-id",
			Usage:     "Destination or source id. Required.",
			Required:  true,
			QueryPath: "entityId",
		},
	},
	Action:          handleMappingTemplatesList,
	HideHelpCommand: true,
}

func handleMappingTemplatesList(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacy.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := oursprivacy.MappingTemplateListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.MappingTemplates.List(ctx, params, options...)
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
		Title:          "mapping-templates list",
		Transform:      transform,
	})
}
