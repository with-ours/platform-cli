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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"includeAllowedEvents:\n" +
			"  - string\n" +
			"includeConsentSettings:\n" +
			"  - string\n" +
			"includeDestinations:\n" +
			"  - string\n" +
			"includeExternalAllowedEventData:\n" +
			"  - string\n" +
			"includeGlobalDispatchCenters:\n" +
			"  - string\n" +
			"includeMappings:\n" +
			"  - string\n" +
			"includeReplaySettings:\n" +
			"  - string\n" +
			"includeSources:\n" +
			"  - string\n" +
			"includeTagManagerTags:\n" +
			"  - string\n" +
			"includeTagManagerTriggers:\n" +
			"  - string\n" +
			"includeTagManagerVariables:\n" +
			"  - string\n" +
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
