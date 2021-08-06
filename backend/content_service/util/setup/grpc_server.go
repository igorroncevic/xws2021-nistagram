package setup

import (
	"context"
	"crypto/tls"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/igorroncevic/xws2021-nistagram/common"
	"github.com/igorroncevic/xws2021-nistagram/common/grpc_common"
	"github.com/igorroncevic/xws2021-nistagram/common/interceptors"
	"github.com/igorroncevic/xws2021-nistagram/common/logger"
	protopb "github.com/igorroncevic/xws2021-nistagram/common/proto"
	"github.com/igorroncevic/xws2021-nistagram/common/tracer"
	"github.com/igorroncevic/xws2021-nistagram/content_service/controllers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"time"
)

func GRPCServer(db *gorm.DB, customLogger *logger.Logger) {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", grpc_common.Content_service_address)
	if err != nil {
		customLogger.ToStdoutAndFile("Content GRPC Server", "Couldn't listen to "+grpc_common.Content_service_address, logger.Fatal)
		return
	}

	jwtManager := common.NewJWTManager("somesecretkey", 60*time.Minute)
	rbacInterceptor := interceptors.NewRBACInterceptor(db, jwtManager, customLogger)

	// Create a gRPC server object
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			rbacInterceptor.Authorize(),
			grpc_prometheus.UnaryServerInterceptor),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.MaxSendMsgSize(4<<30), // Default: 1024 * 1024 * 4 = 4MB -> Override to 4GBs
		grpc.MaxRecvMsgSize(4<<30), // Default: 1024 * 1024 * 4 = 4MB -> Override to 4GBs
	)

	server, err := controllers.NewServer(db, jwtManager, customLogger)
	if err != nil {
		customLogger.ToStdoutAndFile("Content GRPC Server", "Failed to create server", logger.Fatal)
		return
	}

	protopb.RegisterContentServer(s, server)
	customLogger.ToStdoutAndFile("Content GRPC Server", "Serving gRPC on "+grpc_common.Content_service_address, logger.Info)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc_common.CreateGrpcConnection(grpc_common.Content_service_address)
	if err != nil {
		// TODO: Graceful shutdown
		customLogger.ToStdoutAndFile("Content GRPC Server", "Couldn't serve gRPC on "+grpc_common.Content_service_address, logger.Fatal)
		return
	}

	gatewayMux := runtime.NewServeMux()
	err = protopb.RegisterContentHandler(context.Background(), gatewayMux, conn)
	if err != nil {
		customLogger.ToStdoutAndFile("Content GRPC Server", "Failed to register gateway", logger.Fatal)
	}

	grpc_prometheus.Register(s)
	gatewayMux.HandlePath("GET", "/metrics", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		promhttp.Handler().ServeHTTP(w, r)
	})

	c := common.SetupCors()
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	gwServer := &http.Server{
		Addr:    grpc_common.Content_gateway_address,
		Handler: tracer.TracingWrapper(c.Handler(gatewayMux)),
	}

	customLogger.ToStdoutAndFile("Content GRPC Server", "Serving gRPC-Gateway on "+grpc_common.Content_gateway_address, logger.Info)
	//log.Fatalln(gwServer.ListenAndServeTLS("./../common/sslFile/gateway.crt", "./../common/sslFile/gateway.key"))
	log.Fatalln(gwServer.ListenAndServe())
}
