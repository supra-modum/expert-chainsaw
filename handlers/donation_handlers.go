package handlers

import (
	"expert-chainsaw/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserDonations(c *gin.Context, db *gorm.DB) {
	userID := c.Param("user_id")
	var donations []models.Donation
	db.Where("user_id = ?", userID).Find(&donations)
	c.JSON(200, donations)
}

func UpdateDonation(c *gin.Context, db *gorm.DB) {
	donationID := c.Param("id")

	var updateData struct {
		Amount float64 `json:"amount"`
		Sent   bool    `json:"sent"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var donation models.Donation
	if err := db.First(&donation, donationID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Donation not found"})
		return
	}

	donation.Amount = updateData.Amount
	donation.Sent = updateData.Sent
	if result := db.Save(&donation); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, donation)
}

func AddDonation(c *gin.Context, db *gorm.DB) {
	var newDonation struct {
		UserID        uint    `json:"user_id"`
		FundraisingID uint    `json:"fundraising_id"`
		Amount        float64 `json:"amount"`
		Sent          bool    `json:"sent"`
	}

	if err := c.ShouldBindJSON(&newDonation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	donation := models.Donation{
		UserID:        newDonation.UserID,
		FundraisingID: newDonation.FundraisingID,
		Amount:        newDonation.Amount,
		Sent:          newDonation.Sent,
	}

	if result := db.Create(&donation); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, donation)
}
