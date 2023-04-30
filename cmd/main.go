package main

import (
	"context"
	"github.com/hmuriyMax/vote/internal/api/vote-service"
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

	voteService := vote.NewVoteService(
		cypherService,
	)

	app := server.NewRestServer(vote_service.NewImplementation(voteService), "8080", "5001")
	err = app.Start(ctx)
	if err != nil {
		log.Panicln(err)
	}
}
