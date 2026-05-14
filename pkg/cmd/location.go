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

var locationsList = cli.Command{
	Name:            "list",
	Usage:           "List every location for this account. Not paginated — each account has a small\nmap-count limit (single digits in practice) so the response always fits in a\nsingle page. Requires scope: maps:list",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleLocationsList,
	HideHelpCommand: true,
}

var locationsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new location (map embed). All address fields are optional and can be\nfilled in later via PATCH. Returns the slim entity with the server-assigned `id`\nso callers can immediately request `GET /rest/v1/locations/{id}/embed-code`.\nRequires scope: maps:create",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "additional-address",
			BodyPath: "additionalAddresses",
		},
		&requestflag.Flag[any]{
			Name:     "center",
			BodyPath: "center",
		},
		&requestflag.Flag[*string]{
			Name:     "city",
			BodyPath: "city",
		},
		&requestflag.Flag[*string]{
			Name:     "country",
			BodyPath: "country",
		},
		&requestflag.Flag[*string]{
			Name:     "custom-domain",
			BodyPath: "customDomain",
		},
		&requestflag.Flag[*float64]{
			Name:     "latitude",
			BodyPath: "latitude",
		},
		&requestflag.Flag[*string]{
			Name:     "line1",
			BodyPath: "line1",
		},
		&requestflag.Flag[*string]{
			Name:     "line2",
			BodyPath: "line2",
		},
		&requestflag.Flag[*float64]{
			Name:     "longitude",
			BodyPath: "longitude",
		},
		&requestflag.Flag[*string]{
			Name:     "map-name",
			BodyPath: "mapName",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "phone-number",
			BodyPath: "phoneNumber",
		},
		&requestflag.Flag[*string]{
			Name:     "state",
			BodyPath: "state",
		},
		&requestflag.Flag[*string]{
			Name:     "website-link-text",
			BodyPath: "websiteLinkText",
		},
		&requestflag.Flag[*string]{
			Name:     "website-url",
			BodyPath: "websiteUrl",
		},
		&requestflag.Flag[*string]{
			Name:     "zip",
			BodyPath: "zip",
		},
	},
	Action:          handleLocationsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"additional-address": {
		&requestflag.InnerFlag[float64]{
			Name:                  "additional-address.latitude",
			InnerField:            "latitude",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[float64]{
			Name:                  "additional-address.longitude",
			InnerField:            "longitude",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.city",
			InnerField:            "city",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.country",
			InnerField:            "country",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.line1",
			InnerField:            "line1",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.line2",
			InnerField:            "line2",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.name",
			InnerField:            "name",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.phone-number",
			InnerField:            "phoneNumber",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.state",
			InnerField:            "state",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.website-link-text",
			InnerField:            "websiteLinkText",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.website-url",
			InnerField:            "websiteUrl",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.zip",
			InnerField:            "zip",
			OuterIsArrayOfObjects: true,
		},
	},
})

var locationsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Partially update a location. Only the fields you send are changed.\n`additionalAddresses` is replaced wholesale when sent — partial item updates are\nnot merged. The map's computed center is recalculated on every PATCH from the\nlatest coordinates. Requires scope: maps:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[any]{
			Name:     "additional-address",
			BodyPath: "additionalAddresses",
		},
		&requestflag.Flag[any]{
			Name:     "center",
			BodyPath: "center",
		},
		&requestflag.Flag[*string]{
			Name:     "city",
			BodyPath: "city",
		},
		&requestflag.Flag[*string]{
			Name:     "country",
			BodyPath: "country",
		},
		&requestflag.Flag[*string]{
			Name:     "custom-domain",
			BodyPath: "customDomain",
		},
		&requestflag.Flag[*float64]{
			Name:     "latitude",
			BodyPath: "latitude",
		},
		&requestflag.Flag[*string]{
			Name:     "line1",
			BodyPath: "line1",
		},
		&requestflag.Flag[*string]{
			Name:     "line2",
			BodyPath: "line2",
		},
		&requestflag.Flag[*float64]{
			Name:     "longitude",
			BodyPath: "longitude",
		},
		&requestflag.Flag[*string]{
			Name:     "map-name",
			BodyPath: "mapName",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "phone-number",
			BodyPath: "phoneNumber",
		},
		&requestflag.Flag[*string]{
			Name:     "state",
			BodyPath: "state",
		},
		&requestflag.Flag[*string]{
			Name:     "website-link-text",
			BodyPath: "websiteLinkText",
		},
		&requestflag.Flag[*string]{
			Name:     "website-url",
			BodyPath: "websiteUrl",
		},
		&requestflag.Flag[*string]{
			Name:     "zip",
			BodyPath: "zip",
		},
	},
	Action:          handleLocationsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"additional-address": {
		&requestflag.InnerFlag[float64]{
			Name:                  "additional-address.latitude",
			InnerField:            "latitude",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[float64]{
			Name:                  "additional-address.longitude",
			InnerField:            "longitude",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.city",
			InnerField:            "city",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.country",
			InnerField:            "country",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.line1",
			InnerField:            "line1",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.line2",
			InnerField:            "line2",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.name",
			InnerField:            "name",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.phone-number",
			InnerField:            "phoneNumber",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.state",
			InnerField:            "state",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.website-link-text",
			InnerField:            "websiteLinkText",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.website-url",
			InnerField:            "websiteUrl",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "additional-address.zip",
			InnerField:            "zip",
			OuterIsArrayOfObjects: true,
		},
	},
})

var locationsEmbedCode = cli.Command{
	Name:    "embed-code",
	Usage:   "Generate the paste-ready HTML embed snippet for a location. The response is a\nsingle self-contained HTML string (a `<style>` block + `<div>` wrapping an\n`<iframe>` pointed at the maps CDN for the current stage, plus an optional\nJSON-LD `<script>`). Customize the render with the optional query params\n(`color`, `theme`, `colorScheme`, `mapStyle`, `includeAddressBox`, `zoom`,\n`includeControls`, `includeSEOSchema`); all have sane defaults. Requires scope:\nmaps:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "color",
			Usage:     "Brand color used in the embedded map UI. Any CSS color string.",
			Default:   "#B33E3E",
			QueryPath: "color",
		},
		&requestflag.Flag[string]{
			Name:      "color-scheme",
			Usage:     "Light or dark color scheme.",
			Default:   "light",
			QueryPath: "colorScheme",
		},
		&requestflag.Flag[bool]{
			Name:      "include-address-box",
			Usage:     "Render the address sidebar overlay next to the map. Send `false` to hide.",
			QueryPath: "includeAddressBox",
		},
		&requestflag.Flag[string]{
			Name:      "include-controls",
			Usage:     "Whether the embed renders map controls. `accessible` enables keyboard-navigable controls.",
			Default:   "yes",
			QueryPath: "includeControls",
		},
		&requestflag.Flag[bool]{
			Name:      "include-seo-schema",
			Usage:     "Emit a `schema.org` Place JSON-LD block alongside the iframe so search engines can index the location.",
			QueryPath: "includeSEOSchema",
		},
		&requestflag.Flag[string]{
			Name:      "map-style",
			Usage:     "Base map style.",
			Default:   "Monochrome",
			QueryPath: "mapStyle",
		},
		&requestflag.Flag[string]{
			Name:      "theme",
			Usage:     "Visual theme variant.",
			Default:   "default",
			QueryPath: "theme",
		},
		&requestflag.Flag[int64]{
			Name:      "zoom",
			Usage:     "Initial map zoom level (Google-style 1–20).",
			Default:   11,
			QueryPath: "zoom",
		},
	},
	Action:          handleLocationsEmbedCode,
	HideHelpCommand: true,
}

func handleLocationsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Locations.List(ctx, options...)
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
		Title:          "locations list",
		Transform:      transform,
	})
}

func handleLocationsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.LocationNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Locations.New(ctx, params, options...)
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
		Title:          "locations create",
		Transform:      transform,
	})
}

func handleLocationsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.LocationUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Locations.Update(
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
		Title:          "locations update",
		Transform:      transform,
	})
}

func handleLocationsEmbedCode(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.LocationEmbedCodeParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Locations.EmbedCode(
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
		Title:          "locations embed-code",
		Transform:      transform,
	})
}
