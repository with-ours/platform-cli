// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
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
			"--bot-control-mode", "botControlMode",
			"--bot-score-threshold", "0",
			"--exclude-request-context=true",
			"--name", "name",
			"--probabilistic-identity", "{}",
			"--project-api-key", "projectAPIKey",
			"--redirect-url", "redirectUrl",
			"--selected-account-id", "selectedAccountId",
			"--whitelist-domain", "[{}]",
			"--whitelist-ip", "[{}]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"status: Disabled\n" +
			"botControlMode: botControlMode\n" +
			"botScoreThreshold: 0\n" +
			"excludeRequestContext: true\n" +
			"name: name\n" +
			"probabilisticIdentity: {}\n" +
			"projectAPIKey: projectAPIKey\n" +
			"redirectUrl: redirectUrl\n" +
			"selectedAccountId: selectedAccountId\n" +
			"whitelistDomains:\n" +
			"  - {}\n" +
			"whitelistIps:\n" +
			"  - {}\n")
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
