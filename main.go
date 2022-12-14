package main

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

// Context
type ContextValues struct {
	UserID string
}

type contextValuesKey string

const ContextValuesKey = contextValuesKey("ContextValuesKey")

func GetContextValues(c echo.Context) *ContextValues {
	values, ok := c.Get(string(ContextValuesKey)).(*ContextValues)
	if values == nil || !ok {
		values = &ContextValues{}
		c.Set(string(ContextValuesKey), values)
	}
	return values
}

// Custom Middleware
func setUserIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Request().Header.Get("X-USER-ID")
		ctx := GetContextValues(c)
		ctx.UserID = userID
		return next(c)
	}
}

// Handler
func GetUser(c echo.Context) error {
	ctx := GetContextValues(c)
	return c.String(http.StatusOK, ctx.UserID)
}

func initEcho() *echo.Echo {
	e := echo.New()
	e.Use(setUserIDMiddleware)
	e.GET("/", GetUser)
	return e
}

func main() {
	e := initEcho()
	e.Logger.Fatal(e.Start(":1323"))
}
