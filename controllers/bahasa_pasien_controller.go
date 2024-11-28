package controllers

import (
	"master/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Show(c *fiber.Ctx) error {
	id := c.Params("id")

	var bahasa models.BahasaPasien

	if err := models.DB.Where("id = ?", id).First(&bahasa).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Bahasa tidak ditemukan.",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	return c.JSON(bahasa)
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

	id := c.Params("id")

	var bahasa models.BahasaPasien

	if err := c.BodyParser(&bahasa); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(&bahasa).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Id tidak ditemukan.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Bahasa berhasil diubah.",
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
