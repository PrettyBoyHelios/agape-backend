package main

import (
	"github.com/PrettyBoyHelios/agape-backend/controllers"
	"github.com/PrettyBoyHelios/agape-backend/firebase"
	"github.com/gin-gonic/gin"
)

func main() {
	fbAdmin := firebase.NewFirebaseAdmin()
	agape := controllers.NewAgapeController(*fbAdmin)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "agape",
		})
	})
	r.POST("/create", agape.CreateCouple)
	r.POST("/join", agape.JoinCouple)
	r.GET("/users", agape.GetUsers)
	r.GET("/couples/:uid", agape.GetUserCouples)
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
