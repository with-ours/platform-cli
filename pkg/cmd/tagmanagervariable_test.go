// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestTagManagerVariablesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-variables", "list",
			"--max-items", "10",
			"--tag-manager-id", "x",
			"--cursor", "cursor",
			"--limit", "25",
		)
	})
}

func TestTagManagerVariablesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-variables", "create",
			"--name", "name",
			"--parameters", "{foo: bar}",
			"--tag-manager-id", "x",
			"--type", "type",
			"--variable", "Variable",
			"--default-value", "{foo: bar}",
			"--enabled=true",
			"--look-up-table", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"parameters:\n" +
			"  foo: bar\n" +
			"tagManagerId: x\n" +
			"type: type\n" +
			"Variable: Variable\n" +
			"defaultValue:\n" +
			"  foo: bar\n" +
			"enabled: true\n" +
			"lookUpTable:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tag-manager-variables", "create",
		)
	})
}

func TestTagManagerVariablesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-variables", "retrieve",
			"--id", "id",
		)
	})
}

func TestTagManagerVariablesUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-variables", "update",
			"--id", "id",
			"--default-value", "{foo: bar}",
			"--enabled=true",
			"--look-up-table", "{foo: bar}",
			"--name", "name",
			"--parameters", "{foo: bar}",
			"--type", "type",
			"--variable", "Variable",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"defaultValue:\n" +
			"  foo: bar\n" +
			"enabled: true\n" +
			"lookUpTable:\n" +
			"  foo: bar\n" +
			"name: name\n" +
			"parameters:\n" +
			"  foo: bar\n" +
			"type: type\n" +
			"Variable: Variable\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tag-manager-variables", "update",
			"--id", "id",
		)
	})
}

func TestTagManagerVariablesDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-variables", "delete",
			"--id", "id",
		)
	})
}

func TestTagManagerVariablesTypes(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-variables", "types",
		)
	})
}
