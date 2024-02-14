package main

import (
	"context"

	"github.com/IcaroSilvaFK/rinha-backend-02/application/database"
)

type User struct {
	ID      int
	Limit   int
	Balance int
}

func main() {
	db := database.NewDatabaseConnection()

	users := []User{
		{
			ID:      1,
			Limit:   100000,
			Balance: 0,
		},
		{
			ID:      2,
			Limit:   80000,
			Balance: 0,
		},
		{
			ID:      3,
			Limit:   1000000,
			Balance: 0,
		},
		{
			ID:      4,
			Limit:   10000000,
			Balance: 0,
		},
		{
			ID:      5,
			Limit:   500000,
			Balance: 0,
		},
	}

	ctx := context.Background()

	tx, _ := db.Begin(ctx)

	for _, user := range users {
		tx.Query(ctx, "INSERT INTO clients (id,accountLimit, initialBalance) VALUES ($1,$2, $3)", user.ID, user.Limit, user.Balance)
	}

	tx.Commit(ctx)

}
