package handlers

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EndtoEndSuite struct {
	suite.Suite
}

// TestEndToEndSuite runs the end-to-end test suite.
func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndtoEndSuite))
}

func (s *EndtoEndSuite) TestPostHandler() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:1323/posts")
	s.Equal(http.StatusOK, r.StatusCode)
}

func (s *EndtoEndSuite) TestPostNoResult() {
	// Test that the endpoint returns an empty slice when the index is out of range.

	c := http.Client{}
	r, _ := c.Get("http://localhost:1323/post/999")
	s.Equal(http.StatusOK, r.StatusCode)

	b, _ := io.ReadAll(r.Body)
	s.JSONEq(`{"status": "ok", "data": []}`, string(b))
}
