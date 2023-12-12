package grpc

import (
	"context"

	"google.golang.org/grpc"
)

type streamWrapper struct {
	grpc.ServerStream
	ctx context.Context
}

func newStreamWrapper(stream grpc.ServerStream) *streamWrapper {
	ctx := stream.Context()
	return &streamWrapper{
		stream,
		ctx,
	}
}

func (w *streamWrapper) Context() context.Context {
	return w.ctx
}

func (w *streamWrapper) SetContext(ctx context.Context) {
	w.ctx = ctx
}
