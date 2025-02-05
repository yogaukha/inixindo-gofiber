package delivery

import (
	"strconv"
	"strings"

	"yogaukha-gofiber/domain"
	"yogaukha-gofiber/internal"

	"github.com/gofiber/fiber/v2"
)

type CourseHandler struct {
	MawUsecase domain.CourseUsecase
}

func NewCourseHandler(c fiber.Router, uc domain.CourseUsecase) {
	handler := &CourseHandler{
		MawUsecase: uc,
	}
	c.Get("/courses", handler.FetchAll)
	c.Get("/courses/:id", handler.FetchOneByID)
	c.Post("/courses", handler.Save)
	c.Put("/courses/:id", handler.Edit)
	c.Delete("/courses/:id/:username", handler.Delete)
}

func (mawh *CourseHandler) FetchAll(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return internal.ReturnTheResponse(c, false, int(400), "Bad Request", nil)
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		return internal.ReturnTheResponse(c, false, int(400), "Bad Request", nil)
	}
	listMaw, rowLen, err := mawh.MawUsecase.FetchAll(c, page, size)
	if rowLen <= 0 {
		return internal.ReturnTheResponse(c, true, int(404), "Record not Found", nil)
	}
	if err != nil {
		return internal.ReturnTheResponse(c, true, int(404), err.Error(), nil)
	}

	return internal.ReturnTheResponse(c, false, int(200), "", listMaw)
}

func (mawh *CourseHandler) FetchOneByID(c *fiber.Ctx) error {
	res, err := mawh.MawUsecase.FetchOneByID(c)
	if err != nil {
		return internal.ReturnTheResponse(c, true, int(404), err.Error(), nil)
	}

	return internal.ReturnTheResponse(c, false, int(200), "", res)
}

func (mawh *CourseHandler) Save(c *fiber.Ctx) error {
	res, err := mawh.MawUsecase.Save(c)
	if err != nil {
		return internal.ReturnTheResponse(c, true, int(500), err.Error(), nil)
	}
	return internal.ReturnTheResponse(c, false, int(200), "", res)
}

func (mawh *CourseHandler) Edit(c *fiber.Ctx) error {
	res, err := mawh.MawUsecase.Edit(c)
	if err != nil {
		errMessage := err.Error()
		if strings.Contains(errMessage, "not found") {
			return internal.ReturnTheResponse(c, true, int(404), err.Error(), nil)
		} else {
			return internal.ReturnTheResponse(c, true, int(500), err.Error(), nil)
		}
	}
	return internal.ReturnTheResponse(c, false, int(200), "", res)
}

func (mawh *CourseHandler) Delete(c *fiber.Ctx) error {
	err := mawh.MawUsecase.Delete(c)
	if err != nil {
		return internal.ReturnTheResponse(c, true, int(500), err.Error(), nil)
	}
	return internal.ReturnTheResponse(c, false, int(200), "Deleted succesfully", nil)
}
