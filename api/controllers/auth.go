package controllers

import (
	"github.com/ZiplEix/super_snake/api/request_models"
	"github.com/ZiplEix/super_snake/api/services"
	"github.com/ZiplEix/super_snake/api/validation"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var req request_models.LoginReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errorResponse{Error: err.Error()})
	}

	err := validation.Login(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errorResponse{Error: err.Error()})
	}

	token, user, err := services.Login(req)
	if err != nil {
		return c.Status(err.(services.ServiceError).Code).JSON(errorResponse{
			Error: err.Error(),
		})
	}

	// set the token on the cookies
	c.Cookie(&fiber.Cookie{
		Name:  "jwt",
		Value: token,
	})
	c.Cookie(&fiber.Cookie{
		Name:  "user",
		Value: user.Name,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
		"user": fiber.Map{
			"email": user.Email,
			"name":  user.Name,
			"id":    user.ID,
		},
	})
}

func Register(c *fiber.Ctx) error {
	var req request_models.RegisterReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errorResponse{Error: err.Error()})
	}

	err := validation.Register(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errorResponse{Error: err.Error()})
	}

	token, user, err := services.Register(req)
	if err != nil {
		return c.Status(err.(services.ServiceError).Code).JSON(errorResponse{
			Error: err.Error(),
		})
	}

	// set the token on the cookies
	c.Cookie(&fiber.Cookie{
		Name:  "jwt",
		Value: token,
	})
	c.Cookie(&fiber.Cookie{
		Name:  "user",
		Value: user.Name,
	})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
		"user": fiber.Map{
			"email": user.Email,
			"name":  user.Name,
			"id":    user.ID,
		},
	})
}
