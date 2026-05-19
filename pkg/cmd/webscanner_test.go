// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestWebScannersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web-scanners", "list",
		)
	})
}

func TestWebScannersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web-scanners", "create",
			"--root-domain", "x",
			"--excluded-pattern", "[string]",
			"--included-url", "[string]",
			"--name", "name",
			"--status", "Disabled",
			"--url-limit", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"rootDomain: x\n" +
			"excludedPatterns:\n" +
			"  - string\n" +
			"includedUrls:\n" +
			"  - string\n" +
			"name: name\n" +
			"status: Disabled\n" +
			"urlLimit: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"web-scanners", "create",
		)
	})
}

func TestWebScannersRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web-scanners", "retrieve",
			"--id", "id",
		)
	})
}

func TestWebScannersUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web-scanners", "update",
			"--id", "id",
			"--excluded-pattern", "[string]",
			"--included-url", "[string]",
			"--name", "name",
			"--root-domain", "rootDomain",
			"--status", "Disabled",
			"--url-limit", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"excludedPatterns:\n" +
			"  - string\n" +
			"includedUrls:\n" +
			"  - string\n" +
			"name: name\n" +
			"rootDomain: rootDomain\n" +
			"status: Disabled\n" +
			"urlLimit: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"web-scanners", "update",
			"--id", "id",
		)
	})
}

func TestWebScannersDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web-scanners", "delete",
			"--id", "id",
		)
	})
}

func TestWebScannersTrigger(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web-scanners", "trigger",
			"--id", "id",
		)
	})
}
