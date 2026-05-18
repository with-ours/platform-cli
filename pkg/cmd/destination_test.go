// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestDestinationsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"destinations", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "25",
			"--status", "Disabled",
			"--type", "AWSEventBridge",
		)
	})
}

func TestDestinationsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"destinations", "create",
			"--type", "Audiohook",
			"--name", "name",
			"--settings", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"type: Audiohook\n" +
			"name: name\n" +
			"settings: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"destinations", "create",
		)
	})
}

func TestDestinationsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"destinations", "retrieve",
			"--id", "id",
		)
	})
}

func TestDestinationsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"destinations", "update",
			"--id", "id",
			"--hashing-salt", "hashingSalt",
			"--limited-to-source-id", "[string]",
			"--name", "name",
			"--settings", "{}",
			"--status", "Disabled",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"hashingSalt: hashingSalt\n" +
			"limitedToSourceIds:\n" +
			"  - string\n" +
			"name: name\n" +
			"settings: {}\n" +
			"status: Disabled\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"destinations", "update",
			"--id", "id",
		)
	})
}

func TestDestinationsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"destinations", "delete",
			"--id", "id",
		)
	})
}
