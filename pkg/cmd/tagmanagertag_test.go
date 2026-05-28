// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestTagManagerTagsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-tags", "list",
			"--max-items", "10",
			"--tag-manager-id", "x",
			"--cursor", "cursor",
			"--limit", "25",
		)
	})
}

func TestTagManagerTagsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-tags", "create",
			"--fire-trigger-id", "string",
			"--name", "name",
			"--parameters", "{foo: bar}",
			"--tag-manager-id", "x",
			"--type", "type",
			"--block-trigger-id", "[string]",
			"--enabled=true",
			"--priority", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fireTriggerIds:\n" +
			"  - string\n" +
			"name: name\n" +
			"parameters:\n" +
			"  foo: bar\n" +
			"tagManagerId: x\n" +
			"type: type\n" +
			"blockTriggerIds:\n" +
			"  - string\n" +
			"enabled: true\n" +
			"priority: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tag-manager-tags", "create",
		)
	})
}

func TestTagManagerTagsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-tags", "retrieve",
			"--id", "id",
		)
	})
}

func TestTagManagerTagsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-tags", "update",
			"--id", "id",
			"--block-trigger-id", "[string]",
			"--enabled=true",
			"--fire-trigger-id", "string",
			"--name", "name",
			"--parameters", "{foo: bar}",
			"--priority", "0",
			"--type", "type",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"blockTriggerIds:\n" +
			"  - string\n" +
			"enabled: true\n" +
			"fireTriggerIds:\n" +
			"  - string\n" +
			"name: name\n" +
			"parameters:\n" +
			"  foo: bar\n" +
			"priority: 0\n" +
			"type: type\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tag-manager-tags", "update",
			"--id", "id",
		)
	})
}

func TestTagManagerTagsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-tags", "delete",
			"--id", "id",
		)
	})
}

func TestTagManagerTagsTypes(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-tags", "types",
		)
	})
}
