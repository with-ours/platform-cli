// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/ours-privacy-platform-cli/internal/mocktest"
)

func TestRestV1GlobalDispatchCentersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:global-dispatch-centers", "create",
	)
}

func TestRestV1GlobalDispatchCentersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:global-dispatch-centers", "retrieve",
		"--id", "id",
	)
}

func TestRestV1GlobalDispatchCentersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:global-dispatch-centers", "update",
		"--id", "id",
	)
}

func TestRestV1GlobalDispatchCentersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:global-dispatch-centers", "list",
	)
}

func TestRestV1GlobalDispatchCentersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:global-dispatch-centers", "delete",
		"--id", "id",
	)
}
