// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
	"github.com/with-ours/platform-cli/internal/requestflag"
)

func TestDefaultMappingsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"default-mappings", "list",
		)
	})
}

func TestDefaultMappingsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"default-mappings", "retrieve",
			"--id", "id",
		)
	})
}

func TestDefaultMappingsReplace(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"default-mappings", "replace",
			"--id", "id",
			"--mapping", "{map: map, property: property, modification: CamelCase}",
			"--is-enabled=true",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(defaultMappingsReplace)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"default-mappings", "replace",
			"--id", "id",
			"--mapping.map", "map",
			"--mapping.property", "property",
			"--mapping.modification", "CamelCase",
			"--is-enabled=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"mappings:\n" +
			"  - map: map\n" +
			"    property: property\n" +
			"    modification: CamelCase\n" +
			"isEnabled: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"default-mappings", "replace",
			"--id", "id",
		)
	})
}
