package services

import (
	"context"
	"database/sql"
	"library-api-author/internal/commons/response"
	"library-api-author/internal/models"
	"library-api-author/internal/params"
	"library-api-author/internal/repositories"
	"time"
)

type AuthorService interface {
	CreateAuthor(ctx context.Context, req *params.AuthorRequest) *response.CustomError
	GetDetailAuthor(ctx context.Context, id uint64) (*params.AuthorResponse, *response.CustomError)
	UpdateAuthor(ctx context.Context, id uint64, req *params.AuthorRequest) *response.CustomError
	DeleteAuthor(ctx context.Context, id uint64) *response.CustomError
	GetAllAuthors(ctx context.Context, pagination *models.Pagination) ([]*params.AuthorResponse, *response.CustomError)
}

type AuthorServiceImpl struct {
	DB               *sql.DB
	AuthorRepository repositories.AuthorRepository
}

func NewAuthorService(db *sql.DB, authorRepository repositories.AuthorRepository) AuthorService {
	return &AuthorServiceImpl{
		DB:               db,
		AuthorRepository: authorRepository,
	}
}

func (service *AuthorServiceImpl) CreateAuthor(ctx context.Context, req *params.AuthorRequest) *response.CustomError {
	tx, err := service.DB.Begin()
	if err != nil {
		return response.GeneralError("Failed Connection to database errors:" + err.Error())
	}
	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var author = models.Author{
		UserID:    req.UserID,
		Name:      req.Name,
		Bio:       req.Bio,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = service.AuthorRepository.CreateAuthor(ctx, tx, &author)

	if err != nil {
		return response.GeneralError(err.Error())
	}

	return nil
}

func (service *AuthorServiceImpl) GetDetailAuthor(ctx context.Context, id uint64) (*params.AuthorResponse, *response.CustomError) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, response.GeneralError("Failed Connection to database errors: " + err.Error())
	}
	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	author, err := service.AuthorRepository.FindAuthorByID(ctx, tx, id)
	if err != nil {
		return nil, response.NotFoundError("author not found")
	}

	AuthorResponse := &params.AuthorResponse{
		ID:        author.ID,
		UserID:    author.UserID,
		Name:      author.Name,
		Bio:       author.Bio,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}

	return AuthorResponse, nil
}

func (service *AuthorServiceImpl) UpdateAuthor(ctx context.Context, id uint64, req *params.AuthorRequest) *response.CustomError {
	tx, err := service.DB.Begin()
	if err != nil {
		return response.GeneralError("Failed to connect to the database: %s", err.Error())
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	author := models.Author{
		ID:        id,
		UserID:    req.UserID,
		Name:      req.Name,
		Bio:       req.Bio,
		UpdatedAt: time.Now(),
	}

	err = service.AuthorRepository.UpdateAuthor(ctx, tx, &author)
	if err != nil {
		tx.Rollback()
		return response.GeneralError("Failed to update author: %s", err.Error())
	}

	return nil
}

func (service *AuthorServiceImpl) DeleteAuthor(ctx context.Context, id uint64) *response.CustomError {
	tx, err := service.DB.Begin()
	if err != nil {
		return response.GeneralError("Failed to connect to the database: " + err.Error())
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = service.AuthorRepository.DeleteAuthor(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return response.GeneralError("Failed to delete author: " + err.Error())
	}

	return nil
}

func (service *AuthorServiceImpl) GetAllAuthors(ctx context.Context, pagination *models.Pagination) ([]*params.AuthorResponse, *response.CustomError) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, response.GeneralError("Failed to connect to the database: " + err.Error())
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	pagination.Offset = (pagination.Page - 1) * pagination.PageSize

	authors, err := service.AuthorRepository.GetAllAuthors(ctx, tx, pagination)
	if err != nil {
		return nil, response.GeneralError("Failed to fetch authors: " + err.Error())
	}

	AuthorResponses := make([]*params.AuthorResponse, len(authors))
	for i, author := range authors {
		AuthorResponses[i] = &params.AuthorResponse{
			ID:        author.ID,
			UserID:    author.UserID,
			Name:      author.Name,
			Bio:       author.Bio,
			CreatedAt: author.CreatedAt,
			UpdatedAt: author.UpdatedAt,
		}
	}

	pagination.PageCount = (pagination.TotalCount + pagination.PageSize - 1) / pagination.PageSize

	return AuthorResponses, nil
}
