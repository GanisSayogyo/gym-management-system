package database

import (
	"context"
	"fmt"
	"net"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/GanisSayogyo/gym-management-system/config"
)

// NewPool membuat connection pool ke PostgreSQL.
func NewPool(ctx context.Context) (*pgxpool.Pool, error) {
	// Mengambil konfigurasi database dari environment.
	host := config.GetEnv("DB_HOST", "localhost")
	port := config.GetEnv("DB_PORT", "5432")
	user := config.GetEnv("DB_USER", "postgres")
	password := config.GetEnv("DB_PASSWORD", "")
	dbName := config.GetEnv("DB_NAME", "gym_management")
	sslMode := config.GetEnv("DB_SSLMODE", "disable")

	// Menyusun connection string.
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, password),
		Host:   net.JoinHostPort(host, port),
		Path:   dbName,
	}

	query := dsn.Query()
	query.Set("sslmode", sslMode)
	dsn.RawQuery = query.Encode()

	// Membuat connection pool.
	pool, err := pgxpool.New(ctx, dsn.String())
	if err != nil {
		return nil, fmt.Errorf("gagal membuat koneksi database: %w", err)
	}

	// Memastikan database benar-benar dapat dihubungi.
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("gagal menghubungi database: %w", err)
	}

	return pool, nil
}
