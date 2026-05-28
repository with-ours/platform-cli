// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/with-ours/platform-cli/internal/mocktest"
)

func TestTagManagerAssetFoldersCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tag-manager-asset-folders", "create",
			"--asset-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--asset-type", "tagManagerTag",
			"--folder-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"assetId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"assetType: tagManagerTag\n" +
			"folderId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tag-manager-asset-folders", "create",
		)
	})
}
