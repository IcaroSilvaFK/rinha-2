package controllers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/IcaroSilvaFK/rinha-backend-02/application/controllers/views"
	"github.com/IcaroSilvaFK/rinha-backend-02/application/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionController struct {
	db *pgxpool.Pool
}

type TransactionControllerInterface interface {
	CreateNewTransaction(w http.ResponseWriter, r *http.Request)
	ListTransactions(w http.ResponseWriter, r *http.Request)
}

var ctx = context.Background()

func NewTransactionController(
	db *pgxpool.Pool,
) TransactionControllerInterface {

	return &TransactionController{
		db,
	}
}

func (tc *TransactionController) CreateNewTransaction(w http.ResponseWriter, r *http.Request) {

	clientId, _ := strconv.Atoi(r.PathValue("userId"))

	bt, err := io.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var c models.ClientModel

	if err := tc.db.QueryRow(ctx, "SELECT accountLimit, initialBalance FROM clients WHERE id = $1", clientId).Scan(&c.Limit, &c.Balance); err != nil {

		w.WriteHeader(http.StatusNotFound)
		return
	}

	var t models.TransactionModel

	if err := json.Unmarshal(bt, &t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if t.Type == "d" && c.Limit > c.Balance-t.Value {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if t.Type == "c" {
		c.Balance += t.Value
	} else {

		c.Balance -= t.Value
	}

	t.UserId = clientId

	tx, _ := tc.db.Begin(ctx)
	tx.Exec(ctx,
		"INSERT INTO transactions (description, value, userId, type) VALUES ($1, $2, $3, $4)",
		t.Description, t.Value, t.UserId, t.Type,
	)

	tx.Exec(ctx,
		"UPDATE clients SET initialBalance = $1 WHERE id = $2",
		c.Balance, t.UserId,
	)

	err = tx.Commit(ctx)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ot := views.NewCreateTransactionOutputView(c)

	bt, _ = json.Marshal(ot)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bt)
}

func (tc *TransactionController) ListTransactions(w http.ResponseWriter, r *http.Request) {

	clientId, _ := strconv.Atoi(r.PathValue("userId"))

	var c models.ClientModel

	if err := tc.db.QueryRow(ctx, "SELECT * FROM clients WHERE id = $1", clientId).Scan(&c.ID, &c.Limit, &c.Balance); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	rows, _ := tc.db.Query(ctx, "SELECT description,value,userId,type,createdAt FROM transactions WHERE userId = $1 ORDER BY createdAt DESC LIMIT 10", clientId)

	defer rows.Close()

	var transactions []models.TransactionModel

	for rows.Next() {
		var t models.TransactionModel

		if err := rows.Scan(&t.Description, &t.Value, &t.UserId, &t.Type, &t.CreateAt); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		transactions = append(transactions, t)
	}

	result := views.NewListTransactionOutputView(c, transactions)

	response, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
