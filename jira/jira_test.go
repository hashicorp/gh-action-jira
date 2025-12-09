// Copyright IBM Corp. 2020, 2025
// SPDX-License-Identifier: MPL-2.0

package jira

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateURL(t *testing.T) {
	url := generateURL("https://org.atlassian.net/", "/rest/api/3/search", url.Values{
		"jql": {"project = FOO"},
	})
	require.Equal(t, "https://org.atlassian.net/rest/api/3/search?jql=project+%3D+FOO", url)
}
