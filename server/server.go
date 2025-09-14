package main

import (
	_ "io/ioutil"
	_ "os"

	"github.com/gofiber/fiber/v2"
	"nadhi.dev/binaries/server/api"
)

func RegisterRoutes(app *fiber.App) {
    app.Get("/packages", api.GetPackages)
	app.Get("/packages/search", api.SearchPackages)
}