package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type TestRunner struct {
	Path string
}

func To(path string) *TestRunner {
	return &TestRunner{Path: path}
}

func (t *TestRunner) buildApp() *fiber.App {
	return fiber.New()
}

// Get dispatches a new GET request to the test instance.
func (t *TestRunner) Get() *http.Response {
	request := httptest.NewRequest("GET", t.Path, nil)
	response, _ := t.buildApp().Test(request)

	return response
}

// PostRaw dispatches a new POST request to the test instance with a given raw string.
func (t *TestRunner) PostRaw(body string) *http.Response {
	request := httptest.NewRequest("POST", t.Path, strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	response, _ := t.buildApp().Test(request)

	return response
}

// Post dispatches a new POST request to the test instance.
func (t *TestRunner) Post(body interface{}) *http.Response {
	requestBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	request := httptest.NewRequest("POST", t.Path, bytes.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")
	response, _ := t.buildApp().Test(request)

	return response
}
