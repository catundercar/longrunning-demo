package longrunning

import (
	"context"
	"github.com/catundercar/longrunning-demo/api/google/longrunning"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	extLongrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"testing"
)

var grpcListen = ":8080"
var httpListen = ":8081"

func TestServiceRouter(t *testing.T) {
	go StartGRPC(t)
	go StartGateway(t)

	select {}
}

func StartGRPC(t *testing.T) {
	lis, err := net.Listen("tcp", grpcListen)
	if err != nil {
		t.Fatal("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	extLongrunning.RegisterOperationsServer(s, &server{})
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0", grpcListen)
	log.Fatal(s.Serve(lis))
}

func StartGateway(t *testing.T) {
	gwmux := runtime.NewServeMux()
	if err := longrunning.RegisterOperationsHandlerFromEndpoint(context.Background(), gwmux, "127.0.0.1"+grpcListen, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}); err != nil {
		t.Fatal(err)
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0", httpListen)
	log.Fatalln(http.ListenAndServe("0.0.0.0"+httpListen, http.HandlerFunc(func(w http.ResponseWriter, rq *http.Request) {
		log.Println("Input URL:", rq.URL.String())
		gwmux.ServeHTTP(w, rq)
	})))
}
