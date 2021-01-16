package main

import (
	"crypto/tls"
	"crypto/x509"
	"grpc-server/pb"
	"io/ioutil"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	// _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	// _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	// _ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	// _ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

func main() {
	// 自签证书
	// creds, err := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server_no_passwd.key")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// 双向证书
	cert, _ := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterEmployeeServiceServer(rpcServer, &employeeService{})

	lis, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(lis)

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	// 	fmt.Println(request)
	// 	rpcServer.ServeHTTP(writer, request)
	// })
	// httpServer := &http.Server{
	// 	Addr:    ":8081",
	// 	Handler: mux,
	// }
	// httpServer.ListenAndServeTLS("keys/server.crt", "keys/server_no_passwd.key")
}
