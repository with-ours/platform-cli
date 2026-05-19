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
			"--max-items", "10",
			"--entity-id", "00000000-0000-0000-0000-000000000000",
			"--cursor", "cursor",
			"--limit", "1000",
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
			"--entity-id", "entityId",
			"--insert-after-idx", "0",
			"--is-enabled=true",
			"--logic", "{AND: [{}], condition: {operator: Is, property: property, value: value}, NOT: {}, OR: [{}]}",
			"--mapping", "{map: map, property: property, modification: CamelCase}",
			"--name", "name",
			"--template-id", "templateId",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(mappingsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mappings", "create",
			"--allowed-event-id", "allowedEventId",
			"--destination-id", "destinationId",
			"--entity-id", "entityId",
			"--insert-after-idx", "0",
			"--is-enabled=true",
			"--logic.and", "[{}]",
			"--logic.condition", "{operator: Is, property: property, value: value}",
			"--logic.not", "{}",
			"--logic.or", "[{}]",
			"--mapping.map", "map",
			"--mapping.property", "property",
			"--mapping.modification", "CamelCase",
			"--name", "name",
			"--template-id", "templateId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"allowedEventId: allowedEventId\n" +
			"destinationId: destinationId\n" +
			"entityId: entityId\n" +
			"insertAfterIdx: 0\n" +
			"isEnabled: true\n" +
			"logic:\n" +
			"  AND:\n" +
			"    - {}\n" +
			"  condition:\n" +
			"    operator: Is\n" +
			"    property: property\n" +
			"    value: value\n" +
			"  NOT: {}\n" +
			"  OR:\n" +
			"    - {}\n" +
			"mappings:\n" +
			"  - map: map\n" +
			"    property: property\n" +
			"    modification: CamelCase\n" +
			"name: name\n" +
			"templateId: templateId\n")
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
			"--is-enabled=true",
			"--logic", "{AND: [{}], condition: {operator: Is, property: property, value: value}, NOT: {}, OR: [{}]}",
			"--mapping", "{map: map, property: property, modification: CamelCase}",
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
			"--is-enabled=true",
			"--logic.and", "[{}]",
			"--logic.condition", "{operator: Is, property: property, value: value}",
			"--logic.not", "{}",
			"--logic.or", "[{}]",
			"--mapping.map", "map",
			"--mapping.property", "property",
			"--mapping.modification", "CamelCase",
			"--name", "name",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"isEnabled: true\n" +
			"logic:\n" +
			"  AND:\n" +
			"    - {}\n" +
			"  condition:\n" +
			"    operator: Is\n" +
			"    property: property\n" +
			"    value: value\n" +
			"  NOT: {}\n" +
			"  OR:\n" +
			"    - {}\n" +
			"mappings:\n" +
			"  - map: map\n" +
			"    property: property\n" +
			"    modification: CamelCase\n" +
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

func TestMappingsReorder(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mappings", "reorder",
			"--uuid", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"uuids:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"mappings", "reorder",
		)
	})
}
