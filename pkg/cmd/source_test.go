// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestSourcesCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"sources", "create",
	)
}

func TestSourcesRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"sources", "retrieve",
		"--id", "id",
	)
}

func TestSourcesUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"sources", "update",
		"--id", "id",
	)
}

func TestSourcesList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"sources", "list",
	)
}

func TestSourcesDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"sources", "delete",
		"--id", "id",
	)
}
