package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Athlevo/Booking-Athlevo/config"
	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/Athlevo/Booking-Athlevo/service"
	"github.com/Athlevo/Booking-Athlevo/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	// Initialize PostgreSQL storage
	storage, err := postgres.NewPostgresStorage(cfg)
	if err != nil {
		log.Fatalf("failed to initialize storage: %v", err)
	}

	// Initialize gRPC server
	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Register booking services
	booking.RegisterBookingPersonalServiceServer(s, service.NewBookingPersonalService(storage))
	booking.RegisterBookingGroupServiceServer(s, service.NewBookingGroupService(storage))
	booking.RegisterBookingCoachServiceServer(s, service.NewBookingCoachService(storage))

	// Register subscription services
	booking.RegisterSubscriptionPersonalServiceServer(s, service.NewSubscriptionPersonalService(storage))
	booking.RegisterSubscriptionGroupServiceServer(s, service.NewSubscriptionGroupService(storage))
	booking.RegisterSubscriptionCoachServiceServer(s, service.NewSubscriptionCoachService(storage))

	// Register access service
	booking.RegisterAccessServiceServer(s, service.NewAccessService(storage))
	booking.RegisterAccessServiceBetaServer(s, service.NewAccessServiceBeta(storage))
	fmt.Println("gRPC server listening on", cfg.GRPCPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
