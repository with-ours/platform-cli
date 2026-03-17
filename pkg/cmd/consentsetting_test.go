// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
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
			"--category", "{}",
			"--name", "name",
			"--region", "{}",
			"--service", "{}",
			"--status", "Disabled",
			"--consent-cookie-name", "consentCookieName",
			"--custom-domain", "{}",
			"--default", "{}",
			"--revision", "0",
			"--skip-blocking-class-name", "[string]",
			"--web-sdk-token", "webSDKToken",
			"--whitelist-domain", "[{}]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"categories:\n" +
			"  - {}\n" +
			"name: name\n" +
			"regions:\n" +
			"  - {}\n" +
			"services:\n" +
			"  - {}\n" +
			"status: Disabled\n" +
			"consentCookieName: consentCookieName\n" +
			"customDomain: {}\n" +
			"default: {}\n" +
			"revision: 0\n" +
			"skipBlockingClassNames:\n" +
			"  - string\n" +
			"webSDKToken: webSDKToken\n" +
			"whitelistDomains:\n" +
			"  - {}\n")
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
