// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
	"github.com/with-ours/platform-cli/internal/requestflag"
)

func TestVersionsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"versions", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--is-published", "true",
			"--limit", "25",
			"--name-contains", "nameContains",
			"--notes-contains", "notesContains",
		)
	})
}

func TestVersionsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"versions", "create",
			"--include-allowed-event", "[string]",
			"--include-consent-setting", "[string]",
			"--include-data-governance-event", "[string]",
			"--include-data-governance-rule", "[string]",
			"--include-destination", "[string]",
			"--include-experiment", "[string]",
			"--include-experiment-setting", "[string]",
			"--include-experiment-variant", "[string]",
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
			"includeDataGovernanceEvents:\n" +
			"  - string\n" +
			"includeDataGovernanceRules:\n" +
			"  - string\n" +
			"includeDestinations:\n" +
			"  - string\n" +
			"includeExperiments:\n" +
			"  - string\n" +
			"includeExperimentSettings:\n" +
			"  - string\n" +
			"includeExperimentVariants:\n" +
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

func TestVersionsPublish(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"versions", "publish",
			"--id", "id",
		)
	})
}

func TestVersionsSnapshot(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"versions", "snapshot",
			"--id", "id",
		)
	})
}

func TestVersionsDiff(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"versions", "diff",
			"--id", "draft",
			"--against", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestVersionsRevert(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"versions", "revert",
			"--id", "draft",
			"--entity", "{id: id, collection: allowedEvents}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(versionsRevert)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"versions", "revert",
			"--id", "draft",
			"--entity.id", "id",
			"--entity.collection", "allowedEvents",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"entities:\n" +
			"  - id: id\n" +
			"    collection: allowedEvents\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"versions", "revert",
			"--id", "draft",
		)
	})
}

func TestVersionsAbandon(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"versions", "abandon",
			"--id", "draft",
		)
	})
}
