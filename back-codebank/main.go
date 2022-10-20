package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/RafaelEstevam/fullCycle/back-codebank/domain"
	"github.com/RafaelEstevam/fullCycle/back-codebank/infrastructure/repository"
	"github.com/RafaelEstevam/fullCycle/back-codebank/usecase"
	_ "github.com/lib/pq"
)

func main() {
	db := setupDb()
	defer db.Close()
	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.CVV = 123
	cc.ExpirationMonth = 10
	cc.ExpirationYear = 30
	cc.Limit = 3000
	cc.Name = "Bradesco"
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)
	err := repo.CreateCreditCard(*cc)

	if err != nil {
		fmt.Println(err)
	}
}

func setupTransactionUsecase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	usecase := usecase.NewUseCaseTransaction(transactionRepository)
	return usecase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "db", "5432", "postgres", "root", "codebank")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Erro connection")
	}

	return db

}
