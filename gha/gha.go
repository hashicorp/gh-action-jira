// Copyright IBM Corp. 2020, 2025
// SPDX-License-Identifier: MPL-2.0

package gha

import (
	"fmt"
	"os"
)

const (
	githubOutputFile = "GITHUB_OUTPUT"
)

func SetOutput(key, value string) error {
	// Set the action's output by writing to a file
	// See https://docs.github.com/en/actions/using-workflows/workflow-commands-for-github-actions#setting-an-output-parameter
	f, err := os.OpenFile(os.Getenv(githubOutputFile), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = fmt.Fprintf(f, "%s=%s\n", key, value); err != nil {
		return err
	}

	return nil
}
