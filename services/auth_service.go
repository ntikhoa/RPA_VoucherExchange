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
	Register(registerDTO dto.RegisterDTO, roleID uint) error
	Login(loginDTO dto.LoginDTO, roleID uint) (entities.Account, error)
}

type authService struct {
	accountRepo  repositories.AccountRepo
	providerRepo repositories.ProviderRepo
}

func NewAuthService(accountRepo repositories.AccountRepo,
	providerRepo repositories.ProviderRepo) AuthService {
	return &authService{
		accountRepo:  accountRepo,
		providerRepo: providerRepo,
	}
}

func (s *authService) Register(registerDTO dto.RegisterDTO, roleID uint) error {
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

	//account exist
	if _, err := s.accountRepo.FindByUsername(registerDTO.Username); err == nil {
		log.Println(err)
		return custom_error.NewConflictError(constants.USERNAME_DUPLICATE_ERROR)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(registerDTO.Password), 14)
	if err != nil {
		return err
	}

	account := entities.NewAccount(registerDTO, string(bytes), roleID)

	return s.accountRepo.Create(account)
}

func (s *authService) Login(loginDTO dto.LoginDTO, roleID uint) (entities.Account, error) {
	account, err := s.accountRepo.FindByUsername(loginDTO.Username)
	if err != nil {
		log.Println("username")
		return account, custom_error.NewUnauthorizedError(constants.CREDENTIAL_ERROR)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.HashedPassword), []byte(loginDTO.Password)); err != nil {
		log.Println("password")
		return account, custom_error.NewUnauthorizedError(constants.CREDENTIAL_ERROR)
	}

	if account.RoleID != roleID {
		return account, custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}

	account.IssueAt = time.Now()
	if err := s.accountRepo.Update(account); err != nil {
		return entities.Account{}, err
	}

	return account, nil
}
