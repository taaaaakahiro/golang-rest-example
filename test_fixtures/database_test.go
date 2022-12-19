package testfixtures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleSuccess(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"RememberMe", "remember_me"},
		{"requestID", "request_id"},
		{"HTTPRequest", "http_request"},
		{"HTML5Script", "html5_script"},
	}

	for _, test := range tests {
		res := camelToSnake(test.in)
		assert.Equal(t, test.expected, res)
	}
}
