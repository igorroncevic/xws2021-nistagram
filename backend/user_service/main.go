package main

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/controllers"
	userspb "github.com/david-drvar/xws2021-nistagram/user_service/proto"
	"github.com/david-drvar/xws2021-nistagram/user_service/util/setup"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	otgo "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"io"
	"log"
	"net"
	"net/http"
)

func main() {
	db := common.InitDatabase()
	err := setup.FillDatabase(db)
	if err != nil {
		panic("Cannot setup database tables. Error message: " + err.Error())
	}

	SetupGRPCServer(db)
}

type server struct {
	userspb.UnimplementedUsersServer
	controller *controllers.UserGrpcController
	tracer otgo.Tracer
	closer io.Closer
}

func NewServer(db *gorm.DB) (*server, error) {
	controller, _ := controllers.NewUserController(db)
	tracer, closer := tracer.Init("userService")
	otgo.SetGlobalTracer(tracer)
	return &server{
		controller: controller,
		tracer: tracer,
		closer: closer,
	}, nil
}

func SetupGRPCServer(db *gorm.DB){
	grpcServerPort := "localhost:8091"
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", grpcServerPort)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	server, err := NewServer(db)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// Attach the Greeter service to the server
	userspb.RegisterUsersServer(s, server)
	// Serve gRPC server
	log.Println("Serving gRPC on " + grpcServerPort)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		grpcServerPort,
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(
				grpc_opentracing.WithTracer(otgo.GlobalTracer()),
			),
		),
		grpc.WithStreamInterceptor(
			grpc_opentracing.StreamClientInterceptor(
				grpc_opentracing.WithTracer(otgo.GlobalTracer()),
			),
		),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = userspb.RegisterUsersHandler(context.Background(), gwmux, conn)
	 if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServerPort := "localhost:8001"
	gwServer := &http.Server{
		Addr:    gwServerPort,
		Handler: tracer.TracingWrapper(gwmux),
	}

	log.Println("Serving gRPC-Gateway on " + gwServerPort)
	log.Fatalln(gwServer.ListenAndServe())
}
