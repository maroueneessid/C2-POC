package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AuthFn(ctx context.Context) (context.Context, error) {
	method, ok := grpc.Method(ctx)
	if ok {
		if strings.HasPrefix(method, "/AssetService") {
			return ctx, nil
		}
	}

	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "missing auth token")
	}

	if token != "supersecretmanagertoken" {
		fmt.Println("[!] Unauthorized access to gRPC Method:", method)
		return nil, status.Error(codes.Unauthenticated, "invalid auth token")
	}

	return ctx, nil
}
