package postgresRepo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/viciousvs/blog-microservices/post/config"
	"log"
	"sync"
	"time"
)

var postgresDB *PostgresDB
var initOnce sync.Once

type PostgresDB struct {
	*pgxpool.Pool
}

func NewPostgresDB(config config.PostgresConfig) *PostgresDB {
	initOnce.Do(func() {
		var err error
		postgresDB, err = newPostgresRepository(config)
		if err != nil {
			log.Fatalf("cannot connect to postgres db: %v", err)
		}
	})

	return postgresDB
}

func newPostgresRepository(config config.PostgresConfig) (*PostgresDB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DatabaseName)

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	err = postgresDB.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &PostgresDB{pool}, nil
}
