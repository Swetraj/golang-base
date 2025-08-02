package auth

import (
	"context"
	"errors"
	"github.com/Swetraj/golang-base/internal/domain/auth"
	"github.com/Swetraj/golang-base/internal/emails"
	"github.com/Swetraj/golang-base/internal/helpers"
	"github.com/Swetraj/golang-base/internal/validations"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
	"time"
)

func (u *userService) Register(ctx context.Context, email string) error {

	isUnique, err := validations.IsUniqueValue("users", "email", email)
	if err != nil {
		return err
	}

	if !isUnique {
		return errors.New("email already taken")
	}

	pwd, _ := helpers.RandomString(10)

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)

	user := &auth.User{
		Email:        email,
		PasswordHash: string(hashPassword),
		Provider:     auth.ProviderEmail,
	}

	err = u.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	token := helpers.GenerateResetToken()

	reset := &auth.VerificationToken{
		Token:     token,
		UserID:    user.ID,
		Used:      false,
		Type:      auth.VerificationEmail,
		ExpiresAt: time.Now().Add(30 * time.Minute),
	}

	err = u.tokenService.repo.Create(ctx, reset)
	if err != nil {
		return err
	}

	go u.SendEmail(user.Email, reset.Token)
	return nil
}

func (u *userService) Login(ctx context.Context, email string, password string) (*auth.User, error) {
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

func (u *userService) GetUserByEmail(ctx context.Context, email string) (*auth.User, error) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("email is incorrect")
	}

	return user, nil
}

func (u *userService) GetUserById(ctx context.Context, id uint) (*auth.User, error) {
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
	token, err := u.tokenService.repo.GetByToken(ctx, tokenString)
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

	return nil
}

func (u *userService) SendEmail(email string, token string) {
	var emailStruct struct {
		Subject string
		Message string
	}
	emailStruct.Subject = "Activate your account"
	resetLink := os.Getenv("FRONTEND_URL") + "/auth/signup?link=" + token
	emailStruct.Message = strings.ReplaceAll(emails.PasswordResetTemplate, "{{RESET_LINK}}", resetLink)

	if err := helpers.SendMail(email, emailStruct.Subject, emailStruct.Message); err != nil {
		log.Printf("Failed to send activation email to %s: %v", email, err)
	}
}
