package customcontext

import (
	"context"
)

type CorrelationIdKey struct{}

// Extracts the correlation ID from the context if exists or return an empty string
func ExtractCorrelationId(ctx context.Context) string {
	if correlationId, ok := ctx.Value(CorrelationIdKey{}).(string); ok {
		return correlationId
	}
	return ""
}

func WithCorrelationId(ctx context.Context, correlationId string) context.Context {
	return context.WithValue(ctx, CorrelationIdKey{}, correlationId)
}
