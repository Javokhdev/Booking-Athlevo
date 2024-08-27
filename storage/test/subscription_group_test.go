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

func TestSubscriptionGroupRepo(t *testing.T) {
	db := createDBConnection(t)
	defer db.Close(context.Background())

	subscriptionRepo := postgres.NewSubscriptionGroupRepo(db)

	gymID := createGym(t, db)
	defer deleteGym(t, db, gymID)

	coachID := uuid.New().String() // Replace with a valid coach ID if needed

	t.Run("CreateSubscriptionGroup", func(t *testing.T) {
		req := &booking.CreateSubscriptionGroupRequest{
			SubscriptionGroup: &booking.SubscriptionGroup{
				GymId:       gymID,
				CoachId:     coachID,
				Type:        "Group Fitness",
				Description: "High-intensity interval training",
				Price:       50,
				Capacity:    20,
				Time:        time.Now().Format(time.RFC3339), // Using RFC3339 for consistency
				Duration:    60,                              // In minutes
				Count:       10,
			},
		}

		createdSubscription, err := subscriptionRepo.CreateSubscriptionGroup(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription)
		assert.NotEmpty(t, createdSubscription.Id)
		assert.Equal(t, req.SubscriptionGroup.GymId, createdSubscription.GymId)
		assert.Equal(t, req.SubscriptionGroup.CoachId, createdSubscription.CoachId)
		assert.Equal(t, req.SubscriptionGroup.Type, createdSubscription.Type)
		// ... (Assert other fields)

		// Cleanup
		defer deleteSubscriptionGroup(t, db, createdSubscription.Id)
	})

	t.Run("GetSubscriptionGroup", func(t *testing.T) {
		// Create a subscription to retrieve
		createReq := &booking.CreateSubscriptionGroupRequest{
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
		createdSubscription, err := subscriptionRepo.CreateSubscriptionGroup(context.Background(), createReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription)

		getReq := &booking.GetSubscriptionGroupRequest{
			Id: createdSubscription.Id,
		}

		retrievedSubscription, err := subscriptionRepo.GetSubscriptionGroup(context.Background(), getReq)
		assert.NoError(t, err)
		assert.NotNil(t, retrievedSubscription)
		assert.Equal(t, createdSubscription.Id, retrievedSubscription.Id)
		assert.Equal(t, createdSubscription.GymId, retrievedSubscription.GymId)
		assert.Equal(t, createdSubscription.CoachId, retrievedSubscription.CoachId)
		// ... (Assert other fields)

		// Cleanup
		defer deleteSubscriptionGroup(t, db, createdSubscription.Id)
	})

	t.Run("UpdateSubscriptionGroup", func(t *testing.T) {
		createReq := &booking.CreateSubscriptionGroupRequest{
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
		createdSubscription, err := subscriptionRepo.CreateSubscriptionGroup(context.Background(), createReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription)

		updateReq := &booking.UpdateSubscriptionGroupRequest{
			SubscriptionGroup: &booking.SubscriptionGroup{
				Id:          createdSubscription.Id,
				GymId:       gymID,
				CoachId:     coachID,
				Type:        "Updated Group Fitness",
				Description: "Updated description",
				Price:       75,
				Capacity:    25,
				Time:        time.Now().Add(time.Hour * 24).Format(time.RFC3339), // Updated time
				Duration:    90,
				Count:       15,
			},
		}

		updatedSubscription, err := subscriptionRepo.UpdateSubscriptionGroup(context.Background(), updateReq)
		assert.NoError(t, err)
		assert.NotNil(t, updatedSubscription)
		assert.Equal(t, updateReq.SubscriptionGroup.Id, updatedSubscription.Id)
		assert.Equal(t, updateReq.SubscriptionGroup.Type, updatedSubscription.Type)
		assert.Equal(t, updateReq.SubscriptionGroup.Description, updatedSubscription.Description)
		assert.Equal(t, updateReq.SubscriptionGroup.Price, updatedSubscription.Price)
		assert.Equal(t, updateReq.SubscriptionGroup.Capacity, updatedSubscription.Capacity)
		assert.Equal(t, updateReq.SubscriptionGroup.Time, updatedSubscription.Time)
		assert.Equal(t, updateReq.SubscriptionGroup.Duration, updatedSubscription.Duration)
		assert.Equal(t, updateReq.SubscriptionGroup.Count, updatedSubscription.Count)

		// Cleanup
		defer deleteSubscriptionGroup(t, db, createdSubscription.Id)
	})

	t.Run("DeleteSubscriptionGroup", func(t *testing.T) {
		createReq := &booking.CreateSubscriptionGroupRequest{
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
		createdSubscription, err := subscriptionRepo.CreateSubscriptionGroup(context.Background(), createReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription)

		deleteReq := &booking.DeleteSubscriptionGroupRequest{
			Id: createdSubscription.Id,
		}

		err = subscriptionRepo.DeleteSubscriptionGroup(context.Background(), deleteReq)
		assert.NoError(t, err)

		// Try to get the deleted subscription
		getReq := &booking.GetSubscriptionGroupRequest{
			Id: createdSubscription.Id,
		}
		_, err = subscriptionRepo.GetSubscriptionGroup(context.Background(), getReq)
		assert.ErrorIs(t, err, pgx.ErrNoRows)
	})

	t.Run("ListSubscriptionGroup", func(t *testing.T) {
		// Create multiple subscriptions for the same gym
		createReq1 := &booking.CreateSubscriptionGroupRequest{
			SubscriptionGroup: &booking.SubscriptionGroup{
				GymId:       gymID,
				CoachId:     coachID,
				Type:        "Group Fitness 1",
				Description: "Description 1",
				Price:       50,
				Capacity:    20,
				Time:        time.Now().Format(time.RFC3339),
				Duration:    60,
				Count:       10,
			},
		}
		createdSubscription1, err := subscriptionRepo.CreateSubscriptionGroup(context.Background(), createReq1)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription1)

		createReq2 := &booking.CreateSubscriptionGroupRequest{
			SubscriptionGroup: &booking.SubscriptionGroup{
				GymId:       gymID,
				CoachId:     coachID,
				Type:        "Group Fitness 2",
				Description: "Description 2",
				Price:       75,
				Capacity:    25,
				Time:        time.Now().Add(time.Hour * 24).Format(time.RFC3339),
				Duration:    90,
				Count:       15,
			},
		}
		createdSubscription2, err := subscriptionRepo.CreateSubscriptionGroup(context.Background(), createReq2)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription2)

		listReq := &booking.ListSubscriptionGroupRequest{
			GymId: gymID,
		}

		listResponse, err := subscriptionRepo.ListSubscriptionGroup(context.Background(), listReq)
		assert.NoError(t, err)
		assert.NotNil(t, listResponse)
		assert.GreaterOrEqual(t, len(listResponse.SubscriptionGroup), 2) // At least 2 subscriptions

		// Cleanup
		defer deleteSubscriptionGroup(t, db, createdSubscription1.Id)
		defer deleteSubscriptionGroup(t, db, createdSubscription2.Id)
	})
}

func deleteSubscriptionGroup(t *testing.T, db *pgx.Conn, id string) {
	// _, err := db.Exec(context.Background(), "DELETE FROM subscription_group WHERE id = $1", id)
	// assert.NoError(t, err)
}
