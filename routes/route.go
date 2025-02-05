package routes

import (
	course_delivery "yogaukha-gofiber/app/course/delivery"
	course_repo "yogaukha-gofiber/app/course/repository"
	course_usecase "yogaukha-gofiber/app/course/usecase"
	"yogaukha-gofiber/configs"

	"github.com/gofiber/fiber/v2"
)

func RouteRegister(app *fiber.App, config configs.Config) {
	// api versioning, check config.yaml file to set api versioning
	ver := app.Group("/api/" + config.ApiVersion)

	// call each modules to register to fiber routing
	cr := course_repo.NewCourseRepository(configs.DBConn)
	cu := course_usecase.NewCourseUsecase(cr)
	course_delivery.NewCourseHandler(ver, cu)
}
