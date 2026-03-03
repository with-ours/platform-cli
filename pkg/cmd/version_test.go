// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestVersionsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"versions", "create",
		"--api-key", "string",
		"--include-allowed-event", "[string]",
		"--include-consent-setting", "[string]",
		"--include-destination", "[string]",
		"--include-external-allowed-event-data", "[string]",
		"--include-global-dispatch-center", "[string]",
		"--include-mapping", "[string]",
		"--include-replay-setting", "[string]",
		"--include-source", "[string]",
		"--include-tag-manager-tag", "[string]",
		"--include-tag-manager-trigger", "[string]",
		"--include-tag-manager-variable", "[string]",
		"--name", "name",
		"--notes", "notes",
	)
}

func TestVersionsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"versions", "retrieve",
		"--api-key", "string",
		"--id", "id",
	)
}

func TestVersionsUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"versions", "update",
		"--api-key", "string",
		"--id", "id",
		"--name", "name",
		"--notes", "notes",
	)
}

func TestVersionsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"versions", "list",
		"--api-key", "string",
	)
}
