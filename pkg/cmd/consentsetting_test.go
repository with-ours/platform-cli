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
	)
}

func TestConsentSettingsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"consent-settings", "retrieve",
		"--id", "id",
	)
}

func TestConsentSettingsUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"consent-settings", "update",
		"--id", "id",
	)
}

func TestConsentSettingsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"consent-settings", "list",
	)
}

func TestConsentSettingsDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"consent-settings", "delete",
		"--id", "id",
	)
}
