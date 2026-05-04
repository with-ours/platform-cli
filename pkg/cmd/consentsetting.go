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

var consentSettingsCreate = cli.Command{
	Name:            "create",
	Usage:           "Create a new consent settings record. POST takes no request body — the server\ninitializes the record with defaults (Disabled status, opt-out default rule,\nEnglish translations, necessary/analytics/advertising categories, no regions, no\nwhitelisted domains). Configure the record afterward with PATCH (partial update)\nor PUT (full replacement). Returns the same shape as GET so you can read the\nserver-assigned `id`, default rule, and categories without a follow-up fetch.\nRequires scope: consentSettings:create",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleConsentSettingsCreate,
	HideHelpCommand: true,
}

var consentSettingsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Find a single consent setting by ID. Requires scope: consentSettings:find",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleConsentSettingsRetrieve,
	HideHelpCommand: true,
}

var consentSettingsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Partially update a consent setting. Send only the fields you want to change —\nevery field is optional and unspecified fields are preserved. List-valued fields\n(services, categories, regions) are replaced wholesale when sent. Requires\nscope: consentSettings:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "category",
			Usage:    "Replace the entire categories list. Omit to leave existing categories untouched.",
			BodyPath: "categories",
		},
		&requestflag.Flag[*string]{
			Name:     "consent-cookie-name",
			Usage:    "Set or clear the consent cookie name.",
			BodyPath: "consentCookieName",
		},
		&requestflag.Flag[*string]{
			Name:     "custom-domain",
			Usage:    "Set or clear the custom CDN domain.",
			BodyPath: "customDomain",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "default",
			Usage:    "Replace the default rule wholesale. Omit to leave it untouched.",
			BodyPath: "default",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Rename the consent settings record.",
			BodyPath: "name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "region",
			Usage:    "Replace the entire regions list. Omit to leave it untouched. To change one region, send the full regions array with that region's rule modified.",
			BodyPath: "regions",
		},
		&requestflag.Flag[*float64]{
			Name:     "revision",
			Usage:    "Bump the revision counter to re-prompt users.",
			BodyPath: "revision",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "service",
			Usage:    "Replace the entire services list. Omit to leave existing services untouched.",
			BodyPath: "services",
		},
		&requestflag.Flag[any]{
			Name:     "skip-blocking-class-name",
			Usage:    "Replace the skipBlockingClassNames list. Pass null/[] to clear.",
			BodyPath: "skipBlockingClassNames",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Toggle Enabled/Disabled without re-sending the rest of the config.",
			BodyPath: "status",
		},
		&requestflag.Flag[*string]{
			Name:     "web-sdk-token",
			Usage:    "Set or clear the WebSource pixel link. A non-null token must be a valid WebSource of yours.",
			BodyPath: "webSDKToken",
		},
		&requestflag.Flag[any]{
			Name:     "whitelist-domain",
			Usage:    "Replace the allowlist. Pass null/[] to clear.",
			BodyPath: "whitelistDomains",
		},
	},
	Action:          handleConsentSettingsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"category": {
		&requestflag.InnerFlag[string]{
			Name:       "category.label",
			Usage:      "Human-readable label shown next to the toggle in the preferences modal.",
			InnerField: "label",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "category.priority",
			Usage:      "Sort key. Lower numbers render first. Server re-stamps to 0..N on write — send any integer, gaps and duplicates are ironed out.",
			InnerField: "priority",
		},
		&requestflag.InnerFlag[string]{
			Name:       "category.value",
			Usage:      `Stable identifier referenced by services and translation sections. Conventionally lowercase (e.g. "necessary", "analytics", "advertising").`,
			InnerField: "value",
		},
	},
	"default": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "default.categories",
			Usage:      "Per-category default config for this rule. Every category defined in the top-level `categories[].value` should have an entry here.",
			InnerField: "categories",
		},
		&requestflag.InnerFlag[string]{
			Name:       "default.language",
			Usage:      "BCP 47 default language for this rule. Must have a matching entry in `translations`. Examples: \"en\", \"en-US\", \"es\", \"de\".",
			InnerField: "language",
		},
		&requestflag.InnerFlag[string]{
			Name:       "default.mode",
			Usage:      "opt_in: scripts blocked until user accepts (GDPR style). opt_out: scripts run by default until user rejects (CCPA style).",
			InnerField: "mode",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "default.translations",
			Usage:      "All UI copy, keyed by language. Must include an entry whose `language` matches the rule's `language` field.",
			InnerField: "translations",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "default.autoblock-unknown",
			Usage:      "When true, scripts not classified by services[] are blocked until the user opts in.",
			InnerField: "autoblockUnknown",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "default.auto-show",
			Usage:      "When true, the consent modal auto-opens on page load.",
			InnerField: "autoShow",
		},
		&requestflag.InnerFlag[any]{
			Name:       "default.auto-show-dismiss-config",
			Usage:      "Threshold config for autoShowDismissMode (page count or seconds).",
			InnerField: "autoShowDismissConfig",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "default.auto-show-dismiss-mode",
			Usage:      "How the modal is treated as dismissed (never, after_pages, after_seconds).",
			InnerField: "autoShowDismissMode",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "default.disable-page-interaction",
			Usage:      "When true, the rest of the page is locked behind a backdrop until the user chooses.",
			InnerField: "disablePageInteraction",
		},
		&requestflag.InnerFlag[any]{
			Name:       "default.gui-options",
			Usage:      "Visual options for the modals (layout/position/colors).",
			InnerField: "guiOptions",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "default.hide-from-bots",
			Usage:      "When true, the modal is suppressed for known bot user agents.",
			InnerField: "hideFromBots",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "default.show-vendors-in-preferences",
			Usage:      "When true, the per-service list (services[]) is rendered inside the preferences modal.",
			InnerField: "showVendorsInPreferences",
		},
	},
	"region": {
		&requestflag.InnerFlag[string]{
			Name:       "region.region-code",
			Usage:      "Region this rule applies to. Use ISO 3166-1 alpha-2 country code (\"US\", \"DE\", \"BR\") or country-subdivision code (\"US-CA\", \"GB-ENG\", \"CA-ON\"). Each region code may appear in only one rule across `regions[]`.",
			InnerField: "regionCode",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "region.rule",
			InnerField: "rule",
		},
		&requestflag.InnerFlag[any]{
			Name:       "region.additional-regions",
			Usage:      "Other region codes that should reuse this rule. Same code-format rules as `regionCode`. Cannot include `regionCode` itself, cannot duplicate, cannot overlap with another rule's regions.",
			InnerField: "additionalRegions",
		},
	},
	"service": {
		&requestflag.InnerFlag[string]{
			Name:       "service.internal-notes",
			Usage:      "Internal notes shown to admins in the dashboard. Not user-facing.",
			InnerField: "internalNotes",
		},
		&requestflag.InnerFlag[string]{
			Name:       "service.label",
			Usage:      "Display name for this service in the preferences modal.",
			InnerField: "label",
		},
		&requestflag.InnerFlag[any]{
			Name:       "service.additional-categories",
			Usage:      "Extra category values this service belongs to. Each must match a `categories[].value`.",
			InnerField: "additionalCategories",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "service.category",
			Usage:      "Primary category value this service belongs to. Must match one of the top-level `categories[].value` entries.",
			InnerField: "category",
		},
		&requestflag.InnerFlag[any]{
			Name:       "service.domain-patterns",
			Usage:      "Domains/paths this service matches. Patterns matching the CMP's own scripts (e.g. cdn.oursprivacy.com/cmp-init) are rejected to prevent the CMP blocking itself — use a more specific path like cdn.oursprivacy.com/main.js to block a specific script.",
			InnerField: "domainPatterns",
		},
	},
})

var consentSettingsList = cli.Command{
	Name:            "list",
	Usage:           "List all consent settings. Requires scope: consentSettings:list",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleConsentSettingsList,
	HideHelpCommand: true,
}

var consentSettingsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a consent setting. Requires scope: consentSettings:delete",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleConsentSettingsDelete,
	HideHelpCommand: true,
}

func handleConsentSettingsCreate(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ConsentSettings.New(ctx, options...)
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
		Title:          "consent-settings create",
		Transform:      transform,
	})
}

func handleConsentSettingsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ConsentSettings.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "consent-settings retrieve",
		Transform:      transform,
	})
}

func handleConsentSettingsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomwithoursplatformsdkgo.ConsentSettingUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ConsentSettings.Update(
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
		Title:          "consent-settings update",
		Transform:      transform,
	})
}

func handleConsentSettingsList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ConsentSettings.List(ctx, options...)
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
		Title:          "consent-settings list",
		Transform:      transform,
	})
}

func handleConsentSettingsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ConsentSettings.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "consent-settings delete",
		Transform:      transform,
	})
}
