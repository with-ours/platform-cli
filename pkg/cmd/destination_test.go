// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestDestinationsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"destinations", "create",
		"--api-key", "string",
		"--type", "AWSEventBridge",
		"--name", "name",
	)
}

func TestDestinationsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"destinations", "retrieve",
		"--api-key", "string",
		"--id", "id",
	)
}

func TestDestinationsUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"destinations", "update",
		"--api-key", "string",
		"--id", "id",
		"--status", "Disabled",
		"--facebook-conversion-api-key", "facebookConversionAPIKey",
		"--facebook-pixel-id", "facebookPixelId",
		"--g4-analytics-api-key", "g4AnalyticsApiKey",
		"--g4-analytics-measurement-id", "g4AnalyticsMeasurementId",
		"--g4-analytics-track-on-page=true",
		"--hashing-salt", "hashingSalt",
		"--http-destination-url", "httpDestinationUrl",
		"--limited-to-source-id", "[string]",
		"--manager-google-customer-id", "managerGoogleCustomerId",
		"--name", "name",
		"--project-api-key", "projectAPIKey",
		"--project-token", "projectToken",
		"--selected-account-id", "selectedAccountId",
		"--settings", "{foo: bar}",
	)
}

func TestDestinationsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"destinations", "list",
		"--api-key", "string",
	)
}

func TestDestinationsDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"destinations", "delete",
		"--api-key", "string",
		"--id", "id",
	)
}
