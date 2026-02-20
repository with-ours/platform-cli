// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
	"github.com/with-ours/platform-cli/internal/requestflag"
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
		"--category", "[{description: description, destinationIds: [string], logic: {}, name: name, priority: 0}]",
		"--is-enabled=true",
		"--name", "name",
		"--notes", "notes",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(globalDispatchCentersUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-dispatch-centers", "update",
		"--id", "id",
		"--category.description", "description",
		"--category.destination-ids", "[string]",
		"--category.logic", "{}",
		"--category.name", "name",
		"--category.priority", "0",
		"--is-enabled=true",
		"--name", "name",
		"--notes", "notes",
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
