// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestVersionsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"versions", "create",
	)
}

func TestVersionsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"versions", "retrieve",
		"--id", "id",
	)
}

func TestVersionsUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"versions", "update",
		"--id", "id",
	)
}

func TestVersionsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"versions", "list",
	)
}
