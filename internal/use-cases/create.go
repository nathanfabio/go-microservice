package use_cases

import (
	"log"

	"github.com/nathanfabio/go-microservice/internal/entities"
)

type createCategory struct {
	//db
}

func NewCreateCategory() *createCategory {
	return &createCategory{}
}

func (c *createCategory) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	//todo: persist entity to db
	log.Println(category)


	return nil
}