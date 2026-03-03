// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestConsentSettingsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"consent-settings", "create",
		"--api-key", "string",
	)
}

func TestConsentSettingsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"consent-settings", "retrieve",
		"--api-key", "string",
		"--id", "id",
	)
}

func TestConsentSettingsUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"consent-settings", "update",
		"--api-key", "string",
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
}

func TestConsentSettingsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"consent-settings", "list",
		"--api-key", "string",
	)
}

func TestConsentSettingsDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"consent-settings", "delete",
		"--api-key", "string",
		"--id", "id",
	)
}
