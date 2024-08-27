package service

import (
	"context"
	"fmt"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/Athlevo/Booking-Athlevo/storage"
)

// SubscriptionPersonalService implements the gRPC server for subscription personal-related operations.
type SubscriptionPersonalService struct {
	storage storage.StorageI
	booking.UnimplementedSubscriptionPersonalServiceServer
}

// NewSubscriptionPersonalService creates a new SubscriptionPersonalService instance.
func NewSubscriptionPersonalService(storage storage.StorageI) *SubscriptionPersonalService {
	return &SubscriptionPersonalService{
		storage: storage,
	}
}

// CreateSubscriptionPersonal handles the CreateSubscriptionPersonal gRPC request.
func (s *SubscriptionPersonalService) CreateSubscriptionPersonal(ctx context.Context, req *booking.CreateSubscriptionPersonalRequest) (*booking.SubscriptionPersonal, error) {
	subscription, err := s.storage.SubscriptionPersonal().CreateSubscriptionPersonal(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create personal subscription: %w", err)
	}
	return subscription, nil
}

// GetSubscriptionPersonal handles the GetSubscriptionPersonal gRPC request.
func (s *SubscriptionPersonalService) GetSubscriptionPersonal(ctx context.Context, req *booking.GetSubscriptionPersonalRequest) (*booking.SubscriptionPersonal, error) {
	subscription, err := s.storage.SubscriptionPersonal().GetSubscriptionPersonal(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get personal subscription: %w", err)
	}
	return subscription, nil
}

// UpdateSubscriptionPersonal handles the UpdateSubscriptionPersonal gRPC request.
func (s *SubscriptionPersonalService) UpdateSubscriptionPersonal(ctx context.Context, req *booking.UpdateSubscriptionPersonalRequest) (*booking.SubscriptionPersonal, error) {
	subscription, err := s.storage.SubscriptionPersonal().UpdateSubscriptionPersonal(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update personal subscription: %w", err)
	}
	return subscription, nil
}

// DeleteSubscriptionPersonal handles the DeleteSubscriptionPersonal gRPC request.
func (s *SubscriptionPersonalService) DeleteSubscriptionPersonal(ctx context.Context, req *booking.DeleteSubscriptionPersonalRequest) (*booking.Empty, error) {
	err := s.storage.SubscriptionPersonal().DeleteSubscriptionPersonal(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to delete personal subscription: %w", err)
	}
	return &booking.Empty{}, nil
}

// ListSubscriptionPersonal handles the ListSubscriptionPersonal gRPC request.
func (s *SubscriptionPersonalService) ListSubscriptionPersonal(ctx context.Context, req *booking.ListSubscriptionPersonalRequest) (*booking.ListSubscriptionPersonalResponse, error) {
	subscriptions, err := s.storage.SubscriptionPersonal().ListSubscriptionPersonal(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to list personal subscriptions: %w", err)
	}
	return subscriptions, nil
}
