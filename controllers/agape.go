package controllers

import (
	"github.com/PrettyBoyHelios/agape-backend/firebase"
	"github.com/gin-gonic/gin"
)

type AgapeController struct {
	admin firebase.FirebaseAdmin
}

func NewAgapeController(admin firebase.FirebaseAdmin) *AgapeController {
	a := new(AgapeController)
	a.admin = admin
	return a
}
func (a *AgapeController) CreateCouple(c *gin.Context) {

}

func (a *AgapeController) GetUsers(c *gin.Context) {
	users := a.admin.GetUsers()
	c.JSON(200, users)
}

func (a *AgapeController) GetUserCouples(c *gin.Context) {
	uid := c.Query("uid")
	couples := a.admin.GetUserCouples(uid)
	c.JSON(200, couples)
}
