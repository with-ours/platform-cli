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

var experimentVariantsList = cli.Command{
	Name:    "list",
	Usage:   "List variants for a specific parent experiment. Requires the `experimentId`\nquery parameter — variants are always scoped to a single experiment. Supports\ncursor pagination via `limit` and `cursor`; SDK runtimes that need the full set\nin one request can pass `?limit=100`. Requires scope: experiment:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "experiment-id",
			Usage:     "Required. List variants belonging to this parent experiment.",
			Required:  true,
			QueryPath: "experimentId",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor from pagination.nextCursor in the previous response. Do not decode or modify it. Malformed cursors return 400 Bad Request.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[*int64]{
			Name:      "limit",
			Usage:     "Maximum number of variants to return. Defaults to 200; values below 1 are clamped to 1 and values above 200 are clamped to 200. Variants per experiment are capped at 200 server-side, so a single request returns the full set.",
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleExperimentVariantsList,
	HideHelpCommand: true,
}

var experimentVariantsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new experiment variant. Requires scope: experiment:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "experiment-id",
			Usage:    "Parent experiment ID that will own this new variant.",
			Required: true,
			BodyPath: "experimentId",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable name for the new variant.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[int64]{
			Name:     "weight",
			Usage:    "Traffic weight for this variant as a percentage (0–100). Treatment weights are percentages of the split and must total 99% or less to leave room for the control; the control variant is the remainder (100 − Σ treatment weights, always ≥ 1%) and is maintained automatically.",
			Required: true,
			BodyPath: "weight",
		},
		&requestflag.Flag[any]{
			Name:     "dom-modification",
			Usage:    "Required for DOM modification variants. Omit for redirect variants. Each entry is `{selector, action, value}`.",
			BodyPath: "domModifications",
		},
		&requestflag.Flag[*bool]{
			Name:     "is-control",
			Usage:    "Mark this variant as the experiment control. Defaults to `false`. The API rejects the request with 409 if the experiment already has a control variant. Every experiment keeps exactly one control whose weight is the auto-derived remainder of the split, so the control cannot be cleared while it would leave the treatments at 100% (no room for a control). DELETE on the control returns 409.",
			BodyPath: "isControl",
		},
		&requestflag.Flag[*string]{
			Name:     "redirect-url",
			Usage:    "Required for redirect variants. Use either a site-relative path such as `/pricing-v2` or an absolute `https://` URL. Cross-origin `http://` URLs are rejected. Omit for DOM modification variants.",
			BodyPath: "redirectUrl",
		},
		&requestflag.Flag[*string]{
			Name:     "variant-type",
			Usage:    "Variant delivery mechanism. `dom_modifications` mutates the current page in-place at SDK runtime — use it for copy/style/image/HTML changes that keep visitors on the same URL (headline copy tests, button color, hero image swap). `redirect` routes the visitor to a different URL entirely — use it for landing-page A/B tests, alternate pricing pages, or any test where the *page itself* is the variable. They are not interchangeable: a redirect variant cannot also tweak DOM, and a dom_modifications variant cannot send the visitor elsewhere.",
			BodyPath: "variantType",
		},
	},
	Action:          handleExperimentVariantsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"dom-modification": {
		&requestflag.InnerFlag[string]{
			Name:                  "dom-modification.action",
			Usage:                 "Mutation to apply when the selector matches.",
			InnerField:            "action",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "dom-modification.selector",
			Usage:                 "CSS selector for the element to modify at runtime. PREFER specific selectors that match exactly one element: an `id` (`#hero-headline`), a stable `data-*` attribute (`[data-testid=\"hero-headline\"]`), or a unique class/structural chain (`section.hero > h1.headline`). AVOID bare tag selectors like `h1`, `button`, or `img` — modern pages usually contain several, and the runtime applies the mutation to ONLY THE FIRST match, which silently picks the wrong element. If you only have a tag name, scope it with the nearest unique ancestor (e.g. `main h1`, `header nav a:first-of-type`).",
			InnerField:            "selector",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[any]{
			Name:                  "dom-modification.attribute",
			Usage:                 "Use this for `setAttribute` to avoid JSON-stringifying `{name: value}` yourself. Ignored for other actions.",
			InnerField:            "attribute",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[any]{
			Name:                  "dom-modification.styles",
			Usage:                 "Use this for `setStyle` to avoid JSON-stringifying `{property: value}` yourself. Ignored for other actions.",
			InnerField:            "styles",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "dom-modification.value",
			Usage:                 "Canonical action payload. For `setText` / `setHtml` / `customCss` / `customJs` / `setImage` / `insertBefore` / `insertAfter` this is the literal text/HTML/CSS/JS/URL. For `setStyle` and `setAttribute` it is a JSON-stringified `{key: value}` object — or you can supply the structured `styles` / `attribute` field instead and the server will normalize.",
			InnerField:            "value",
			OuterIsArrayOfObjects: true,
		},
	},
})

var experimentVariantsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single experiment variant by ID. Requires scope: experiment:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleExperimentVariantsRetrieve,
	HideHelpCommand: true,
}

var experimentVariantsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Partially update an experiment variant. Only the fields you send are changed.\nRequires scope: experiment:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[any]{
			Name:     "dom-modification",
			Usage:    "Updated declarative DOM mutations. Sending this field replaces the prior list — partial-array merging is not supported.",
			BodyPath: "domModifications",
		},
		&requestflag.Flag[*bool]{
			Name:     "is-control",
			Usage:    "Promote or demote this variant as the control. Promoting a second variant while another already has `isControl: true` is rejected with 409. Demoting the control (`isControl: false`) is rejected when it would leave the treatments at 100% — there must always be room for a control.",
			BodyPath: "isControl",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "Updated variant name.",
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "redirect-url",
			Usage:    "Updated redirect URL for redirect variants. Use either a site-relative path such as `/pricing-v2` or an absolute `https://` URL. Cross-origin `http://` URLs are rejected.",
			BodyPath: "redirectUrl",
		},
		&requestflag.Flag[*string]{
			Name:     "variant-type",
			Usage:    "Updated variant delivery mechanism. `dom_modifications` mutates the current page in-place; `redirect` sends the visitor to a different URL — pick based on whether the *page* or the *content* is the variable. Changing this also requires updating the matching payload field (`redirectUrl` or `domModifications`).",
			BodyPath: "variantType",
		},
		&requestflag.Flag[*int64]{
			Name:     "weight",
			Usage:    "Updated traffic weight as a percentage (0–100). The control variant weight is derived from the treatments (it is the remainder of the split) and cannot be set directly.",
			BodyPath: "weight",
		},
	},
	Action:          handleExperimentVariantsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"dom-modification": {
		&requestflag.InnerFlag[string]{
			Name:                  "dom-modification.action",
			Usage:                 "Mutation to apply when the selector matches.",
			InnerField:            "action",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "dom-modification.selector",
			Usage:                 "CSS selector for the element to modify at runtime. PREFER specific selectors that match exactly one element: an `id` (`#hero-headline`), a stable `data-*` attribute (`[data-testid=\"hero-headline\"]`), or a unique class/structural chain (`section.hero > h1.headline`). AVOID bare tag selectors like `h1`, `button`, or `img` — modern pages usually contain several, and the runtime applies the mutation to ONLY THE FIRST match, which silently picks the wrong element. If you only have a tag name, scope it with the nearest unique ancestor (e.g. `main h1`, `header nav a:first-of-type`).",
			InnerField:            "selector",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[any]{
			Name:                  "dom-modification.attribute",
			Usage:                 "Use this for `setAttribute` to avoid JSON-stringifying `{name: value}` yourself. Ignored for other actions.",
			InnerField:            "attribute",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[any]{
			Name:                  "dom-modification.styles",
			Usage:                 "Use this for `setStyle` to avoid JSON-stringifying `{property: value}` yourself. Ignored for other actions.",
			InnerField:            "styles",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "dom-modification.value",
			Usage:                 "Canonical action payload. For `setText` / `setHtml` / `customCss` / `customJs` / `setImage` / `insertBefore` / `insertAfter` this is the literal text/HTML/CSS/JS/URL. For `setStyle` and `setAttribute` it is a JSON-stringified `{key: value}` object — or you can supply the structured `styles` / `attribute` field instead and the server will normalize.",
			InnerField:            "value",
			OuterIsArrayOfObjects: true,
		},
	},
})

var experimentVariantsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an experiment variant. Requires scope: experiment:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleExperimentVariantsDelete,
	HideHelpCommand: true,
}

func handleExperimentVariantsList(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.ExperimentVariantListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.ExperimentVariants.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "experiment-variants list",
			Transform:      transform,
		})
	} else {
		iter := client.ExperimentVariants.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "experiment-variants list",
			Transform:      transform,
		})
	}
}

func handleExperimentVariantsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.ExperimentVariantNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ExperimentVariants.New(ctx, params, options...)
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
		Title:          "experiment-variants create",
		Transform:      transform,
	})
}

func handleExperimentVariantsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ExperimentVariants.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "experiment-variants retrieve",
		Transform:      transform,
	})
}

func handleExperimentVariantsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.ExperimentVariantUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ExperimentVariants.Update(
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
		Title:          "experiment-variants update",
		Transform:      transform,
	})
}

func handleExperimentVariantsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ExperimentVariants.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "experiment-variants delete",
		Transform:      transform,
	})
}
