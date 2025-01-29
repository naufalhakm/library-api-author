package factory

import (
	"database/sql"
	"library-api-author/internal/controllers"
	"library-api-author/internal/repositories"
	"library-api-author/internal/services"
)

type Provider struct {
	AuthorProvider controllers.AuthorController
}

func InitFactory(db *sql.DB) *Provider {

	authorRepo := repositories.NewAuthorRepository()
	authorService := services.NewAuthorService(db, authorRepo)
	AuthorController := controllers.NewAuthorController(authorService)

	return &Provider{
		AuthorProvider: AuthorController,
	}
}
