package controller

import (
	"fmt"
	"lecture/go-mvc/model"
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

func (p *Controller) GetName(c *gin.Context) {
	name := c.Param("name")
	message := p.md.GetNamePerson(name)
	fmt.Println(message)
	c.JSON(200, gin.H{
		"message": "ok",
		"body": message,
	})
}

func (p *Controller) GetPnum(c *gin.Context) {
	pnum := c.Param("pnum")
	c.JSON(200, gin.H{
		"message": "ok",
		"body": p.md.GetPnumPerson(pnum),
	})
}

func (p *Controller) Post(c *gin.Context) {
	name := c.PostForm("name")
	age, _ := strconv.Atoi(c.PostForm("age"))
	pnum := c.PostForm("pnum")
	c.JSON(200, gin.H{
		"message": "ok",
		"body": p.md.PostNewperson(name, pnum, age),
	})
}

func (p *Controller) DeletePnum(c *gin.Context) {
	pnum := c.Param("pnum")
	c.JSON(200, gin.H{
		"message": "ok",
		"body": p.md.DeletePnump(pnum),
	})
}

func (p *Controller) PutPnum(c *gin.Context) {
	pnum := c.Param("pnum")
	age, _ := strconv.Atoi(c.PostForm("age"))
	c.JSON(200, gin.H{
		"message": "ok",
		"body": p.md.PutPnumAge(pnum, age),
	})
}