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

var webScannerRulesList = cli.Command{
	Name:    "list",
	Usage:   "List suppression rules for a single web scanner. Requires the `scannerId` query\nparameter — rules are always scoped to a parent scanner. Not paginated; the\nper-scanner rule count is bounded. Requires scope: webScanner:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "scanner-id",
			Usage:     "The web scanner whose suppression rules should be returned.",
			Required:  true,
			QueryPath: "scannerId",
		},
	},
	Action:          handleWebScannerRulesList,
	HideHelpCommand: true,
}

var webScannerRulesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a suppression rule on a web scanner. Auth is enforced against the parent\nscanner via `webScanner:update`. At least one of `cookiePatterns`,\n`domainPatterns`, or `scriptPatterns` should be set for the rule to match\nanything; omitted pattern arrays default to `[]`. Requires scope:\nwebScanner:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "User-friendly name for the suppression rule.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[int64]{
			Name:     "priority",
			Usage:    "Rule priority (1–10,000). Lower numbers are evaluated first when multiple rules match.",
			Required: true,
			BodyPath: "priority",
		},
		&requestflag.Flag[string]{
			Name:     "scanner-id",
			Usage:    "The web scanner this rule belongs to.",
			Required: true,
			BodyPath: "scannerId",
		},
		&requestflag.Flag[[]string]{
			Name:     "cookie-pattern",
			Usage:    "Glob patterns matched against cookie names (e.g. `_ga*`). Max 100 entries. When sent on PATCH, replaces the existing list wholesale.",
			BodyPath: "cookiePatterns",
		},
		&requestflag.Flag[[]string]{
			Name:     "domain-pattern",
			Usage:    "Glob patterns matched against cookie domain / script hostname (e.g. `*.google-analytics.com`). Max 100 entries. When sent on PATCH, replaces the existing list wholesale.",
			BodyPath: "domainPatterns",
		},
		&requestflag.Flag[*string]{
			Name:     "notes",
			Usage:    "Free-form notes about why this rule exists or what it covers. Trimmed server-side; empty strings become `null`.",
			BodyPath: "notes",
		},
		&requestflag.Flag[*string]{
			Name:     "reason",
			Usage:    "Why this rule was added. Surfaced in audit views. Send `null` to clear an existing reason on patch.",
			BodyPath: "reason",
		},
		&requestflag.Flag[[]string]{
			Name:     "script-pattern",
			Usage:    "Glob patterns matched against full script URLs (e.g. `https://www.googletagmanager.com/gtm.js?id=*`). Max 100 entries. When sent on PATCH, replaces the existing list wholesale.",
			BodyPath: "scriptPatterns",
		},
	},
	Action:          handleWebScannerRulesCreate,
	HideHelpCommand: true,
}

var webScannerRulesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single web scanner rule by ID. Requires scope: webScanner:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleWebScannerRulesRetrieve,
	HideHelpCommand: true,
}

var webScannerRulesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Partially update a suppression rule. Only the fields you send are changed.\nList-valued fields (`cookiePatterns`, `domainPatterns`, `scriptPatterns`) are\nreplaced wholesale when sent. Requires scope: webScanner:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[[]string]{
			Name:     "cookie-pattern",
			Usage:    "Glob patterns matched against cookie names (e.g. `_ga*`). Max 100 entries. When sent on PATCH, replaces the existing list wholesale.",
			BodyPath: "cookiePatterns",
		},
		&requestflag.Flag[[]string]{
			Name:     "domain-pattern",
			Usage:    "Glob patterns matched against cookie domain / script hostname (e.g. `*.google-analytics.com`). Max 100 entries. When sent on PATCH, replaces the existing list wholesale.",
			BodyPath: "domainPatterns",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "notes",
			Usage:    "Free-form notes about why this rule exists or what it covers. Trimmed server-side; empty strings become `null`.",
			BodyPath: "notes",
		},
		&requestflag.Flag[int64]{
			Name:     "priority",
			BodyPath: "priority",
		},
		&requestflag.Flag[*string]{
			Name:     "reason",
			Usage:    "Why this rule was added. Surfaced in audit views. Send `null` to clear an existing reason on patch.",
			BodyPath: "reason",
		},
		&requestflag.Flag[[]string]{
			Name:     "script-pattern",
			Usage:    "Glob patterns matched against full script URLs (e.g. `https://www.googletagmanager.com/gtm.js?id=*`). Max 100 entries. When sent on PATCH, replaces the existing list wholesale.",
			BodyPath: "scriptPatterns",
		},
	},
	Action:          handleWebScannerRulesUpdate,
	HideHelpCommand: true,
}

var webScannerRulesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a web scanner rule. Requires scope: webScanner:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleWebScannerRulesDelete,
	HideHelpCommand: true,
}

func handleWebScannerRulesList(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.WebScannerRuleListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WebScannerRules.List(ctx, params, options...)
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
		Title:          "web-scanner-rules list",
		Transform:      transform,
	})
}

func handleWebScannerRulesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.WebScannerRuleNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WebScannerRules.New(ctx, params, options...)
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
		Title:          "web-scanner-rules create",
		Transform:      transform,
	})
}

func handleWebScannerRulesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.WebScannerRules.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "web-scanner-rules retrieve",
		Transform:      transform,
	})
}

func handleWebScannerRulesUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.WebScannerRuleUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WebScannerRules.Update(
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
		Title:          "web-scanner-rules update",
		Transform:      transform,
	})
}

func handleWebScannerRulesDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.WebScannerRules.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "web-scanner-rules delete",
		Transform:      transform,
	})
}
