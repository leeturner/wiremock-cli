package wiremock

import (
	"fmt"
	"io"
	"net/http"
)

const userAgentHeader = "User-Agent"
const acceptHeader = "Accept"
const defaultAdminPrefix = "__admin"

type Wiremock struct {
	host        string
	port        string
	adminPrefix string
}

func Init(host string, port string) *Wiremock {
	wm := &Wiremock{
		port:        port,
		host:        host,
		adminPrefix: defaultAdminPrefix,
	}
	return wm
}

func (wm *Wiremock) WithAdminPrefix(newAdminPrefix string) {
	wm.adminPrefix = newAdminPrefix
}

func (wm *Wiremock) performRequest(path string, method string) (body string, err error) {
	url := fmt.Sprintf("%s:%s/%s%s", wm.host, wm.port, wm.adminPrefix, path)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set(userAgentHeader, "Wiremock CLI")
	req.Header.Set(acceptHeader, "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	return string(b), nil
}

// Mapping endpoints

func (wm *Wiremock) GetMappings(id string, limit int, offset int) (body string, err error) {
	if id != "" {
		return wm.performRequest(fmt.Sprintf("/mappings/%s", id), http.MethodGet)
	} else {
		return wm.performRequest(fmt.Sprintf("/mappings?limit=%d&offset=%d", limit, offset), http.MethodGet)
	}
}

func (wm *Wiremock) DeleteMappings(id string) (body string, err error) {
	if id != "" {
		return wm.performRequest(fmt.Sprintf("/mappings/%s", id), http.MethodDelete)
	} else {
		return wm.performRequest("/mappings", http.MethodDelete)
	}
}

func (wm *Wiremock) ResetMappings() (body string, err error) {
	return wm.performRequest("/mappings/reset", http.MethodPost)
}

func (wm *Wiremock) SaveMappings() (body string, err error) {
	return wm.performRequest("/mappings/save", http.MethodPost)
}

func (wm *Wiremock) ImportMappings() (body string, err error) {
	return wm.performRequest("/mappings/import", http.MethodPost)
}

// Scenario endpoints

func (wm *Wiremock) GetScenarios() (body string, err error) {
	return wm.performRequest("/scenarios", http.MethodGet)
}

// Requests endpoints

func (wm *Wiremock) GetRequests(id string, limit int) (body string, err error) {
	if id != "" {
		return wm.performRequest(fmt.Sprintf("/requests/%s", id), http.MethodGet)
	} else {
		return wm.performRequest(fmt.Sprintf("/requests?limit=%d", limit), http.MethodGet)
	}
}

func (wm *Wiremock) GetUnmatchedRequests() (body string, err error) {
	return wm.performRequest("/requests/unmatched", http.MethodGet)
}

func (wm *Wiremock) DeleteRequests(id string) (body string, err error) {
	if id != "" {
		return wm.performRequest(fmt.Sprintf("/requests/%s", id), http.MethodDelete)
	} else {
		return wm.performRequest("/requests", http.MethodDelete)
	}
}

// Recordings endpoints

func (wm *Wiremock) GetRecordingsStatus() (body string, err error) {
	return wm.performRequest("/recordings/status", http.MethodGet)
}

// System endpoints

func (wm *Wiremock) Reset() (body string, err error) {
	return wm.performRequest("/reset", http.MethodPost)
}

func (wm *Wiremock) Shutdown() (body string, err error) {
	return wm.performRequest("/shutdown", http.MethodPost)
}

func (wm *Wiremock) Version() (body string, err error) {
	return wm.performRequest("/version", http.MethodGet)
}
