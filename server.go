package main

import (
	"github.com/gin-gonic/gin"
	"github.com/viky1sr/latihan_go-lang.git/config"
	"github.com/viky1sr/latihan_go-lang.git/pkg"
	"github.com/viky1sr/latihan_go-lang.git/routes"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	SetupRouter()
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	if pkg.GodotEnv("GO_ENV") != "production" && pkg.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if pkg.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	routes.InitUserRoute(db, r)

	return r
}
