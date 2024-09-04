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

func TestBookingCoachRepo(t *testing.T) {
	db := createDBConnection(t)
	defer db.Close(context.Background())

	bookingRepo := postgres.NewBookingCoachRepo(db)
	subscriptionRepo := postgres.NewSubscriptionCoachRepo(db) // For creating subscriptions

	gymID := createGym(t, db)
	defer deleteGym(t, db, gymID)

	coachID := uuid.New().String() // Replace with a valid coach ID if needed
	userID := uuid.New().String()  // Replace with a valid user ID if needed

	t.Run("CreateBookingCoach", func(t *testing.T) {
		// Create a subscription first
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

		req := &booking.CreateBookingCoachRequest{
			BookingCoach: &booking.BookingCoach{
				UserId:         userID,
				SubscriptionId: createdSubscription.Id,
				Payment:        150,
				AccessStatus:   "pending",
				StartDate:      time.Now().Format(time.RFC3339),
				Count:          1,
			},
		}

		createdBooking, err := bookingRepo.CreateBookingCoach(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, createdBooking)
		assert.NotEmpty(t, createdBooking.Id)
		assert.Equal(t, req.BookingCoach.UserId, createdBooking.UserId)
		assert.Equal(t, req.BookingCoach.SubscriptionId, createdBooking.SubscriptionId)
		// ... (Assert other fields)

		// Cleanup
		defer deleteBookingCoach(t, db, createdBooking.Id)
		defer deleteSubscriptionCoach(t, db, createdSubscription.Id)
	})

	t.Run("GetBookingCoach", func(t *testing.T) {
		// Create a subscription first
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

		// Create a booking to retrieve
		createBookingReq := &booking.CreateBookingCoachRequest{
			BookingCoach: &booking.BookingCoach{
				UserId:         userID,
				SubscriptionId: createdSubscription.Id,
				Payment:        150,
				AccessStatus:   "pending",
				StartDate:      time.Now().Format(time.RFC3339),
				Count:          1,
			},
		}
		createdBooking, err := bookingRepo.CreateBookingCoach(context.Background(), createBookingReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdBooking)

		getReq := &booking.GetBookingCoachRequest{
			Id: createdBooking.Id,
		}

		retrievedBooking, err := bookingRepo.GetBookingCoach(context.Background(), getReq)
		assert.NoError(t, err)
		assert.NotNil(t, retrievedBooking)
		assert.Equal(t, createdBooking.Id, retrievedBooking.Id)
		assert.Equal(t, createdBooking.UserId, retrievedBooking.UserId)
		assert.Equal(t, createdBooking.SubscriptionId, retrievedBooking.SubscriptionId)
		// ... (Assert other fields)

		// Cleanup
		defer deleteBookingCoach(t, db, createdBooking.Id)
		defer deleteSubscriptionCoach(t, db, createdSubscription.Id)
	})

	t.Run("UpdateBookingCoach", func(t *testing.T) {
		// Create a subscription first
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

		// Create a booking to update
		createBookingReq := &booking.CreateBookingCoachRequest{
			BookingCoach: &booking.BookingCoach{
				UserId:         userID,
				SubscriptionId: createdSubscription.Id,
				Payment:        150,
				AccessStatus:   "pending",
				StartDate:      time.Now().Format(time.RFC3339),
				Count:          1,
			},
		}
		createdBooking, err := bookingRepo.CreateBookingCoach(context.Background(), createBookingReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdBooking)

		updateReq := &booking.UpdateBookingCoachRequest{
			BookingCoach: &booking.BookingCoach{
				Id:             createdBooking.Id,
				UserId:         userID,
				SubscriptionId: createdSubscription.Id,
				Payment:        200,
				AccessStatus:   "granted",
				StartDate:      time.Now().Add(time.Hour * 24).Format(time.RFC3339), // Updated start date
				Count:          2,
			},
		}

		updatedBooking, err := bookingRepo.UpdateBookingCoach(context.Background(), updateReq)
		assert.NoError(t, err)
		assert.NotNil(t, updatedBooking)
		assert.Equal(t, updateReq.BookingCoach.Id, updatedBooking.Id)
		assert.Equal(t, updateReq.BookingCoach.Payment, updatedBooking.Payment)
		assert.Equal(t, updateReq.BookingCoach.AccessStatus, updatedBooking.AccessStatus)
		assert.Equal(t, updateReq.BookingCoach.StartDate, updatedBooking.StartDate)
		assert.Equal(t, updateReq.BookingCoach.Count, updatedBooking.Count)

		// Cleanup
		defer deleteBookingCoach(t, db, createdBooking.Id)
		defer deleteSubscriptionCoach(t, db, createdSubscription.Id)
	})

	t.Run("DeleteBookingCoach", func(t *testing.T) {
		// Create a subscription first
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

		// Create a booking to delete
		createBookingReq := &booking.CreateBookingCoachRequest{
			BookingCoach: &booking.BookingCoach{
				UserId:         userID,
				SubscriptionId: createdSubscription.Id,
				Payment:        150,
				AccessStatus:   "pending",
				StartDate:      time.Now().Format(time.RFC3339),
				Count:          1,
			},
		}
		createdBooking, err := bookingRepo.CreateBookingCoach(context.Background(), createBookingReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdBooking)

		deleteReq := &booking.DeleteBookingCoachRequest{
			Id: createdBooking.Id,
		}

		err = bookingRepo.DeleteBookingCoach(context.Background(), deleteReq)
		assert.NoError(t, err)

		// Try to get the deleted booking
		getReq := &booking.GetBookingCoachRequest{
			Id: createdBooking.Id,
		}
		_, err = bookingRepo.GetBookingCoach(context.Background(), getReq)
		assert.ErrorIs(t, err, pgx.ErrNoRows)

		// Cleanup
		defer deleteSubscriptionCoach(t, db, createdSubscription.Id)
	})

	t.Run("ListBookingCoach", func(t *testing.T) {
		// Create a subscription first
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

		// Create multiple bookings for testing filtering
		createBookingReq1 := &booking.CreateBookingCoachRequest{
			BookingCoach: &booking.BookingCoach{
				UserId:         userID,
				SubscriptionId: createdSubscription.Id,
				Payment:        150,
				AccessStatus:   "pending",
				StartDate:      time.Now().Format(time.RFC3339),
				Count:          1,
			},
		}
		createdBooking1, err := bookingRepo.CreateBookingCoach(context.Background(), createBookingReq1)
		assert.NoError(t, err)
		assert.NotNil(t, createdBooking1)

		createBookingReq2 := &booking.CreateBookingCoachRequest{
			BookingCoach: &booking.BookingCoach{
				UserId:         uuid.New().String(), // Different user ID
				SubscriptionId: createdSubscription.Id,
				Payment:        200,
				AccessStatus:   "granted",
				StartDate:      time.Now().Add(time.Hour * 24).Format(time.RFC3339),
				Count:          2,
			},
		}
		createdBooking2, err := bookingRepo.CreateBookingCoach(context.Background(), createBookingReq2)
		assert.NoError(t, err)
		assert.NotNil(t, createdBooking2)

		// Test listing all bookings
		listReq := &booking.ListBookingCoachRequest{}
		listResponse, err := bookingRepo.ListBookingCoach(context.Background(), listReq)
		assert.NoError(t, err)
		assert.NotNil(t, listResponse)
		assert.GreaterOrEqual(t, len(listResponse.BookingCoach), 2) // At least 2 bookings

		// Test filtering by user ID
		listReq = &booking.ListBookingCoachRequest{
			UserId: userID,
		}
		listResponse, err = bookingRepo.ListBookingCoach(context.Background(), listReq)
		assert.NoError(t, err)
		assert.NotNil(t, listResponse)
		assert.LessOrEqual(t, 1, len(listResponse.BookingCoach)) // Only 1 booking for this user

		// Test filtering by subscription ID
		listReq = &booking.ListBookingCoachRequest{
			SubscriptionId: createdSubscription.Id,
		}
		listResponse, err = bookingRepo.ListBookingCoach(context.Background(), listReq)
		assert.NoError(t, err)
		assert.NotNil(t, listResponse)
		assert.Equal(t, 2, len(listResponse.BookingCoach)) // 2 bookings for this subscription

		// Cleanup
		defer deleteBookingCoach(t, db, createdBooking1.Id)
		defer deleteBookingCoach(t, db, createdBooking2.Id)
		defer deleteSubscriptionCoach(t, db, createdSubscription.Id)
	})
}

func deleteBookingCoach(t *testing.T, db *pgx.Conn, id string) {
	// _, err := db.Exec(context.Background(), "DELETE FROM booking_coach WHERE id = $1", id)
	// assert.NoError(t, err)
}
