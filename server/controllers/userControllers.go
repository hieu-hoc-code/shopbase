package controllers

import "github.com/gofiber/fiber"

func GetHome(c *fiber.Ctx) error{
		return c.SendString("Hello, World 👋!")
}