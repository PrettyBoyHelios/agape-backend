package controllers

import (
	"encoding/json"
	"github.com/PrettyBoyHelios/agape-backend/firebase"
	"github.com/PrettyBoyHelios/agape-backend/models"
	app_errors "github.com/PrettyBoyHelios/agape-backend/models/errors"
	"github.com/PrettyBoyHelios/agape-backend/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"time"
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
	var input models.CreateCoupleInput
	body := c.Request.Body
	data, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(500, err)
		return
	}
	_ = json.Unmarshal(data, &input)

	userCouples := a.admin.GetUserCouples(input.CreatorID)

	if len(userCouples) == 0 {
		log.Println("can create couple!")
		newCouple := models.Couple{
			CreatorID:   input.CreatorID,
			CreatedAt:   time.Now().Unix(),
			PairingCode: utils.RandStringBytes(6),
			Status:      models.INCOMPLETE,
		}
		coupleId := a.admin.CreateCouple(newCouple)
		newCouple.CoupleID = coupleId
		c.JSON(200, newCouple)
	} else {
		c.JSON(401, app_errors.AppError{Error: app_errors.ALREADY_IN_COUPLE.Error()})
		return
	}
}

func (a *AgapeController) JoinCouple(c *gin.Context) {
	// TODO Validate user account exists before creating and joining
	var input models.JoinCoupleInput
	body := c.Request.Body
	data, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(500, err)
		return
	}
	_ = json.Unmarshal(data, &input)

	currentCouple, id := a.admin.GetCoupleByPairingCode(input.PairCode)

	if currentCouple.CoupleID == "" {
		currentCouple.CoupleID = input.ID
		currentCouple.Status = models.READY
		res := a.admin.Update(currentCouple, id)
		c.JSON(200, res)
	} else {
		c.JSON(401, app_errors.AppError{Error: app_errors.ALREADY_IN_COUPLE.Error()})
		return
	}
}

func (a *AgapeController) GetUsers(c *gin.Context) {
	users := a.admin.GetUsers()
	c.JSON(200, users)
}

func (a *AgapeController) GetUserCouples(c *gin.Context) {
	uid := c.Param("uid")
	couples := a.admin.GetUserCouples(uid)
	c.JSON(200, couples)
}
