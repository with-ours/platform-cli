// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestShortLinksList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"short-links", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "25",
			"--name-contains", "nameContains",
			"--status", "Disabled",
		)
	})
}

func TestShortLinksCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"short-links", "create",
			"--name", "Spring Sale QR",
			"--qr", "{}",
			"--redirect-url", "https://example.com/spring",
			"--utm", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Spring Sale QR\n" +
			"qr: {}\n" +
			"redirectUrl: https://example.com/spring\n" +
			"utm: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"short-links", "create",
		)
	})
}

func TestShortLinksRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"short-links", "retrieve",
			"--id", "id",
		)
	})
}

func TestShortLinksUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"short-links", "update",
			"--id", "id",
			"--name", "name",
			"--qr", "{}",
			"--redirect-url", "redirectUrl",
			"--status", "status",
			"--utm", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"qr: {}\n" +
			"redirectUrl: redirectUrl\n" +
			"status: status\n" +
			"utm: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"short-links", "update",
			"--id", "id",
		)
	})
}

func TestShortLinksDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"short-links", "delete",
			"--id", "id",
		)
	})
}

func TestShortLinksResults(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"short-links", "results",
			"--id", "id",
			"--from", "2026-06-01",
			"--to", "2026-06-30",
			"--exclude-bots=true",
			"--granularity", "DAILY",
		)
	})
}
