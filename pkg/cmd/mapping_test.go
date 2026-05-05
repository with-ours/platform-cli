// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
	"github.com/with-ours/platform-cli/internal/requestflag"
)

func TestMappingsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mappings", "list",
			"--entity-id", "00000000-0000-0000-0000-000000000000",
		)
	})
}

func TestMappingsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mappings", "create",
			"--allowed-event-id", "allowedEventId",
			"--destination-id", "destinationId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"allowedEventId: allowedEventId\n" +
			"destinationId: destinationId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"mappings", "create",
		)
	})
}

func TestMappingsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mappings", "retrieve",
			"--id", "id",
		)
	})
}

func TestMappingsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mappings", "update",
			"--id", "id",
			"--logic", "{}",
			"--mapping", "{map: map, property: property, modification: modification}",
			"--name", "name",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(mappingsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mappings", "update",
			"--id", "id",
			"--logic", "{}",
			"--mapping.map", "map",
			"--mapping.property", "property",
			"--mapping.modification", "modification",
			"--name", "name",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"logic: {}\n" +
			"mappings:\n" +
			"  - map: map\n" +
			"    property: property\n" +
			"    modification: modification\n" +
			"name: name\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"mappings", "update",
			"--id", "id",
		)
	})
}

func TestMappingsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mappings", "delete",
			"--id", "id",
		)
	})
}
