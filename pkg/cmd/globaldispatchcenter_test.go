// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestGlobalDispatchCentersCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-dispatch-centers", "create",
	)
}

func TestGlobalDispatchCentersRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-dispatch-centers", "retrieve",
		"--id", "id",
	)
}

func TestGlobalDispatchCentersUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-dispatch-centers", "update",
		"--id", "id",
	)
}

func TestGlobalDispatchCentersList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-dispatch-centers", "list",
	)
}

func TestGlobalDispatchCentersDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-dispatch-centers", "delete",
		"--id", "id",
	)
}
