package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type EmployeeRepo interface {
	Create(employee entities.Employee) error
	Update(employee entities.Employee) error
	FindByUsername(username string) (entities.Employee, error)
	FindEmployee(employeeID uint, providerID uint) (entities.Employee, error)
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

func (repo *employeeRepo) Update(employee entities.Employee) error {
	return repo.db.Save(&employee).Error
}

func (repo *employeeRepo) FindByUsername(username string) (entities.Employee, error) {
	employee := entities.Employee{}
	err := repo.db.
		Where(&entities.Employee{Username: username}).
		First(&employee).
		Error

	return employee, err
}

func (repo *employeeRepo) FindEmployee(employeeID uint, providerID uint) (entities.Employee, error) {
	employee := entities.Employee{}
	err := repo.db.
		//Model(&employee).
		Where(&entities.Employee{
			Model: gorm.Model{
				ID: employeeID,
			},
			ProviderID: providerID,
		}).
		Find(&employee).
		Error

	return employee, err
}
