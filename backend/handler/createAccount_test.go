package handler

import (
	"io"
	"testing"
)

func TestHandleCreateUserAccount(t *testing.T) {
	type testCaseRequest struct {
		requestBody          io.Reader
		expectedStatus       int
		expectedResponseBody string
	}

}
