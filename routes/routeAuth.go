package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viky1sr/latihan_go-lang.git/controller"
	"github.com/viky1sr/latihan_go-lang.git/repository"
	"github.com/viky1sr/latihan_go-lang.git/service"
	"gorm.io/gorm"
)

func InitAuthRoute(db *gorm.DB, r *gin.Engine) {
	var (
		userRepository repository.UserRepository = repository.NewUserRepository(db)
		jwtService     service.JWTService        = service.NewJwtService()
		authService    service.AuthService       = service.NewAuthService(userRepository)
		authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	)

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
}
