package main

import (
	"github.com/gin-gonic/gin"
	"github.com/viky1sr/latihan_go-lang.git/config"
	"github.com/viky1sr/latihan_go-lang.git/controller"
	"github.com/viky1sr/latihan_go-lang.git/middleware"
	"github.com/viky1sr/latihan_go-lang.git/repository"
	"github.com/viky1sr/latihan_go-lang.git/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJwtService()
	userService    service.UserService       = service.NewUserService(userRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoute := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoute.GET("/user-profile", userController.Profile)
		userRoute.PUT("/user-update", userController.Update)
	}

	r.Run(":5000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
