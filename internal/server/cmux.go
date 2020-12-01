package server

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/pronuu/lincoln/api/vote"

	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

func Cmux() {
	// Create the main listener.
	l, err := net.Listen("tcp", ":23456")
	if err != nil {
		log.Fatal(err)
	}

	// Create a cmux.
	m := cmux.New(l)

	// Match connections in order:
	// First grpc, then HTTP, and otherwise Go RPC/TCP.
	grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())
	trpcL := m.Match(cmux.Any()) // Any means anything that is not yet matched.

	// Create your protocol servers.
	grpcS := grpc.NewServer()

	vote.RegisterVoteServiceServer(grpcS)

	httpS := &http.Server{
		Handler: &helloHTTP1Handler{},
	}

	trpcS := rpc.NewServer()
	trpcS.Register(&ExampleRPCRcvr{})

	// Use the muxed listeners for your servers.
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)
	go trpcS.Accept(trpcL)

	// Start serving!
	m.Serve()
}
