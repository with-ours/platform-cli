// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestVersionsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"versions", "create",
			"--include-allowed-event", "[{}]",
			"--include-consent-setting", "[{}]",
			"--include-destination", "[{}]",
			"--include-external-allowed-event-data", "[{}]",
			"--include-global-dispatch-center", "[{}]",
			"--include-mapping", "[{}]",
			"--include-replay-setting", "[{}]",
			"--include-source", "[{}]",
			"--include-tag-manager-tag", "[{}]",
			"--include-tag-manager-trigger", "[{}]",
			"--include-tag-manager-variable", "[{}]",
			"--name", "name",
			"--notes", "notes",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"includeAllowedEvents:\n" +
			"  - {}\n" +
			"includeConsentSettings:\n" +
			"  - {}\n" +
			"includeDestinations:\n" +
			"  - {}\n" +
			"includeExternalAllowedEventData:\n" +
			"  - {}\n" +
			"includeGlobalDispatchCenters:\n" +
			"  - {}\n" +
			"includeMappings:\n" +
			"  - {}\n" +
			"includeReplaySettings:\n" +
			"  - {}\n" +
			"includeSources:\n" +
			"  - {}\n" +
			"includeTagManagerTags:\n" +
			"  - {}\n" +
			"includeTagManagerTriggers:\n" +
			"  - {}\n" +
			"includeTagManagerVariables:\n" +
			"  - {}\n" +
			"name: name\n" +
			"notes: notes\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"versions", "create",
		)
	})
}

func TestVersionsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"versions", "retrieve",
			"--id", "id",
		)
	})
}

func TestVersionsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"versions", "update",
			"--id", "id",
			"--name", "name",
			"--notes", "notes",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"notes: notes\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"versions", "update",
			"--id", "id",
		)
	})
}

func TestVersionsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"versions", "list",
		)
	})
}
