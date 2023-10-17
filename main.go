package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// init
	r := echo.New()
	// routess
	r.GET("/", Home)
	r.GET("/get-user", GetUser)

	r.Logger.Fatal(r.Start(":9000"))
}

func Home(ctx echo.Context) error {
	data := "Hello Adrian"
	return ctx.String(http.StatusOK, data)
}

func GetUser(ctx echo.Context) error {
	data := "Test API Get User"
	return ctx.String(http.StatusOK, data)
}
