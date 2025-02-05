package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gorm.io/gorm"

	"yogaukha-gofiber/domain"
)

type courseRepository struct {
	DB *gorm.DB
}

func NewCourseRepository(DB *gorm.DB) domain.CourseRepository {
	return &courseRepository{DB}
}

func (r *courseRepository) FetchAll(c *fiber.Ctx) (paged paginate.Page, rowLen int64, err error) {
	var mawdomain []domain.Course
	pg := paginate.New()
	result := r.DB.Find(&mawdomain)
	pager := pg.With(result).Request(c.Request()).Response(&mawdomain)
	return pager, result.RowsAffected, result.Error
}

func (r *courseRepository) FetchOneByID(c *fiber.Ctx, id uint) (res domain.Course, err error) {
	result := r.DB.First(&res, id)
	return res, result.Error
}

func (r *courseRepository) Save(c *fiber.Ctx) (res domain.Course, err error) {
	maw := new(domain.Course)
	if err1 := c.BodyParser(maw); err1 != nil {
		return
	}
	r.DB.Create(&maw)
	return *maw, nil
}

func (r *courseRepository) Edit(c *fiber.Ctx, id uint) (res domain.Course, err error) {
	var maw domain.Course
	result := r.DB.First(&maw, id)
	if result.Error != nil {
		return domain.Course{}, result.Error
	}
	if err1 := c.BodyParser(&maw); err1 != nil {
		return domain.Course{}, err1
	}
	r.DB.Save(&maw)
	return maw, nil
}

func (r *courseRepository) Delete(c *fiber.Ctx, id uint) (err error) {
	var maw domain.Course
	result := r.DB.Where("id = ?", id).Delete(&maw)
	if result.Error != nil {
		return result.Error
	}
	return
}
