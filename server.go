package main

import (
	helmet "github.com/danielkov/gin-helmet"
	"github.com/sirupsen/logrus"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
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
	app := SetupRouter()
	logrus.Fatal(app.Run(":" + pkg.GodotEnv("GO_PORT")))
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	r.Use(helmet.Default())
	r.Use(gzip.Gzip(gzip.BestCompression))

	routes.InitAuthRoute(db, r)
	routes.InitUserRoute(db, r)

	return r
}
