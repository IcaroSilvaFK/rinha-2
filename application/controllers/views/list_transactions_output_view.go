package views

import (
	"time"

	"github.com/IcaroSilvaFK/rinha-backend-02/application/models"
)

type Balance struct {
	Total       int       `json:"total"`
	DateExtract time.Time `json:"data_extrato"`
	Limit       int       `json:"limite"`
}

type LastTransactions struct {
	Description string    `json:"descricao"`
	Value       int       `json:"valor"`
	Type        string    `json:"tipo"`
	CreatedAt   time.Time `json:"realizada_em"`
}

type ListTransactionsOutputView struct {
	Balance          Balance            `json:"saldo"`
	LastTransactions []LastTransactions `json:"ultimas_transacoes"`
}

func NewListTransactionOutputView(
	c models.ClientModel, ts []models.TransactionModel,
) ListTransactionsOutputView {

	l := ListTransactionsOutputView{
		Balance: Balance{
			Total:       c.Balance,
			DateExtract: time.Now(),
			Limit:       c.Limit,
		},
		LastTransactions: []LastTransactions{},
	}

	for _, t := range ts {

		l.LastTransactions = append(l.LastTransactions, LastTransactions{
			Description: t.Description,
			Value:       t.Value,
			Type:        t.Type,
			CreatedAt:   t.CreateAt,
		})
	}

	return l
}
