package services

import (
	"github.com/doteneff/bijakbudget-api/internal/models"
	"github.com/doteneff/bijakbudget-api/internal/repositories"
)

type CategoryService interface {
	CreateCategory(category *models.Category) error
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id string) (*models.Category, error)
	UpdateCategory(id string, categoryData *models.Category) error
	DeleteCategory(id string) error
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) CreateCategory(category *models.Category) error {
	return s.repo.Create(category)
}

func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) GetCategoryByID(id string) (*models.Category, error) {
	return s.repo.FindByID(id)
}

func (s *categoryService) UpdateCategory(id string, categoryData *models.Category) error {
	// First check if exists
	existingCategory, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// Apply updates
	existingCategory.Name = categoryData.Name
	existingCategory.Icon = categoryData.Icon
	existingCategory.Color = categoryData.Color
	existingCategory.MonthlyLimit = categoryData.MonthlyLimit
	existingCategory.IsIncome = categoryData.IsIncome

	return s.repo.Update(existingCategory)
}

func (s *categoryService) DeleteCategory(id string) error {
	return s.repo.Delete(id)
}
