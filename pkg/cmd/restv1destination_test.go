// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/ours-privacy-platform-cli/internal/mocktest"
)

func TestRestV1DestinationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:destinations", "create",
	)
}

func TestRestV1DestinationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:destinations", "retrieve",
		"--id", "id",
	)
}

func TestRestV1DestinationsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:destinations", "update",
		"--id", "id",
	)
}

func TestRestV1DestinationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:destinations", "list",
	)
}

func TestRestV1DestinationsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:destinations", "delete",
		"--id", "id",
	)
}
