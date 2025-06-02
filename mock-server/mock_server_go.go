package mockserver

import (
	"embed"
	"net/http"
	"net/http/httptest"
)

var MockFS embed.FS

type MockServer struct {
	server *httptest.Server
	mux    *http.ServeMux
}

func NewMockServer() *MockServer {
	mux := http.NewServeMux()
	s := httptest.NewServer(mux)
	return &MockServer{server: s, mux: mux}
}

func (ms *MockServer) AddJSONHandler(path string, filePath string, statusCode int) {
	ms.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		data, err := MockFS.ReadFile(filePath)
		if err != nil {
			http.Error(w, "Mock file not found: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		if _, err := w.Write(data); err != nil {
			http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
		}
	})
}

func (ms *MockServer) URL() string {
	return ms.server.URL
}

func (ms *MockServer) Close() {
	ms.server.Close()
}
