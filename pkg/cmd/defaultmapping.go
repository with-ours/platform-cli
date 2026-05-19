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

var defaultMappingsList = cli.Command{
	Name:            "list",
	Usage:           "List every stored default mapping for the account, one per destination that has\never written a default. Destinations that have not yet written a default mapping\ndo not appear here. Use `GET /rest/v1/default-mappings/{destinationId}` to fetch\nthe hydrated would-be row for a specific destination. Requires scope:\nmapping:list",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleDefaultMappingsList,
	HideHelpCommand: true,
}

var defaultMappingsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch the destination's default mapping by destination id. Returns a hydrated\nrow with empty `mappings[]` when no default mapping has been written yet (so\ncallers do not need to handle a 404-vs-200 branch). Each destination has at most\none default mapping. Requires scope: mapping:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleDefaultMappingsRetrieve,
	HideHelpCommand: true,
}

var defaultMappingsReplace = requestflag.WithInnerFlags(cli.Command{
	Name:    "replace",
	Usage:   "Upsert the destination default mapping. Always replaces `mappings[]` wholesale\n(default mappings have no merge-partial semantic). Default mappings cannot have\ncustom `logic`; the field is not accepted on this endpoint. Requires scope:\nmapping:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "mapping",
			Usage:    "Property mappings to persist as the destination default. Use `GET /rest/v1/mapping-templates?entityId={destinationId}` to discover the valid `property` values.",
			Required: true,
			BodyPath: "mappings",
		},
		&requestflag.Flag[*bool]{
			Name:     "is-enabled",
			Usage:    "Toggle the default mapping on/off. Defaults to `true` when omitted. `null` is treated as omitted.",
			BodyPath: "isEnabled",
		},
	},
	Action:          handleDefaultMappingsReplace,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"mapping": {
		&requestflag.InnerFlag[string]{
			Name:       "mapping.map",
			InnerField: "map",
		},
		&requestflag.InnerFlag[string]{
			Name:       "mapping.property",
			InnerField: "property",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "mapping.modification",
			Usage:      `Allowed values: "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs", "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash", "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender", "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone", "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null", "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase".`,
			InnerField: "modification",
		},
	},
})

func handleDefaultMappingsList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.DefaultMappings.List(ctx, options...)
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
		Title:          "default-mappings list",
		Transform:      transform,
	})
}

func handleDefaultMappingsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacy.NewClient(getDefaultRequestOptions(cmd)...)
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
	_, err = client.DefaultMappings.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "default-mappings retrieve",
		Transform:      transform,
	})
}

func handleDefaultMappingsReplace(ctx context.Context, cmd *cli.Command) error {
	client := oursprivacy.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := oursprivacy.DefaultMappingReplaceParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.DefaultMappings.Replace(
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
		Title:          "default-mappings replace",
		Transform:      transform,
	})
}
