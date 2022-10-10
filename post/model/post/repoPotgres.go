package post

import (
	"context"
	"errors"
	"fmt"
	UUID "github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/viciousvs/blog-microservices/post/config"
	"github.com/viciousvs/blog-microservices/post/storage/postgresRepo"
	"log"
	"time"
)

type repoPostgres struct {
	pgSource *postgresRepo.PostgresDB
}

func (r *repoPostgres) getNewUUID() string {
	uuid := UUID.New()
	return uuid.String()
}
func NewRepoPostgres(config config.PostgresConfig) Repository {
	return &repoPostgres{pgSource: postgresRepo.NewPostgresDB(config)}
}

func (r *repoPostgres) Create(ctx context.Context, post Post) (string, error) {
	post.UUID = r.getNewUUID()
	now := time.Now().Unix()
	post.CreatedAt, post.UpdatedAt = now, now

	ct, err := r.pgSource.Exec(ctx, `insert into post (id, title, content, created_at, deleted_at) values($1, $2, $3, $4, $5)`,
		post.Title, post.Content, post.CreatedAt, post.UpdatedAt)
	if err != nil {
		return "", fmt.Errorf("cannot run Exec insert into, err:%v", err)
	}
	if r := ct.RowsAffected(); r != 1 {
		return "", fmt.Errorf("count if RowsAffected != 1")
	}
	return post.UUID, nil
}

func (r *repoPostgres) GetAll(ctx context.Context) ([]*Post, error) {
	posts := make([]*Post, 0)
	rows, err := r.pgSource.Query(ctx, `select * from post`)
	if err != nil {
		if err == pgx.ErrNoRows {
			return posts, fmt.Errorf("No rows, err: %v", err)
		}
		return posts, fmt.Errorf("cannot run Query select *, err:%v", err)
	}
	defer func() {
		if err := rows.Close; err != nil {
			log.Printf("close row error: %v", err)
		}
	}()

	for rows.Next() {
		var p Post
		err := rows.Scan(&p.UUID, &p.Title, &p.Content, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return posts, fmt.Errorf("cannot scan to post, err: %v", err)
		}
		posts = append(posts, &p)
	}
	if err := rows.Err(); err != nil {
		return posts, fmt.Errorf("rows.Err(), err:%v", err)
	}
	return posts, nil
}

func (r *repoPostgres) GetById(ctx context.Context, uuid string) (Post, error) {
	var p Post
	err := r.pgSource.QueryRow(ctx, `select * from post where id=$1`, uuid).
		Scan(&p.UUID, &p.Title, &p.Content, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return p, fmt.Errorf("cannot run query row post, err: %v", err)
	}
	return p, nil
}

func (r *repoPostgres) Update(ctx context.Context, post Post) (string, error) {
	if post.Title == "" && post.Content == "" {
		return "", errors.New("nothing to update")
	}
	tr, err := r.pgSource.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer func() {
		if err != nil {
			// rollback
		}
	}()
	if post.Title == "" {
		_, err := tr.Exec(ctx, `update post set title=$1 where id=$2`, post.Title, post.UUID)
		if err != nil {
			return "", err
		}
	}
	if post.Content == "" {
		_, err := tr.Exec(ctx, `update post set content=$1 where id=$2`, post.Content, post.UUID)
		if err != nil {
			return "", err
		}
	}

	now := time.Now().Unix()
	_, err = tr.Exec(ctx, `update post set updated_at=$1 where id=$2`, now, post.UUID)
	if err != nil {
		return "", err
	}

	err = tr.Commit(ctx)
	if err != nil {
		return "", fmt.Errorf("Cannot commit update transaction, err: %v", err)
	}
	return post.UUID, nil
}

func (r *repoPostgres) Delete(ctx context.Context, uuid string) (string, error) {
	ct, err := r.pgSource.Exec(ctx, `delete from post where id=$1`, uuid)
	if err != nil {
		return "", fmt.Errorf("cannot run Exec insert into, err:%v", err)
	}
	if r := ct.RowsAffected(); r != 1 {
		return "", fmt.Errorf("count if RowsAffected != 1")
	}
	return uuid, nil
}
