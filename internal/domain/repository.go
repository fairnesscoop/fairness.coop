package domain

import "fairness.coop/victor/internal/domain/entities"

type CollectionRepository interface {
	GetAll() []*entities.Collection
	GetById(id int) *entities.Collection
}
