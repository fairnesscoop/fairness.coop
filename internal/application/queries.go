package application

import (
	"fmt"

	"fairness.coop/victor/internal/domain"
	"fairness.coop/victor/internal/domain/entities"
)

func GetAllCollections(ctn domain.Container) []*entities.Collection {
	return ctn.GetCollectionRepository().GetAll()
}

func GetCollectionById(ctn domain.Container, id int) (*entities.Collection, error) {
	collection := ctn.GetCollectionRepository().GetById(id)

	if collection == nil {
		return nil, fmt.Errorf("collection with id %d does not exist", id)
	}

	return collection, nil
}

func GetFileByPathAndCollectionId(ctn domain.Container, path string, collectionId int) (*entities.File, error) {
	collection, err := GetCollectionById(ctn, collectionId)

	if err != nil {
		return nil, err
	}

	file := collection.FindFileByPath(path)

	if file == nil {
		return nil, fmt.Errorf("file with path %s does not exist in collection %s", path, collection.Name)
	}

	return file, nil
}
