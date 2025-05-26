package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/LizaMeytner/foru/forum-service/config"
	grpcDelivery "github.com/LizaMeytner/foru/forum-service/internal/delivery/grpc"
	"github.com/LizaMeytner/foru/forum-service/internal/repository/postgres"
	"github.com/LizaMeytner/foru/forum-service/internal/usecase"
	"github.com/LizaMeytner/foru/forum-service/proto"
	"google.golang.org/grpc"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Инициализация подключения к базе данных
	db, err := postgres.NewPostgresDB(cfg.Postgres)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Инициализация репозиториев
	postRepo := postgres.NewPostRepository(db)
	commentRepo := postgres.NewCommentRepository(db)
	chatRepo := postgres.NewChatRepository(db)

	// Инициализация сценариев использования
	postUseCase := usecase.NewPostUseCase(postRepo)
	commentUseCase := usecase.NewCommentUseCase(commentRepo)
	chatUseCase := usecase.NewChatUseCase(chatRepo)

	// Инициализация gRPC сервера
	grpcServer := grpc.NewServer()
	forumService := grpcDelivery.NewForumService(postUseCase, commentUseCase, chatUseCase)
	proto.RegisterForumServiceServer(grpcServer, forumService)

	// Запуск gRPC сервера
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}

	go func() {
		log.Printf("Запуск gRPC сервера на порту %d", cfg.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Ошибка работы сервера: %v", err)
		}
	}()

	// Ожидание сигнала прерывания
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Корректное завершение работы
	log.Println("Завершение работы сервера...")
	grpcServer.GracefulStop()
}
