package middleware

import (
	"context"
	"gemini/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// validateAPIKey checks the context for the API key and validates it
func validateAPIKey(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "missing metadata")
	}

	values, ok := md["api-key"]
	if !ok || len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "API key is not provided")
	}

	apiKey := values[0]
	if apiKey != util.Config.ApiKey {
		return status.Errorf(codes.Unauthenticated, "invalid API key")
	}

	return nil
}

func ApiKeyAuthenticationInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if err := validateAPIKey(ss.Context()); err != nil {
		return err
	}

	// Proceed with the handler if the API key is valid
	return handler(srv, ss)
}
