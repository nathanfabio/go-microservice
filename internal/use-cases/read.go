package use_cases


import (
	"github.com/nathanfabio/go-microservice/internal/entities"
	"github.com/nathanfabio/go-microservice/internal/repositories"
)

type listCategory struct {
	repository repositories.ICategoryRepository
}

func NewlistCategory(repository repositories.ICategoryRepository) *listCategory {
	return &listCategory{repository}
}

func (c *listCategory) Execute() ([]*entities.Category, error) {

	categories, err := c.repository.List()
	if err != nil {
		return nil, err
	}

	return categories, nil
}