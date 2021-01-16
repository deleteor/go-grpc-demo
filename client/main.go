package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"gprc-client/pb"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// 自签证书
	// creds, err := credentials.NewClientTLSFromFile("keys/server.crt", "thefools.ml")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// 双向证书

	cert, _ := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewEmployeeServiceClient(conn)
	res, err := client.GetEmployee(context.Background(), &pb.EmployeeRequest{Id: 1})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
