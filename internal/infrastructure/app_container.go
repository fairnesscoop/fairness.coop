package infrastructure

import (
	"os"
	"path/filepath"

	"fairness.coop/victor/internal/domain"
	"fairness.coop/victor/internal/domain/entities"
)

type AppContainer struct {
	collectionRepository *InMemoryCollectionRepository
	templateEngine       *FileTemplateEngine
}

func (ctn *AppContainer) GetCollectionRepository() domain.CollectionRepository {
	return ctn.collectionRepository
}

func (ctn *AppContainer) GetTemplateEngine() domain.TemplateEngine {
	return ctn.templateEngine
}

func NewContainer() (domain.Container, error) {
	collections, err := loadCollections()

	if err != nil {
		return nil, err
	}

	collectionRepository := NewInMemoryCollectionRepository(collections)

	templateEngine, err := NewFileTemplateEngine()

	if err != nil {
		return nil, err
	}

	return &AppContainer{collectionRepository: collectionRepository, templateEngine: templateEngine}, nil
}

func loadCollections() ([]*entities.Collection, error) {
	collections := []*entities.Collection{
		entities.NewCollection(
			1,
			"Blog posts",
			"posts",
			"content/blog",
		),
	}

	for _, col := range collections {
		err := filepath.Walk(col.Path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			pathInCollection, err := filepath.Rel(col.Path, path)

			if err != nil {
				return err
			}

			col.AddFile(entities.NewFile(col, pathInCollection))

			return nil
		})

		if err != nil {
			return nil, err
		}
	}

	return collections, nil
}
