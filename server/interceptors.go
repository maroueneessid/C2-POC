package main

import (
	"bytes"
	"context"
	"log"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AuthFn(ctx context.Context) (context.Context, error) {
	method, ok := grpc.Method(ctx)
	if ok && strings.HasPrefix(method, "/AssetService") {
		return ctx, nil
	}

	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "missing auth token")
	}

	found := false
	for _, opToken := range OpConf.Operators {
		if bytes.Equal([]byte(token), []byte(opToken)) {
			found = true
			break
		}
	}

	if !found {
		log.Println("[!] Unauthorized access to gRPC Method:", method)
		return nil, status.Error(codes.Unauthenticated, "invalid auth token")
	}

	return ctx, nil
}
