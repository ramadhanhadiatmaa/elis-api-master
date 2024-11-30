package controllers

import (
	"master/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ShowKel(c *fiber.Ctx) error {
	var kelurahan []models.Kelurahan

	if err := models.DB.Find(&kelurahan).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal memuat data",
			"error":   err.Error(),
		})
	}

	if len(kelurahan) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
		})
	}

	return c.JSON(kelurahan)
}

func IndexKel(c *fiber.Ctx) error {
	id := c.Params("kd_kel")
	var kelurahan models.Kelurahan

	if err := models.DB.First(&kelurahan, "kd_kel = ?", id).Error; err != nil {
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

	return c.JSON(kelurahan)
}

func CreateKel(c *fiber.Ctx) error {
	var kelurahan models.Kelurahan

	if err := c.BodyParser(&kelurahan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	if err := models.DB.Create(&kelurahan).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal untuk menyimpan data",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Data berhasil ditambahkan",
		"data":    kelurahan,
	})
}

func UpdateKel(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("kd_kel"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID format",
		})
	}

	var kelurahan models.Kelurahan
	if err := models.DB.First(&kelurahan, id).Error; err != nil {
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

	var updateKel models.Kelurahan
	if err := c.BodyParser(&updateKel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	if updateKel.Kd != 0 && updateKel.Kd != id {
		if err := models.DB.First(&models.Kelurahan{}, updateKel.Kd).Error; err == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "ID yang diubah telah digunakan",
			})
		}
	}

	if err := models.DB.Model(&kelurahan).Updates(&updateKel).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal untuk mengubah data",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diubah",
	})
}

func DeleteKel(c *fiber.Ctx) error {
	id := c.Params("kd_kel")

	if models.DB.Delete(&models.Kelurahan{}, id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan atau sudah dihapus",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil menghapus data",
	})
}
