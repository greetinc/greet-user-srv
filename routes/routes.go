package routes

import (
	"greet-user-srv/configs"

	"github.com/greetinc/greet-middlewares/middlewares"
	"github.com/labstack/echo/v4"

	h_user "greet-user-srv/handlers/profile"
	r_user "greet-user-srv/repositories/profile"
	s_user "greet-user-srv/services/profile"
)

var (
	DB = configs.InitDB()

	JWT   = middlewares.NewJWTService()
	userR = r_user.NewUserRepository(DB)
	userS = s_user.NewUserService(userR, JWT)
	userH = h_user.NewUserHandler(userS)
)

func New() *echo.Echo {

	e := echo.New()
	v1 := e.Group("api/v1")
	userview := v1.Group("/userview", middlewares.AuthorizeJWT(JWT))
	{
		userview.GET("", userH.UserPreview)
	}
	user := v1.Group("/user", middlewares.AuthorizeJWT(JWT))
	{
		user.GET("", userH.Get)
		user.PUT("/:id", userH.Update)
	}
	return e
}
