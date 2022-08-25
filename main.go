package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Flavor struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

var flavors = []Flavor{
	{
		ID:       "1",
		Name:     "Vanilla",
		Quantity: 54,
	},
	{
		ID:       "2",
		Name:     "Chocolate",
		Quantity: 13,
	},
}

func getAllFlavors(c *gin.Context) {
	c.JSON(http.StatusOK, flavors)
}

func getFlavor(c *gin.Context) {
	id := c.Param("ID")
	for _, v := range flavors {
		if v.ID == id {
			c.JSON(http.StatusOK, v)
			return
		}
	}
}

func updateFlavor(c *gin.Context) {
	var updateFlavor Flavor

	if err := c.BindJSON(&updateFlavor); err != nil {
		c.JSON(http.StatusOK, gin.H{"ErrorMessage": "Invalid input!"})
		//c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for index, flavor := range flavors {
		if flavor.ID == updateFlavor.ID {
			flavor.Name = updateFlavor.Name
			flavor.Quantity = updateFlavor.Quantity

			flavors[index] = flavor
			c.JSON(http.StatusOK, flavors)
			return
		}
	}

}

func addFlavor(c *gin.Context) {
	var newFlavor Flavor
	if err := c.BindJSON(&newFlavor); err != nil {
		c.JSON(http.StatusOK, gin.H{"ErrorMessage": "Invalid input!"})
		return
	}

	flavors = append(flavors, newFlavor)
	c.JSON(http.StatusOK, newFlavor)
}

func deleteFlavor(c *gin.Context) {
	var deleteFlavor Flavor
	if err := c.BindJSON(&deleteFlavor); err != nil {
		c.JSON(http.StatusOK, gin.H{"ErrorMessage": "Invalid input!"})
		return
	}

	for index, flavor := range flavors {
		if flavor.ID == deleteFlavor.ID {
			flavors = append(flavors[:index], flavors[index+1:]...)
			c.JSON(http.StatusOK, flavor)
			return
		}
	}
}

func main() {
	r := gin.Default()
	// Quick test!
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Read
	r.GET("/flavors", getAllFlavors)
	r.GET("/flavors/:ID", getFlavor)
	// Update
	r.PUT("/flavors", updateFlavor)
	// Add
	r.POST("/flavors", addFlavor)
	// Delete
	r.DELETE("/flavors", deleteFlavor)

	r.Run()
}
