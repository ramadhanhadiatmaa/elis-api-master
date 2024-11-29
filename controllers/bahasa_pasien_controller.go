package controllers

import (
	"master/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Show(c *fiber.Ctx) error {
	var bahasa []models.BahasaPasien

	if err := models.DB.Find(&bahasa).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal memuat data",
			"error":   err.Error(),
		})
	}

	if len(bahasa) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
		})
	}

	return c.JSON(bahasa)
}

func Index(c *fiber.Ctx) error {
	id := c.Params("id")
	var bahasa models.BahasaPasien

	if err := models.DB.First(&bahasa, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal memuat data",
			"error":   err.Error(),
		})
	}

	return c.JSON(bahasa)
}

func Create(c *fiber.Ctx) error {
	var bahasa models.BahasaPasien

	if err := c.BodyParser(&bahasa); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	if err := models.DB.Create(&bahasa).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal untuk menyimpan",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Data berhasil ditambahkan",
		"data":    bahasa,
	})
}

func Update(c *fiber.Ctx) error {
	// Parse and validate the ID from the URL
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID format",
		})
	}

	var bahasa models.BahasaPasien
	if err := models.DB.First(&bahasa, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal memuat data",
			"error":   err.Error(),
		})
	}

	var updatedBahasa models.BahasaPasien
	if err := c.BodyParser(&updatedBahasa); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	if updatedBahasa.Id != 0 && updatedBahasa.Id != id {
		if err := models.DB.First(&models.BahasaPasien{}, updatedBahasa.Id).Error; err == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "ID yang diubah telah digunakan",
			})
		}
	}

	if err := models.DB.Model(&bahasa).Updates(&updatedBahasa).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal untuk mengubah data",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diubah",
	})
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if models.DB.Delete(&models.BahasaPasien{}, id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan atau sudah dihapus",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil menghapus data",
	})
}
