package database

import (
	"database/sql"
	"gabrielmaurici/first-go-api/internal/entity"
)

type PostDb struct {
	DB *sql.DB
}

func NewPostDb(db *sql.DB) *PostDb {
	return &PostDb{
		DB: db,
	}
}

func (p *PostDb) Get(id string) (*entity.Post, error) {
	post := &entity.Post{}

	stmt, err := p.DB.Prepare("SELECT id, title, body from posts where id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	if err := row.Scan(post.Id, post.Title, post.Body); err != nil {
		return nil, err
	}

	return post, nil
}

func (p *PostDb) Save(post *entity.Post) error {
	stmt, err := p.DB.Prepare("INSERT INTO posts(id, title, body) VALEUS(?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(post.Id, post.Title, post.Body)
	if err != nil {
		return err
	}

	return nil
}
