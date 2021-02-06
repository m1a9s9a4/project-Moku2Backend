package route

import (
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
	"github.com/m1a9s9a4/api"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(echoMw.Logger())
	e.Use(echoMw.CORS())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	v1 := e.Group("/api/v1")
	{
		v1.GET("/", api.Hello())
		v1.GET("/members/", api.GetOnlineMembers())
		v1.POST("/token/create/", api.CreateToken())
		v1.POST("/token/verify/", api.VerifyToken())
	}

	return e
}
