// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gha

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetOutputCI(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "" {
		t.Skip("Skipping as not running within a real action")
	}

	require.NoError(t, SetOutput("FOO_OUT", os.Getenv("FOO")+"test"))
	require.NoError(t, SetOutput("bar_out", os.Getenv("BAR")+"test"))
}

func TestSetOutput(t *testing.T) {
	tmpDir := t.TempDir()

	for name, tc := range map[string]struct {
		create   bool
		contents string
		outputs  [][2]string
		expected string
	}{
		"base": {
			create:   true,
			contents: "",
			outputs: [][2]string{
				{"foo", "bar"},
			},
			expected: "foo=bar\n",
		},
		"appending": {
			create:   true,
			contents: "foo=bar\n",
			outputs: [][2]string{
				{"bar", "baz"},
			},
			expected: `foo=bar
bar=baz
`,
		},
		"creating": {
			create:   false,
			contents: "",
			outputs: [][2]string{
				{"foo", "bar"},
			},
			expected: "foo=bar\n",
		},
		"multiple": {
			create: true,
			contents: `1=1
2=2
`,
			outputs: [][2]string{
				{"3", "3"},
				{"4", "4"},
			},
			expected: `1=1
2=2
3=3
4=4
`,
		},
	} {
		t.Run(name, func(t *testing.T) {
			filePath := path.Join(tmpDir, name)
			t.Setenv(githubOutputFile, filePath)

			if tc.create {
				f, err := os.Create(filePath)
				require.NoError(t, err)

				_, err = fmt.Fprint(f, tc.contents)
				require.NoError(t, err)
			}

			for _, output := range tc.outputs {
				err := SetOutput(output[0], output[1])
				require.NoError(t, err)
			}

			actual, err := os.ReadFile(filePath)
			require.NoError(t, err)

			require.Equal(t, tc.expected, string(actual))
		})
	}
}
