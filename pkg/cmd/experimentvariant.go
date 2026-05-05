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
	Usage:   "List variants for a specific parent experiment. Requires the `experimentId`\nquery parameter — variants are always scoped to a single experiment. Requires\nscope: experiment:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "experiment-id",
			Usage:     "Required. List variants belonging to this parent experiment.",
			Required:  true,
			QueryPath: "experimentId",
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
			Usage:    "Traffic weight to assign to this variant. Weights are relative shares; the runtime normalizes by their sum. Must be a positive integer in the range 1..1_000_000.",
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
			Usage:    "Mark this variant as the experiment control. Defaults to `false`. The API rejects the request with 409 if the experiment already has a control variant — to swap controls, first PATCH the existing control to clear `isControl`, then create or PATCH the new one with `isControl: true`. The auto-generated control variant created with each new experiment can be replaced this way. DELETE on the control returns 409.",
			BodyPath: "isControl",
		},
		&requestflag.Flag[*string]{
			Name:     "redirect-url",
			Usage:    "Required for redirect variants. Use either a site-relative path such as `/pricing-v2` or an absolute `https://` URL. Cross-origin `http://` URLs are rejected. Omit for DOM modification variants.",
			BodyPath: "redirectUrl",
		},
		&requestflag.Flag[*string]{
			Name:     "variant-type",
			Usage:    "Variant type to create. Use `redirect` for redirect tests or `dom_modifications` for on-page changes.",
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
			Usage:                 "CSS selector used to find the element to modify on the page at runtime.",
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
			Usage:    "Promote or demote this variant as the control. Promoting a second variant while another already has `isControl: true` is rejected with 409 — clear the existing control first.",
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
			Usage:    "Updated variant type — `redirect` or `dom_modifications`. Changing this also requires updating the matching payload field (`redirectUrl` or `domModifications`).",
			BodyPath: "variantType",
		},
		&requestflag.Flag[*int64]{
			Name:     "weight",
			Usage:    "Updated traffic weight relative to other variants. Must be a positive integer in the range 1..1_000_000.",
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
			Usage:                 "CSS selector used to find the element to modify on the page at runtime.",
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

	params := githubcomwithoursplatformsdkgo.ExperimentVariantListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ExperimentVariants.List(ctx, params, options...)
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
		Title:          "experiment-variants list",
		Transform:      transform,
	})
}

func handleExperimentVariantsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomwithoursplatformsdkgo.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := githubcomwithoursplatformsdkgo.ExperimentVariantNewParams{}

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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomwithoursplatformsdkgo.ExperimentVariantUpdateParams{}

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
