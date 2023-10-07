package main

import (
	"contact-manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	contacts := make([]models.Contact, 0)

	contacts = append(contacts, models.Contact{Name: models.Name{FirstName: "Eli", LastName: "Smith"}, Age: 22})

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Pong!")
	})
	server.GET("/contact", func(c *gin.Context) {
		firstName := c.Query("firstName")
		lastName := c.Query("lastName")

		if firstName == "" && lastName == "" {
			c.String(http.StatusNotFound, "Could not find a matching contact!")
			return
		}

		for _, thisContact := range contacts {
			if thisContact.Name.FirstName == firstName || thisContact.Name.LastName == lastName {
				c.JSON(http.StatusOK, gin.H{
					"contact": thisContact,
				})
				return
			}
		}

		c.String(http.StatusNotFound, "Could not find a matching contact!")
	})

	server.Run()
}
