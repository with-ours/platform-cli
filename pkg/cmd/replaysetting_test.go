// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestReplaySettingsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"replay-settings", "create",
		"--api-key", "string",
		"--custom-domain", "customDomain",
		"--name", "name",
		"--status", "Disabled",
		"--whitelist-domain", "[string]",
	)
}

func TestReplaySettingsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"replay-settings", "retrieve",
		"--api-key", "string",
		"--id", "id",
	)
}

func TestReplaySettingsUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"replay-settings", "update",
		"--api-key", "string",
		"--id", "id",
		"--custom-domain", "customDomain",
		"--name", "name",
		"--status", "Disabled",
		"--whitelist-domain", "[string]",
	)
}

func TestReplaySettingsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"replay-settings", "list",
		"--api-key", "string",
	)
}

func TestReplaySettingsDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"replay-settings", "delete",
		"--api-key", "string",
		"--id", "id",
	)
}
