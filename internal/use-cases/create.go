package use_cases

import (
	"log"

	"github.com/nathanfabio/go-microservice/internal/entities"
	"github.com/nathanfabio/go-microservice/internal/repositories"
)

type createCategory struct {
	repository repositories.ICategoryRepository
}

func NewCreateCategory(repository repositories.ICategoryRepository) *createCategory {
	return &createCategory{repository}
}

func (c *createCategory) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	//todo: persist entity to db
	log.Println(category)

	err = c.repository.Save(category)
	if err != nil {
		return err
	}

	return nil
}