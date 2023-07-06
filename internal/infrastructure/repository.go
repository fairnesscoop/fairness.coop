package infrastructure

import "fairness.coop/victor/internal/domain/entities"

type InMemoryCollectionRepository struct {
	collections []*entities.Collection
}

func NewInMemoryCollectionRepository(collections []*entities.Collection) *InMemoryCollectionRepository {
	return &InMemoryCollectionRepository{collections: collections}
}

func (r *InMemoryCollectionRepository) GetAll() []*entities.Collection {
	return r.collections
}

func (r *InMemoryCollectionRepository) GetById(id int) *entities.Collection {
	for _, col := range r.collections {
		if col.Id == id {
			return col
		}
	}

	return nil
}
