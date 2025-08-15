package services

import (
	"context"
	"errors"
	"github.com/Swetraj/golang-base/internal/domain/constants"
	"github.com/Swetraj/golang-base/internal/domain/model"
	"github.com/Swetraj/golang-base/internal/domain/repository"
	"github.com/Swetraj/golang-base/internal/domain/service"
	"github.com/Swetraj/golang-base/internal/helpers"
	"github.com/Swetraj/golang-base/internal/pkg/emails"
	"github.com/Swetraj/golang-base/internal/pkg/validations"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
	"time"
)

type tokenService struct {
	repo repository.VerificationTokenRepository
}

type userService struct {
	db           *gorm.DB
	repo         repository.UserRepository
	tokenService repository.VerificationTokenRepository
}

func NewUserService(
	db *gorm.DB,
	repo repository.UserRepository, tokenService repository.VerificationTokenRepository,
) service.UserService {
	return &userService{db, repo, tokenService}
}

func NewTokenService(repo repository.VerificationTokenRepository) service.VerificationService {
	return &tokenService{repo}
}

func (u *userService) Register(ctx context.Context, email string) error {
	tx := u.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Ensure rollback on panic or error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	isUnique, err := validations.IsUniqueValue("users", "email", email)
	if err != nil {
		return err
	}

	if !isUnique {
		return errors.New("email already taken")
	}

	pwd, _ := helpers.RandomString(10)

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)

	user := &model.User{
		Email:        email,
		PasswordHash: string(hashPassword),
		Provider:     constants.ProviderEmail,
	}

	err = u.repo.CreateWithTx(ctx, tx, user)
	if err != nil {
		tx.Rollback()
		return err
	}

	token := helpers.GenerateResetToken()

	reset := &model.VerificationToken{
		Token:     token,
		UserID:    user.ID,
		Used:      false,
		Type:      constants.VerificationEmail,
		ExpiresAt: time.Now().Add(30 * time.Minute),
	}

	err = u.tokenService.CreateWithTx(ctx, tx, reset)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	go u.SendEmail(user.Email, reset.Token)
	return nil
}

func (u *userService) Login(ctx context.Context, email string, password string) (*model.User, error) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("email or password is incorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("email or password is incorrect")
	}

	return user, nil
}

func (u *userService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("email is incorrect")
	}

	return user, nil
}

func (u *userService) GetUserById(ctx context.Context, id uint) (*model.User, error) {
	user, err := u.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (u *userService) ResetPassword(ctx context.Context, tokenString string, pwd string) error {
	token, err := u.tokenService.GetByToken(ctx, tokenString)
	if err != nil {
		return err
	}

	user, err := u.GetUserById(ctx, token.UserID)
	if err != nil {
		return err
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	user.PasswordHash = string(hashPassword)
	err = u.repo.Update(ctx, user)
	if err != nil {
		return err
	}

	token.Used = true
	err = u.tokenService.Update(ctx, token)
	if err != nil {
		return err
	}

	return nil
}
func (t tokenService) UpdateToken(ctx context.Context, token *model.VerificationToken) error {
	err := t.repo.Update(ctx, token)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) SendEmail(email string, token string) {
	var emailStruct struct {
		Subject string
		Message string
	}
	emailStruct.Subject = "Activate your account"
	resetLink := os.Getenv("FRONTEND_URL") + "/api/auth/reset?link=" + token
	emailStruct.Message = strings.ReplaceAll(emails.PasswordResetTemplate, "{{RESET_LINK}}", resetLink)
	emailStruct.Message = strings.ReplaceAll(emailStruct.Message, "{{COMPANY_NAME}}", os.Getenv("COMPANY_NAME"))
	emailStruct.Message = strings.ReplaceAll(emailStruct.Message, "{{EMAIL}}", email)

	if err := helpers.SendMail(email, emailStruct.Subject, emailStruct.Message); err != nil {
		log.Printf("Failed to send activation email to %s: %v", email, err)
	}
}
