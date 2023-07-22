package service

import (
	InforLog "Golang/log/infor"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	config "Golang/config"

	gormdb "Golang/adapter/gormdb"

	utils "Golang/utils"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func Run() {
	InforLog.PrintLog(fmt.Sprintf("echo.New call"))
	e := echo.New()
	connectString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.Getconfig().Postgresql.Host,
		config.Getconfig().Postgresql.User,
		config.Getconfig().Postgresql.Pass,
		config.Getconfig().Postgresql.Db,
		config.Getconfig().Postgresql.Port,
	)
	var err error

	gormDb, err := gorm.Open(postgres.Open(connectString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	gormdb.Start(gormDb)
	g := e.Group("/api")
	g.POST("/login", nil)
	e.GET("/swageer/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":" + config.Getconfig().Port))

}
func Checktoken(context echo.Context) error {
	authercations := context.Request().Header.Get("Authorization")
	user := utils.ParseToken(authercations)
	userID := user.Claims.(jwt.MapClaims)["userID"].(string)
	if !utils.GetToken(authercations, userID) {
		return context.String(http.StatusForbidden, "token awrong")
	}
	return nil
}
