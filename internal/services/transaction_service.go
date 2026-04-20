package services

import (
	"github.com/doteneff/bijakbudget-api/internal/models"
	"github.com/doteneff/bijakbudget-api/internal/repositories"
)

type TransactionService interface {
	CreateTransaction(transaction *models.Transaction) error
	GetAllTransactions() ([]models.Transaction, error)
	GetTransactionByID(id string) (*models.Transaction, error)
	UpdateTransaction(id string, data *models.Transaction) error
	DeleteTransaction(id string) error
}

type transactionService struct {
	repo repositories.TransactionRepository
}

func NewTransactionService(repo repositories.TransactionRepository) TransactionService {
	return &transactionService{repo}
}

func (s *transactionService) CreateTransaction(transaction *models.Transaction) error {
	// Possible business logic here (e.g. check if category exists, valid amount, deduct limit)
	return s.repo.Create(transaction)
}

func (s *transactionService) GetAllTransactions() ([]models.Transaction, error) {
	return s.repo.FindAll()
}

func (s *transactionService) GetTransactionByID(id string) (*models.Transaction, error) {
	return s.repo.FindByID(id)
}

func (s *transactionService) UpdateTransaction(id string, data *models.Transaction) error {
	existingTransaction, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	existingTransaction.Title = data.Title
	existingTransaction.Amount = data.Amount
	existingTransaction.Type = data.Type
	existingTransaction.CategoryID = data.CategoryID
	existingTransaction.Date = data.Date
	existingTransaction.Note = data.Note

	return s.repo.Update(existingTransaction)
}

func (s *transactionService) DeleteTransaction(id string) error {
	return s.repo.Delete(id)
}
