package handlers

import (
	"expert-chainsaw/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllFundraisings(c *gin.Context, db *gorm.DB) {
	var fundraisings []models.Fundraising
	db.Find(&fundraisings)
	c.JSON(200, fundraisings)
}

func GetFundraising(c *gin.Context, db *gorm.DB, id uint) {
	var fundraising models.Fundraising
	if err := db.First(&fundraising, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fundraising not found"})
		return
	}
	c.JSON(200, fundraising)
}

func CreateFundraising(c *gin.Context, db *gorm.DB) {
	var fundraising models.Fundraising
	if err := c.ShouldBindJSON(&fundraising); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Create(&fundraising); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, fundraising)
}

func UpdateFundraising(c *gin.Context, db *gorm.DB, id uint) {
	var fundraising models.Fundraising
	if err := db.First(&fundraising, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fundraising not found"})
		return
	}

	var updatedFundraising models.Fundraising
	if err := c.ShouldBindJSON(&updatedFundraising); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&fundraising).Updates(updatedFundraising)

	c.JSON(http.StatusOK, fundraising)
}

func DeleteFundraising(c *gin.Context, db *gorm.DB, id uint) {
	var fundraising models.Fundraising
	if err := db.First(&fundraising, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fundraising not found"})
		return
	}

	db.Delete(&fundraising)
	c.JSON(http.StatusOK, gin.H{"message": "Fundraising deleted"})
}
