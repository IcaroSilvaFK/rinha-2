package models

import "time"

type TransactionModel struct {
	Value       int `json:"valor"`
	UserId      int
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	CreateAt    time.Time `json:"realizada_em"`
}

func NewTransactionModel(
	description, transactionType string, value, userId int,
) *TransactionModel {
	return &TransactionModel{
		Description: description,
		Value:       value,
		Type:        transactionType,
		UserId:      userId,
		CreateAt:    time.Now(),
	}
}
