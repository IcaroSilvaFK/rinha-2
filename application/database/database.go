package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
// single sync.Once
// host.docker.internal
// docker = "postgresql://admin:admin@host.docker.internal:5434/rinha?sslmode=disable"
)

/*
user=jack password=secret host=pg.example.com port=5432 dbname=mydb sslmode=verify-ca pool_max_conns=10

	docker.for.mac.host.internal
*/
func NewDatabaseConnection() *pgxpool.Pool {

	dbConf, err := pgxpool.ParseConfig("user=admin password=admin host=docker.for.mac.host.internal port=5432 dbname=rinha sslmode=disable")

	if err != nil {
		panic(err)
	}

	// dbConf.MaxConns = 50
	// dbConf.MinConns = 20

	db, err := pgxpool.NewWithConfig(context.Background(), dbConf)

	if err != nil {
		panic(err)
	}

	// generateTables(db)

	return db
}

// func generateTables(db *pgxpool.Pool) {

// 	_, err := db.Exec(context.Background(), `
// 		CREATE TABLE IF NOT EXISTS clients (
// 			id SERIAL PRIMARY KEY,
// 			accountLimit NUMERIC NOT NULL,
// 			initialBalance NUMERIC NOT NULL
// 		);

// 		CREATE TABLE IF NOT EXISTS transactions (
// 			id SERIAL PRIMARY KEY,
// 			description VARCHAR(255),
// 			value INTEGER,
// 			userId INTEGER,
// 			type VARCHAR(2),
// 			createdAt TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 			FOREIGN KEY (userId) REFERENCES clients(id)
// 		);
// 	`)

// 	if err != nil {
// 		log.Fatalf("Error creating tables: %v", err)
// 	}

// }
