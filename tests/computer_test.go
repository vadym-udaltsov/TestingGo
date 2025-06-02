package tests

import (
	_struct "TestingGo/internal/model/struct"
	"encoding/json"
	"github.com/dailymotion/allure-go"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestGetComputersFromMockServer(t *testing.T) {
	allure.Test(t,
		allure.Description("Test /sam/v2/orgs/1/computers using mock server and validate response structure"),
		allure.Action(func() {
			allure.Step(allure.Description("Open and read mock JSON file"), allure.Action(func() {
				cwd, err := os.Getwd()
				if err != nil {
					t.Fatalf("Failed to get working directory: %v", err)
				}

				t.Logf("Current working dir: %s", cwd)

				jsonPath := filepath.Join(cwd, "internal", "model", "mock_data", "computers_response.json")
				jsonFile, err := os.Open(jsonPath)
				if err != nil {
					t.Fatalf("Failed to open JSON file: %v", err)
				}

				if err != nil {
					t.Fatalf("Failed to open JSON file: %v", err)
				}
				defer func() {
					if cerr := jsonFile.Close(); cerr != nil {
						t.Errorf("Failed to close JSON file: %v", cerr)
					}
				}()

				mockJSON, err := io.ReadAll(jsonFile)
				if err != nil {
					t.Fatalf("Failed to read JSON file: %v", err)
				}

				err = allure.AddAttachment("Mock JSON", "application/json", mockJSON)
				if err != nil {
					return
				}

				allure.Step(allure.Description("Start mock server and register handler"), allure.Action(func() {
					server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						if r.URL.Path != "/sam/v2/orgs/1/computers" {
							http.NotFound(w, r)
							return
						}
						w.Header().Set("Content-Type", "application/json")
						if _, err := w.Write(mockJSON); err != nil {
							t.Errorf("Failed to write mock response: %v", err)
						}
					}))
					defer server.Close()

					allure.Step(allure.Description("Send GET request to mock server"), allure.Action(func() {
						resp, err := http.Get(server.URL + "/sam/v2/orgs/1/computers")
						if err != nil {
							t.Fatalf("Failed to send GET request: %v", err)
						}
						defer func() {
							if cerr := resp.Body.Close(); cerr != nil {
								t.Errorf("Failed to close response body: %v", cerr)
							}
						}()

						if resp.StatusCode != http.StatusOK {
							t.Fatalf("Expected status 200, got %d", resp.StatusCode)
						}

						var result _struct.ComputerList
						if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
							t.Fatalf("Failed to decode JSON response: %v", err)
						}

						if body, err := json.MarshalIndent(result, "", "  "); err == nil {
							err := allure.AddAttachment("Parsed Response", "application/json", body)
							if err != nil {
								return
							}
						}

						allure.Step(allure.Description("Validate response fields"), allure.Action(func() {
							if result.Count == 0 {
								t.Errorf("Expected Count > 0, got 0")
							}
							if len(result.Values) == 0 {
								t.Errorf("Expected non-empty list of computers, but got empty")
							}
							if result.Values[0].ID == "" {
								t.Errorf("Expected first computer to have non-empty ID")
							}
						}))
					}))
				}))
			}))
		}))
}
