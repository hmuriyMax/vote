package main

import (
	"context"
	"github.com/hmuriyMax/vote/internal/api/vote-service"
	"github.com/hmuriyMax/vote/internal/repo/db"
	"github.com/hmuriyMax/vote/internal/service/auth"
	"github.com/hmuriyMax/vote/internal/service/cypher"
	"github.com/hmuriyMax/vote/internal/service/server"
	"github.com/hmuriyMax/vote/internal/service/vote"
	"log"
)

const port = 8080

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	cypherService, err := cypher.NewCypherService()
	if err != nil {
		log.Fatalf("error creating cypher service: %w", err)
	}

	repo := db.NewPostgres(nil)

	authService := auth.NewAuthService(repo)

	voteService := vote.NewVoteService(
		authService,
		repo,
	)

	app := server.NewRestServer(vote_service.NewImplementation(
		voteService,
		authService,
		cypherService,
	), "7001", "5300")
	err = app.Start(ctx)
	if err != nil {
		log.Panicln(err)
	}
}
