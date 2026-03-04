package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type greeterServer struct {
	UnimplementedGreeterServer
}

func (g *greeterServer) SayHello(ctx context.Context, req *wrapperspb.StringValue) (*wrapperspb.StringValue, error) {
	name := req.GetValue()
	if name == "" {
		name = "world"
	}
	return &wrapperspb.StringValue{Value: fmt.Sprintf("Hello, %s!", name)}, nil
}

func (g *greeterServer) SayHelloStream(req *wrapperspb.StringValue, stream Greeter_SayHelloStreamServer) error {
	name := req.GetValue()
	if name == "" {
		name = "world"
	}
	for i := 1; i <= 3; i++ {
		msg := fmt.Sprintf("Hello, %s! (%d)", name, i)
		if err := stream.Send(&wrapperspb.StringValue{Value: msg}); err != nil {
			return err
		}
		time.Sleep(300 * time.Millisecond)
	}
	return nil
}

type httpResponse struct {
	Message string `json:"message"`
}

func main() {
	grpcAddr := ":9090"
	httpAddr := ":8080"

	grpcLis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("grpc listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	RegisterGreeterServer(grpcServer, &greeterServer{})
	reflection.Register(grpcServer)

	httpMux := http.NewServeMux()
	httpMux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, httpResponse{Message: "ok"}, http.StatusOK)
	})
	httpMux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, httpResponse{Message: "healthy"}, http.StatusOK)
	})
	httpMux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		if name == "" {
			name = "world"
		}
		writeJSON(w, httpResponse{Message: fmt.Sprintf("Hello, %s!", name)}, http.StatusOK)
	})
	httpMux.HandleFunc("POST /echo", func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			Message string `json:"message"`
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			writeJSON(w, httpResponse{Message: "invalid json"}, http.StatusBadRequest)
			return
		}
		if payload.Message == "" {
			payload.Message = "empty"
		}
		writeJSON(w, httpResponse{Message: payload.Message}, http.StatusOK)
	})

	httpServer := &http.Server{
		Addr:              httpAddr,
		Handler:           httpMux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	errCh := make(chan error, 2)

	go func() {
		log.Printf("gRPC listening on %s", grpcAddr)
		errCh <- grpcServer.Serve(grpcLis)
	}()

	go func() {
		log.Printf("HTTP listening on %s", httpAddr)
		errCh <- httpServer.ListenAndServe()
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigCh:
		log.Printf("shutdown on %s", sig)
	case err := <-errCh:
		log.Printf("server error: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grpcServer.GracefulStop()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("http shutdown: %v", err)
	}
}

func writeJSON(w http.ResponseWriter, body any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
