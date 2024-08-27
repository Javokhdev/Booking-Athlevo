package test

import (
	"context"
	"testing"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/Athlevo/Booking-Athlevo/storage/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func TestSubscriptionCoachRepo(t *testing.T) {
	db := createDBConnection(t)
	defer db.Close(context.Background())

	subscriptionRepo := postgres.NewSubscriptionCoachRepo(db)

	gymID := createGym(t, db)
	defer deleteGym(t, db, gymID)

	coachID := uuid.New().String() // Replace with a valid coach ID if needed

	t.Run("CreateSubscriptionCoach", func(t *testing.T) {
		req := &booking.CreateSubscriptionCoachRequest{
			SubscriptionCoach: &booking.SubscriptionCoach{
				GymId:       gymID,
				CoachId:     coachID,
				Type:        "Personal Coaching",
				Description: "One-on-one coaching sessions",
				Price:       150,
				Duration:    60, // In minutes
			},
		}

		createdSubscription, err := subscriptionRepo.CreateSubscriptionCoach(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription)
		assert.NotEmpty(t, createdSubscription.Id)
		assert.Equal(t, req.SubscriptionCoach.GymId, createdSubscription.GymId)
		assert.Equal(t, req.SubscriptionCoach.CoachId, createdSubscription.CoachId)
		assert.Equal(t, req.SubscriptionCoach.Type, createdSubscription.Type)
		// ... (Assert other fields)

		// Cleanup
		defer deleteSubscriptionCoach(t, db, createdSubscription.Id)
	})

	t.Run("GetSubscriptionCoach", func(t *testing.T) {
		// Create a subscription to retrieve
		createReq := &booking.CreateSubscriptionCoachRequest{
			SubscriptionCoach: &booking.SubscriptionCoach{
				GymId:       gymID,
				CoachId:     coachID,
				Type:        "Personal Coaching",
				Description: "One-on-one coaching sessions",
				Price:       150,
				Duration:    60, // In minutes
			},
		}
		createdSubscription, err := subscriptionRepo.CreateSubscriptionCoach(context.Background(), createReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription)

		getReq := &booking.GetSubscriptionCoachRequest{
			Id: createdSubscription.Id,
		}

		retrievedSubscription, err := subscriptionRepo.GetSubscriptionCoach(context.Background(), getReq)
		assert.NoError(t, err)
		assert.NotNil(t, retrievedSubscription)
		assert.Equal(t, createdSubscription.Id, retrievedSubscription.Id)
		assert.Equal(t, createdSubscription.GymId, retrievedSubscription.GymId)
		assert.Equal(t, createdSubscription.CoachId, retrievedSubscription.CoachId)
		// ... (Assert other fields)

		// Cleanup
		defer deleteSubscriptionCoach(t, db, createdSubscription.Id)
	})

	t.Run("UpdateSubscriptionCoach", func(t *testing.T) {
		createReq := &booking.CreateSubscriptionCoachRequest{
			SubscriptionCoach: &booking.SubscriptionCoach{
				GymId:       gymID,
				CoachId:     coachID,
				Type:        "Personal Coaching",
				Description: "One-on-one coaching sessions",
				Price:       150,
				Duration:    60, // In minutes
			},
		}
		createdSubscription, err := subscriptionRepo.CreateSubscriptionCoach(context.Background(), createReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription)

		updateReq := &booking.UpdateSubscriptionCoachRequest{
			SubscriptionCoach: &booking.SubscriptionCoach{
				Id:          createdSubscription.Id,
				GymId:       gymID,
				CoachId:     coachID,
				Type:        "Updated Coaching",
				Description: "Updated description",
				Price:       200,
				Duration:    90, // In minutes
			},
		}

		updatedSubscription, err := subscriptionRepo.UpdateSubscriptionCoach(context.Background(), updateReq)
		assert.NoError(t, err)
		assert.NotNil(t, updatedSubscription)
		assert.Equal(t, updateReq.SubscriptionCoach.Id, updatedSubscription.Id)
		assert.Equal(t, updateReq.SubscriptionCoach.Type, updatedSubscription.Type)
		assert.Equal(t, updateReq.SubscriptionCoach.Description, updatedSubscription.Description)
		assert.Equal(t, updateReq.SubscriptionCoach.Price, updatedSubscription.Price)
		assert.Equal(t, updateReq.SubscriptionCoach.Duration, updatedSubscription.Duration)

		// Cleanup
		defer deleteSubscriptionCoach(t, db, createdSubscription.Id)
	})

	t.Run("DeleteSubscriptionCoach", func(t *testing.T) {
		createReq := &booking.CreateSubscriptionCoachRequest{
			SubscriptionCoach: &booking.SubscriptionCoach{
				GymId:       gymID,
				CoachId:     coachID,
				Type:        "Personal Coaching",
				Description: "One-on-one coaching sessions",
				Price:       150,
				Duration:    60, // In minutes
			},
		}
		createdSubscription, err := subscriptionRepo.CreateSubscriptionCoach(context.Background(), createReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription)

		deleteReq := &booking.DeleteSubscriptionCoachRequest{
			Id: createdSubscription.Id,
		}

		err = subscriptionRepo.DeleteSubscriptionCoach(context.Background(), deleteReq)
		assert.NoError(t, err)

		// Try to get the deleted subscription
		getReq := &booking.GetSubscriptionCoachRequest{
			Id: createdSubscription.Id,
		}
		_, err = subscriptionRepo.GetSubscriptionCoach(context.Background(), getReq)
		assert.ErrorIs(t, err, pgx.ErrNoRows)
	})

	t.Run("ListSubscriptionCoach", func(t *testing.T) {
		// Create multiple subscriptions for the same gym
		createReq1 := &booking.CreateSubscriptionCoachRequest{
			SubscriptionCoach: &booking.SubscriptionCoach{
				GymId:       gymID,
				CoachId:     coachID,
				Type:        "Personal Coaching 1",
				Description: "Description 1",
				Price:       150,
				Duration:    60,
			},
		}
		createdSubscription1, err := subscriptionRepo.CreateSubscriptionCoach(context.Background(), createReq1)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription1)

		createReq2 := &booking.CreateSubscriptionCoachRequest{
			SubscriptionCoach: &booking.SubscriptionCoach{
				GymId:       gymID,
				CoachId:     coachID,
				Type:        "Personal Coaching 2",
				Description: "Description 2",
				Price:       200,
				Duration:    90,
			},
		}
		createdSubscription2, err := subscriptionRepo.CreateSubscriptionCoach(context.Background(), createReq2)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription2)

		listReq := &booking.ListSubscriptionCoachRequest{
			GymId: gymID,
		}

		listResponse, err := subscriptionRepo.ListSubscriptionCoach(context.Background(), listReq)
		assert.NoError(t, err)
		assert.NotNil(t, listResponse)
		assert.GreaterOrEqual(t, len(listResponse.SubscriptionCoach), 2) // At least 2 subscriptions

		// Cleanup
		defer deleteSubscriptionCoach(t, db, createdSubscription1.Id)
		defer deleteSubscriptionCoach(t, db, createdSubscription2.Id)
	})
}

func deleteSubscriptionCoach(t *testing.T, db *pgx.Conn, id string) {
	// _, err := db.Exec(context.Background(), "DELETE FROM subscription_coach WHERE id = $1", id)
	// assert.NoError(t, err)
}
