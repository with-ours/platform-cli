// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/ours-privacy-platform-cli/internal/mocktest"
)

func TestRestV1ConsentSettingsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:consent-settings", "create",
	)
}

func TestRestV1ConsentSettingsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:consent-settings", "retrieve",
		"--id", "id",
	)
}

func TestRestV1ConsentSettingsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:consent-settings", "update",
		"--id", "id",
	)
}

func TestRestV1ConsentSettingsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:consent-settings", "list",
	)
}

func TestRestV1ConsentSettingsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:consent-settings", "delete",
		"--id", "id",
	)
}
