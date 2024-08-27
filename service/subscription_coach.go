package service

import (
	"context"
	"fmt"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/Athlevo/Booking-Athlevo/storage"
)

// SubscriptionCoachService implements the gRPC server for subscription coach-related operations.
type SubscriptionCoachService struct {
	storage storage.StorageI
	booking.UnimplementedSubscriptionCoachServiceServer
}

// NewSubscriptionCoachService creates a new SubscriptionCoachService instance.
func NewSubscriptionCoachService(storage storage.StorageI) *SubscriptionCoachService {
	return &SubscriptionCoachService{
		storage: storage,
	}
}

// CreateSubscriptionCoach handles the CreateSubscriptionCoach gRPC request.
func (s *SubscriptionCoachService) CreateSubscriptionCoach(ctx context.Context, req *booking.CreateSubscriptionCoachRequest) (*booking.SubscriptionCoach, error) {
	subscription, err := s.storage.SubscriptionCoach().CreateSubscriptionCoach(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create coach subscription: %w", err)
	}
	return subscription, nil
}

// GetSubscriptionCoach handles the GetSubscriptionCoach gRPC request.
func (s *SubscriptionCoachService) GetSubscriptionCoach(ctx context.Context, req *booking.GetSubscriptionCoachRequest) (*booking.SubscriptionCoach, error) {
	subscription, err := s.storage.SubscriptionCoach().GetSubscriptionCoach(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get coach subscription: %w", err)
	}
	return subscription, nil
}

// UpdateSubscriptionCoach handles the UpdateSubscriptionCoach gRPC request.
func (s *SubscriptionCoachService) UpdateSubscriptionCoach(ctx context.Context, req *booking.UpdateSubscriptionCoachRequest) (*booking.SubscriptionCoach, error) {
	subscription, err := s.storage.SubscriptionCoach().UpdateSubscriptionCoach(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update coach subscription: %w", err)
	}
	return subscription, nil
}

// DeleteSubscriptionCoach handles the DeleteSubscriptionCoach gRPC request.
func (s *SubscriptionCoachService) DeleteSubscriptionCoach(ctx context.Context, req *booking.DeleteSubscriptionCoachRequest) (*booking.Empty, error) {
	err := s.storage.SubscriptionCoach().DeleteSubscriptionCoach(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to delete coach subscription: %w", err)
	}
	return &booking.Empty{}, nil
}

// ListSubscriptionCoach handles the ListSubscriptionCoach gRPC request.
func (s *SubscriptionCoachService) ListSubscriptionCoach(ctx context.Context, req *booking.ListSubscriptionCoachRequest) (*booking.ListSubscriptionCoachResponse, error) {
	subscriptions, err := s.storage.SubscriptionCoach().ListSubscriptionCoach(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to list coach subscriptions: %w", err)
	}
	return subscriptions, nil
}
