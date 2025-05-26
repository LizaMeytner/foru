package core

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserExists      = errors.New("user already exists")
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInternal        = errors.New("internal server error")
)

type AuthCore struct {
	rdb *redis.Client
}

func NewAuthCore(rdb *redis.Client) *AuthCore {
	return &AuthCore{rdb: rdb}
}

func (c *AuthCore) Register(ctx context.Context, email, password string) (string, error) {
	// Проверка существования пользователя с обработкой ошибок Redis
	exists, err := c.rdb.Exists(ctx, "user:"+email).Result()
	if err != nil {
		return "", ErrInternal
	}
	if exists == 1 {
		return "", ErrUserExists
	}

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", ErrInternal
	}

	// Сохранение в Redis с TTL (24 часа)
	err = c.rdb.HSet(ctx, "user:"+email, map[string]interface{}{
		"email":    email,
		"password": string(hashedPassword),
	}).Err()
	if err != nil {
		return "", ErrInternal
	}

	// Устанавливаем срок жизни записи
	if err := c.rdb.Expire(ctx, "user:"+email, 24*time.Hour).Err(); err != nil {
		return "", ErrInternal
	}

	return c.GenerateToken(email)
}

func (c *AuthCore) Login(ctx context.Context, email, password string) (string, error) {
	// Получаем хеш пароля с проверкой ошибок
	storedHash, err := c.rdb.HGet(ctx, "user:"+email, "password").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrUserNotFound
		}
		return "", ErrInternal
	}

	// Сравнение паролей
	if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password)); err != nil {
		return "", ErrInvalidPassword
	}

	return c.GenerateToken(email)
}

func (c *AuthCore) GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your-secret-key")) // Замените на env-переменную
}
