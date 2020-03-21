package main

import (
    "net/http"
    "os"
    "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"example.com/go-echo-cloud-run/controllers"
)

func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", func(c echo.Context) error {
		helloController := controllers.New()
		text := helloController.Get_main()
        return c.String(http.StatusOK, text)
    })

    e.GET("/:param", func(c echo.Context) error {
		param := c.Param("param")
		helloController := controllers.New()
		text := helloController.Get_with_param(param)
        return c.String(http.StatusOK, text)
	})
	
	e.GET("/bye", func(c echo.Context) error {
		helloController := controllers.New()
		text := helloController.Get_goodbye()
        return c.String(http.StatusOK, text)
    })

    e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}