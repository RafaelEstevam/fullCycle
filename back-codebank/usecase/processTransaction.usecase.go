package usecase

import (
	"time"

	"github.com/RafaelEstevam/fullCycle/back-codebank/domain"
	"github.com/RafaelEstevam/fullCycle/back-codebank/dto"
)

type UseCaseTransaction struct {
	TransactionRepository domain.TransactionRepository
}

func NewUseCaseTransaction(transactionRepository domain.TransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{TransactionRepository: transactionRepository}
}

func (useCaseTransaction UseCaseTransaction) ProcessTransaction(transactionDto dto.Transaction) (domain.Transaction, error) {
	creditCard := useCaseTransaction.setCreditCard(transactionDto)
	ccBalanceLimit, err := useCaseTransaction.TransactionRepository.GetCreditCard(*creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}
	creditCard.ID = ccBalanceLimit.ID
	creditCard.Limit = ccBalanceLimit.Limit
	creditCard.Balance = ccBalanceLimit.Balance
	transaction := useCaseTransaction.newTransaction(transactionDto, ccBalanceLimit)
	transaction.ProcessAndValidate(creditCard)
	err = useCaseTransaction.TransactionRepository.SaveTransaction(*transaction, *creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}
	return *transaction, nil
}

func (UseCaseTransaction UseCaseTransaction) setCreditCard(transactionDto dto.Transaction) *domain.CreditCard {
	creditCard := domain.NewCreditCard()
	creditCard.Name = transactionDto.Name
	creditCard.Number = transactionDto.Number
	creditCard.ExpirationMonth = transactionDto.ExpirationMonth
	creditCard.ExpirationYear = transactionDto.ExpirationYear
	creditCard.CVV = transactionDto.CVV
	return creditCard
}

func (UseCaseTransaction UseCaseTransaction) newTransaction(transactionDto dto.Transaction, creditCard domain.CreditCard) *domain.Transaction {
	transaction := domain.NewTransaction()
	transaction.ID = creditCard.ID
	transaction.Amount = transactionDto.Amount
	transaction.Store = transactionDto.Store
	transaction.Description = transactionDto.Description
	transaction.CreatedAt = time.Now()
	return transaction
}
