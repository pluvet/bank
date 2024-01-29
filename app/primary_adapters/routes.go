package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.POST("/user", CreateUser)
	r.POST("/account/:id/deposit", AccountDeposit)
	r.POST("/account/:id/withdraw", AccountWithdraw)
}
