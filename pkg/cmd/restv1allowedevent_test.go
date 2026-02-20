// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/ours-privacy-platform-cli/internal/mocktest"
)

func TestRestV1AllowedEventsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:allowed-events", "create",
	)
}

func TestRestV1AllowedEventsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:allowed-events", "retrieve",
		"--id", "id",
	)
}

func TestRestV1AllowedEventsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:allowed-events", "list",
	)
}

func TestRestV1AllowedEventsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:allowed-events", "delete",
		"--id", "id",
	)
}
