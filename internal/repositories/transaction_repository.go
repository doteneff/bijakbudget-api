package repositories

import (
	"github.com/doteneff/bijakbudget-api/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	FindAll() ([]models.Transaction, error)
	FindByID(id string) (*models.Transaction, error)
	Update(transaction *models.Transaction) error
	Delete(id string) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) Create(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) FindAll() ([]models.Transaction, error) {
	var transactions []models.Transaction
	// Preload the related Category and Member to have full context in JSON responses
	err := r.db.Preload("Category").Preload("Member").Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) FindByID(id string) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Category").Preload("Member").First(&transaction, "id = ?", id).Error
	return &transaction, err
}

func (r *transactionRepository) Update(transaction *models.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r *transactionRepository) Delete(id string) error {
	return r.db.Delete(&models.Transaction{}, "id = ?", id).Error
}
