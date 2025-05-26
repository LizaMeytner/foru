package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"auth-service/internal/model"
	"auth-service/internal/repository"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidToken       = errors.New("invalid token")
)

type AuthUseCase struct {
	repo          repository.UserRepository
	secretKey     []byte
	tokenExpiry   time.Duration
	refreshExpiry time.Duration
}

type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func NewAuthUseCase(repo repository.UserRepository, secretKey string, tokenExpiry, refreshExpiry time.Duration) *AuthUseCase {
	return &AuthUseCase{
		repo:          repo,
		secretKey:     []byte(secretKey),
		tokenExpiry:   tokenExpiry,
		refreshExpiry: refreshExpiry,
	}
}

func (uc *AuthUseCase) Register(ctx context.Context, username, password, email string) (*model.User, error) {
	// Проверяем, существует ли пользователь с таким username или email
	userByUsername, _ := uc.repo.GetByUsername(ctx, username)
	userByEmail, _ := uc.repo.GetByEmail(ctx, email)
	if userByUsername != nil || userByEmail != nil {
		return nil, ErrUserAlreadyExists
	}

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
		Email:        email,
		Role:         "user",
	}

	err = uc.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *AuthUseCase) Login(ctx context.Context, username, password string) (string, string, error) {
	user, err := uc.repo.GetByUsername(ctx, username)
	if err != nil || user == nil {
		return "", "", ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", "", ErrInvalidCredentials
	}

	accessToken, refreshToken, err := uc.generateTokens(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (uc *AuthUseCase) RefreshToken(ctx context.Context, refreshToken string) (string, string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return uc.secretKey, nil
	})

	if err != nil || !token.Valid {
		return "", "", ErrInvalidToken
	}

	user, err := uc.repo.GetByID(ctx, claims.UserID)
	if err != nil || user == nil {
		return "", "", ErrInvalidToken
	}

	accessToken, newRefreshToken, err := uc.generateTokens(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, newRefreshToken, nil
}

func (uc *AuthUseCase) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return uc.secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

func (uc *AuthUseCase) generateTokens(user *model.User) (string, string, error) {
	// Generate access token
	accessClaims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(uc.tokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(uc.secretKey)
	if err != nil {
		return "", "", err
	}

	// Generate refresh token
	refreshClaims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(uc.refreshExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(uc.secretKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}
