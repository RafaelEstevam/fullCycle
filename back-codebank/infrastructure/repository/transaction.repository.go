package repository

import (
	"database/sql"

	"github.com/RafaelEstevam/fullCycle/back-codebank/domain"
)

type TransactionRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{db: db}
}

func (transactionRepositoryDb *TransactionRepositoryDb) SaveTransaction(transaction domain.Transaction, creditCard domain.CreditCard) error {
	stmt, err := transactionRepositoryDb.db.Prepare(`insert into transactions (id, credit_card_id, amount, status, description, store, created_at) values ($1, $2, $3, $4, $5, $6, $7)`)

	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		transaction.ID,
		transaction.CreditCardId,
		transaction.Amount,
		transaction.Status,
		transaction.Description,
		transaction.Store,
		transaction.CreatedAt,
	)
	if err != nil {
		return err
	}

	if transaction.Status == "approved" {
		err = transactionRepositoryDb.updateBalance(creditCard)
		if err != nil {
			return err
		}
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}

func (transactionRepositoryDb *TransactionRepositoryDb) updateBalance(creditCard domain.CreditCard) error {
	_, err := transactionRepositoryDb.db.Exec(`update credit_cards set balance = $1 where id = $2`, creditCard.Balance, creditCard.ID)
	if err != nil {
		return err
	}
	return nil
}

func (transactionRepositoryDb *TransactionRepositoryDb) CreateCreditCard(creditCard domain.CreditCard) error {

	stmt, err := transactionRepositoryDb.db.Prepare(`insert into credit_cards (id, name, number, expiration_month, expiration_year, cvv, balance, balance_limit) values ($1, $2, $3,$4,$5,$6,$7,$8)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		creditCard.ID,
		creditCard.Name,
		creditCard.Numner,
		creditCard.ExpirationMonth,
		creditCard.ExpirationYear,
		creditCard.CVV,
		creditCard.Balance,
		creditCard.Limit,
	)

	if err != nil {
		return nil
	}

	err = stmt.Close()

	if err != nil {
		return nil
	}

	return nil
}
