package service

import (
	"Golang/usecase"

	"github.com/labstack/echo/v4"
)

type commonhandler struct {
	Interactor *usecase.Interactor
}

func AppV1Login(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.Login(context)
	}

}
func AppV1Register(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.Register(context)
	}

}
func AppV1Logout(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.Logout(context)
	}

}
func AppV1VerifyToken(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.VerifyToken(context)
	}

}
