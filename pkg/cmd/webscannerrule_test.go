// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestWebScannerRulesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web-scanner-rules", "list",
			"--scanner-id", "x",
		)
	})
}

func TestWebScannerRulesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web-scanner-rules", "create",
			"--name", "x",
			"--priority", "1",
			"--scanner-id", "x",
			"--cookie-pattern", "string",
			"--domain-pattern", "string",
			"--notes", "notes",
			"--reason", "approved",
			"--script-pattern", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: x\n" +
			"priority: 1\n" +
			"scannerId: x\n" +
			"cookiePatterns:\n" +
			"  - string\n" +
			"domainPatterns:\n" +
			"  - string\n" +
			"notes: notes\n" +
			"reason: approved\n" +
			"scriptPatterns:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"web-scanner-rules", "create",
		)
	})
}

func TestWebScannerRulesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web-scanner-rules", "retrieve",
			"--id", "id",
		)
	})
}

func TestWebScannerRulesUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web-scanner-rules", "update",
			"--id", "id",
			"--cookie-pattern", "string",
			"--domain-pattern", "string",
			"--name", "x",
			"--notes", "notes",
			"--priority", "1",
			"--reason", "approved",
			"--script-pattern", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"cookiePatterns:\n" +
			"  - string\n" +
			"domainPatterns:\n" +
			"  - string\n" +
			"name: x\n" +
			"notes: notes\n" +
			"priority: 1\n" +
			"reason: approved\n" +
			"scriptPatterns:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"web-scanner-rules", "update",
			"--id", "id",
		)
	})
}

func TestWebScannerRulesDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"web-scanner-rules", "delete",
			"--id", "id",
		)
	})
}
