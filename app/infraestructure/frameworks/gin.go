package frameworks

import (
	"github.com/gin-gonic/gin"
	"github.com/pluvet/bank/app/infraestructure/database"
	"github.com/pluvet/bank/app/primaryadapters/controllers"
)

type GinFramework struct{}

func (g *GinFramework) Run() {
	r := gin.Default()
	database.DBConnect()
	controllers.Routes(r)
	r.Run()
}
