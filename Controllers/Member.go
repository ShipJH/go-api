package Controllers

import (
	"net/http"

	"../Models"
	"github.com/gin-gonic/gin"
)

func GetMembers(c *gin.Context) {
	var member []Models.Member
	err := Models.GetAllMembers(&member)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, member)
	}
}

func CreateAMember(c *gin.Context) {
	var member Models.Member
	c.BindJSON(&member)
	err := Models.CreateAMember(&member)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, member)
	}
}

func GetAMember(c *gin.Context) {
	id := c.Params.ByName("id")
	var member Models.Member
	err := Models.GetAMember(&member, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, member)
	}
}

func UpdateAMember(c *gin.Context) {
	var member Models.Member
	id := c.Params.ByName("id")
	err := Models.GetAMember(&member, id)
	if err != nil {
		c.JSON(http.StatusNotFound, member)
	}
	c.BindJSON(&member)
	err = Models.UpdateAMember(&member, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, member)
	}
}

func DeleteAMember(c *gin.Context) {
	var member Models.Member
	id := c.Params.ByName("id")
	err := Models.DeleteAMember(&member, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
