// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package format

import (
	"bytes"
	"os/exec"
	"strings"
)

// GitHubToJira executes `docker run -i --rm pandoc/core` to convert from GitHub
// flavoured markdown to Jira markdown
func GitHubToJira(s string) (string, error) {
	cmd := exec.Command("docker", "run", "--interactive", "--rm", "pandoc/core:2.11.0.4", "--from=gfm", "--to=jira")
	cmd.Stdin = strings.NewReader(s)
	out := bytes.Buffer{}
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
