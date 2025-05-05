package service

import (
	"bytes"
	"context"
	"errors"
	"text/template"
	"time"

	"github.com/elingsuryo/movie-app/config"
	"github.com/elingsuryo/movie-app/internal/entity"
	"github.com/elingsuryo/movie-app/internal/http/dto"
	"github.com/elingsuryo/movie-app/internal/repository"
	"github.com/elingsuryo/movie-app/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type UserService interface {
	Login(ctx context.Context, username string, password string) (*entity.JwtCustomClaims, error)
	Register(ctx context.Context, req dto.UserRegisterRequest) error
	GetAll(ctx context.Context) ([]entity.User, error)
	GetByID(ctx context.Context, id int64) (*entity.User, error)
	Create(ctx context.Context, req dto.UserCreateRequest) error
	Update(ctx context.Context, req dto.UserUpdateRequest) error
	Delete(ctx context.Context, user *entity.User) error
	ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) error
	// RequestResetPassword(ctx context.Context, username string) error
	VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) error
}

// GenerateAccessToken implements TokenService.
func (u userService) GenerateAccessToken(ctx context.Context, claims entity.JwtCustomClaims) (string, error) {
	panic("unimplemented")
}

type userService struct {
	cfg            *config.Config
	userRepository repository.UserRepository
}

func NewUserService(cfg *config.Config, userRepository repository.UserRepository) UserService {
	return &userService{cfg, userRepository}
}

func (s *userService) Login(ctx context.Context, username string, password string) (*entity.JwtCustomClaims, error) {
	user, err := s.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("username atau password salah")
	}

	if user.IsVerified == 0 {
		return nil, errors.New("email belum diverifikasi")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("username atau password salah")
	}

	// if user.Password != password {
	// 	return nil, errors.New("username atau password salah")
	// }

	expiredTime := time.Now().Add(time.Minute * 10)

	claims := &entity.JwtCustomClaims{
		Username: user.Username,
		FUllName: user.FullName,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "movie-app",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	return claims, nil
}

func (s *userService) Register(ctx context.Context, req dto.UserRegisterRequest) error {
	user := new(entity.User)
	user.Username = req.Username
	user.FullName = req.FullName
	user.Role = "Administrator"
	user.ResetPasswordToken = utils.RandomString(20)
	user.VerifyEmailToken = utils.RandomString(20)

	exist, err := s.userRepository.GetByUsername(ctx, req.Username)
	if err == nil && exist != nil {
		return errors.New("username sudah terdaftar")
	}

templatePath := "./template/email/verify-email.html"
tmpl, err := template.ParseFiles(templatePath)
if err != nil {
	return err
}

var replacerEmail = struct {
	Token string
}{
	Token: user.VerifyEmailToken,
}

var body bytes.Buffer
if err := tmpl.Execute(&body, &replacerEmail); err != nil {
	return err
}


m := gomail.NewMessage()
m.SetHeader("From", s.cfg.SMTPConfig.Username)
m.SetHeader("To", user.Username)
m.SetHeader("Subject", "Reset Password Request !")
m.SetBody("text/html", body.String())

d := gomail.NewDialer(
	s.cfg.SMTPConfig.Host,
	int(s.cfg.SMTPConfig.Port),
	s.cfg.SMTPConfig.Username,
	s.cfg.SMTPConfig.Password,
)

if err := d.DialAndSend(m); err != nil {
	return err
}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return s.userRepository.Create(ctx, user)
}
 
func (s userService)GetAll(ctx context.Context) ([]entity.User, error) {
	return s.userRepository.GetAll(ctx)
}

func (s userService) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	return s.userRepository.GetByID(ctx, id)
}

func (s userService) Create(ctx context.Context, req dto.UserCreateRequest) error{
	User := &entity.User{
		Username: req.Username,
		Password: req.Password,	
		FullName: req.FullName,
		Role: req.Role,
	}
	exist, err := s.userRepository.GetByUsername(ctx, req.Username)
	if err == nil && exist != nil {
		return errors.New("username sudah terdaftar")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	User.Password = string(hashedPassword)

	return s.userRepository.Create(ctx, User)
}


func (s userService) Update(ctx context.Context, req dto.UserUpdateRequest) error {
	user, err := s.userRepository.GetByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.FullName != "" {
		user.FullName = req.FullName
	}
	if req.Role != "" {
		user.Role = req.Role
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	return s.userRepository.Update(ctx, user)
}	

func (s userService) Delete(ctx context.Context, user *entity.User) error{
	return s.userRepository.Delete(ctx, user)
}

func (s userService) ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) error {
	user, err := s.userRepository.GetByResetPasswordToken(ctx, req.Token)
	if err != nil {
		return errors.New("token anda salah")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.userRepository.Update(ctx, user)
 
}

func (s *userService) RequestResetPassword(ctx context.Context, username string) error {
	user, err := s.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return errors.New("username tidak ditemukan")
	}
	
templatePath := "./template/email/reset-password.html"
tmpl, err := template.ParseFiles(templatePath)
if err != nil {
	return err
}

var replacerEmail = struct {
	Token string
}{
	Token: user.ResetPasswordToken,
}

var body bytes.Buffer
if err := tmpl.Execute(&body, &replacerEmail); err != nil {
	return err
}


m := gomail.NewMessage()
m.SetHeader("From", s.cfg.SMTPConfig.Username)
m.SetHeader("To", user.Username)
m.SetHeader("Subject", "Reset Password Request !")
m.SetBody("text/html", body.String())

d := gomail.NewDialer(
	s.cfg.SMTPConfig.Host,
	int(s.cfg.SMTPConfig.Port),
	s.cfg.SMTPConfig.Username,
	s.cfg.SMTPConfig.Password,
)

if err := d.DialAndSend(m); err != nil {
	return err
}
return nil
}

func (s *userService) VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) error {
	user, err := s.userRepository.GetByVerifyEmailToken(ctx, req.Token)
	if err != nil {
		return errors.New("token anda salah")
	}

	user.IsVerified = 1
	return s.userRepository.Update(ctx, user)
}