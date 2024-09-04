package storage

import (
	"context"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
)

// StorageI defines the interface for interacting with the storage layer.
type StorageI interface {
	BookingPersonal() BookingPersonalRepoI
	BookingGroup() BookingGroupRepoI
	BookingCoach() BookingCoachRepoI

	SubscriptionPersonal() SubscriptionPersonalRepoI
	SubscriptionGroup() SubscriptionGroupRepoI
	SubscriptionCoach() SubscriptionCoachRepoI

	Access() AccessRepoI

	AccessBeta() AccessRepoBetaI
}

// BookingPersonalRepoI defines methods for interacting with personal bookings.
type BookingPersonalRepoI interface {
	CreateBookingPersonal(ctx context.Context, req *booking.CreateBookingPersonalRequest) (*booking.BookingPersonal, error)
	GetBookingPersonal(ctx context.Context, req *booking.GetBookingPersonalRequest) (*booking.BookingPersonal, error)
	UpdateBookingPersonal(ctx context.Context, req *booking.UpdateBookingPersonalRequest) (*booking.BookingPersonal, error)
	DeleteBookingPersonal(ctx context.Context, req *booking.DeleteBookingPersonalRequest) error
	ListBookingPersonal(ctx context.Context, req *booking.ListBookingPersonalRequest) (*booking.ListBookingPersonalResponse, error)
}

// BookingGroupRepoI defines methods for interacting with group bookings.
type BookingGroupRepoI interface {
	CreateBookingGroup(ctx context.Context, req *booking.CreateBookingGroupRequest) (*booking.BookingGroup, error)
	GetBookingGroup(ctx context.Context, req *booking.GetBookingGroupRequest) (*booking.BookingGroup, error)
	UpdateBookingGroup(ctx context.Context, req *booking.UpdateBookingGroupRequest) (*booking.BookingGroup, error)
	DeleteBookingGroup(ctx context.Context, req *booking.DeleteBookingGroupRequest) error
	ListBookingGroup(ctx context.Context, req *booking.ListBookingGroupRequest) (*booking.ListBookingGroupResponse, error)
}

// BookingCoachRepoI defines methods for interacting with coach bookings.
type BookingCoachRepoI interface {
	CreateBookingCoach(ctx context.Context, req *booking.CreateBookingCoachRequest) (*booking.BookingCoach, error)
	GetBookingCoach(ctx context.Context, req *booking.GetBookingCoachRequest) (*booking.BookingCoach, error)
	UpdateBookingCoach(ctx context.Context, req *booking.UpdateBookingCoachRequest) (*booking.BookingCoach, error)
	DeleteBookingCoach(ctx context.Context, req *booking.DeleteBookingCoachRequest) error
	ListBookingCoach(ctx context.Context, req *booking.ListBookingCoachRequest) (*booking.ListBookingCoachResponse, error)
}

// SubscriptionPersonalRepoI defines methods for interacting with personal subscriptions.
type SubscriptionPersonalRepoI interface {
	CreateSubscriptionPersonal(ctx context.Context, req *booking.CreateSubscriptionPersonalRequest) (*booking.SubscriptionPersonal, error)
	GetSubscriptionPersonal(ctx context.Context, req *booking.GetSubscriptionPersonalRequest) (*booking.SubscriptionPersonal, error)
	UpdateSubscriptionPersonal(ctx context.Context, req *booking.UpdateSubscriptionPersonalRequest) (*booking.SubscriptionPersonal, error)
	DeleteSubscriptionPersonal(ctx context.Context, req *booking.DeleteSubscriptionPersonalRequest) error
	ListSubscriptionPersonal(ctx context.Context, req *booking.ListSubscriptionPersonalRequest) (*booking.ListSubscriptionPersonalResponse, error)
}

// SubscriptionGroupRepoI defines methods for interacting with group subscriptions.
type SubscriptionGroupRepoI interface {
	CreateSubscriptionGroup(ctx context.Context, req *booking.CreateSubscriptionGroupRequest) (*booking.SubscriptionGroup, error)
	GetSubscriptionGroup(ctx context.Context, req *booking.GetSubscriptionGroupRequest) (*booking.SubscriptionGroup, error)
	UpdateSubscriptionGroup(ctx context.Context, req *booking.UpdateSubscriptionGroupRequest) (*booking.SubscriptionGroup, error)
	DeleteSubscriptionGroup(ctx context.Context, req *booking.DeleteSubscriptionGroupRequest) error
	ListSubscriptionGroup(ctx context.Context, req *booking.ListSubscriptionGroupRequest) (*booking.ListSubscriptionGroupResponse, error)
}

// SubscriptionCoachRepoI defines methods for interacting with coach subscriptions.
type SubscriptionCoachRepoI interface {
	CreateSubscriptionCoach(ctx context.Context, req *booking.CreateSubscriptionCoachRequest) (*booking.SubscriptionCoach, error)
	GetSubscriptionCoach(ctx context.Context, req *booking.GetSubscriptionCoachRequest) (*booking.SubscriptionCoach, error)
	UpdateSubscriptionCoach(ctx context.Context, req *booking.UpdateSubscriptionCoachRequest) (*booking.SubscriptionCoach, error)
	DeleteSubscriptionCoach(ctx context.Context, req *booking.DeleteSubscriptionCoachRequest) error
	ListSubscriptionCoach(ctx context.Context, req *booking.ListSubscriptionCoachRequest) (*booking.ListSubscriptionCoachResponse, error)
}

// AccessRepoI defines methods for interacting with access records.
type AccessRepoI interface {
	CreateAccessPersonal(ctx context.Context, req *booking.CreateAccessPersonalRequest) (*booking.AccessPersonal, error)
	ListAccessPersonal(ctx context.Context, req *booking.ListAccessPersonalRequest) (*booking.ListAccessPersonalResponse, error)

	CreateAccessGroup(ctx context.Context, req *booking.CreateAccessGroupRequest) (*booking.AccessGroup, error)
	ListAccessGroup(ctx context.Context, req *booking.ListAccessGroupRequest) (*booking.ListAccessGroupResponse, error)

	CreateAccessCoach(ctx context.Context, req *booking.CreateAccessCoachRequest) (*booking.AccessCoach, error)
	ListAccessCoach(ctx context.Context, req *booking.ListAccessCoachRequest) (*booking.ListAccessCoachResponse, error)
}

type AccessRepoBetaI interface {
	CheckUserAccess(ctx context.Context, req *booking.AccessBetaPersonalRequest) (*booking.AccessBetaPersonalResponse, error)
}
