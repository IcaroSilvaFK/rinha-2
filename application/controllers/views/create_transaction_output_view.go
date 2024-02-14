package views

import "github.com/IcaroSilvaFK/rinha-backend-02/application/models"

type CreateTransactionOutputView struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}

func NewCreateTransactionOutputView(
	c models.ClientModel,
) CreateTransactionOutputView {

	return CreateTransactionOutputView{
		Limit:   c.Limit,
		Balance: c.Balance,
	}
}
