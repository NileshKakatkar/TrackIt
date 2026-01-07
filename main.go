package main

import (
	"log"
	"reflect"
	"trackit/app/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	engine.Reload(true)
	engine.AddFuncMap(map[string]interface{}{
		"hasField": HasField,
	})
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	route.SetupUserRoutes(app)
	log.Fatal(app.Listen(":8082"))
}

func HasField(data interface{}, name string) bool {
	rv := reflect.ValueOf(data)
	if rv.Kind() != reflect.Map {
		return false
	} else if _, ok := data.(fiber.Map)[name]; ok {
		return true
	}
	return false
}
