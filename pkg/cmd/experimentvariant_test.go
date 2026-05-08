// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
	"github.com/with-ours/platform-cli/internal/requestflag"
)

func TestExperimentVariantsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-variants", "list",
			"--max-items", "10",
			"--experiment-id", "08524dc8-5289-48e8-bf40-b3a7cfa6ca0a",
			"--cursor", "cursor",
			"--limit", "200",
		)
	})
}

func TestExperimentVariantsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-variants", "create",
			"--experiment-id", "x",
			"--name", "Variant B",
			"--weight", "50",
			"--dom-modification", "[{action: customCss, selector: '#hero-headline', attribute: {}, styles: [{property: background-color, value: '#10B981'}], value: Start your free trial}]",
			"--is-control=false",
			"--redirect-url", "https://www.example.com/pricing-v2",
			"--variant-type", "dom_modifications",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(experimentVariantsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-variants", "create",
			"--experiment-id", "x",
			"--name", "Variant B",
			"--weight", "50",
			"--dom-modification.action", "customCss",
			"--dom-modification.selector", "#hero-headline",
			"--dom-modification.attribute", "{}",
			"--dom-modification.styles", "[{property: background-color, value: '#10B981'}]",
			"--dom-modification.value", "Start your free trial",
			"--is-control=false",
			"--redirect-url", "https://www.example.com/pricing-v2",
			"--variant-type", "dom_modifications",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"experimentId: x\n" +
			"name: Variant B\n" +
			"weight: 50\n" +
			"domModifications:\n" +
			"  - action: customCss\n" +
			"    selector: '#hero-headline'\n" +
			"    attribute: {}\n" +
			"    styles:\n" +
			"      - property: background-color\n" +
			"        value: '#10B981'\n" +
			"    value: Start your free trial\n" +
			"isControl: false\n" +
			"redirectUrl: https://www.example.com/pricing-v2\n" +
			"variantType: dom_modifications\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiment-variants", "create",
		)
	})
}

func TestExperimentVariantsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-variants", "retrieve",
			"--id", "id",
		)
	})
}

func TestExperimentVariantsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-variants", "update",
			"--id", "id",
			"--dom-modification", "[{action: customCss, selector: '#hero-headline', attribute: {}, styles: [{property: background-color, value: '#10B981'}], value: Start your free trial}]",
			"--is-control=true",
			"--name", "name",
			"--redirect-url", "redirectUrl",
			"--variant-type", "dom_modifications",
			"--weight", "50",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(experimentVariantsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-variants", "update",
			"--id", "id",
			"--dom-modification.action", "customCss",
			"--dom-modification.selector", "#hero-headline",
			"--dom-modification.attribute", "{}",
			"--dom-modification.styles", "[{property: background-color, value: '#10B981'}]",
			"--dom-modification.value", "Start your free trial",
			"--is-control=true",
			"--name", "name",
			"--redirect-url", "redirectUrl",
			"--variant-type", "dom_modifications",
			"--weight", "50",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"domModifications:\n" +
			"  - action: customCss\n" +
			"    selector: '#hero-headline'\n" +
			"    attribute: {}\n" +
			"    styles:\n" +
			"      - property: background-color\n" +
			"        value: '#10B981'\n" +
			"    value: Start your free trial\n" +
			"isControl: true\n" +
			"name: name\n" +
			"redirectUrl: redirectUrl\n" +
			"variantType: dom_modifications\n" +
			"weight: 50\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiment-variants", "update",
			"--id", "id",
		)
	})
}

func TestExperimentVariantsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-variants", "delete",
			"--id", "id",
		)
	})
}
