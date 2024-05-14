// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package format

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMDToJira(t *testing.T) {
	const testFolder = "testdata"
	files, err := os.ReadDir(testFolder)
	require.NoError(t, err)
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() || !strings.HasSuffix(fileName, ".md") {
			continue
		}

		t.Run(fileName, func(t *testing.T) {
			input, err := os.ReadFile(filepath.Join(testFolder, fileName))
			require.NoError(t, err)
			expected, err := os.ReadFile(filepath.Join(testFolder, fileName[:len(fileName)-3]+".jira"))
			require.NoError(t, err)

			output, err := GitHubToJira(string(input))
			require.NoError(t, err)
			assert.Equal(t, string(expected), output, "mismatched output for %q", fileName)
		})
	}

}
