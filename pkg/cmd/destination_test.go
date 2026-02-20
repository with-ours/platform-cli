// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestDestinationsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"destinations", "create",
	)
}

func TestDestinationsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"destinations", "retrieve",
		"--id", "id",
	)
}

func TestDestinationsUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"destinations", "update",
		"--id", "id",
	)
}

func TestDestinationsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"destinations", "list",
	)
}

func TestDestinationsDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"destinations", "delete",
		"--id", "id",
	)
}
