// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestAllowedEventsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "allowed-events", "create",
			"--api-key", "string",
			"--name", "name",
			"--destination-ids", "[string]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"destinationIds:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "allowed-events", "create",
			"--api-key", "string",
		)
	})
}

func TestAllowedEventsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "allowed-events", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestAllowedEventsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "allowed-events", "list",
			"--api-key", "string",
		)
	})
}

func TestAllowedEventsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "allowed-events", "delete",
			"--api-key", "string",
			"--id", "id",
		)
	})
}
