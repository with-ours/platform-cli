// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
	"github.com/with-ours/platform-cli/internal/requestflag"
)

func TestGlobalDispatchCentersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"global-dispatch-centers", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "25",
		)
	})
}

func TestGlobalDispatchCentersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"global-dispatch-centers", "create",
			"--is-enabled=true",
			"--name", "name",
			"--notes", "notes",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"isEnabled: true\n" +
			"name: name\n" +
			"notes: notes\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"global-dispatch-centers", "create",
		)
	})
}

func TestGlobalDispatchCentersRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"global-dispatch-centers", "retrieve",
			"--id", "id",
		)
	})
}

func TestGlobalDispatchCentersUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"global-dispatch-centers", "update",
			"--id", "id",
			"--category", "[{description: description, destinationIds: [string], logic: {}, name: name, priority: 0}]",
			"--is-enabled=true",
			"--name", "name",
			"--notes", "notes",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(globalDispatchCentersUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"global-dispatch-centers", "update",
			"--id", "id",
			"--category.description", "description",
			"--category.destination-ids", "[string]",
			"--category.logic", "{}",
			"--category.name", "name",
			"--category.priority", "0",
			"--is-enabled=true",
			"--name", "name",
			"--notes", "notes",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"categories:\n" +
			"  - description: description\n" +
			"    destinationIds:\n" +
			"      - string\n" +
			"    logic: {}\n" +
			"    name: name\n" +
			"    priority: 0\n" +
			"isEnabled: true\n" +
			"name: name\n" +
			"notes: notes\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"global-dispatch-centers", "update",
			"--id", "id",
		)
	})
}

func TestGlobalDispatchCentersDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"global-dispatch-centers", "delete",
			"--id", "id",
		)
	})
}
