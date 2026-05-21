// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestSourcesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sources", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "25",
			"--name-contains", "nameContains",
			"--status", "Disabled",
			"--type", "AlchemerWebhook",
		)
	})
}

func TestSourcesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sources", "create",
			"--type", "AlchemerWebhook",
			"--name", "name",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"type: AlchemerWebhook\n" +
			"name: name\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"sources", "create",
		)
	})
}

func TestSourcesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sources", "retrieve",
			"--id", "id",
		)
	})
}

func TestSourcesUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sources", "update",
			"--id", "id",
			"--bot-control-mode", "botControlMode",
			"--bot-score-threshold", "0",
			"--exclude-request-context=true",
			"--name", "name",
			"--probabilistic-identity", "{}",
			"--project-api-key", "projectAPIKey",
			"--redirect-url", "redirectUrl",
			"--selected-account-id", "selectedAccountId",
			"--status", "status",
			"--whitelist-domain", "[string]",
			"--whitelist-ip", "[string]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"botControlMode: botControlMode\n" +
			"botScoreThreshold: 0\n" +
			"excludeRequestContext: true\n" +
			"name: name\n" +
			"probabilisticIdentity: {}\n" +
			"projectAPIKey: projectAPIKey\n" +
			"redirectUrl: redirectUrl\n" +
			"selectedAccountId: selectedAccountId\n" +
			"status: status\n" +
			"whitelistDomains:\n" +
			"  - string\n" +
			"whitelistIps:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"sources", "update",
			"--id", "id",
		)
	})
}

func TestSourcesDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sources", "delete",
			"--id", "id",
		)
	})
}

func TestSourcesTokens(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sources", "tokens",
			"--id", "id",
		)
	})
}
