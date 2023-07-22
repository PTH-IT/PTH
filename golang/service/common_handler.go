package service

import (
	"Golang/usecase"

	"github.com/labstack/echo/v4"
)

type commonhandler struct {
	Interactor *usecase.Interactor
}

func AppV1GetUsers(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.GetUser(context)
	}

}
