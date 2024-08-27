package service

import (
	"context"
	"fmt"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/Athlevo/Booking-Athlevo/storage"
)

// BookingCoachService implements the gRPC server for access-related operations.
type BookingCoachService struct {
	storage storage.StorageI
	booking.UnimplementedBookingCoachServiceServer
}

// NewBookingCoachService creates a new BookingCoachService instance.
func NewBookingCoachService(storage storage.StorageI) *BookingCoachService {
	return &BookingCoachService{
		storage: storage,
	}
}

// CreateBookingCoach handles the CreateBookingCoach gRPC request.
func (s *BookingCoachService) CreateBookingCoach(ctx context.Context, req *booking.CreateBookingCoachRequest) (*booking.BookingCoach, error) {
	booking, err := s.storage.BookingCoach().CreateBookingCoach(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create coach booking: %w", err)
	}
	return booking, nil
}

// GetBookingCoach handles the GetBookingCoach gRPC request.
func (s *BookingCoachService) GetBookingCoach(ctx context.Context, req *booking.GetBookingCoachRequest) (*booking.BookingCoach, error) {
	booking, err := s.storage.BookingCoach().GetBookingCoach(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get coach booking: %w", err)
	}
	return booking, nil
}

// UpdateBookingCoach handles the UpdateBookingCoach gRPC request.
func (s *BookingCoachService) UpdateBookingCoach(ctx context.Context, req *booking.UpdateBookingCoachRequest) (*booking.BookingCoach, error) {
	booking, err := s.storage.BookingCoach().UpdateBookingCoach(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update coach booking: %w", err)
	}
	return booking, nil
}

// DeleteBookingCoach handles the DeleteBookingCoach gRPC request.
func (s *BookingCoachService) DeleteBookingCoach(ctx context.Context, req *booking.DeleteBookingCoachRequest) (*booking.Empty, error) {
	err := s.storage.BookingCoach().DeleteBookingCoach(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to delete coach booking: %w", err)
	}
	return &booking.Empty{}, nil
}

// ListBookingCoach handles the ListBookingCoach gRPC request.
func (s *BookingCoachService) ListBookingCoach(ctx context.Context, req *booking.ListBookingCoachRequest) (*booking.ListBookingCoachResponse, error) {
	bookings, err := s.storage.BookingCoach().ListBookingCoach(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to list coach bookings: %w", err)
	}
	return bookings, nil
}
