package service

import (
	"context"
	"fmt"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/Athlevo/Booking-Athlevo/storage"
)

// AccessService implements the gRPC server for access-related operations.
type AccessService struct {
	storage                                  storage.StorageI
	booking.UnimplementedAccessServiceServer // Embed the unimplemented server
}

// NewAccessService creates a new AccessService instance.
func NewAccessService(storage storage.StorageI) *AccessService {
	return &AccessService{
		storage: storage,
	}
}

// CreateAccessPersonal handles the CreateAccessPersonal gRPC request.
func (s *AccessService) CreateAccessPersonal(ctx context.Context, req *booking.CreateAccessPersonalRequest) (*booking.AccessPersonal, error) {
	access, err := s.storage.Access().CreateAccessPersonal(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create personal access record: %w", err)
	}
	return access, nil
}

// ListAccessPersonal handles the ListAccessPersonal gRPC request.
func (s *AccessService) ListAccessPersonal(ctx context.Context, req *booking.ListAccessPersonalRequest) (*booking.ListAccessPersonalResponse, error) {
	accesses, err := s.storage.Access().ListAccessPersonal(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to list personal access records: %w", err)
	}
	return accesses, nil
}

// CreateAccessGroup handles the CreateAccessGroup gRPC request.
func (s *AccessService) CreateAccessGroup(ctx context.Context, req *booking.CreateAccessGroupRequest) (*booking.AccessGroup, error) {
	access, err := s.storage.Access().CreateAccessGroup(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create group access record: %w", err)
	}
	return access, nil
}

// ListAccessGroup handles the ListAccessGroup gRPC request.
func (s *AccessService) ListAccessGroup(ctx context.Context, req *booking.ListAccessGroupRequest) (*booking.ListAccessGroupResponse, error) {
	accesses, err := s.storage.Access().ListAccessGroup(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to list group access records: %w", err)
	}
	return accesses, nil
}

// CreateAccessCoach handles the CreateAccessCoach gRPC request.
func (s *AccessService) CreateAccessCoach(ctx context.Context, req *booking.CreateAccessCoachRequest) (*booking.AccessCoach, error) {
	access, err := s.storage.Access().CreateAccessCoach(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create coach access record: %w", err)
	}
	return access, nil
}

// ListAccessCoach handles the ListAccessCoach gRPC request.
func (s *AccessService) ListAccessCoach(ctx context.Context, req *booking.ListAccessCoachRequest) (*booking.ListAccessCoachResponse, error) {
	accesses, err := s.storage.Access().ListAccessCoach(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to list coach access records: %w", err)
	}
	return accesses, nil
}
