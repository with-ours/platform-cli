// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
	"github.com/with-ours/platform-cli/internal/requestflag"
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
			"--category", "[{description: description, destinationIds: [string], logic: {AND: [{AND: [{condition: {operator: Contains, property: property, value: value}}], condition: {operator: Contains, property: property, value: value}, NOT: {condition: {operator: Contains, property: property, value: value}}, OR: [{condition: {operator: Contains, property: property, value: value}}]}], condition: {operator: Contains, property: property, value: value}, NOT: {AND: [{condition: {operator: Contains, property: property, value: value}}], condition: {operator: Contains, property: property, value: value}, NOT: {condition: {operator: Contains, property: property, value: value}}, OR: [{condition: {operator: Contains, property: property, value: value}}]}, OR: [{AND: [{condition: {operator: Contains, property: property, value: value}}], condition: {operator: Contains, property: property, value: value}, NOT: {condition: {operator: Contains, property: property, value: value}}, OR: [{condition: {operator: Contains, property: property, value: value}}]}]}, name: name, priority: 0}]",
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
			"--category.logic", "{AND: [{AND: [{condition: {operator: Contains, property: property, value: value}}], condition: {operator: Contains, property: property, value: value}, NOT: {condition: {operator: Contains, property: property, value: value}}, OR: [{condition: {operator: Contains, property: property, value: value}}]}], condition: {operator: Contains, property: property, value: value}, NOT: {AND: [{condition: {operator: Contains, property: property, value: value}}], condition: {operator: Contains, property: property, value: value}, NOT: {condition: {operator: Contains, property: property, value: value}}, OR: [{condition: {operator: Contains, property: property, value: value}}]}, OR: [{AND: [{condition: {operator: Contains, property: property, value: value}}], condition: {operator: Contains, property: property, value: value}, NOT: {condition: {operator: Contains, property: property, value: value}}, OR: [{condition: {operator: Contains, property: property, value: value}}]}]}",
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
			"        - AND:\n" +
			"            - condition:\n" +
			"                operator: Contains\n" +
			"                property: property\n" +
			"                value: value\n" +
			"          condition:\n" +
			"            operator: Contains\n" +
			"            property: property\n" +
			"            value: value\n" +
			"          NOT:\n" +
			"            condition:\n" +
			"              operator: Contains\n" +
			"              property: property\n" +
			"              value: value\n" +
			"          OR:\n" +
			"            - condition:\n" +
			"                operator: Contains\n" +
			"                property: property\n" +
			"                value: value\n" +
			"      condition:\n" +
			"        operator: Contains\n" +
			"        property: property\n" +
			"        value: value\n" +
			"      NOT:\n" +
			"        AND:\n" +
			"          - condition:\n" +
			"              operator: Contains\n" +
			"              property: property\n" +
			"              value: value\n" +
			"        condition:\n" +
			"          operator: Contains\n" +
			"          property: property\n" +
			"          value: value\n" +
			"        NOT:\n" +
			"          condition:\n" +
			"            operator: Contains\n" +
			"            property: property\n" +
			"            value: value\n" +
			"        OR:\n" +
			"          - condition:\n" +
			"              operator: Contains\n" +
			"              property: property\n" +
			"              value: value\n" +
			"      OR:\n" +
			"        - AND:\n" +
			"            - condition:\n" +
			"                operator: Contains\n" +
			"                property: property\n" +
			"                value: value\n" +
			"          condition:\n" +
			"            operator: Contains\n" +
			"            property: property\n" +
			"            value: value\n" +
			"          NOT:\n" +
			"            condition:\n" +
			"              operator: Contains\n" +
			"              property: property\n" +
			"              value: value\n" +
			"          OR:\n" +
			"            - condition:\n" +
			"                operator: Contains\n" +
			"                property: property\n" +
			"                value: value\n" +
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
