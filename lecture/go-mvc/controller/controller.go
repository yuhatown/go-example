package controller

import (
	"lecture/go-mvc/model"

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
	c.JSON(200, gin.H{
		"message": "ok",
		"body": p.md.GetNamePerson(c.PostForm("name")),
	})
}

func (p *Controller) GetPnum(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
		"body": p.md.GetPnumPerson(c.PostForm("pnum")),
	})
}

func (p *Controller) Post(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
		"body": p.md.PostNewperson(c.PostForm("name"), c.PostForm("pnum")),
	})
}

func (p *Controller) DeletePnum(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
		"body": p.md.DeletePnump(c.PostForm("pnum")),
	})
}

func (p *Controller) PutPnum(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
		"body": p.md.PutPnumAge(c.PostForm("pnum"), c.PostForm("age")),
	})
}