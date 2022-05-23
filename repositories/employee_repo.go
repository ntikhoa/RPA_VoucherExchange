package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type EmployeeRepo interface {
	Create(employee entities.Employee) error
	FindByUsername(username string) (entities.Employee, error)
}

type employeeRepo struct {
	db *gorm.DB
}

func NewEmployeeRepo(db *gorm.DB) EmployeeRepo {
	return &employeeRepo{
		db: db,
	}
}

func (repo *employeeRepo) Create(employee entities.Employee) error {
	return repo.db.Create(&employee).Error
}

func (repo *employeeRepo) FindByUsername(username string) (entities.Employee, error) {
	employee := entities.Employee{}
	err := repo.db.
		Where("username = ?", username).
		First(&employee).
		Error

	return employee, err
}
