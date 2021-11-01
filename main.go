package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// userData represents data about a reord userData.
type userData struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Sex    string  `json:"sex"`
	Mobile float64 `json:"mobile"`
}

// userDatas slice to seed record userData data.
var userDatas = []userData{
	{ID: "1", Name: "abc", Sex: "M", Mobile: 986578249},
	{ID: "2", Name: "sdg", Sex: "F", Mobile: 846841156},
	{ID: "3", Name: "hff", Sex: "M", Mobile: 768364578},
}

// getUserData responds with the list of all userdata as JSON.
func getUserData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, userDatas)
}

// postAlbums adds an user Data from JSON received in the request body.
func postuserDatas(c *gin.Context) {
	var newuserData userData

	// Call BindJSON to bind the received JSON to
	// newData.
	if err := c.ShouldBindJSON(&newuserData); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return

	}

	// Add the new album to the slice.
	userDatas = append(userDatas, newuserData)
	c.IndentedJSON(http.StatusCreated, newuserData)
	c.JSON(200, gin.H{
		"error": false})
}

//Put user data. the Put code is below
func putuserDatas(c *gin.Context) {
	id := c.Param("id")
	var newuserData userData
	if err := c.ShouldBindJSON(&newuserData); err != nil {
		//log.Fatalf(err.Error())
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return

	}
	for i, u := range userDatas {
		if u.ID == id {
			userDatas[i].Name = newuserData.Name
			userDatas[i].Sex = newuserData.Sex
			userDatas[i].Mobile = newuserData.Mobile

			c.JSON(200, gin.H{
				"error": false,
			})
			return

		}
	}

	log.Println("User Data not found")
	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})
}

//Delete user data by ID is below the code
func deleteuserDatas(c *gin.Context) {
	id := c.Param("id")

	for i, u := range userDatas {
		if u.ID == id {
			userDatas = append(userDatas[:i], userDatas[i+1:]...)

			c.JSON(200, gin.H{
				"error": false})
			return

		}
	}
	log.Println("User Data not found")
	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})
}

func main() {
	router := gin.Default()
	router.GET("/userDatas", getUserData)
	router.POST("/userDatas", postuserDatas)
	router.PUT("/userDatas/:id", putuserDatas)
	router.DELETE("/userDatas/:id", deleteuserDatas)

	router.Run("localhost:8080")
}
