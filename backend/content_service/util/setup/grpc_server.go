package setup

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	"github.com/david-drvar/xws2021-nistagram/common/interceptors"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/controllers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"time"
)

func GRPCServer(db *gorm.DB) {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", grpc_common.Content_service_address)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	jwtManager := common.NewJWTManager("somesecretkey", 15 * time.Minute)
	authInterceptor := interceptors.NewAuthInterceptor(jwtManager)

	// Create a gRPC server object
	s := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.Unary()),
    )

	server, err := controllers.NewServer(db)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// Attach the Greeter service to the server
	protopb.RegisterContentServer(s, server)
	// Serve gRPC server
	log.Println("Serving gRPC on " + grpc_common.Content_service_address)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc_common.CreateGrpcConnection(grpc_common.Content_service_address)
	if err != nil {
		log.Fatalln(err) // TODO: Graceful shutdown
		return
	}

	gatewayMux := runtime.NewServeMux()
	// Register Greeter
	err = protopb.RegisterContentHandler(context.Background(), gatewayMux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    grpc_common.Content_gateway_address,
		Handler: tracer.TracingWrapper(gatewayMux),
	}

	log.Println("Serving gRPC-Gateway on " + grpc_common.Content_gateway_address)
	log.Fatalln(gwServer.ListenAndServe())
}
