package domain

type Container interface {
	GetCollectionRepository() CollectionRepository
	GetTemplateEngine() TemplateEngine
}
