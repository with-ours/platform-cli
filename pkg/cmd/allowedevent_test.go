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
		"--name", "name",
		"--destination-ids", "[string]",
	)
}

func TestAllowedEventsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"allowed-events", "retrieve",
		"--id", "id",
	)
}

func TestAllowedEventsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"allowed-events", "list",
	)
}

func TestAllowedEventsDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"allowed-events", "delete",
		"--id", "id",
	)
}
