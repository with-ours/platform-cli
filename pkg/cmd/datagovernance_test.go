// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
	"github.com/with-ours/platform-cli/internal/requestflag"
)

func TestDataGovernanceList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"data-governance", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "25",
		)
	})
}

func TestDataGovernanceCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"data-governance", "create",
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
			"data-governance", "create",
		)
	})
}

func TestDataGovernanceRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"data-governance", "retrieve",
			"--id", "id",
		)
	})
}

func TestDataGovernanceUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"data-governance", "update",
			"--id", "id",
			"--category", "[{description: description, destinationIds: [string], logic: {AND: [{}], condition: {operator: Is, property: property, value: value}, NOT: {}, OR: [{}]}, name: name, priority: 0}]",
			"--is-enabled=true",
			"--name", "name",
			"--notes", "notes",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(dataGovernanceUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"data-governance", "update",
			"--id", "id",
			"--category.description", "description",
			"--category.destination-ids", "[string]",
			"--category.logic", "{AND: [{}], condition: {operator: Is, property: property, value: value}, NOT: {}, OR: [{}]}",
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
			"    logic:\n" +
			"      AND:\n" +
			"        - {}\n" +
			"      condition:\n" +
			"        operator: Is\n" +
			"        property: property\n" +
			"        value: value\n" +
			"      NOT: {}\n" +
			"      OR:\n" +
			"        - {}\n" +
			"    name: name\n" +
			"    priority: 0\n" +
			"isEnabled: true\n" +
			"name: name\n" +
			"notes: notes\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"data-governance", "update",
			"--id", "id",
		)
	})
}

func TestDataGovernanceDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"data-governance", "delete",
			"--id", "id",
		)
	})
}
