package rpcserver

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type Server struct{}

//LogursWrapper implemented the gRPC LoggerV2 interface
type LogursWrapper struct {
	logrus.Entry
	verboseModeDisabled bool
}

// V reports whether verbosity level l is at least the requested verbose level.
func (l *LogursWrapper) V(i int) bool {
	return !l.verboseModeDisabled
}

func NewGrpcServer() *grpc.Server {
	// Logrus entry is used, allowing pre-definition of certain fields by the user.
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []grpc_logrus.Option{}
	// Make sure that log statements internal to gRPC library are logged using the logrus Logger as well.
	//grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	logrusEntry.Logger.Level = logrus.InfoLevel
	grpclog.SetLoggerV2(&LogursWrapper{
		Entry:               *logrusEntry,
		verboseModeDisabled: true,
	})

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry, opts...),
		),
	)
	return grpcServer
}
