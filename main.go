package main

import (
	"ESClient/apis"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Static("/static", "assets")

	e.GET("/search", apis.Search)
	e.PUT("/save", apis.Save)

	// 启动服务
	e.Logger.Fatal(e.Start(":3000"))
}
