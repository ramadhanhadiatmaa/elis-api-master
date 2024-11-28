package controllers

import (
	"master/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Show(c *fiber.Ctx) error {
	var bahasa []models.BahasaPasien
	models.DB.Find(&bahasa)

	return c.Status(fiber.StatusOK).JSON(bahasa)
}

func Create(c *fiber.Ctx) error {

	var bahasa models.BahasaPasien
	if err := c.BodyParser(&bahasa); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(&bahasa).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Bahasa berhasil ditambahkan",
	})
}

func Update(c *fiber.Ctx) error {
	// Get the ID from the URL parameters
	id := c.Params("id")

	// Convert the ID from string to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID format",
		})
	}

	var bahasa models.BahasaPasien

	// Parse the request body into the struct
	if err := c.BodyParser(&bahasa); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	// Ensure the ID from the path is used
	bahasa.Id = intID

	// Update the record and check for affected rows
	if result := models.DB.Model(&models.BahasaPasien{}).Where("id = ?", intID).Updates(&bahasa); result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Record not found or no changes made.",
		})
	} else if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update record.",
			"error":   result.Error.Error(),
		})
	}

	// Return success response
	return c.JSON(fiber.Map{
		"message": "Record updated successfully.",
	})
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")
	var bahasa models.BahasaPasien

	if models.DB.Where("id = ?", id).Delete(&bahasa).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Bahasa tidak dapat dihapus.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil menghapus bahasa.",
	})
}
