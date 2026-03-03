// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestAllowedEventsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"allowed-events", "create",
		"--api-key", "string",
		"--name", "name",
		"--destination-ids", "[string]",
	)
}

func TestAllowedEventsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"allowed-events", "retrieve",
		"--api-key", "string",
		"--id", "id",
	)
}

func TestAllowedEventsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"allowed-events", "list",
		"--api-key", "string",
	)
}

func TestAllowedEventsDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"allowed-events", "delete",
		"--api-key", "string",
		"--id", "id",
	)
}
