package main

import (
	"ESClient/apis"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Static("/static", "assets") //       static/index.html

	e.GET("/search", apis.Search)
	e.GET("/search_simple", apis.SearchSimple)
	e.PUT("/save", apis.Save)
	e.GET("/download", apis.Download)

	// 启动服务
	e.Logger.Fatal(e.Start(":3001"))
}
