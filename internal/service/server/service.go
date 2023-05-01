package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	auth "github.com/hmuriyMax/vote/internal/pb/auth_service"
	vote "github.com/hmuriyMax/vote/internal/pb/vote_service"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

// RestServer реализует сервер REST для сервиса голосования
type RestServer struct {
	httpServer  *http.Server
	grpcServer  *grpc.Server
	grpcPort    string
	voteService vote.VoteServiceServer
	authService auth.AuthServiceServer
}

// NewRestServer отлично подходит для создания RestServer
func NewRestServer(voteService vote.VoteServiceServer, httpPort string, grpcPort string) RestServer {
	router := gin.Default()
	rs := RestServer{
		httpServer: &http.Server{
			Addr:    ":" + httpPort,
			Handler: router,
		},
		grpcServer:  grpc.NewServer(),
		grpcPort:    grpcPort,
		voteService: voteService,
	}

	vote.RegisterVoteServiceServer(rs.grpcServer, voteService)

	// Регистрация маршрутов
	router.POST("/vote", rs.vote)
	router.POST("/log", rs.login)
	router.POST("/reg", rs.register)

	return rs
}

// Start запускает сервер
func (s RestServer) Start(ctx context.Context) error {
	localCtx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	errChan := make(chan error, 1)
	go func(ctx context.Context) {
		err := s.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			errChan <- fmt.Errorf("failed to start httpServer: %w", err)
		}
	}(localCtx)

	go func(ctx context.Context) {
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", s.grpcPort))
		if err != nil {
			errChan <- fmt.Errorf("failed to listen: %w", err)
		}

		err = s.grpcServer.Serve(lis)
		if err != nil {
			errChan <- fmt.Errorf("failed to start grpcServer: %w", err)
		}
	}(localCtx)
	log.Printf("started server at %s", s.httpServer.Addr)
	select {
	case <-ctx.Done():
		err := s.httpServer.Shutdown(context.Background())
		if err != nil {
			return err
		}
		s.grpcServer.GracefulStop()
		log.Panicln("server gracefully stopped")
	case err := <-errChan:
		return fmt.Errorf("failed to start: %w", err)
	}
	return nil
}
