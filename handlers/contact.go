package handlers

import (
	"log"

	"github.com/F-hatopuku/homepage-app/models"
	"github.com/gin-gonic/gin"
)

func ContactHandler(c *gin.Context) {

	var request models.ContactRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"sucess":  false,
			"message": err.Error(),
		})
		return
	}

	err := sendEmail(
		request.Name,
		request.Email,
		request.Message,
	)

	if err != nil {
		log.Printf("Email sending failed: %v", err)

		c.JSON(500, gin.H{
			"sucess":  false,
			"message": "Email sending failed. Please try again",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Your message has been sent successfully",
	})
}
