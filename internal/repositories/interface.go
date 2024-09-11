package repositories

import "github.com/nathanfabio/go-microservice/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
}