package routes

import (
	"fmt"
	controllers "myapp/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {

	echoServer := echo.New()
	echoServer.Use(middleware.Logger())  // Logger
	echoServer.Use(middleware.Recover()) // Recover

	//middlewares:
	echoServer.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method}, uri=${uri}, status=${status}\n",
		Skipper: func(c echo.Context) bool {
			return false
		},
	}))
	echoServer.Pre(middleware.AddTrailingSlash())
	echoServer.Use(middleware.CORS())
	fmt.Println("Authentication Service")

	echoServer.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Running Test Go Program")
	})
	//API Route
	apiRoute := echoServer.Group("/api")

	//User Route
	userRoute := apiRoute.Group("/users")
	userRoute.GET("/", controllers.UserGetAllUsers)

	//User Route
	authRoute := apiRoute.Group("/auth")
	authRoute.POST("/signup/", controllers.AuthSignUp)
	authRoute.POST("/login/", controllers.AuthLogin)
	authRoute.POST("/checktoken/", controllers.AuthCheckToken)

	// apiGroup.Use(api.TokenMiddleware)

	return echoServer
}
