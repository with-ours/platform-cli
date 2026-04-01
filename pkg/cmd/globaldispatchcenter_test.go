// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestGlobalDispatchCentersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
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
			"--category", "[{}]",
			"--is-enabled=true",
			"--name", "name",
			"--notes", "notes",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"categories:\n" +
			"  - {}\n" +
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

func TestGlobalDispatchCentersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"global-dispatch-centers", "list",
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
