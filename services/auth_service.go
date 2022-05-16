package services

import (
	"errors"

	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/custom_error"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(registerDTO dto.RegisterDTO) error
	Login()
}

type authService struct {
	employeeRepo repositories.EmployeeRepo
	providerRepo repositories.ProviderRepo
}

func NewAuthService(employeeRepo repositories.EmployeeRepo, providerRepo repositories.ProviderRepo) AuthService {
	return &authService{
		employeeRepo: employeeRepo,
		providerRepo: providerRepo,
	}
}

func (s *authService) Register(registerDTO dto.RegisterDTO) error {
	if registerDTO.Password != registerDTO.ConfirmedPassword {
		return custom_error.NewBadRequestError(constants.CONFIRMED_PASSWORD_ERROR)
	}

	if _, err := s.providerRepo.FindProviderByID(registerDTO.ProviderID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return custom_error.NewConflictError(constants.INVALID_PROVIDER_ID_ERROR)
		}
		return err
	}

	//employee exist
	if _, err := s.employeeRepo.FindEmployeeByUsername(registerDTO.Username); err == nil {
		return custom_error.NewConflictError(constants.USERNAME_DUPLICATE_ERROR)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(registerDTO.Password), 14)
	if err != nil {
		return err
	}

	employee := registerDTO.ToEmployeeEntity()
	employee.HashedPassword = string(bytes)

	return s.employeeRepo.CreateEmployee(employee)
}

func (s *authService) Login() {

}
