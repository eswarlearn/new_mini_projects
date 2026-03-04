package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/eswar/grpc-doctor-booking/internal/config"
	"github.com/eswar/grpc-doctor-booking/internal/server"
	doctorv1 "github.com/eswar/grpc-doctor-booking/gen/go/doctor/v1"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func() { _ = logger.Sync() }()

	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("config load failed", zap.Error(err))
	}

	grpcAddr := fmt.Sprintf(":%s", cfg.GRPCPort)
	gwAddr := fmt.Sprintf(":%s", cfg.GatewayPort)

	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		logger.Fatal("grpc listen failed", zap.Error(err))
	}

	grpcServer := grpc.NewServer()
	doctorv1.RegisterDoctorServiceServer(grpcServer, server.NewDoctorServer())

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	go func() {
		logger.Info("gRPC server started", zap.String("addr", grpcAddr))
		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatal("grpc serve failed", zap.Error(err))
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	if err := doctorv1.RegisterDoctorServiceHandlerFromEndpoint(
		ctx,
		mux,
		grpcAddr,
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	); err != nil {
		logger.Fatal("gateway register failed", zap.Error(err))
	}

	gwServer := &http.Server{
		Addr:    gwAddr,
		Handler: mux,
	}

	go func() {
		logger.Info("gateway started", zap.String("addr", gwAddr))
		if err := gwServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("gateway serve failed", zap.Error(err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	logger.Info("shutdown initiated")

	grpcServer.GracefulStop()

	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()
	_ = gwServer.Shutdown(ctxShutdown)
}
