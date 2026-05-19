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

var webScannersList = cli.Command{
	Name:            "list",
	Usage:           "List every web scanner for this account. Not paginated — accounts have a small\nnumber of scanners in practice, so the response always fits in a single page.\nRequires scope: webScanner:list",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleWebScannersList,
	HideHelpCommand: true,
}

var webScannersCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new web scanner for a root domain. A first scan is enqueued\nautomatically after creation on a best-effort basis. `rootDomain` is required;\nmissing, empty, or malformed values are rejected as HTTP 400. Everything else\nfalls back to defaults (`status: Enabled`, `urlLimit: 100`, no excluded\npatterns, no extra seed URLs). The returned entity is the created scanner row\nand may not yet reflect async scan-state changes. Requires scope:\nwebScanner:create",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "root-domain",
			Usage:    "Root domain to crawl (e.g. `example.com`). Required on create. Missing or empty values fail request validation as HTTP 400. Present-but-malformed values are rejected as HTTP 400 with the validation reason in `details`.",
			Required: true,
			BodyPath: "rootDomain",
		},
		&requestflag.Flag[any]{
			Name:     "excluded-pattern",
			Usage:    "URL glob patterns to skip during crawl. Max 100 entries.",
			BodyPath: "excludedPatterns",
		},
		&requestflag.Flag[any]{
			Name:     "included-url",
			Usage:    "Additional seed URLs to include as crawl entry points. Each must be an http(s) URL. Max 100 entries.",
			BodyPath: "includedUrls",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    `Allowed values: "Disabled", "Enabled".`,
			BodyPath: "status",
		},
		&requestflag.Flag[*float64]{
			Name:     "url-limit",
			Usage:    "Maximum URLs to crawl per scan (1–20,000). Defaults to 100 when omitted.",
			BodyPath: "urlLimit",
		},
	},
	Action:          handleWebScannersCreate,
	HideHelpCommand: true,
}

var webScannersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single web scanner by ID. Requires scope: webScanner:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleWebScannersRetrieve,
	HideHelpCommand: true,
}

var webScannersUpdate = cli.Command{
	Name:    "update",
	Usage:   "Partially update a web scanner. Only the fields you send are changed; omitted\nfields keep their current value. List-valued fields (`excludedPatterns`,\n`includedUrls`) are replaced wholesale when sent. If `rootDomain` is provided\nand malformed, the request is rejected as HTTP 400. Use\n`POST /rest/v1/web-scanners/{id}/trigger` to start a new scan after edits.\nRequires scope: webScanner:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[any]{
			Name:     "excluded-pattern",
			BodyPath: "excludedPatterns",
		},
		&requestflag.Flag[any]{
			Name:     "included-url",
			BodyPath: "includedUrls",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "root-domain",
			Usage:    "Replace the scanner root domain. When provided, malformed values are rejected as HTTP 400 with the validation reason in `details`.",
			BodyPath: "rootDomain",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    `Allowed values: "Disabled", "Enabled".`,
			BodyPath: "status",
		},
		&requestflag.Flag[*float64]{
			Name:     "url-limit",
			BodyPath: "urlLimit",
		},
	},
	Action:          handleWebScannersUpdate,
	HideHelpCommand: true,
}

var webScannersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a web scanner. Associated suppression rules are deleted in the same\noperation. Requires scope: webScanner:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleWebScannersDelete,
	HideHelpCommand: true,
}

var webScannersTrigger = cli.Command{
	Name:    "trigger",
	Usage:   "Manually kick off a new scan for this web scanner. The request body is empty (or\n`{}`). A successful response means the enqueue request was accepted; because the\nscan starts asynchronously, the returned entity may still reflect pre-trigger\nvalues for fields like `scanStatus` and `lastScanStartedAt`. The trigger is\nrate-limited: a 409 is returned if another scan is already in flight, the\nper-account cooldown has not elapsed, or the trigger backend rejects the\nrequest; the reason is in the response `error` field. Requires scope:\nwebScanner:trigger",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleWebScannersTrigger,
	HideHelpCommand: true,
}

func handleWebScannersList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.WebScanners.List(ctx, options...)
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
		Title:          "web-scanners list",
		Transform:      transform,
	})
}

func handleWebScannersCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.WebScannerNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WebScanners.New(ctx, params, options...)
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
		Title:          "web-scanners create",
		Transform:      transform,
	})
}

func handleWebScannersRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.WebScanners.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "web-scanners retrieve",
		Transform:      transform,
	})
}

func handleWebScannersUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.WebScannerUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WebScanners.Update(
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
		Title:          "web-scanners update",
		Transform:      transform,
	})
}

func handleWebScannersDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.WebScanners.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "web-scanners delete",
		Transform:      transform,
	})
}

func handleWebScannersTrigger(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.WebScanners.Trigger(ctx, cmd.Value("id").(string), options...)
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
		Title:          "web-scanners trigger",
		Transform:      transform,
	})
}
