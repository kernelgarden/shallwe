package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

type HelloControllers struct {
}

func (c HelloControllers) Init(g *echo.Group) {
	g.GET("", c.Get)
}

func (HelloControllers) Get(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "hello world")
}
