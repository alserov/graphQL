package db

import (
	"context"
	"database/sql"
	"github.com/alserov/graphQL/graph/model"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func New() (*sql.DB, error) {
	var err error

	dts := "host=localhost port=5432 password=1787 user=postgres dbname=template1 sslmode=disable"
	DB, err = sql.Open("postgres", dts)
	if err != nil {
		return nil, err
	}

	if err = DB.Ping(); err != nil {
		return nil, err
	}

	return DB, nil
}

type Repository interface {
	Create(ctx context.Context, video model.Video) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context) ([]*model.Video, error)
}

func NewRepo() Repository {
	return &repo{
		db: DB,
	}
}

type repo struct {
	db *sql.DB
}

func (r *repo) Create(ctx context.Context, video model.Video) error {
	query := `INSERT INTO videos (id,title,url,authorId) VALUES ($1,$2,$3,$4)`
	_, err := r.db.Query(query, video.ID, video.Title, video.URL, video.Author.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM videos WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Get(ctx context.Context) ([]*model.Video, error) {
	query := `SELECT id,title,url FROM videos`

	var videos []*model.Video

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var video model.Video
		err = rows.Scan(&video.ID, &video.Title, &video.URL)
		if err != nil {
			return nil, err
		}
		videos = append(videos, &video)
	}

	return videos, nil
}
