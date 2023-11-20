package service

import (
	InforLog "Golang/log/infor"
	"Golang/usecase"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	db := gormdb.NewGormDb()
	db.Start(gormDb)
	userRepository := gormdb.NewUser()
	referrance := usecase.NewReferrance(userRepository)
	interactor := usecase.NewInteractor(db, referrance)

	api := commonhandler{
		Interactor: &interactor,
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.POST("/login", AppV1Login(api))
	e.POST("/register", AppV1Register(api))
	g := e.Group("/api")
	g.POST("/verifytoken", AppV1VerifyToken(api))
	g.POST("/logout", AppV1Logout(api))
	g.POST("/addinformation", AppV1Logout(api))
	g.POST("/editinformation", AppV1Logout(api))
	e.GET("/swagger/*", echoSwagger.WrapHandler)
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
