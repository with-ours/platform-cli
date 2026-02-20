// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/ours-privacy-platform-cli/internal/mocktest"
)

func TestRestV1VersionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:versions", "create",
	)
}

func TestRestV1VersionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:versions", "retrieve",
		"--id", "id",
	)
}

func TestRestV1VersionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:versions", "update",
		"--id", "id",
	)
}

func TestRestV1VersionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rest:v1:versions", "list",
	)
}
