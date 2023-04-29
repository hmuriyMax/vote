package main

import (
	"flag"
	"fmt"
	"github.com/hmuriyMax/vote/internal/api"
	"github.com/hmuriyMax/vote/internal/pb/vote_service"
	"github.com/hmuriyMax/vote/internal/service/cypher"
	"github.com/hmuriyMax/vote/internal/service/vote"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = 8080

func main() {
	//ctx, cancelFunc := context.WithCancel(context.Background())
	//defer cancelFunc()

	cypherService, err := cypher.NewCypherService()
	if err != nil {
		log.Fatalf("error creating cypher service: %w", err)
	}

	voteService := vote.NewVoteService(
		cypherService,
	)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	vote_service.RegisterVoteServiceServer(grpcServer, api.NewImplementation(voteService))
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("error serving: %v", err)
	}

}
