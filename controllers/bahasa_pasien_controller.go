package controllers

import (
	"master/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

// Update allows updating a record, including the ID
func Update(c *fiber.Ctx) error {
	// Get the current ID from the URL parameters
	id := c.Params("id")

	// Convert the ID from string to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID format",
		})
	}

	var bahasa models.BahasaPasien

	// Check if the record with the current ID exists
	if err := models.DB.First(&bahasa, intID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Bahasa tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal memuat data",
			"error":   err.Error(),
		})
	}

	// Parse the request body into the struct
	var updatedBahasa models.BahasaPasien
	if err := c.BodyParser(&updatedBahasa); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	// Allow updating the ID to a new unique ID
	if updatedBahasa.Id != 0 && updatedBahasa.Id != bahasa.Id {
		// Check if the new ID already exists
		var existingRecord models.BahasaPasien
		if err := models.DB.First(&existingRecord, updatedBahasa.Id).Error; err == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "ID yang diubah telah digunakan",
			})
		}
	}

	// Update the record
	if err := models.DB.Model(&bahasa).Updates(&updatedBahasa).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal untuk mengubah bahasa",
			"error":   err.Error(),
		})
	}

	// Return success response
	return c.JSON(fiber.Map{
		"message": "Bahasa berhasil diubah",
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
