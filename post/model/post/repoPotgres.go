package post

import (
	"context"
	"fmt"
	UUID "github.com/google/uuid"
	"github.com/viciousvs/blog-microservices/post/config"
	"github.com/viciousvs/blog-microservices/post/storage/postgresRepo"
	"log"
	"strconv"
	"strings"
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
		, post.Title, post.Content, post.CreatedAt, post.UpdatedAt)
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
		return posts, fmt.Errorf("cannot run Query select *, err:%v",err)
	}
	defer func() {
		if err:= rows.Close; err != nil{
			log.Printf("close row error: %v", err)
		}
	}()

	for rows.Next(){
		var p Post
		err := rows.Scan(&p.UUID, &p.Title, &p.Content, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return posts, fmt.Errorf("cannot scan to post, err: %v",err)
		}
		posts = append(posts, &p)
	}
	if err:=rows.Err(); err != nil{
		return posts, fmt.Errorf("rows.Err(), err:%v", err)
	}
	return posts, nil
}

func (r *repoPostgres) GetById(ctx context.Context, uuid string) (Post, error) {
	var p Post
	err := r.pgSource.QueryRow(ctx, `select * from post where id=$1`,uuid).
		Scan(&p.UUID, &p.Title, &p.Content, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return p, fmt.Errorf("cannot run query row post, err: %v", err)
	}
	return p, nil
}

func (r *repoPostgres) Update(ctx context.Context, post Post) (string, error) {
	tmpCreatedAt := post.CreatedAt
	query := `update post set`
	queryCounter := 1
	if strings.Trim(post.Title, " ") != ""{
		query += " title=$" + strconv.Itoa(queryCounter)
		queryCounter++
	}
	if strings.Trim(post.Content, " ") != ""{
		query += " content=$" + strconv.Itoa(queryCounter)
		queryCounter++
	}

	r.pgSource.BeginTx()

	ct, err := r.pgSource.Exec(ctx, `update post set `, uuid)
	if err != nil {
		post.CreatedAt = tmpCreatedAt
		return "", fmt.Errorf("cannot run Exec insert into, err:%v", err)
	}
	if r := ct.RowsAffected(); r != 1 {
		post.CreatedAt = tmpCreatedAt
		return "", fmt.Errorf("count if RowsAffected != 1")
	}
	return uuid, nil
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
