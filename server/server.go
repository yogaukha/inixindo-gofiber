package server

import (
	"log"
	"time"

	"yogaukha-gofiber/configs"
	"yogaukha-gofiber/internal"
	"yogaukha-gofiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run() {
	app := fiber.New()
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		app.Use(func(c *fiber.Ctx) error {
			return internal.ReturnTheResponse(c, true, int(500), "Can not init the timezone", nil)
		})
	}
	time.Local = loc // -> this is setting the global timezone

	config, err := configs.LoadConfig("./configs")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// db init
	configs.InitDB(config)

	// load Middlewares
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	// decode the token from frontend
	// app.Use(func(c *fiber.Ctx) error {
	// 	return internal.DecodeToken(c)
	// })

	// register route in another package
	routes.RouteRegister(app, config)

	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return internal.ReturnTheResponse(c, true, int(404), "Not Found", nil)
	})

	// Here we go!
	log.Fatalln(app.Listen(config.Server.Host + ":" + config.Server.Port))

}
