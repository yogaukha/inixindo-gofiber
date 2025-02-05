package usecase

import (
	"strconv"

	"yogaukha-gofiber/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

type courseUsecase struct {
	courseRepo domain.CourseRepository
}

func NewCourseUsecase(cr domain.CourseRepository) domain.CourseUsecase {
	return &courseUsecase{
		courseRepo: cr,
	}
}

func (u *courseUsecase) FetchAll(c *fiber.Ctx, page int, size int) (paged paginate.Page, len int64, err error) {
	paged, len, err = u.courseRepo.FetchAll(c)
	if err != nil {
		return paginate.Page{}, 0, err
	}

	return
}

func (u *courseUsecase) FetchOneByID(c *fiber.Ctx) (res domain.Course, err error) {
	id, err1 := strconv.Atoi(c.Params("id"))
	if err1 != nil {
		return
	}
	res, err = u.courseRepo.FetchOneByID(c, uint(id))

	return
}

func (u *courseUsecase) Save(c *fiber.Ctx) (res domain.Course, err error) {
	res, err = u.courseRepo.Save(c)
	return
}

func (u *courseUsecase) Edit(c *fiber.Ctx) (res domain.Course, err error) {
	id, err1 := strconv.Atoi(c.Params("id"))
	if err1 != nil {
		return
	}
	res, err = u.courseRepo.Edit(c, uint(id))
	return
}

func (u *courseUsecase) Delete(c *fiber.Ctx) (err error) {
	id, err1 := strconv.Atoi(c.Params("id"))
	if err1 != nil {
		return
	}
	err = u.courseRepo.Delete(c, uint(id))
	return
}
