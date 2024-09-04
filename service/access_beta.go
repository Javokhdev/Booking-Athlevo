package service

import (
	"context"
	"fmt"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/Athlevo/Booking-Athlevo/storage"
)

// ... (other service structs and methods) ...

// AccessServiceBeta implements the gRPC server for access beta-related operations.
type AccessServiceBeta struct {
	storage storage.StorageI
	booking.UnimplementedAccessServiceBetaServer
}

// NewAccessServiceBeta creates a new AccessServiceBeta instance.
func NewAccessServiceBeta(storage storage.StorageI) *AccessServiceBeta {
	return &AccessServiceBeta{
		storage: storage,
	}
}

// CheckUserAccess handles the CheckUserAccess gRPC request for Access Beta Personal.
func (s *AccessServiceBeta) CheckUserAccess(ctx context.Context, req *booking.AccessBetaPersonalRequest) (*booking.AccessBetaPersonalResponse, error) {
	response, err := s.storage.AccessBeta().CheckUserAccess(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to check user access: %w", err)
	}
	return response, nil
}
