// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestAttributionInitial(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"attribution", "initial",
			"--event-name", "purchase",
			"--from", "2026-05-01",
			"--to", "2026-06-30",
			"--attribution-type", "INITIAL",
			"--utm-campaign", "x",
			"--utm-content", "x",
			"--utm-medium", "x",
			"--utm-name", "x",
			"--utm-source", "x",
			"--utm-term", "x",
		)
	})
}

func TestAttributionLastTouch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"attribution", "last-touch",
			"--event-name", "purchase",
			"--from", "2026-05-01",
			"--to", "2026-06-30",
			"--attribution-type", "INITIAL",
			"--utm-campaign", "x",
			"--utm-content", "x",
			"--utm-medium", "x",
			"--utm-name", "x",
			"--utm-source", "x",
			"--utm-term", "x",
		)
	})
}

func TestAttributionConversion(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"attribution", "conversion",
			"--attribution-model", "FIRST_TOUCH",
			"--event-name", "purchase",
			"--from", "2026-06-01",
			"--to", "2026-06-30",
			"--limit", "1",
			"--lookback-window", "SEVEN_DAYS",
			"--web-source-id", "x",
		)
	})
}

func TestAttributionAudienceConversion(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"attribution", "audience-conversion",
			"--event-name", "purchase",
			"--from", "2026-05-01",
			"--to", "2026-06-30",
			"--attribution-window", "IN_RANGE",
			"--value-property", "revenue",
		)
	})
}

func TestAttributionUtmComparison(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"attribution", "utm-comparison",
			"--combos", `[{"utmSource":"google","utmMedium":"cpc"},{"utmSource":"meta"}]`,
			"--event-name", "purchase",
			"--from", "2026-06-01",
			"--to", "2026-06-30",
		)
	})
}
