package repositories

import (
	"context"
	"database/sql"
	"errors"
	"library-api-author/internal/models"
)

type AuthorRepository interface {
	CreateAuthor(ctx context.Context, tx *sql.Tx, author *models.Author) error
	FindAuthorByID(ctx context.Context, tx *sql.Tx, id uint64) (*models.Author, error)
	UpdateAuthor(ctx context.Context, tx *sql.Tx, author *models.Author) error
	DeleteAuthor(ctx context.Context, tx *sql.Tx, id uint64) error
	GetAllAuthors(ctx context.Context, tx *sql.Tx, pagination *models.Pagination) ([]*models.Author, error)
}

type AuthorRepositoryImpl struct {
}

func NewAuthorRepository() AuthorRepository {
	return &AuthorRepositoryImpl{}
}

func (repository *AuthorRepositoryImpl) CreateAuthor(ctx context.Context, tx *sql.Tx, author *models.Author) error {
	query := `INSERT INTO authors (user_id, name, bio, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	response, err := tx.ExecContext(ctx, query, author.UserID, author.Name, author.Bio, author.CreatedAt, author.UpdatedAt)
	if err != nil || response == nil {
		return errors.New("Failed to create a author, transaction rolled back. Reason: " + err.Error())
	}

	return nil
}

func (repository *AuthorRepositoryImpl) FindAuthorByID(ctx context.Context, tx *sql.Tx, id uint64) (*models.Author, error) {
	query := "SELECT id, user_id, name, bio, created_at, updated_at FROM authors WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var author = models.Author{}
	if rows.Next() {
		err := rows.Scan(&author.ID, &author.UserID, &author.Name, &author.Bio, &author.CreatedAt, &author.UpdatedAt)
		if err != nil {
			return nil, err
		}
		return &author, nil
	} else {
		return nil, errors.New("author is not found")
	}
}

func (repository *AuthorRepositoryImpl) UpdateAuthor(ctx context.Context, tx *sql.Tx, author *models.Author) error {
	query := `UPDATE authors SET user_id = $1, name = $2, bio = $3, updated_at = $4 WHERE id = $5`

	_, err := tx.ExecContext(ctx, query,
		author.UserID,
		author.Name,
		author.Bio,
		author.UpdatedAt,
		author.ID,
	)
	if err != nil {
		return errors.New("Failed to update a author, transaction rolled back. Reason: " + err.Error())
	}
	return nil
}

func (repository *AuthorRepositoryImpl) DeleteAuthor(ctx context.Context, tx *sql.Tx, id uint64) error {
	SQL := `DELETE FROM authors WHERE id = $1`

	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return errors.New("Failed to update a author, transaction rolled back. Reason: " + err.Error())
	}
	return nil
}

func (repository *AuthorRepositoryImpl) GetAllAuthors(ctx context.Context, tx *sql.Tx, pagination *models.Pagination) ([]*models.Author, error) {
	query := `SELECT id, user_id, name, bio, created_at, updated_at FROM authors ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, query, pagination.PageSize, pagination.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []*models.Author
	for rows.Next() {
		var author models.Author
		err := rows.Scan(&author.ID, &author.UserID, &author.Name, &author.Bio, &author.CreatedAt, &author.UpdatedAt)
		if err != nil {
			return nil, err
		}

		authors = append(authors, &author)
	}
	return authors, nil
}
