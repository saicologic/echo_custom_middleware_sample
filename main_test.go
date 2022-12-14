package main

import (
	"net/http/httptest"
	"strings"
	"testing"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCustomHandler(t *testing.T) {
	t.Run("Add userID", func(t *testing.T) {
		// Request Header Params
		userID := "12345678"

		// Request
		method := echo.GET
		target := "/"
		body := strings.NewReader("")
		req := httptest.NewRequest(method, target, body)

		// Add Header
		req.Header.Set("X-USER-ID", userID)
		rec := httptest.NewRecorder()

		// Serve
		e := initEcho()
		e.ServeHTTP(rec, req)

		// Result
		ctx := e.AcquireContext()
		defer e.ReleaseContext(ctx)
		c := GetContextValues(ctx)
		res := rec.Result()

		assert.Equal(t, 200, res.StatusCode)
		assert.Equal(t, userID, rec.Body.String())
		assert.Equal(t, userID, c.UserID)
	})
}
