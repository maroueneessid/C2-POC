package utils

import (
	"crypto/x509"
	_ "embed"
	"fmt"

	"google.golang.org/grpc/credentials"
)

var (
	server_crt = "utils/cert/server/server.crt"
	server_key = "utils/cert/server/server.key"
)

//go:embed cert/server/server.crt
var cert []byte

func SimpleClientTLS() (credentials.TransportCredentials, error) {

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		return nil, fmt.Errorf("failed to add certificate to cert pool")
	}

	creds := credentials.NewClientTLSFromCert(certPool, "localhost")

	return creds, nil
}

func SimpleServerTLS() (credentials.TransportCredentials, error) {
	creds, err := credentials.NewServerTLSFromFile(server_crt, server_key)

	if err != nil {
		return nil, fmt.Errorf("failed to load server certificate or key: %v", err)
	}

	return creds, nil
}
