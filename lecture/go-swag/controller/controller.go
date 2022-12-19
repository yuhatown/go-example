package controller

import (
	"go-example/lecture/go-swag/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{md : rep}
	return r, nil
}

// GetOK godoc
// @Summary call GetOK, return ok by json
// @Description api test를 위한 기능
// @name GetOK
// @Accept json
// @Produce json
// @Param name path string true "User name"
// @Router /acc/v01/ok [get]
// @Success 200 {object} Controller
func (p *Controller) GetOK(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

// Post godoc
// @Summary call Post, return ok by json
// @Description api test를 위한 기능
// @name Post
// @Accept json
// @Produce json
// @Param name path string true "User name"
// @Param age path string true "User age"
// @Param pnum path string true "User pum"
// @Router /acc/v01/post [post]
// @Success 200 {object} Controller
func (p *Controller) Post(c *gin.Context) {
	name := c.PostForm("name")
	age, _ := strconv.Atoi(c.PostForm("age"))
	pnum := c.PostForm("pnum")
	c.JSON(200, gin.H{
		"message": "ok",
		"body": p.md.PostNewperson(name, pnum, age),
	})
}

// PutName godoc
// @Summary call PutName, return ok by json
// @Description api test를 위한 기능
// @name PutName
// @Accept json
// @Produce json
// @Param name path string true "User name"
// @Param age path string true "User age"
// @Router /acc/v01/PutName [put]
// @Success 200 {object} Controller
func (p *Controller) PutName(c *gin.Context) {
	name := c.Param("pnum")
	age, _ := strconv.Atoi(c.PostForm("age"))
	c.JSON(200, gin.H{
		"message": "ok",
		"body": p.md.PutNameAge(name, age),
	})
}