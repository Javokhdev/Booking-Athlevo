package service

import (
	"context"
	"fmt"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/Athlevo/Booking-Athlevo/storage"
)

// BookingPersonalService implements the gRPC server for booking personal-related operations.
type BookingPersonalService struct {
	storage storage.StorageI
	booking.UnimplementedBookingPersonalServiceServer
}

// NewBookingPersonalService creates a new BookingPersonalService instance.
func NewBookingPersonalService(storage storage.StorageI) *BookingPersonalService {
	return &BookingPersonalService{
		storage: storage,
	}
}

// CreateBookingPersonal handles the CreateBookingPersonal gRPC request.
func (s *BookingPersonalService) CreateBookingPersonal(ctx context.Context, req *booking.CreateBookingPersonalRequest) (*booking.BookingPersonal, error) {
	booking, err := s.storage.BookingPersonal().CreateBookingPersonal(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create personal booking: %w", err)
	}
	return booking, nil
}

// GetBookingPersonal handles the GetBookingPersonal gRPC request.
func (s *BookingPersonalService) GetBookingPersonal(ctx context.Context, req *booking.GetBookingPersonalRequest) (*booking.BookingPersonal, error) {
	booking, err := s.storage.BookingPersonal().GetBookingPersonal(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get personal booking: %w", err)
	}
	return booking, nil
}

// UpdateBookingPersonal handles the UpdateBookingPersonal gRPC request.
func (s *BookingPersonalService) UpdateBookingPersonal(ctx context.Context, req *booking.UpdateBookingPersonalRequest) (*booking.BookingPersonal, error) {
	booking, err := s.storage.BookingPersonal().UpdateBookingPersonal(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update personal booking: %w", err)
	}
	return booking, nil
}

// DeleteBookingPersonal handles the DeleteBookingPersonal gRPC request.
func (s *BookingPersonalService) DeleteBookingPersonal(ctx context.Context, req *booking.DeleteBookingPersonalRequest) (*booking.Empty, error) {
	err := s.storage.BookingPersonal().DeleteBookingPersonal(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to delete personal booking: %w", err)
	}
	return &booking.Empty{}, nil
}

// ListBookingPersonal handles the ListBookingPersonal gRPC request.
func (s *BookingPersonalService) ListBookingPersonal(ctx context.Context, req *booking.ListBookingPersonalRequest) (*booking.ListBookingPersonalResponse, error) {
	bookings, err := s.storage.BookingPersonal().ListBookingPersonal(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to list personal bookings: %w", err)
	}
	return bookings, nil
}
