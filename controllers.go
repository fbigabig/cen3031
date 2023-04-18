package controllers

import (
	"../models"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

func Placeholder(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	return c.JSON(user)
}
