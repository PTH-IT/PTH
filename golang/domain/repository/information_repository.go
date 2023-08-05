package repository

import "Golang/domain/model"

type InformatonRepository interface {
	AddInformation(information model.Information) error
	UpdateInformation(information model.Information) error
	GetInformation() *model.Information
}
