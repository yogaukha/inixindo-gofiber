package domain

import (
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

type Course struct {
	// do not use gorm.Model, but type it manually
	ID             uint   `json:"id"`
	CourseName     string `gorm:"size:400" json:"course_name"`
	Description    string `json:"description"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	InstructorName string `gorm:"size:400" json:"instructor_name"`
	CourseFile     string `json:"course_file"`
	Status         string `gorm:"size:15" json:"status"`
}

type CourseUsecase interface {
	FetchAll(c *fiber.Ctx, page int, size int) (paginate.Page, int64, error)
	FetchOneByID(c *fiber.Ctx) (Course, error)
	Save(c *fiber.Ctx) (Course, error)
	Edit(c *fiber.Ctx) (Course, error)
	Delete(c *fiber.Ctx) error
}

type CourseRepository interface {
	FetchAll(c *fiber.Ctx) (paginate.Page, int64, error)
	FetchOneByID(c *fiber.Ctx, id uint) (Course, error)
	Save(c *fiber.Ctx) (Course, error)
	Edit(c *fiber.Ctx, id uint) (Course, error)
	Delete(c *fiber.Ctx, id uint) error
}
