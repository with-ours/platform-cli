// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
	"github.com/with-ours/platform-cli/internal/requestflag"
)

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
			"--status", "Disabled",
			"--bot-control-mode", "Allow",
			"--bot-score-threshold", "0",
			"--exclude-request-context=true",
			"--name", "name",
			"--probabilistic-identity", "{enabled: true, matchWindowMinutes: 1, maxMatchesPerIp: 1}",
			"--project-api-key", "projectAPIKey",
			"--redirect-url", "redirectUrl",
			"--selected-account-id", "selectedAccountId",
			"--whitelist-domain", "[{}]",
			"--whitelist-ip", "[string]",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(sourcesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sources", "update",
			"--id", "id",
			"--status", "Disabled",
			"--bot-control-mode", "Allow",
			"--bot-score-threshold", "0",
			"--exclude-request-context=true",
			"--name", "name",
			"--probabilistic-identity.enabled=true",
			"--probabilistic-identity.match-window-minutes", "1",
			"--probabilistic-identity.max-matches-per-ip", "1",
			"--project-api-key", "projectAPIKey",
			"--redirect-url", "redirectUrl",
			"--selected-account-id", "selectedAccountId",
			"--whitelist-domain", "[{}]",
			"--whitelist-ip", "[string]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"status: Disabled\n" +
			"botControlMode: Allow\n" +
			"botScoreThreshold: 0\n" +
			"excludeRequestContext: true\n" +
			"name: name\n" +
			"probabilisticIdentity:\n" +
			"  enabled: true\n" +
			"  matchWindowMinutes: 1\n" +
			"  maxMatchesPerIp: 1\n" +
			"projectAPIKey: projectAPIKey\n" +
			"redirectUrl: redirectUrl\n" +
			"selectedAccountId: selectedAccountId\n" +
			"whitelistDomains:\n" +
			"  - {}\n" +
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

func TestSourcesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sources", "list",
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
