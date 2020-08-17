package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleMain(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/main", nil)
	s.main().ServeHTTP(rec, req)

	assert.Equal(t, rec.Body.String(), "main")
}
