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
			"--default", "{categories: [{key: key, value: {enabled: true, autoDisableOnGPC: true, readOnly: true, reloadPage: true}}], language: language, mode: opt_in, translations: [{language: language, value: {consentModal: {acceptAllBtn: acceptAllBtn, description: description, rejectAllBtn: rejectAllBtn, showPreferencesBtn: showPreferencesBtn, title: title, closeIconLabel: closeIconLabel, footer: footer, gpcNotification: gpcNotification, privacyPolicyLabel: privacyPolicyLabel, privacyPolicyUrl: privacyPolicyUrl, revisionMessage: revisionMessage, termsOfServiceLabel: termsOfServiceLabel, termsOfServiceUrl: termsOfServiceUrl}, preferencesModal: {acceptAllBtn: acceptAllBtn, closeIconLabel: closeIconLabel, rejectAllBtn: rejectAllBtn, savePreferencesBtn: savePreferencesBtn, sections: [{description: description, title: title, cookieTable: {body: [{key: key, value: value}], headers: [{key: key, value: value}], caption: caption}, linkedCategory: linkedCategory}], title: title, serviceCounterLabel: serviceCounterLabel}}}], autoShow: true, autoShowDismissConfig: {pageCount: 0, seconds: 0}, autoShowDismissMode: after_pages, disablePageInteraction: true, guiOptions: {consentModal: {buttonLayout: AcceptAllNecessaryPreferences, equalWeightButtons: true, flipButtons: true, layout: bar, position: bottom, showCloseIcon: true}, cssVariables: {buttonBorderRadius: buttonBorderRadius, cookieCategoryBlockBg: cookieCategoryBlockBg, cookieCategoryBlockHoverBg: cookieCategoryBlockHoverBg, footerBg: footerBg, footerColor: footerColor, footerLinkColor: footerLinkColor, footerLinkHoverColor: footerLinkHoverColor, modalBg: modalBg, modalBorderRadius: modalBorderRadius, primaryButtonBg: primaryButtonBg, primaryButtonColor: primaryButtonColor, primaryButtonHoverBg: primaryButtonHoverBg, primaryButtonHoverColor: primaryButtonHoverColor, primaryTextColor: primaryTextColor, secondaryButtonBg: secondaryButtonBg, secondaryButtonColor: secondaryButtonColor, secondaryButtonHoverBg: secondaryButtonHoverBg, secondaryButtonHoverColor: secondaryButtonHoverColor, secondaryTextColor: secondaryTextColor, toggleOnBg: toggleOnBg}, preferencesModal: {buttonLayout: AcceptAllRejectAllSave, equalWeightButtons: true, flipButtons: true, layout: bar, position: left}}, hideFromBots: true, showVendorsInPreferences: true}",
			"--name", "name",
			"--region", "{regionCode: regionCode, rule: {categories: [{key: key, value: {enabled: true, autoDisableOnGPC: true, readOnly: true, reloadPage: true}}], language: language, mode: opt_in, translations: [{language: language, value: {consentModal: {acceptAllBtn: acceptAllBtn, description: description, rejectAllBtn: rejectAllBtn, showPreferencesBtn: showPreferencesBtn, title: title, closeIconLabel: closeIconLabel, footer: footer, gpcNotification: gpcNotification, privacyPolicyLabel: privacyPolicyLabel, privacyPolicyUrl: privacyPolicyUrl, revisionMessage: revisionMessage, termsOfServiceLabel: termsOfServiceLabel, termsOfServiceUrl: termsOfServiceUrl}, preferencesModal: {acceptAllBtn: acceptAllBtn, closeIconLabel: closeIconLabel, rejectAllBtn: rejectAllBtn, savePreferencesBtn: savePreferencesBtn, sections: [{description: description, title: title, cookieTable: {body: [{key: key, value: value}], headers: [{key: key, value: value}], caption: caption}, linkedCategory: linkedCategory}], title: title, serviceCounterLabel: serviceCounterLabel}}}], autoShow: true, autoShowDismissConfig: {pageCount: 0, seconds: 0}, autoShowDismissMode: after_pages, disablePageInteraction: true, guiOptions: {consentModal: {buttonLayout: AcceptAllNecessaryPreferences, equalWeightButtons: true, flipButtons: true, layout: bar, position: bottom, showCloseIcon: true}, cssVariables: {buttonBorderRadius: buttonBorderRadius, cookieCategoryBlockBg: cookieCategoryBlockBg, cookieCategoryBlockHoverBg: cookieCategoryBlockHoverBg, footerBg: footerBg, footerColor: footerColor, footerLinkColor: footerLinkColor, footerLinkHoverColor: footerLinkHoverColor, modalBg: modalBg, modalBorderRadius: modalBorderRadius, primaryButtonBg: primaryButtonBg, primaryButtonColor: primaryButtonColor, primaryButtonHoverBg: primaryButtonHoverBg, primaryButtonHoverColor: primaryButtonHoverColor, primaryTextColor: primaryTextColor, secondaryButtonBg: secondaryButtonBg, secondaryButtonColor: secondaryButtonColor, secondaryButtonHoverBg: secondaryButtonHoverBg, secondaryButtonHoverColor: secondaryButtonHoverColor, secondaryTextColor: secondaryTextColor, toggleOnBg: toggleOnBg}, preferencesModal: {buttonLayout: AcceptAllRejectAllSave, equalWeightButtons: true, flipButtons: true, layout: bar, position: left}}, hideFromBots: true, showVendorsInPreferences: true}, additionalRegions: [string]}",
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
			"--default.translations", "[{language: language, value: {consentModal: {acceptAllBtn: acceptAllBtn, description: description, rejectAllBtn: rejectAllBtn, showPreferencesBtn: showPreferencesBtn, title: title, closeIconLabel: closeIconLabel, footer: footer, gpcNotification: gpcNotification, privacyPolicyLabel: privacyPolicyLabel, privacyPolicyUrl: privacyPolicyUrl, revisionMessage: revisionMessage, termsOfServiceLabel: termsOfServiceLabel, termsOfServiceUrl: termsOfServiceUrl}, preferencesModal: {acceptAllBtn: acceptAllBtn, closeIconLabel: closeIconLabel, rejectAllBtn: rejectAllBtn, savePreferencesBtn: savePreferencesBtn, sections: [{description: description, title: title, cookieTable: {body: [{key: key, value: value}], headers: [{key: key, value: value}], caption: caption}, linkedCategory: linkedCategory}], title: title, serviceCounterLabel: serviceCounterLabel}}}]",
			"--default.auto-show=true",
			"--default.auto-show-dismiss-config", "{pageCount: 0, seconds: 0}",
			"--default.auto-show-dismiss-mode", "after_pages",
			"--default.disable-page-interaction=true",
			"--default.gui-options", "{consentModal: {buttonLayout: AcceptAllNecessaryPreferences, equalWeightButtons: true, flipButtons: true, layout: bar, position: bottom, showCloseIcon: true}, cssVariables: {buttonBorderRadius: buttonBorderRadius, cookieCategoryBlockBg: cookieCategoryBlockBg, cookieCategoryBlockHoverBg: cookieCategoryBlockHoverBg, footerBg: footerBg, footerColor: footerColor, footerLinkColor: footerLinkColor, footerLinkHoverColor: footerLinkHoverColor, modalBg: modalBg, modalBorderRadius: modalBorderRadius, primaryButtonBg: primaryButtonBg, primaryButtonColor: primaryButtonColor, primaryButtonHoverBg: primaryButtonHoverBg, primaryButtonHoverColor: primaryButtonHoverColor, primaryTextColor: primaryTextColor, secondaryButtonBg: secondaryButtonBg, secondaryButtonColor: secondaryButtonColor, secondaryButtonHoverBg: secondaryButtonHoverBg, secondaryButtonHoverColor: secondaryButtonHoverColor, secondaryTextColor: secondaryTextColor, toggleOnBg: toggleOnBg}, preferencesModal: {buttonLayout: AcceptAllRejectAllSave, equalWeightButtons: true, flipButtons: true, layout: bar, position: left}}",
			"--default.hide-from-bots=true",
			"--default.show-vendors-in-preferences=true",
			"--name", "name",
			"--region.region-code", "regionCode",
			"--region.rule", "{categories: [{key: key, value: {enabled: true, autoDisableOnGPC: true, readOnly: true, reloadPage: true}}], language: language, mode: opt_in, translations: [{language: language, value: {consentModal: {acceptAllBtn: acceptAllBtn, description: description, rejectAllBtn: rejectAllBtn, showPreferencesBtn: showPreferencesBtn, title: title, closeIconLabel: closeIconLabel, footer: footer, gpcNotification: gpcNotification, privacyPolicyLabel: privacyPolicyLabel, privacyPolicyUrl: privacyPolicyUrl, revisionMessage: revisionMessage, termsOfServiceLabel: termsOfServiceLabel, termsOfServiceUrl: termsOfServiceUrl}, preferencesModal: {acceptAllBtn: acceptAllBtn, closeIconLabel: closeIconLabel, rejectAllBtn: rejectAllBtn, savePreferencesBtn: savePreferencesBtn, sections: [{description: description, title: title, cookieTable: {body: [{key: key, value: value}], headers: [{key: key, value: value}], caption: caption}, linkedCategory: linkedCategory}], title: title, serviceCounterLabel: serviceCounterLabel}}}], autoShow: true, autoShowDismissConfig: {pageCount: 0, seconds: 0}, autoShowDismissMode: after_pages, disablePageInteraction: true, guiOptions: {consentModal: {buttonLayout: AcceptAllNecessaryPreferences, equalWeightButtons: true, flipButtons: true, layout: bar, position: bottom, showCloseIcon: true}, cssVariables: {buttonBorderRadius: buttonBorderRadius, cookieCategoryBlockBg: cookieCategoryBlockBg, cookieCategoryBlockHoverBg: cookieCategoryBlockHoverBg, footerBg: footerBg, footerColor: footerColor, footerLinkColor: footerLinkColor, footerLinkHoverColor: footerLinkHoverColor, modalBg: modalBg, modalBorderRadius: modalBorderRadius, primaryButtonBg: primaryButtonBg, primaryButtonColor: primaryButtonColor, primaryButtonHoverBg: primaryButtonHoverBg, primaryButtonHoverColor: primaryButtonHoverColor, primaryTextColor: primaryTextColor, secondaryButtonBg: secondaryButtonBg, secondaryButtonColor: secondaryButtonColor, secondaryButtonHoverBg: secondaryButtonHoverBg, secondaryButtonHoverColor: secondaryButtonHoverColor, secondaryTextColor: secondaryTextColor, toggleOnBg: toggleOnBg}, preferencesModal: {buttonLayout: AcceptAllRejectAllSave, equalWeightButtons: true, flipButtons: true, layout: bar, position: left}}, hideFromBots: true, showVendorsInPreferences: true}",
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
			"        consentModal:\n" +
			"          acceptAllBtn: acceptAllBtn\n" +
			"          description: description\n" +
			"          rejectAllBtn: rejectAllBtn\n" +
			"          showPreferencesBtn: showPreferencesBtn\n" +
			"          title: title\n" +
			"          closeIconLabel: closeIconLabel\n" +
			"          footer: footer\n" +
			"          gpcNotification: gpcNotification\n" +
			"          privacyPolicyLabel: privacyPolicyLabel\n" +
			"          privacyPolicyUrl: privacyPolicyUrl\n" +
			"          revisionMessage: revisionMessage\n" +
			"          termsOfServiceLabel: termsOfServiceLabel\n" +
			"          termsOfServiceUrl: termsOfServiceUrl\n" +
			"        preferencesModal:\n" +
			"          acceptAllBtn: acceptAllBtn\n" +
			"          closeIconLabel: closeIconLabel\n" +
			"          rejectAllBtn: rejectAllBtn\n" +
			"          savePreferencesBtn: savePreferencesBtn\n" +
			"          sections:\n" +
			"            - description: description\n" +
			"              title: title\n" +
			"              cookieTable:\n" +
			"                body:\n" +
			"                  - key: key\n" +
			"                    value: value\n" +
			"                headers:\n" +
			"                  - key: key\n" +
			"                    value: value\n" +
			"                caption: caption\n" +
			"              linkedCategory: linkedCategory\n" +
			"          title: title\n" +
			"          serviceCounterLabel: serviceCounterLabel\n" +
			"  autoShow: true\n" +
			"  autoShowDismissConfig:\n" +
			"    pageCount: 0\n" +
			"    seconds: 0\n" +
			"  autoShowDismissMode: after_pages\n" +
			"  disablePageInteraction: true\n" +
			"  guiOptions:\n" +
			"    consentModal:\n" +
			"      buttonLayout: AcceptAllNecessaryPreferences\n" +
			"      equalWeightButtons: true\n" +
			"      flipButtons: true\n" +
			"      layout: bar\n" +
			"      position: bottom\n" +
			"      showCloseIcon: true\n" +
			"    cssVariables:\n" +
			"      buttonBorderRadius: buttonBorderRadius\n" +
			"      cookieCategoryBlockBg: cookieCategoryBlockBg\n" +
			"      cookieCategoryBlockHoverBg: cookieCategoryBlockHoverBg\n" +
			"      footerBg: footerBg\n" +
			"      footerColor: footerColor\n" +
			"      footerLinkColor: footerLinkColor\n" +
			"      footerLinkHoverColor: footerLinkHoverColor\n" +
			"      modalBg: modalBg\n" +
			"      modalBorderRadius: modalBorderRadius\n" +
			"      primaryButtonBg: primaryButtonBg\n" +
			"      primaryButtonColor: primaryButtonColor\n" +
			"      primaryButtonHoverBg: primaryButtonHoverBg\n" +
			"      primaryButtonHoverColor: primaryButtonHoverColor\n" +
			"      primaryTextColor: primaryTextColor\n" +
			"      secondaryButtonBg: secondaryButtonBg\n" +
			"      secondaryButtonColor: secondaryButtonColor\n" +
			"      secondaryButtonHoverBg: secondaryButtonHoverBg\n" +
			"      secondaryButtonHoverColor: secondaryButtonHoverColor\n" +
			"      secondaryTextColor: secondaryTextColor\n" +
			"      toggleOnBg: toggleOnBg\n" +
			"    preferencesModal:\n" +
			"      buttonLayout: AcceptAllRejectAllSave\n" +
			"      equalWeightButtons: true\n" +
			"      flipButtons: true\n" +
			"      layout: bar\n" +
			"      position: left\n" +
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
			"            consentModal:\n" +
			"              acceptAllBtn: acceptAllBtn\n" +
			"              description: description\n" +
			"              rejectAllBtn: rejectAllBtn\n" +
			"              showPreferencesBtn: showPreferencesBtn\n" +
			"              title: title\n" +
			"              closeIconLabel: closeIconLabel\n" +
			"              footer: footer\n" +
			"              gpcNotification: gpcNotification\n" +
			"              privacyPolicyLabel: privacyPolicyLabel\n" +
			"              privacyPolicyUrl: privacyPolicyUrl\n" +
			"              revisionMessage: revisionMessage\n" +
			"              termsOfServiceLabel: termsOfServiceLabel\n" +
			"              termsOfServiceUrl: termsOfServiceUrl\n" +
			"            preferencesModal:\n" +
			"              acceptAllBtn: acceptAllBtn\n" +
			"              closeIconLabel: closeIconLabel\n" +
			"              rejectAllBtn: rejectAllBtn\n" +
			"              savePreferencesBtn: savePreferencesBtn\n" +
			"              sections:\n" +
			"                - description: description\n" +
			"                  title: title\n" +
			"                  cookieTable:\n" +
			"                    body:\n" +
			"                      - key: key\n" +
			"                        value: value\n" +
			"                    headers:\n" +
			"                      - key: key\n" +
			"                        value: value\n" +
			"                    caption: caption\n" +
			"                  linkedCategory: linkedCategory\n" +
			"              title: title\n" +
			"              serviceCounterLabel: serviceCounterLabel\n" +
			"      autoShow: true\n" +
			"      autoShowDismissConfig:\n" +
			"        pageCount: 0\n" +
			"        seconds: 0\n" +
			"      autoShowDismissMode: after_pages\n" +
			"      disablePageInteraction: true\n" +
			"      guiOptions:\n" +
			"        consentModal:\n" +
			"          buttonLayout: AcceptAllNecessaryPreferences\n" +
			"          equalWeightButtons: true\n" +
			"          flipButtons: true\n" +
			"          layout: bar\n" +
			"          position: bottom\n" +
			"          showCloseIcon: true\n" +
			"        cssVariables:\n" +
			"          buttonBorderRadius: buttonBorderRadius\n" +
			"          cookieCategoryBlockBg: cookieCategoryBlockBg\n" +
			"          cookieCategoryBlockHoverBg: cookieCategoryBlockHoverBg\n" +
			"          footerBg: footerBg\n" +
			"          footerColor: footerColor\n" +
			"          footerLinkColor: footerLinkColor\n" +
			"          footerLinkHoverColor: footerLinkHoverColor\n" +
			"          modalBg: modalBg\n" +
			"          modalBorderRadius: modalBorderRadius\n" +
			"          primaryButtonBg: primaryButtonBg\n" +
			"          primaryButtonColor: primaryButtonColor\n" +
			"          primaryButtonHoverBg: primaryButtonHoverBg\n" +
			"          primaryButtonHoverColor: primaryButtonHoverColor\n" +
			"          primaryTextColor: primaryTextColor\n" +
			"          secondaryButtonBg: secondaryButtonBg\n" +
			"          secondaryButtonColor: secondaryButtonColor\n" +
			"          secondaryButtonHoverBg: secondaryButtonHoverBg\n" +
			"          secondaryButtonHoverColor: secondaryButtonHoverColor\n" +
			"          secondaryTextColor: secondaryTextColor\n" +
			"          toggleOnBg: toggleOnBg\n" +
			"        preferencesModal:\n" +
			"          buttonLayout: AcceptAllRejectAllSave\n" +
			"          equalWeightButtons: true\n" +
			"          flipButtons: true\n" +
			"          layout: bar\n" +
			"          position: left\n" +
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
