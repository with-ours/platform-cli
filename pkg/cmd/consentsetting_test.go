// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
	"github.com/with-ours/platform-cli/internal/requestflag"
)

func TestConsentSettingsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"consent-settings", "create",
		)
	})
}

func TestConsentSettingsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"consent-settings", "retrieve",
			"--id", "id",
		)
	})
}

func TestConsentSettingsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"consent-settings", "update",
			"--id", "id",
			"--category", "{label: label, priority: 0, value: value}",
			"--default", "{categories: [{key: key, value: {enabled: true, autoDisableOnGPC: true, readOnly: true, reloadPage: true}}], language: language, mode: opt_in, translations: [{language: language, value: {consentModal: {}, preferencesModal: {}}}], autoblockUnknown: true, autoShow: true, autoShowDismissConfig: {}, autoShowDismissMode: autoShowDismissMode, disablePageInteraction: true, guiOptions: {}, hideFromBots: true, showVendorsInPreferences: true}",
			"--name", "name",
			"--region", "{regionCode: regionCode, rule: {categories: [{key: key, value: {enabled: true, autoDisableOnGPC: true, readOnly: true, reloadPage: true}}], language: language, mode: opt_in, translations: [{language: language, value: {consentModal: {}, preferencesModal: {}}}], autoblockUnknown: true, autoShow: true, autoShowDismissConfig: {}, autoShowDismissMode: autoShowDismissMode, disablePageInteraction: true, guiOptions: {}, hideFromBots: true, showVendorsInPreferences: true}, additionalRegions: [string]}",
			"--service", "{internalNotes: internalNotes, label: label, additionalCategories: [string], category: category, domainPatterns: [string]}",
			"--status", "Disabled",
			"--consent-cookie-name", "consentCookieName",
			"--custom-domain", "customDomain",
			"--revision", "0",
			"--skip-blocking-class-name", "[string]",
			"--web-sdk-token", "webSDKToken",
			"--whitelist-domain", "[string]",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(consentSettingsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"consent-settings", "update",
			"--id", "id",
			"--category.label", "label",
			"--category.priority", "0",
			"--category.value", "value",
			"--default.categories", "[{key: key, value: {enabled: true, autoDisableOnGPC: true, readOnly: true, reloadPage: true}}]",
			"--default.language", "language",
			"--default.mode", "opt_in",
			"--default.translations", "[{language: language, value: {consentModal: {}, preferencesModal: {}}}]",
			"--default.autoblock-unknown=true",
			"--default.auto-show=true",
			"--default.auto-show-dismiss-config", "{}",
			"--default.auto-show-dismiss-mode", "autoShowDismissMode",
			"--default.disable-page-interaction=true",
			"--default.gui-options", "{}",
			"--default.hide-from-bots=true",
			"--default.show-vendors-in-preferences=true",
			"--name", "name",
			"--region.region-code", "regionCode",
			"--region.rule", "{categories: [{key: key, value: {enabled: true, autoDisableOnGPC: true, readOnly: true, reloadPage: true}}], language: language, mode: opt_in, translations: [{language: language, value: {consentModal: {}, preferencesModal: {}}}], autoblockUnknown: true, autoShow: true, autoShowDismissConfig: {}, autoShowDismissMode: autoShowDismissMode, disablePageInteraction: true, guiOptions: {}, hideFromBots: true, showVendorsInPreferences: true}",
			"--region.additional-regions", "[string]",
			"--service.internal-notes", "internalNotes",
			"--service.label", "label",
			"--service.additional-categories", "[string]",
			"--service.category", "category",
			"--service.domain-patterns", "[string]",
			"--status", "Disabled",
			"--consent-cookie-name", "consentCookieName",
			"--custom-domain", "customDomain",
			"--revision", "0",
			"--skip-blocking-class-name", "[string]",
			"--web-sdk-token", "webSDKToken",
			"--whitelist-domain", "[string]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"categories:\n" +
			"  - label: label\n" +
			"    priority: 0\n" +
			"    value: value\n" +
			"default:\n" +
			"  categories:\n" +
			"    - key: key\n" +
			"      value:\n" +
			"        enabled: true\n" +
			"        autoDisableOnGPC: true\n" +
			"        readOnly: true\n" +
			"        reloadPage: true\n" +
			"  language: language\n" +
			"  mode: opt_in\n" +
			"  translations:\n" +
			"    - language: language\n" +
			"      value:\n" +
			"        consentModal: {}\n" +
			"        preferencesModal: {}\n" +
			"  autoblockUnknown: true\n" +
			"  autoShow: true\n" +
			"  autoShowDismissConfig: {}\n" +
			"  autoShowDismissMode: autoShowDismissMode\n" +
			"  disablePageInteraction: true\n" +
			"  guiOptions: {}\n" +
			"  hideFromBots: true\n" +
			"  showVendorsInPreferences: true\n" +
			"name: name\n" +
			"regions:\n" +
			"  - regionCode: regionCode\n" +
			"    rule:\n" +
			"      categories:\n" +
			"        - key: key\n" +
			"          value:\n" +
			"            enabled: true\n" +
			"            autoDisableOnGPC: true\n" +
			"            readOnly: true\n" +
			"            reloadPage: true\n" +
			"      language: language\n" +
			"      mode: opt_in\n" +
			"      translations:\n" +
			"        - language: language\n" +
			"          value:\n" +
			"            consentModal: {}\n" +
			"            preferencesModal: {}\n" +
			"      autoblockUnknown: true\n" +
			"      autoShow: true\n" +
			"      autoShowDismissConfig: {}\n" +
			"      autoShowDismissMode: autoShowDismissMode\n" +
			"      disablePageInteraction: true\n" +
			"      guiOptions: {}\n" +
			"      hideFromBots: true\n" +
			"      showVendorsInPreferences: true\n" +
			"    additionalRegions:\n" +
			"      - string\n" +
			"services:\n" +
			"  - internalNotes: internalNotes\n" +
			"    label: label\n" +
			"    additionalCategories:\n" +
			"      - string\n" +
			"    category: category\n" +
			"    domainPatterns:\n" +
			"      - string\n" +
			"status: Disabled\n" +
			"consentCookieName: consentCookieName\n" +
			"customDomain: customDomain\n" +
			"revision: 0\n" +
			"skipBlockingClassNames:\n" +
			"  - string\n" +
			"webSDKToken: webSDKToken\n" +
			"whitelistDomains:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"consent-settings", "update",
			"--id", "id",
		)
	})
}

func TestConsentSettingsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"consent-settings", "list",
		)
	})
}

func TestConsentSettingsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"consent-settings", "delete",
			"--id", "id",
		)
	})
}
