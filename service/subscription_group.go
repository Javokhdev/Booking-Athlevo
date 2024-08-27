package service

import (
	"context"
	"fmt"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/Athlevo/Booking-Athlevo/storage"
)

// SubscriptionGroupService implements the gRPC server for subscription group-related operations.
type SubscriptionGroupService struct {
	storage storage.StorageI
	booking.UnimplementedSubscriptionGroupServiceServer
}

// NewSubscriptionGroupService creates a new SubscriptionGroupService instance.
func NewSubscriptionGroupService(storage storage.StorageI) *SubscriptionGroupService {
	return &SubscriptionGroupService{
		storage: storage,
	}
}

// CreateSubscriptionGroup handles the CreateSubscriptionGroup gRPC request.
func (s *SubscriptionGroupService) CreateSubscriptionGroup(ctx context.Context, req *booking.CreateSubscriptionGroupRequest) (*booking.SubscriptionGroup, error) {
	subscription, err := s.storage.SubscriptionGroup().CreateSubscriptionGroup(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create group subscription: %w", err)
	}
	return subscription, nil
}

// GetSubscriptionGroup handles the GetSubscriptionGroup gRPC request.
func (s *SubscriptionGroupService) GetSubscriptionGroup(ctx context.Context, req *booking.GetSubscriptionGroupRequest) (*booking.SubscriptionGroup, error) {
	subscription, err := s.storage.SubscriptionGroup().GetSubscriptionGroup(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get group subscription: %w", err)
	}
	return subscription, nil
}

// UpdateSubscriptionGroup handles the UpdateSubscriptionGroup gRPC request.
func (s *SubscriptionGroupService) UpdateSubscriptionGroup(ctx context.Context, req *booking.UpdateSubscriptionGroupRequest) (*booking.SubscriptionGroup, error) {
	subscription, err := s.storage.SubscriptionGroup().UpdateSubscriptionGroup(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update group subscription: %w", err)
	}
	return subscription, nil
}

// DeleteSubscriptionGroup handles the DeleteSubscriptionGroup gRPC request.
func (s *SubscriptionGroupService) DeleteSubscriptionGroup(ctx context.Context, req *booking.DeleteSubscriptionGroupRequest) (*booking.Empty, error) {
	err := s.storage.SubscriptionGroup().DeleteSubscriptionGroup(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to delete group subscription: %w", err)
	}
	return &booking.Empty{}, nil
}

// ListSubscriptionGroup handles the ListSubscriptionGroup gRPC request.
func (s *SubscriptionGroupService) ListSubscriptionGroup(ctx context.Context, req *booking.ListSubscriptionGroupRequest) (*booking.ListSubscriptionGroupResponse, error) {
	subscriptions, err := s.storage.SubscriptionGroup().ListSubscriptionGroup(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to list group subscriptions: %w", err)
	}
	return subscriptions, nil
}
