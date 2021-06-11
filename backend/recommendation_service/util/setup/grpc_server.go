package setup

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/recommendation_service/controllers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func GRPCServer(driver neo4j.Driver) {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", grpc_common.Recommendation_service_address)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	server, err := controllers.NewServer(driver)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// Attach the Greeter service to the server
	protopb.RegisterFollowersServer(s, server)
	// Serve gRPC server
	log.Println("Serving gRPC on " + grpc_common.Recommendation_service_address)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc_common.CreateGrpcConnection(grpc_common.Recommendation_service_address)
	if err != nil {
		log.Fatalln(err) // TODO: Graceful shutdown
		return
	}

	gatewayMux := runtime.NewServeMux()
	// Register Greeter
	err = protopb.RegisterFollowersHandler(context.Background(), gatewayMux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	c := common.SetupCors()
	pool := x509.NewCertPool()

	// Here is the certificate provided by the loading client, preferably the root certificate provided by the client.
	addTrust(pool,"./../common/sslFile/gateway.p12")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	gwServer := &http.Server{
		Addr:    grpc_common.Recommendation_gateway_address,
		Handler: tracer.TracingWrapper(c.Handler(gatewayMux)),
/*		TLSConfig: &tls.Config{
			ClientCAs: pool,
			ClientAuth:  tls.RequireAndVerifyClientCert,
		},*/
	}

	log.Println("Serving gRPC-Gateway on " + grpc_common.Recommendation_gateway_address)
	log.Fatalln(gwServer.ListenAndServeTLS("./../common/sslFile/gateway.crt", "./../common/sslFile/gateway.key"))
}

func addTrust(pool*x509.CertPool, path string) {
	aCrt, err := ioutil.ReadFile(path)
	if err!= nil {
		fmt.Println("ReadFile err:",err)
		return
	}
	pool.AppendCertsFromPEM(aCrt)

}
