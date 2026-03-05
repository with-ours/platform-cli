// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestReplaySettingsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "replay-settings", "create",
			"--api-key", "string",
			"--custom-domain", "customDomain",
			"--name", "name",
			"--status", "Disabled",
			"--whitelist-domain", "[string]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"customDomain: customDomain\n" +
			"name: name\n" +
			"status: Disabled\n" +
			"whitelistDomains:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "replay-settings", "create",
			"--api-key", "string",
		)
	})
}

func TestReplaySettingsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "replay-settings", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestReplaySettingsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "replay-settings", "update",
			"--api-key", "string",
			"--id", "id",
			"--custom-domain", "customDomain",
			"--name", "name",
			"--status", "Disabled",
			"--whitelist-domain", "[string]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"customDomain: customDomain\n" +
			"name: name\n" +
			"status: Disabled\n" +
			"whitelistDomains:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "replay-settings", "update",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestReplaySettingsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "replay-settings", "list",
			"--api-key", "string",
		)
	})
}

func TestReplaySettingsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "replay-settings", "delete",
			"--api-key", "string",
			"--id", "id",
		)
	})
}
