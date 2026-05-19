// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestHeatmapPagesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"heatmap-pages", "list",
			"--max-items", "10",
			"--from", "2026-04-01",
			"--to", "2026-04-30",
			"--browser-name", "Chrome",
			"--country", "x",
			"--cursor", "cursor",
			"--limit", "25",
			"--region", "x",
			"--search", "/pricing",
			"--sort-by", "CLICKS",
			"--sort-dir", "ASC",
		)
	})
}

func TestHeatmapPagesSummary(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"heatmap-pages", "summary",
			"--breakpoint", "desktop",
			"--from", "2026-04-01",
			"--page-key", "https://example.com/pricing",
			"--to", "2026-04-30",
			"--browser-name", "Chrome",
			"--country", "x",
			"--region", "x",
		)
	})
}
