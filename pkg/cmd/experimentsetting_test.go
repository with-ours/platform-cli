// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestExperimentSettingsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-settings", "list",
		)
	})
}

func TestExperimentSettingsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-settings", "create",
			"--cookie-name", "_cord_exp",
			"--name", "Default Web Experiment Settings",
			"--whitelist-domain", "[www.example.com, staging.example.com]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"cookieName: _cord_exp\n" +
			"name: Default Web Experiment Settings\n" +
			"whitelistDomains:\n" +
			"  - www.example.com\n" +
			"  - staging.example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiment-settings", "create",
		)
	})
}

func TestExperimentSettingsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-settings", "retrieve",
			"--id", "id",
		)
	})
}

func TestExperimentSettingsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-settings", "update",
			"--id", "id",
			"--cookie-name", "_cord_exp",
			"--name", "Default Web Experiment Settings",
			"--whitelist-domain", "[www.example.com, staging.example.com]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"cookieName: _cord_exp\n" +
			"name: Default Web Experiment Settings\n" +
			"whitelistDomains:\n" +
			"  - www.example.com\n" +
			"  - staging.example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiment-settings", "update",
			"--id", "id",
		)
	})
}

func TestExperimentSettingsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-settings", "delete",
			"--id", "id",
		)
	})
}
