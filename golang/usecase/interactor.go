package usecase

import (
	"fmt"
	"net/http"

	gormdb "Golang/adapter/gormdb"
	"Golang/domain/model"
	"Golang/utils"

	InforLog "Golang/log/infor"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewInteractor(
	// gormDb *gorm.DB,
	referrance Reference,

) Interactor {

	return Interactor{
		// gormDb,
		referrance,
	}
}

type Interactor struct {
	// gormDb     *gorm.DB
	referrance Reference
}

// LoginUser godoc
// @Summary LoginUser
// @Description login username
// @Tags gormDB
// @Accept json
// @Produce json
// @Param user body  model.Login true "model.Login"
// @Success 201 {object} model.Token
// @Failure 400 {object} string
// @Router /gormdb/login [post]
func (i *Interactor) LoginUserGormdb(context echo.Context) error {
	InforLog.PrintLog(fmt.Sprintf("LoginUserGormdb start"))

	var user model.Login
	err := context.Bind(&user)
	if err != nil {
		return context.String(http.StatusBadRequest, "no user")
	}

	result, err := i.referrance.GetUserGormdb(user.UserID, *utils.CryptPassword(user.Password))
	if err != nil {
		return err
	}
	if result == nil {
		return context.String(http.StatusBadRequest, "user no exist")
	}

	tokenString := utils.GenerateToken(result.UserID)
	token := &model.Token{
		Authorization: tokenString,
		Type:          "bearer",
	}
	err = utils.SetToken(tokenString, user.UserID)
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, token)

}

// GetUser godoc
// @Summary GetUser
// @Description get username from token
// @Tags gormDB
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /gormdb/user [get]
func (i *Interactor) GetUser(context echo.Context) error {
	InforLog.PrintLog(fmt.Sprintf("GetUserGormdb start"))

	authercations := context.Request().Header.Get("token")
	user := utils.ParseToken(authercations)
	userID := user.Claims.(jwt.MapClaims)["userID"].(string)
	if !utils.GetToken(authercations, userID) {
		return context.String(http.StatusForbidden, "token awrong")
	}
	return context.JSON(http.StatusOK, userID)

}

// AddUser godoc
// @Summary AddUser
// @Description Add new user to database
// @Tags gormDB
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorizationc"
// @Param token body model.RegisterUser true "model.RegisterUser"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /gormdb/adduser [post]
func (i *Interactor) AddUserGormdb(context echo.Context) error {
	InforLog.PrintLog(fmt.Sprintf("AddUserGormdb start"))

	authercations := context.Request().Header.Get("Authorization")
	user := utils.ParseToken(authercations)
	userID := user.Claims.(jwt.MapClaims)["userID"].(string)
	if !utils.GetToken(authercations, userID) {
		return context.String(http.StatusForbidden, "Authorization awrong")
	}

	var Adduser model.RegisterUser
	err := context.Bind(&Adduser)

	if err != nil {
		return context.String(http.StatusBadRequest, "no user")
	}
	if userID == Adduser.UserID {
		return context.String(http.StatusBadRequest, "user exists")
	}
	cryptPassword := utils.CryptPassword(Adduser.Password)
	err = gormdb.Begin().Error
	if err != nil {
		return err
	}
	err = i.referrance.AddtUserGormdb(Adduser.UserID, *cryptPassword)
	if err != nil {
		return err
	}
	err = gormdb.Commit().Error
	if err != nil {
		return err
	}
	return context.String(http.StatusOK, "susscess")
}

// GetLogout godoc
// @Summary GetLogout
// @Description GetLogout
// @Tags MonggoDB
// @Accept json
// @Produce json
// @Param Content-Type header string true "application/json" default(application/json)
// @Param Content-Length header string true "1000" default(1000)
// @Param Host header string true "localhost" default(localhost)
// @Param Authorization header string true "Authorization"
// @Success 200 {object} string
// @Failure 400 {object} error
// @Router /logout [get]
func (i *Interactor) GetLogout(context echo.Context) error {
	InforLog.PrintLog(fmt.Sprintf("GetLogout start"))

	authercations := context.Request().Header.Get("Authorization")
	InforLog.PrintLog(fmt.Sprintf("authercations = %s", authercations))

	user := utils.ParseToken(authercations)
	InforLog.PrintLog(fmt.Sprintf("user = %v", user))

	userID := user.Claims.(jwt.MapClaims)["userID"].(string)
	InforLog.PrintLog(fmt.Sprintf("userID = %s", userID))
	InforLog.PrintLog(fmt.Sprintf("utils.GetToken"))

	if !utils.GetToken(authercations, userID) {
		return context.String(http.StatusForbidden, "token awrong")
	}
	InforLog.PrintLog(fmt.Sprintf("utils.DeleteToken"))

	if !utils.DeleteToken(authercations, userID) {
		return context.String(http.StatusBadRequest, "Can not delete token")
	}
	InforLog.PrintLog(fmt.Sprintf("StatusOK"))

	return context.String(http.StatusOK, "susscess")
}
