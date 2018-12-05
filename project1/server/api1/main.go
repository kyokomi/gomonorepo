package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"server/api1/hoge"
	"server/common/piyo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		fmt.Println(piyo.Piyo())
		return c.String(http.StatusOK, "Hello, World! api 1"+hoge.Hoge())
	})
	e.Logger.Fatal(e.Start(":1323"))
}
