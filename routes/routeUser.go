package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viky1sr/latihan_go-lang.git/controller"
	"github.com/viky1sr/latihan_go-lang.git/middleware"
	"github.com/viky1sr/latihan_go-lang.git/repository"
	"github.com/viky1sr/latihan_go-lang.git/service"
	"gorm.io/gorm"
)

func InitUserRoute(db *gorm.DB, r *gin.Engine) {
	var (
		userRepository repository.UserRepository = repository.NewUserRepository(db)
		jwtService     service.JWTService        = service.NewJwtService()
		userService    service.UserService       = service.NewUserService(userRepository)
		userController controller.UserController = controller.NewUserController(userService, jwtService)
	)

	userRoute := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoute.GET("/user-profile", userController.Profile)
		userRoute.PUT("/user-update", userController.Update)
	}

}
