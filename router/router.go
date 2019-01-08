package router

import (
	"github.com/labstack/echo"
	"github.com/kernelgarden/shallwe/controllers"
)

func InitRoutes(e *echo.Echo) {
	controllers.HelloControllers{}.Init(e.Group("/hello"))
}