// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestDestinationsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"destinations", "create",
			"--type", "AWSEventBridge",
			"--name", "name",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"type: AWSEventBridge\n" +
			"name: name\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"destinations", "create",
		)
	})
}

func TestDestinationsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"destinations", "retrieve",
			"--id", "id",
		)
	})
}

func TestDestinationsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"destinations", "update",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"status: Disabled\n" +
			"facebookConversionAPIKey: facebookConversionAPIKey\n" +
			"facebookPixelId: facebookPixelId\n" +
			"g4AnalyticsApiKey: g4AnalyticsApiKey\n" +
			"g4AnalyticsMeasurementId: g4AnalyticsMeasurementId\n" +
			"g4AnalyticsTrackOnPage: true\n" +
			"hashingSalt: hashingSalt\n" +
			"httpDestinationUrl: httpDestinationUrl\n" +
			"limitedToSourceIds:\n" +
			"  - string\n" +
			"managerGoogleCustomerId: managerGoogleCustomerId\n" +
			"name: name\n" +
			"projectAPIKey: projectAPIKey\n" +
			"projectToken: projectToken\n" +
			"selectedAccountId: selectedAccountId\n" +
			"settings:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"destinations", "update",
			"--id", "id",
		)
	})
}

func TestDestinationsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"destinations", "list",
		)
	})
}

func TestDestinationsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"destinations", "delete",
			"--id", "id",
		)
	})
}
