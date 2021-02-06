package api

import (
	"context"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo"
	"google.golang.org/api/option"
)

func VerifyToken() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		token := c.Request().Header.Get("token")

		if token == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "トークンが設定されていません")
		}
		ctx := context.Background()
		cd, _ := os.Getwd()
		opt := option.WithCredentialsFile(cd + "/serviceAccountKey.json")
		app, err := firebase.NewApp(ctx, nil, opt)
		if err != nil {
			log.Fatalln("認証システムに失敗しました", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "認証システムの作成に失敗しました")
		}

		auth, err := app.Auth(context.Background())
		if err != nil {
			log.Fatalln("認証に失敗しました", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "認証システムの作成に失敗しました")
		}

		decoded, err := auth.VerifySessionCookieAndCheckRevoked(context.Background(), token)
		if err != nil {
			log.Fatalln("認証トークン期限ぎれ", err)
			return echo.NewHTTPError(http.StatusUnauthorized, "認証に期限が切れています。ログインし直してください")
		}

		return c.JSON(http.StatusOK, decoded)
	}
}
