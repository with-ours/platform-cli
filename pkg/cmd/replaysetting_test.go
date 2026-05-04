// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestReplaySettingsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"replay-settings", "create",
			"--custom-domain", "customDomain",
			"--name", "name",
			"--status", "status",
			"--whitelist-domain", "[{}]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"customDomain: customDomain\n" +
			"name: name\n" +
			"status: status\n" +
			"whitelistDomains:\n" +
			"  - {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"replay-settings", "create",
		)
	})
}

func TestReplaySettingsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"replay-settings", "retrieve",
			"--id", "id",
		)
	})
}

func TestReplaySettingsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"replay-settings", "update",
			"--id", "id",
			"--custom-domain", "customDomain",
			"--name", "name",
			"--status", "status",
			"--whitelist-domain", "[{}]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"customDomain: customDomain\n" +
			"name: name\n" +
			"status: status\n" +
			"whitelistDomains:\n" +
			"  - {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"replay-settings", "update",
			"--id", "id",
		)
	})
}

func TestReplaySettingsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"replay-settings", "list",
		)
	})
}

func TestReplaySettingsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"replay-settings", "delete",
			"--id", "id",
		)
	})
}
