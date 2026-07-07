// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestFunnelsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"funnels", "list",
		)
	})
}

func TestFunnelsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"funnels", "retrieve",
			"--id", "id",
		)
	})
}

func TestFunnelsResults(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"funnels", "results",
			"--id", "id",
			"--from", "2026-06-01",
			"--to", "2026-06-30",
			"--attribution-type", "INITIAL",
			"--device-type", "DESKTOP",
			"--utm-campaign", "x",
			"--utm-content", "x",
			"--utm-medium", "x",
			"--utm-name", "x",
			"--utm-source", "x",
			"--utm-term", "x",
		)
	})
}
