package test

import (
	"context"
	"testing"
	"time"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/Athlevo/Booking-Athlevo/storage/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func TestAccessRepo(t *testing.T) {
	db := createDBConnection(t)
	defer db.Close(context.Background())

	accessRepo := postgres.NewAccessRepo(db)

	gymID := createGym(t, db)
	defer deleteGym(t, db, gymID)

	coachID := uuid.New().String() // Replace with a valid coach ID if needed
	userID := uuid.New().String()  // Replace with a valid user ID if needed

	// Test cases for each access type (Personal, Group, Coach)
	testAccessPersonal(t, db, accessRepo, gymID, userID)
	testAccessGroup(t, db, accessRepo, gymID, coachID, userID)
	testAccessCoach(t, db, accessRepo, gymID, coachID, userID)
}

func testAccessPersonal(t *testing.T, db *pgx.Conn, accessRepo *postgres.AccessRepo, gymID, userID string) {
	subscriptionRepo := postgres.NewSubscriptionPersonalRepo(db)
	bookingRepo := postgres.NewBookingPersonalRepo(db)

	// 1. Create a Subscription Personal
	createSubscriptionReq := &booking.CreateSubscriptionPersonalRequest{
		SubscriptionPersonal: &booking.SubscriptionPersonal{
			GymId:       gymID,
			Type:        "Personal Training",
			Description: "One-on-one training sessions",
			Price:       10,
			Duration:    60, // In minutes
			Count:       10,
		},
	}
	createdSubscription, err := subscriptionRepo.CreateSubscriptionPersonal(context.Background(), createSubscriptionReq)
	assert.NoError(t, err)
	assert.NotNil(t, createdSubscription)
	defer deleteSubscriptionPersonal(t, db, createdSubscription.Id)

	// 2. Create a Booking Personal
	createBookingReq := &booking.CreateBookingPersonalRequest{
		BookingPersonal: &booking.BookingPersonal{
			UserId:         userID,
			SubscriptionId: createdSubscription.Id,
			Payment:        100,
			AccessStatus:   "granted", // Set to "granted" for access
			StartDate:      time.Now().Format(time.RFC3339),
			Count:          10,
		},
	}
	createdBooking, err := bookingRepo.CreateBookingPersonal(context.Background(), createBookingReq)
	assert.NoError(t, err)
	assert.NotNil(t, createdBooking)
	defer deleteBookingPersonal(t, db, createdBooking.Id)

	t.Run("CreateAccessPersonal", func(t *testing.T) {
		req := &booking.CreateAccessPersonalRequest{
			AccessPersonal: &booking.AccessPersonal{
				BookingPersonalId: createdBooking.Id,
				Date:              time.Now().Format(time.RFC3339),
			},
		}
		createdAccess, err := accessRepo.CreateAccessPersonal(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, createdAccess)
		assert.Equal(t, req.AccessPersonal.BookingPersonalId, createdAccess.BookingPersonalId)
		assert.Equal(t, req.AccessPersonal.Date, createdAccess.Date)

	})

	t.Run("ListAccessPersonal", func(t *testing.T) {
		req := &booking.ListAccessPersonalRequest{
			BookingPersonalId: createdBooking.Id,
		}
		listResponse, err := accessRepo.ListAccessPersonal(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, listResponse)
		assert.GreaterOrEqual(t, len(listResponse.AccessPersonal), 1) // At least 1 access record
	})
}

func testAccessGroup(t *testing.T, db *pgx.Conn, accessRepo *postgres.AccessRepo, gymID, coachID, userID string) {
	subscriptionRepo := postgres.NewSubscriptionGroupRepo(db)
	bookingRepo := postgres.NewBookingGroupRepo(db)

	// 1. Create a Subscription Group
	createSubscriptionReq := &booking.CreateSubscriptionGroupRequest{
		SubscriptionGroup: &booking.SubscriptionGroup{
			GymId:       gymID,
			CoachId:     coachID,
			Type:        "Group Fitness",
			Description: "High-intensity interval training",
			Price:       50,
			Capacity:    20,
			Time:        time.Now().Format(time.RFC3339),
			Duration:    60,
			Count:       10,
		},
	}
	createdSubscription, err := subscriptionRepo.CreateSubscriptionGroup(context.Background(), createSubscriptionReq)
	assert.NoError(t, err)
	assert.NotNil(t, createdSubscription)
	defer deleteSubscriptionGroup(t, db, createdSubscription.Id)

	// 2. Create a Booking Group
	createBookingReq := &booking.CreateBookingGroupRequest{
		BookingGroup: &booking.BookingGroup{
			UserId:         userID,
			SubscriptionId: createdSubscription.Id,
			Payment:        50,
			AccessStatus:   "granted", // Set to "granted" for access
			StartDate:      time.Now().Format(time.RFC3339),
			Count:          1,
		},
	}
	createdBooking, err := bookingRepo.CreateBookingGroup(context.Background(), createBookingReq)
	assert.NoError(t, err)
	assert.NotNil(t, createdBooking)
	defer deleteBookingGroup(t, db, createdBooking.Id)

	t.Run("CreateAccessGroup", func(t *testing.T) {
		req := &booking.CreateAccessGroupRequest{
			AccessGroup: &booking.AccessGroup{
				BookingGroupId: createdBooking.Id,
				Date:           time.Now().Format(time.RFC3339),
			},
		}
		createdAccess, err := accessRepo.CreateAccessGroup(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, createdAccess)
		assert.Equal(t, req.AccessGroup.BookingGroupId, createdAccess.BookingGroupId)
		assert.Equal(t, req.AccessGroup.Date, createdAccess.Date)
	})

	t.Run("ListAccessGroup", func(t *testing.T) {
		req := &booking.ListAccessGroupRequest{
			BookingGroupId: createdBooking.Id,
		}
		listResponse, err := accessRepo.ListAccessGroup(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, listResponse)
		assert.GreaterOrEqual(t, len(listResponse.AccessGroup), 1) // At least 1 access record
	})
}

func testAccessCoach(t *testing.T, db *pgx.Conn, accessRepo *postgres.AccessRepo, gymID, coachID, userID string) {
	subscriptionRepo := postgres.NewSubscriptionCoachRepo(db)
	bookingRepo := postgres.NewBookingCoachRepo(db)

	// 1. Create a Subscription Coach
	createSubscriptionReq := &booking.CreateSubscriptionCoachRequest{
		SubscriptionCoach: &booking.SubscriptionCoach{
			GymId:       gymID,
			CoachId:     coachID,
			Type:        "Personal Coaching",
			Description: "One-on-one coaching sessions",
			Price:       150,
			Duration:    60, // In minutes
		},
	}
	createdSubscription, err := subscriptionRepo.CreateSubscriptionCoach(context.Background(), createSubscriptionReq)
	assert.NoError(t, err)
	assert.NotNil(t, createdSubscription)
	defer deleteSubscriptionCoach(t, db, createdSubscription.Id)

	// 2. Create a Booking Coach
	createBookingReq := &booking.CreateBookingCoachRequest{
		BookingCoach: &booking.BookingCoach{
			UserId:         userID,
			SubscriptionId: createdSubscription.Id,
			Payment:        150,
			AccessStatus:   "granted", // Set to "granted" for access
			StartDate:      time.Now().Format(time.RFC3339),
			Count:          1,
		},
	}
	createdBooking, err := bookingRepo.CreateBookingCoach(context.Background(), createBookingReq)
	assert.NoError(t, err)
	assert.NotNil(t, createdBooking)
	defer deleteBookingCoach(t, db, createdBooking.Id)

	t.Run("CreateAccessCoach", func(t *testing.T) {
		req := &booking.CreateAccessCoachRequest{
			AccessCoach: &booking.AccessCoach{
				BookingCoachId: createdBooking.Id,
				Date:           time.Now().Format(time.RFC3339),
			},
		}
		createdAccess, err := accessRepo.CreateAccessCoach(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, createdAccess)

		assert.Equal(t, req.AccessCoach.BookingCoachId, createdAccess.BookingCoachId)
		assert.Equal(t, req.AccessCoach.Date, createdAccess.Date)
	})

	t.Run("ListAccessCoach", func(t *testing.T) {
		req := &booking.ListAccessCoachRequest{
			BookingCoachId: createdBooking.Id,
		}
		listResponse, err := accessRepo.ListAccessCoach(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, listResponse)
		assert.GreaterOrEqual(t, len(listResponse.AccessCoach), 1) // At least 1 access record
	})
}
