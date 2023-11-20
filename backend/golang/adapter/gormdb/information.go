package gormdb

import (
	"Golang/domain/model"
	"Golang/domain/repository"
)

func NewInformation() repository.InformatonRepository {
	return informationRepository{}
}

type informationRepository struct {
}

func (repo informationRepository) AddInformation(information model.Information) error {
	return nil
}
func (repo informationRepository) UpdateInformation(information model.Information) error {
	return nil
}
func (repo informationRepository) GetInformation() *model.Information {
	return nil

}
