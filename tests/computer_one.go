package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"

	_struct "TestingGo/internal/model/struct"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type TestSuiteOne struct {
	suite.Suite
}

func (s *TestSuiteOne) TestGetComputersFromMockServer1(t provider.T) {
	var mockJSON []byte
	var result _struct.ComputerList

	t.WithNewStep("Load JSON data from file", func(sCtx provider.StepCtx) {
		cwd, err := os.Getwd()
		sCtx.Require().NoError(err)

		projectRoot := filepath.Dir(cwd)
		jsonPath := filepath.Join(projectRoot, "internal", "model", "mock_data", "computers_response.json")

		jsonFile, err := os.Open(jsonPath)
		sCtx.Require().NoError(err)
		defer jsonFile.Close()

		mockJSON, err = io.ReadAll(jsonFile)
		sCtx.Require().NoError(err)

		// Attach mock JSON
		sCtx.Step(allure.NewSimpleStep("Attach Mock JSON").
			WithAttachments(
				allure.NewAttachment("Mock JSON", allure.JSON, mockJSON),
			))
	})

	t.WithNewStep("Start mock server and verify response", func(sCtx provider.StepCtx) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/sam/v2/orgs/1/computers" {
				http.NotFound(w, r)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write(mockJSON)
			sCtx.Require().NoError(err)
		}))
		defer server.Close()

		resp, err := http.Get(server.URL + "/sam/v2/orgs/1/computers")
		sCtx.Require().NoError(err)
		defer resp.Body.Close()

		sCtx.Require().Equal(http.StatusOK, resp.StatusCode)

		err = json.NewDecoder(resp.Body).Decode(&result)
		sCtx.Require().NoError(err)

		// Attach parsed response
		if body, err := json.MarshalIndent(result, "", " "); err == nil {
			sCtx.Step(allure.NewSimpleStep("Attach response").
				WithAttachments(
					allure.NewAttachment("Parsed response", allure.JSON, body),
				))
		}
	})

	t.WithNewStep("Verify response content", func(sCtx provider.StepCtx) {
		sCtx.Assert().Greater(result.Count, 0)
		sCtx.Assert().Greater(len(result.Values), 0)
		sCtx.Assert().NotEmpty(result.Values[0].ID)
	})
}
