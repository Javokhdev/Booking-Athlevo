package service

import (
	"context"
	"fmt"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/Athlevo/Booking-Athlevo/storage"
)

// BookingGroupService implements the gRPC server for booking group-related operations.
type BookingGroupService struct {
	storage storage.StorageI
	booking.UnimplementedBookingGroupServiceServer
}

// NewBookingGroupService creates a new BookingGroupService instance.
func NewBookingGroupService(storage storage.StorageI) *BookingGroupService {
	return &BookingGroupService{
		storage: storage,
	}
}

// CreateBookingGroup handles the CreateBookingGroup gRPC request.
func (s *BookingGroupService) CreateBookingGroup(ctx context.Context, req *booking.CreateBookingGroupRequest) (*booking.BookingGroup, error) {
	booking, err := s.storage.BookingGroup().CreateBookingGroup(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create group booking: %w", err)
	}
	return booking, nil
}

// GetBookingGroup handles the GetBookingGroup gRPC request.
func (s *BookingGroupService) GetBookingGroup(ctx context.Context, req *booking.GetBookingGroupRequest) (*booking.BookingGroup, error) {
	booking, err := s.storage.BookingGroup().GetBookingGroup(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get group booking: %w", err)
	}
	return booking, nil
}

// UpdateBookingGroup handles the UpdateBookingGroup gRPC request.
func (s *BookingGroupService) UpdateBookingGroup(ctx context.Context, req *booking.UpdateBookingGroupRequest) (*booking.BookingGroup, error) {
	booking, err := s.storage.BookingGroup().UpdateBookingGroup(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update group booking: %w", err)
	}
	return booking, nil
}

// DeleteBookingGroup handles the DeleteBookingGroup gRPC request.
func (s *BookingGroupService) DeleteBookingGroup(ctx context.Context, req *booking.DeleteBookingGroupRequest) (*booking.Empty, error) {
	err := s.storage.BookingGroup().DeleteBookingGroup(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to delete group booking: %w", err)
	}
	return &booking.Empty{}, nil
}

// ListBookingGroup handles the ListBookingGroup gRPC request.
func (s *BookingGroupService) ListBookingGroup(ctx context.Context, req *booking.ListBookingGroupRequest) (*booking.ListBookingGroupResponse, error) {
	bookings, err := s.storage.BookingGroup().ListBookingGroup(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to list group bookings: %w", err)
	}
	return bookings, nil
}
