package main

import (
	"sendsmsgroup-producer/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	//routes
	routes.SMSGroupRoute(e)
	e.Logger.Fatal(e.Start(":3000"))
}
