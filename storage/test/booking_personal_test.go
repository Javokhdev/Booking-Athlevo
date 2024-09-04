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

func TestBookingPersonalRepo(t *testing.T) {
	db := createDBConnection(t)
	defer db.Close(context.Background())

	bookingRepo := postgres.NewBookingPersonalRepo(db)
	subscriptionRepo := postgres.NewSubscriptionPersonalRepo(db) // For creating subscriptions

	gymID := createGym(t, db)
	defer deleteGym(t, db, gymID)

	userID := uuid.New().String() // Replace with a valid user ID if needed

	t.Run("CreateBookingPersonal", func(t *testing.T) {
		// Create a subscription first
		createSubscriptionReq := &booking.CreateSubscriptionPersonalRequest{
			SubscriptionPersonal: &booking.SubscriptionPersonal{
				GymId:       gymID,
				Type:        "Personal Training",
				Description: "One-on-one training sessions",
				Price:       100,
				Duration:    60, // In minutes
				Count:       10,
			},
		}
		createdSubscription, err := subscriptionRepo.CreateSubscriptionPersonal(context.Background(), createSubscriptionReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription)

		req := &booking.CreateBookingPersonalRequest{
			BookingPersonal: &booking.BookingPersonal{
				UserId:         userID,
				SubscriptionId: createdSubscription.Id,
				Payment:        100,
				AccessStatus:   "pending",
				StartDate:      time.Now().Format(time.RFC3339),
				Count:          1,
			},
		}

		createdBooking, err := bookingRepo.CreateBookingPersonal(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, createdBooking)
		assert.NotEmpty(t, createdBooking.Id)
		assert.Equal(t, req.BookingPersonal.UserId, createdBooking.UserId)
		assert.Equal(t, req.BookingPersonal.SubscriptionId, createdBooking.SubscriptionId)
		// ... (Assert other fields)

		// Cleanup
		defer deleteBookingPersonal(t, db, createdBooking.Id)
		defer deleteSubscriptionPersonal(t, db, createdSubscription.Id)
	})

	t.Run("GetBookingPersonal", func(t *testing.T) {
		// Create a subscription first
		createSubscriptionReq := &booking.CreateSubscriptionPersonalRequest{
			SubscriptionPersonal: &booking.SubscriptionPersonal{
				GymId:       gymID,
				Type:        "Personal Training",
				Description: "One-on-one training sessions",
				Price:       100,
				Duration:    60, // In minutes
				Count:       10,
			},
		}
		createdSubscription, err := subscriptionRepo.CreateSubscriptionPersonal(context.Background(), createSubscriptionReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription)

		// Create a booking to retrieve
		createBookingReq := &booking.CreateBookingPersonalRequest{
			BookingPersonal: &booking.BookingPersonal{
				UserId:         userID,
				SubscriptionId: createdSubscription.Id,
				Payment:        100,
				AccessStatus:   "pending",
				StartDate:      time.Now().Format(time.RFC3339),
				Count:          1,
			},
		}
		createdBooking, err := bookingRepo.CreateBookingPersonal(context.Background(), createBookingReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdBooking)

		getReq := &booking.GetBookingPersonalRequest{
			Id: createdBooking.Id,
		}

		retrievedBooking, err := bookingRepo.GetBookingPersonal(context.Background(), getReq)
		assert.NoError(t, err)
		assert.NotNil(t, retrievedBooking)
		assert.Equal(t, createdBooking.Id, retrievedBooking.Id)
		assert.Equal(t, createdBooking.UserId, retrievedBooking.UserId)
		assert.Equal(t, createdBooking.SubscriptionId, retrievedBooking.SubscriptionId)
		// ... (Assert other fields)

		// Cleanup
		defer deleteBookingPersonal(t, db, createdBooking.Id)
		defer deleteSubscriptionPersonal(t, db, createdSubscription.Id)
	})

	t.Run("UpdateBookingPersonal", func(t *testing.T) {
		// Create a subscription first
		createSubscriptionReq := &booking.CreateSubscriptionPersonalRequest{
			SubscriptionPersonal: &booking.SubscriptionPersonal{
				GymId:       gymID,
				Type:        "Personal Training",
				Description: "One-on-one training sessions",
				Price:       100,
				Duration:    60, // In minutes
				Count:       10,
			},
		}
		createdSubscription, err := subscriptionRepo.CreateSubscriptionPersonal(context.Background(), createSubscriptionReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription)

		// Create a booking to update
		createBookingReq := &booking.CreateBookingPersonalRequest{
			BookingPersonal: &booking.BookingPersonal{
				UserId:         userID,
				SubscriptionId: createdSubscription.Id,
				Payment:        100,
				AccessStatus:   "pending",
				StartDate:      time.Now().Format(time.RFC3339),
				Count:          1,
			},
		}
		createdBooking, err := bookingRepo.CreateBookingPersonal(context.Background(), createBookingReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdBooking)

		updateReq := &booking.UpdateBookingPersonalRequest{
			BookingPersonal: &booking.BookingPersonal{
				Id:             createdBooking.Id,
				UserId:         userID,
				SubscriptionId: createdSubscription.Id,
				Payment:        150,
				AccessStatus:   "granted",
				StartDate:      time.Now().Add(time.Hour * 24).Format(time.RFC3339), // Updated start date
				Count:          2,
			},
		}

		updatedBooking, err := bookingRepo.UpdateBookingPersonal(context.Background(), updateReq)
		assert.NoError(t, err)
		assert.NotNil(t, updatedBooking)
		assert.Equal(t, updateReq.BookingPersonal.Id, updatedBooking.Id)
		assert.Equal(t, updateReq.BookingPersonal.Payment, updatedBooking.Payment)
		assert.Equal(t, updateReq.BookingPersonal.AccessStatus, updatedBooking.AccessStatus)
		assert.Equal(t, updateReq.BookingPersonal.StartDate, updatedBooking.StartDate)
		assert.Equal(t, updateReq.BookingPersonal.Count, updatedBooking.Count)

		// Cleanup
		defer deleteBookingPersonal(t, db, createdBooking.Id)
		defer deleteSubscriptionPersonal(t, db, createdSubscription.Id)
	})

	t.Run("DeleteBookingPersonal", func(t *testing.T) {
		// Create a subscription first
		createSubscriptionReq := &booking.CreateSubscriptionPersonalRequest{
			SubscriptionPersonal: &booking.SubscriptionPersonal{
				GymId:       gymID,
				Type:        "Personal Training",
				Description: "One-on-one training sessions",
				Price:       100,
				Duration:    60, // In minutes
				Count:       10,
			},
		}
		createdSubscription, err := subscriptionRepo.CreateSubscriptionPersonal(context.Background(), createSubscriptionReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription)

		// Create a booking to delete
		createBookingReq := &booking.CreateBookingPersonalRequest{
			BookingPersonal: &booking.BookingPersonal{
				UserId:         userID,
				SubscriptionId: createdSubscription.Id,
				Payment:        100,
				AccessStatus:   "pending",
				StartDate:      time.Now().Format(time.RFC3339),
				Count:          1,
			},
		}
		createdBooking, err := bookingRepo.CreateBookingPersonal(context.Background(), createBookingReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdBooking)

		deleteReq := &booking.DeleteBookingPersonalRequest{
			Id: createdBooking.Id,
		}

		err = bookingRepo.DeleteBookingPersonal(context.Background(), deleteReq)
		assert.NoError(t, err)

		// Try to get the deleted booking
		getReq := &booking.GetBookingPersonalRequest{
			Id: createdBooking.Id,
		}
		_, err = bookingRepo.GetBookingPersonal(context.Background(), getReq)
		assert.ErrorIs(t, err, pgx.ErrNoRows)

		// Cleanup
		defer deleteSubscriptionPersonal(t, db, createdSubscription.Id)
	})

	t.Run("ListBookingPersonal", func(t *testing.T) {
		// Create a subscription first
		createSubscriptionReq := &booking.CreateSubscriptionPersonalRequest{
			SubscriptionPersonal: &booking.SubscriptionPersonal{
				GymId:       gymID,
				Type:        "Personal Training",
				Description: "One-on-one training sessions",
				Price:       100,
				Duration:    60, // In minutes
				Count:       10,
			},
		}
		createdSubscription, err := subscriptionRepo.CreateSubscriptionPersonal(context.Background(), createSubscriptionReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription)

		// Create multiple bookings for testing filtering
		createBookingReq1 := &booking.CreateBookingPersonalRequest{
			BookingPersonal: &booking.BookingPersonal{
				UserId:         userID,
				SubscriptionId: createdSubscription.Id,
				Payment:        100,
				AccessStatus:   "pending",
				StartDate:      time.Now().Format(time.RFC3339),
				Count:          1,
			},
		}
		createdBooking1, err := bookingRepo.CreateBookingPersonal(context.Background(), createBookingReq1)
		assert.NoError(t, err)
		assert.NotNil(t, createdBooking1)

		createBookingReq2 := &booking.CreateBookingPersonalRequest{
			BookingPersonal: &booking.BookingPersonal{
				UserId:         uuid.New().String(), // Different user ID
				SubscriptionId: createdSubscription.Id,
				Payment:        150,
				AccessStatus:   "granted",
				StartDate:      time.Now().Add(time.Hour * 24).Format(time.RFC3339),
				Count:          2,
			},
		}
		createdBooking2, err := bookingRepo.CreateBookingPersonal(context.Background(), createBookingReq2)
		assert.NoError(t, err)
		assert.NotNil(t, createdBooking2)

		// Test listing all bookings
		listReq := &booking.ListBookingPersonalRequest{}
		listResponse, err := bookingRepo.ListBookingPersonal(context.Background(), listReq)
		assert.NoError(t, err)
		assert.NotNil(t, listResponse)
		assert.GreaterOrEqual(t, len(listResponse.BookingPersonal), 2) // At least 2 bookings

		// Test filtering by user ID
		listReq = &booking.ListBookingPersonalRequest{
			UserId: userID,
		}
		listResponse, err = bookingRepo.ListBookingPersonal(context.Background(), listReq)
		assert.NoError(t, err)
		assert.NotNil(t, listResponse)
		assert.LessOrEqual(t, 1, len(listResponse.BookingPersonal)) // Only 1 booking for this user

		// Cleanup
		defer deleteBookingPersonal(t, db, createdBooking1.Id)
		defer deleteBookingPersonal(t, db, createdBooking2.Id)
		defer deleteSubscriptionPersonal(t, db, createdSubscription.Id)
	})
}

func deleteBookingPersonal(t *testing.T, db *pgx.Conn, id string) {
	// _, err := db.Exec(context.Background(), "DELETE FROM booking_personal WHERE id = $1", id)
	// assert.NoError(t, err)
}
