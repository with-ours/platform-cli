// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestTagManagerTriggersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-triggers", "list",
			"--max-items", "10",
			"--tag-manager-id", "x",
			"--cursor", "cursor",
			"--limit", "25",
		)
	})
}

func TestTagManagerTriggersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-triggers", "create",
			"--condition", "{foo: bar}",
			"--name", "name",
			"--parameters", "{foo: bar}",
			"--tag-manager-id", "x",
			"--type", "type",
			"--enabled=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"conditions:\n" +
			"  - foo: bar\n" +
			"name: name\n" +
			"parameters:\n" +
			"  foo: bar\n" +
			"tagManagerId: x\n" +
			"type: type\n" +
			"enabled: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tag-manager-triggers", "create",
		)
	})
}

func TestTagManagerTriggersRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-triggers", "retrieve",
			"--id", "id",
		)
	})
}

func TestTagManagerTriggersUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-triggers", "update",
			"--id", "id",
			"--condition", "{foo: bar}",
			"--enabled=true",
			"--name", "name",
			"--parameters", "{foo: bar}",
			"--type", "type",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"conditions:\n" +
			"  - foo: bar\n" +
			"enabled: true\n" +
			"name: name\n" +
			"parameters:\n" +
			"  foo: bar\n" +
			"type: type\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tag-manager-triggers", "update",
			"--id", "id",
		)
	})
}

func TestTagManagerTriggersDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-triggers", "delete",
			"--id", "id",
		)
	})
}

func TestTagManagerTriggersTypes(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-triggers", "types",
		)
	})
}
