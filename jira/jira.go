// Copyright IBM Corp. 2020, 2025
// SPDX-License-Identifier: MPL-2.0

package jira

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/gh-action-jira/config"
)

func DoRequest(config config.JiraConfig, method, path string, query url.Values, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, generateURL(config.BaseURL, path, query), body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.UserEmail, config.APIToken))))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("API call %s %s failed (%d): %s", method, path, resp.StatusCode, string(bytes))
	}

	return bytes, nil
}

func generateURL(baseURL string, path string, query url.Values) string {
	baseURL = strings.TrimSuffix(baseURL, "/")
	url := fmt.Sprintf("%s%s", baseURL, path)
	queryString := query.Encode()
	if queryString != "" {
		url += fmt.Sprintf("?%s", queryString)
	}
	return url
}
