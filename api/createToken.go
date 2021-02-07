package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo"
)

func CreateToken() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		m := echo.Map{}
		if err := c.Bind(&m); err != nil {
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "IDが設定されていません")
		}
		// interfaceになってるためstringに変換
		idToken := m["idToken"].(string)
		if idToken == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "IDが設定されていません")
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
		expiresIn := time.Hour * 24 * 7
		cookie, err := auth.SessionCookie(context.Background(), idToken, expiresIn)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "認証システムの作成に失敗しました")
		}

		return c.JSON(http.StatusOK, cookie)
	}
}
