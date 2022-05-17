package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type EmployeeRepo interface {
	CreateEmployee(employee entities.Employee) error
	FindEmployeeByUsername(username string) (entities.Employee, error)
}

type employeeRepo struct {
	db *gorm.DB
}

func NewEmployeeRepo(db *gorm.DB) EmployeeRepo {
	return &employeeRepo{
		db: db,
	}
}

func (repo *employeeRepo) CreateEmployee(employee entities.Employee) error {
	return repo.db.Create(&employee).Error
}

func (repo *employeeRepo) FindEmployeeByUsername(username string) (entities.Employee, error) {
	employee := entities.Employee{}
	err := repo.db.
		// Model(&entities.Employee{
		// 	Username: username,
		// }).
		Where("username = ?", username).
		First(&employee).
		Error

	return employee, err
}
