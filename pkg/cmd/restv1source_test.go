// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/ours-privacy-platform-cli/internal/mocktest"
)

func TestRestV1SourcesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:sources", "create",
	)
}

func TestRestV1SourcesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:sources", "retrieve",
		"--id", "id",
	)
}

func TestRestV1SourcesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:sources", "update",
		"--id", "id",
	)
}

func TestRestV1SourcesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:sources", "list",
	)
}

func TestRestV1SourcesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:sources", "delete",
		"--id", "id",
	)
}
