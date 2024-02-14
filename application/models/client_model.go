package models

type ClientModel struct {
	ID      int
	Limit   int
	Balance int
}

func NewClientModel(
	id, limit, balance int,
) *ClientModel {
	return &ClientModel{
		ID:      id,
		Limit:   limit,
		Balance: balance,
	}
}
