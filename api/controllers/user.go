package controllers

import (
	"github.com/ZiplEix/super_snake/api/database"
	"github.com/ZiplEix/super_snake/api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Me(c *fiber.Ctx) error {
	var user models.User

	result := database.Db.First(&user, c.Locals("userId"))
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Database error",
			})
		}
	}

	res := models.UserResponse{
		Email: user.Email,
		Name:  user.Name,
	}

	return c.JSON(res)
}
