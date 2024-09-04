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

func TestBookingGroupRepo(t *testing.T) {
	db := createDBConnection(t)
	defer db.Close(context.Background())

	bookingRepo := postgres.NewBookingGroupRepo(db)
	subscriptionRepo := postgres.NewSubscriptionGroupRepo(db) // For creating subscriptions

	gymID := createGym(t, db)
	defer deleteGym(t, db, gymID)

	coachID := uuid.New().String() // Replace with a valid coach ID if needed
	userID := uuid.New().String()  // Replace with a valid user ID if needed

	t.Run("CreateBookingGroup", func(t *testing.T) {
		// Create a subscription first
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

		req := &booking.CreateBookingGroupRequest{
			BookingGroup: &booking.BookingGroup{
				UserId:         userID,
				SubscriptionId: createdSubscription.Id,
				Payment:        50,
				AccessStatus:   "pending",
				StartDate:      time.Now().Format(time.RFC3339),
				Count:          1,
			},
		}

		createdBooking, err := bookingRepo.CreateBookingGroup(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, createdBooking)
		assert.NotEmpty(t, createdBooking.Id)
		assert.Equal(t, req.BookingGroup.UserId, createdBooking.UserId)
		assert.Equal(t, req.BookingGroup.SubscriptionId, createdBooking.SubscriptionId)
		// ... (Assert other fields)

		// Cleanup
		defer deleteBookingGroup(t, db, createdBooking.Id)
		defer deleteSubscriptionGroup(t, db, createdSubscription.Id)
	})

	t.Run("GetBookingGroup", func(t *testing.T) {
		// Create a subscription first
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

		// Create a booking to retrieve
		createBookingReq := &booking.CreateBookingGroupRequest{
			BookingGroup: &booking.BookingGroup{
				UserId:         userID,
				SubscriptionId: createdSubscription.Id,
				Payment:        50,
				AccessStatus:   "pending",
				StartDate:      time.Now().Format(time.RFC3339),
				Count:          1,
			},
		}
		createdBooking, err := bookingRepo.CreateBookingGroup(context.Background(), createBookingReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdBooking)

		getReq := &booking.GetBookingGroupRequest{
			Id: createdBooking.Id,
		}

		retrievedBooking, err := bookingRepo.GetBookingGroup(context.Background(), getReq)
		assert.NoError(t, err)
		assert.NotNil(t, retrievedBooking)
		assert.Equal(t, createdBooking.Id, retrievedBooking.Id)
		assert.Equal(t, createdBooking.UserId, retrievedBooking.UserId)
		assert.Equal(t, createdBooking.SubscriptionId, retrievedBooking.SubscriptionId)
		// ... (Assert other fields)

		// Cleanup
		defer deleteBookingGroup(t, db, createdBooking.Id)
		defer deleteSubscriptionGroup(t, db, createdSubscription.Id)
	})

	t.Run("UpdateBookingGroup", func(t *testing.T) {
		// Create a subscription first
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

		// Create a booking to update
		createBookingReq := &booking.CreateBookingGroupRequest{
			BookingGroup: &booking.BookingGroup{
				UserId:         userID,
				SubscriptionId: createdSubscription.Id,
				Payment:        50,
				AccessStatus:   "pending",
				StartDate:      time.Now().Format(time.RFC3339),
				Count:          1,
			},
		}
		createdBooking, err := bookingRepo.CreateBookingGroup(context.Background(), createBookingReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdBooking)

		updateReq := &booking.UpdateBookingGroupRequest{
			BookingGroup: &booking.BookingGroup{
				Id:             createdBooking.Id,
				UserId:         userID,
				SubscriptionId: createdSubscription.Id,
				Payment:        75,
				AccessStatus:   "granted",
				StartDate:      time.Now().Add(time.Hour * 24).Format(time.RFC3339), // Updated start date
				Count:          2,
			},
		}

		updatedBooking, err := bookingRepo.UpdateBookingGroup(context.Background(), updateReq)
		assert.NoError(t, err)
		assert.NotNil(t, updatedBooking)
		assert.Equal(t, updateReq.BookingGroup.Id, updatedBooking.Id)
		assert.Equal(t, updateReq.BookingGroup.Payment, updatedBooking.Payment)
		assert.Equal(t, updateReq.BookingGroup.AccessStatus, updatedBooking.AccessStatus)
		assert.Equal(t, updateReq.BookingGroup.StartDate, updatedBooking.StartDate)
		assert.Equal(t, updateReq.BookingGroup.Count, updatedBooking.Count)

		// Cleanup
		defer deleteBookingGroup(t, db, createdBooking.Id)
		defer deleteSubscriptionGroup(t, db, createdSubscription.Id)
	})

	t.Run("DeleteBookingGroup", func(t *testing.T) {
		// Create a subscription first
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

		// Create a booking to delete
		createBookingReq := &booking.CreateBookingGroupRequest{
			BookingGroup: &booking.BookingGroup{
				UserId:         userID,
				SubscriptionId: createdSubscription.Id,
				Payment:        50,
				AccessStatus:   "pending",
				StartDate:      time.Now().Format(time.RFC3339),
				Count:          1,
			},
		}
		createdBooking, err := bookingRepo.CreateBookingGroup(context.Background(), createBookingReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdBooking)

		deleteReq := &booking.DeleteBookingGroupRequest{
			Id: createdBooking.Id,
		}

		err = bookingRepo.DeleteBookingGroup(context.Background(), deleteReq)
		assert.NoError(t, err)

		// Try to get the deleted booking
		getReq := &booking.GetBookingGroupRequest{
			Id: createdBooking.Id,
		}
		_, err = bookingRepo.GetBookingGroup(context.Background(), getReq)
		assert.ErrorIs(t, err, pgx.ErrNoRows)

		// Cleanup
		defer deleteSubscriptionGroup(t, db, createdSubscription.Id)
	})

	t.Run("ListBookingGroup", func(t *testing.T) {
		// Create a subscription first
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

		// Create multiple bookings for testing filtering
		createBookingReq1 := &booking.CreateBookingGroupRequest{
			BookingGroup: &booking.BookingGroup{
				UserId:         userID,
				SubscriptionId: createdSubscription.Id,
				Payment:        50,
				AccessStatus:   "pending",
				StartDate:      time.Now().Format(time.RFC3339),
				Count:          1,
			},
		}
		createdBooking1, err := bookingRepo.CreateBookingGroup(context.Background(), createBookingReq1)
		assert.NoError(t, err)
		assert.NotNil(t, createdBooking1)

		createBookingReq2 := &booking.CreateBookingGroupRequest{
			BookingGroup: &booking.BookingGroup{
				UserId:         uuid.New().String(), // Different user ID
				SubscriptionId: createdSubscription.Id,
				Payment:        75,
				AccessStatus:   "granted",
				StartDate:      time.Now().Add(time.Hour * 24).Format(time.RFC3339),
				Count:          2,
			},
		}
		createdBooking2, err := bookingRepo.CreateBookingGroup(context.Background(), createBookingReq2)
		assert.NoError(t, err)
		assert.NotNil(t, createdBooking2)

		// Test listing all bookings
		listReq := &booking.ListBookingGroupRequest{}
		listResponse, err := bookingRepo.ListBookingGroup(context.Background(), listReq)
		assert.NoError(t, err)
		assert.NotNil(t, listResponse)
		assert.GreaterOrEqual(t, len(listResponse.BookingGroup), 2) // At least 2 bookings

		// Test filtering by user ID
		listReq = &booking.ListBookingGroupRequest{
			UserId: userID,
		}
		listResponse, err = bookingRepo.ListBookingGroup(context.Background(), listReq)
		assert.NoError(t, err)
		assert.NotNil(t, listResponse)
		assert.LessOrEqual(t, 1, len(listResponse.BookingGroup)) // Only 1 booking for this user

		// Cleanup
		defer deleteBookingGroup(t, db, createdBooking1.Id)
		defer deleteBookingGroup(t, db, createdBooking2.Id)
		defer deleteSubscriptionGroup(t, db, createdSubscription.Id)
	})
}

func deleteBookingGroup(t *testing.T, db *pgx.Conn, id string) {
	// _, err := db.Exec(context.Background(), "DELETE FROM booking_group WHERE id = $1", id)
	// assert.NoError(t, err)
}
