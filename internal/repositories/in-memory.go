package repositories

import "github.com/nathanfabio/go-microservice/internal/entities"

type inMemoryCategoryRepository struct {
	db []*entities.Category
}

// NewInMemoryCategoryRepository returns an in-memory implementation of repositories.ICategoryRepository.
func NewInMemoryCategoryRepository() *inMemoryCategoryRepository {
	return &inMemoryCategoryRepository{
		db: make([]*entities.Category, 0),
	}
}


func (r *inMemoryCategoryRepository) Save(category *entities.Category) error {

	r.db = append(r.db, category)
	return nil
}

func (r *inMemoryCategoryRepository) List() ([]*entities.Category, error) {
	return r.db, nil
}