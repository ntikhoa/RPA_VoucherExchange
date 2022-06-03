package services

import (
	"errors"
	"log"
	"time"

	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/custom_error"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(registerDTO dto.RegisterDTO) error
	Login(loginDTO dto.LoginDTO) (entities.Employee, error)
}

type authService struct {
	employeeRepo repositories.EmployeeRepo
	providerRepo repositories.ProviderRepo
}

func NewAuthService(employeeRepo repositories.EmployeeRepo,
	providerRepo repositories.ProviderRepo) AuthService {
	return &authService{
		employeeRepo: employeeRepo,
		providerRepo: providerRepo,
	}
}

func (s *authService) Register(registerDTO dto.RegisterDTO) error {
	if registerDTO.Password != registerDTO.ConfirmedPassword {
		return custom_error.NewBadRequestError(constants.CONFIRMED_PASSWORD_ERROR)
	}

	if _, err := s.providerRepo.FindByID(registerDTO.ProviderID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err)
			return custom_error.NewConflictError(constants.INVALID_PROVIDER_ID_ERROR)
		}
		return err
	}

	//employee exist
	if _, err := s.employeeRepo.FindByUsername(registerDTO.Username); err == nil {
		log.Println(err)
		return custom_error.NewConflictError(constants.USERNAME_DUPLICATE_ERROR)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(registerDTO.Password), 14)
	if err != nil {
		return err
	}

	employee := registerDTO.ToEntity(string(bytes))

	return s.employeeRepo.Create(employee)
}

func (s *authService) Login(loginDTO dto.LoginDTO) (entities.Employee, error) {
	employee, err := s.employeeRepo.FindByUsername(loginDTO.Username)
	if err != nil {
		log.Println("username")
		return employee, custom_error.NewUnauthorizedError(constants.CREDENTIAL_ERROR)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(employee.HashedPassword), []byte(loginDTO.Password)); err != nil {
		log.Println("password")
		return employee, custom_error.NewUnauthorizedError(constants.CREDENTIAL_ERROR)
	}

	employee.IssueAt = time.Now()
	if err := s.employeeRepo.Update(employee); err != nil {
		return entities.Employee{}, err
	}

	return employee, nil
}
