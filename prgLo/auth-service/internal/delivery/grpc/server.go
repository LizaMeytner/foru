package grpc

import (
	"context"
	"time"

	"github.com/LizaMeytner/prgLo/auth-service/internal/model"
	"github.com/LizaMeytner/prgLo/auth-service/internal/service"
	pb "github.com/LizaMeytner/prgLo/proto/auth"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	authService *service.AuthService
}

func NewServer(authService *service.AuthService) *Server {
	return &Server{
		authService: authService,
	}
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user, err := s.authService.Register(ctx, model.UserCreate{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Id:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := s.authService.Login(ctx, model.UserLogin{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Token: token,
	}, nil
}

func (s *Server) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	user, err := s.authService.ValidateToken(req.Token)
	if err != nil {
		return nil, err
	}

	return &pb.ValidateTokenResponse{
		Id:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}
