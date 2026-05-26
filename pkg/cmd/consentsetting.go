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

var consentSettingsList = cli.Command{
	Name:            "list",
	Usage:           "List all consent settings. Requires scope: consentSettings:list",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleConsentSettingsList,
	HideHelpCommand: true,
}

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

var consentSettingsReplace = requestflag.WithInnerFlags(cli.Command{
	Name:    "replace",
	Usage:   "Replace a consent setting. Send the full ConsentSettingsInput body — omitted\noptional fields are reset. Use PATCH for partial updates. Requires scope:\nconsentSettings:update",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "category",
			Usage:    "Top-level consent categories. Server re-stamps `priority` to 0..N.",
			Required: true,
			BodyPath: "categories",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "default",
			Usage:    "Default rule used when the user is not in any region listed in `regions[]`.",
			Required: true,
			BodyPath: "default",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable name shown in the dashboard.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "region",
			Usage:    "Per-region rule overrides. Each `regionCode` must be unique across rules and must not appear in any other rule's `additionalRegions`.",
			Required: true,
			BodyPath: "regions",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "service",
			Usage:    `Per-service entries powering "show vendors" and category-aware blocking. Empty array clears the list.`,
			Required: true,
			BodyPath: "services",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Enabled to serve the CMP, Disabled to take it offline.",
			Required: true,
			BodyPath: "status",
		},
		&requestflag.Flag[*string]{
			Name:     "consent-cookie-name",
			Usage:    `Name of the cookie that stores consent state. Pass null to clear (defaults to "op_consent").`,
			BodyPath: "consentCookieName",
		},
		&requestflag.Flag[*string]{
			Name:     "custom-domain",
			Usage:    "Custom CDN domain for serving the CMP script. Pass null to clear.",
			BodyPath: "customDomain",
		},
		&requestflag.Flag[*float64]{
			Name:     "revision",
			Usage:    "Revision counter. Bump to re-prompt users who already consented.",
			BodyPath: "revision",
		},
		&requestflag.Flag[any]{
			Name:     "skip-blocking-class-name",
			Usage:    "CSS class names that opt scripts out of consent blocking. Each must be a single class token.",
			BodyPath: "skipBlockingClassNames",
		},
		&requestflag.Flag[*string]{
			Name:     "web-sdk-token",
			Usage:    "Pixel of the WebSource this CMP is wired into. Pass null to clear the link.",
			BodyPath: "webSDKToken",
		},
		&requestflag.Flag[any]{
			Name:     "whitelist-domain",
			Usage:    "Allowlist of domains where this CMP runs. Pass null/[] to clear.",
			BodyPath: "whitelistDomains",
		},
	},
	Action:          handleConsentSettingsReplace,
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

var consentSettingsAnalytics = cli.Command{
	Name:    "analytics",
	Usage:   "Time-series consent analytics for a single consent settings record: banner\nviews, opt-ins, opt-outs, close-icon clicks, and derived opt-in/out rates per\nUTC day (or per UTC hour with `granularity=HOURLY`). The window is zero-filled\nso callers get a contiguous series, and rates are person-level\n(`COUNT(DISTINCT visitor_id)`). Use the optional `pagePath` and `region` filters\nto scope to one page or one visitor region; use `compareWithPreviousPeriod=true`\nto also receive the matching prior window. `DAILY` allows a 90-day window;\n`HOURLY` is capped at 14 days. Requires the API-key scope\n`report:consent-analytics` (this endpoint returns consent analytics report data,\nwhich is PHI-bearing and gated separately from consent-settings management).\nRequires scope: report:consent-analytics",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "from",
			Usage:     "Inclusive lower bound of the analytics window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 90 days or fewer (14 for `HOURLY` granularity).",
			Required:  true,
			QueryPath: "from",
		},
		&requestflag.Flag[string]{
			Name:      "to",
			Usage:     "Inclusive upper bound of the analytics window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 90 days or fewer (14 for `HOURLY` granularity).",
			Required:  true,
			QueryPath: "to",
		},
		&requestflag.Flag[bool]{
			Name:      "compare-with-previous-period",
			Usage:     "When `true`, each bucket also returns the matching bucket from the immediately preceding window of equal length (in `previous*` and `percentageChange*` fields). Defaults to `false`.",
			QueryPath: "compareWithPreviousPeriod",
		},
		&requestflag.Flag[string]{
			Name:      "granularity",
			Usage:     "Bucket size for the time-series rollup. `DAILY` (default) buckets per UTC day; `HOURLY` buckets per UTC hour and limits the window to 14 days.",
			QueryPath: "granularity",
		},
		&requestflag.Flag[string]{
			Name:      "page-path",
			Usage:     "Filter to events whose `default_properties.pathname` equals this value (exact match, case-sensitive). Use this to drill into a single page.",
			QueryPath: "pagePath",
		},
		&requestflag.Flag[string]{
			Name:      "regions",
			Usage:     "Filter results to events whose `request_context.country_region_name` is in this set. Pass a single region (e.g. `California`) or a comma-separated list (`California,Texas`). Case-sensitive. Use `GET /rest/v1/consent-settings/{id}/analytics-by-region` to discover the region names available for an account.",
			QueryPath: "regions",
		},
	},
	Action:          handleConsentSettingsAnalytics,
	HideHelpCommand: true,
}

var consentSettingsPageAnalysis = cli.Command{
	Name:    "page-analysis",
	Usage:   "Per-page consent breakdown for one consent settings record, ranked by opt-outs\n(descending). Each row bundles banner views, opt-outs, close-icon clicks, and\nthe derived opt-out rate. Documented exception to the cursor-pagination\nstandard: this is a derived read whose underlying GraphQL contract is\noffset/limit-based; cursors are not used. `search` is a substring match against\n`pathname`; `region` filters to one visitor region. Requires the API-key scope\n`report:consent-page-analysis` (PHI-bearing report data). Requires scope:\nreport:consent-page-analysis",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "from",
			Usage:     "Inclusive lower bound of the analytics window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 90 days or fewer (14 for `HOURLY` granularity).",
			Required:  true,
			QueryPath: "from",
		},
		&requestflag.Flag[string]{
			Name:      "to",
			Usage:     "Inclusive upper bound of the analytics window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 90 days or fewer (14 for `HOURLY` granularity).",
			Required:  true,
			QueryPath: "to",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of pages to return. Defaults to 50.",
			QueryPath: "limit",
		},
		&requestflag.Flag[*int64]{
			Name:      "offset",
			Usage:     "Skip this many top-ranked pages before returning. Use together with `limit` for load-more pagination.",
			QueryPath: "offset",
		},
		&requestflag.Flag[string]{
			Name:      "regions",
			Usage:     "Filter results to events whose `request_context.country_region_name` is in this set. Pass a single region (e.g. `California`) or a comma-separated list (`California,Texas`). Case-sensitive. Use `GET /rest/v1/consent-settings/{id}/analytics-by-region` to discover the region names available for an account.",
			QueryPath: "regions",
		},
		&requestflag.Flag[string]{
			Name:      "search",
			Usage:     "Case-sensitive substring match against `default_properties.pathname`. Wrapped in `%...%` server-side.",
			QueryPath: "search",
		},
	},
	Action:          handleConsentSettingsPageAnalysis,
	HideHelpCommand: true,
}

var consentSettingsAnalyticsByRegion = cli.Command{
	Name:    "analytics-by-region",
	Usage:   "Region-grouped consent totals for the window: banner views, opt-ins, opt-outs,\nclose-icon clicks, and derived rates per visitor `country_region_name`, ranked\nby banner views (descending). Visitors whose region cannot be resolved (e.g. bot\ntraffic, IP geo failure) are bucketed under the literal `Unknown` so per-region\ncounts always sum to the global totals. Use this to discover the region names\nyou can later pass to the `region` filter on\n`GET /rest/v1/consent-settings/{id}/analytics`. Requires the API-key scope\n`report:consent-analytics` (PHI-bearing report data). Requires scope:\nreport:consent-analytics",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "from",
			Usage:     "Inclusive lower bound of the analytics window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 90 days or fewer (14 for `HOURLY` granularity).",
			Required:  true,
			QueryPath: "from",
		},
		&requestflag.Flag[string]{
			Name:      "to",
			Usage:     "Inclusive upper bound of the analytics window, as a UTC calendar day in `YYYY-MM-DD` format. The window between `from` and `to` must be 90 days or fewer (14 for `HOURLY` granularity).",
			Required:  true,
			QueryPath: "to",
		},
	},
	Action:          handleConsentSettingsAnalyticsByRegion,
	HideHelpCommand: true,
}

func handleConsentSettingsList(ctx context.Context, cmd *cli.Command) error {
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

func handleConsentSettingsCreate(ctx context.Context, cmd *cli.Command) error {
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

func handleConsentSettingsReplace(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.ConsentSettingReplaceParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ConsentSettings.Replace(
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
		Title:          "consent-settings replace",
		Transform:      transform,
	})
}

func handleConsentSettingsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.ConsentSettingUpdateParams{}

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

func handleConsentSettingsDelete(ctx context.Context, cmd *cli.Command) error {
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

func handleConsentSettingsAnalytics(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.ConsentSettingAnalyticsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ConsentSettings.Analytics(
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
		Title:          "consent-settings analytics",
		Transform:      transform,
	})
}

func handleConsentSettingsPageAnalysis(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.ConsentSettingPageAnalysisParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ConsentSettings.PageAnalysis(
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
		Title:          "consent-settings page-analysis",
		Transform:      transform,
	})
}

func handleConsentSettingsAnalyticsByRegion(ctx context.Context, cmd *cli.Command) error {
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

	params := oursprivacy.ConsentSettingAnalyticsByRegionParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ConsentSettings.AnalyticsByRegion(
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
		Title:          "consent-settings analytics-by-region",
		Transform:      transform,
	})
}
