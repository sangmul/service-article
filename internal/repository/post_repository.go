package repository

import (
	"article/internal/domain"
	"database/sql"
)

type PostRepository interface {
	Create(post domain.Post) error
	GetAll() ([]domain.Post, error)
	GetWithPagination(limit, offset int) ([]domain.Post, error)
	GetByID(id int) (domain.Post, error)
	Update(id int, post domain.Post) error
	Delete(id int) error
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) Create(post domain.Post) error {
	query := `INSERT INTO post (title, content, category, status) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, post.Title, post.Content, post.Category, post.Status)
	return err
}

func (r *postRepository) GetAll() ([]domain.Post, error) {
	rows, err := r.db.Query(`SELECT id, title, content, category, created_date, updated_date, status FROM post`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		var p domain.Post
		rows.Scan(&p.ID, &p.Title, &p.Content, &p.Category, &p.CreatedDate, &p.UpdatedDate, &p.Status)
		posts = append(posts, p)
	}
	return posts, nil
}

func (r *postRepository) GetWithPagination(limit, offset int) ([]domain.Post, error) {
	query := `SELECT id, title, content, category, created_date, updated_date, status 
              FROM post 
              LIMIT ? OFFSET ?`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		var p domain.Post
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.Category,
			&p.CreatedDate,
			&p.UpdatedDate,
			&p.Status,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}

func (r *postRepository) GetByID(id int) (domain.Post, error) {
	var p domain.Post

	query := `SELECT id, title, content, category, created_date, updated_date, status 
              FROM post WHERE id = ?`

	err := r.db.QueryRow(query, id).Scan(
		&p.ID,
		&p.Title,
		&p.Content,
		&p.Category,
		&p.CreatedDate,
		&p.UpdatedDate,
		&p.Status,
	)

	return p, err
}

func (r *postRepository) Update(id int, post domain.Post) error {
    query := `UPDATE post 
              SET title = ?, content = ?, category = ?, status = ?
              WHERE id = ?`

    result, err := r.db.Exec(query,
        post.Title,
        post.Content,
        post.Category,
        post.Status,
        id,
    )
    if err != nil {
        return err
    }

    rows, _ := result.RowsAffected()
    if rows == 0 {
        return sql.ErrNoRows
    }

    return nil
}

func (r *postRepository) Delete(id int) error {
    query := `DELETE FROM post WHERE id = ?`

    result, err := r.db.Exec(query, id)
    if err != nil {
        return err
    }

    rows, _ := result.RowsAffected()
    if rows == 0 {
        return sql.ErrNoRows
    }

    return nil
}
