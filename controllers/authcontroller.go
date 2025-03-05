package controllers

import (
    "github.com/gofiber/fiber/v3"
)

func Hello (c fiber.Ctx) error {
        // Send a string response to the client
        return c.SendString("Hello, Woajrld 👋!")
    }

