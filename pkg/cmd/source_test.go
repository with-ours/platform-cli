// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
	"github.com/with-ours/platform-cli/internal/requestflag"
)

func TestSourcesCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"sources", "create",
		"--type", "AlchemerWebhook",
		"--name", "name",
	)
}

func TestSourcesRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"sources", "retrieve",
		"--id", "id",
	)
}

func TestSourcesUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(sourcesUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
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
}

func TestSourcesList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"sources", "list",
	)
}

func TestSourcesDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"sources", "delete",
		"--id", "id",
	)
}
