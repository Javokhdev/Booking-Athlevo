package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/Athlevo/Booking-Athlevo/storage/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/stretchr/testify/assert"
)

func createDBConnection(t *testing.T) *pgx.Conn {
	dbCon := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		"sayyidmuhammad", // Replace with your DB user
		"root",           // Replace with your DB password
		"localhost",      // Replace with your DB host
		5432,             // Replace with your DB port
		"postgres",       // Replace with your DB name
	)

	db, err := pgx.Connect(context.Background(), dbCon)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}
	return db
}

func createGym(t *testing.T, db *pgx.Conn) string {
	gymID := uuid.New().String()
	query := `
		INSERT INTO sport_halls (
			id,
			name,
			location,
			longtitude,
			latitude,
			type_sport,
			type_gender,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
	`
	_, err := db.Exec(context.Background(), query,
		gymID,
		"Test Gym",
		"Test Location",
		"0.0",
		"0.0",
		"General Fitness",
		"male",
	)

	if err != nil {
		t.Fatalf("Failed to create gym: %v", err)
	}
	return gymID
}

func deleteGym(t *testing.T, db *pgx.Conn, gymID string) {

}

func TestSubscriptionPersonalRepo(t *testing.T) {
	db := createDBConnection(t)
	defer db.Close(context.Background())

	subscriptionRepo := postgres.NewSubscriptionPersonalRepo(db)

	gymID := createGym(t, db)
	defer deleteGym(t, db, gymID)

	t.Run("CreateSubscriptionPersonal", func(t *testing.T) {
		req := &booking.CreateSubscriptionPersonalRequest{
			SubscriptionPersonal: &booking.SubscriptionPersonal{
				GymId:       gymID,
				Type:        "Personal Training",
				Description: "One-on-one training sessions",
				Price:       100,
				Duration:    60, // In minutes
				Count:       10,
			},
		}

		createdSubscription, err := subscriptionRepo.CreateSubscriptionPersonal(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription)
		assert.NotEmpty(t, createdSubscription.Id)
		assert.Equal(t, req.SubscriptionPersonal.GymId, createdSubscription.GymId)
		assert.Equal(t, req.SubscriptionPersonal.Type, createdSubscription.Type)
		// ... (Assert other fields)

		// Cleanup
		defer deleteSubscriptionPersonal(t, db, createdSubscription.Id)
	})

	t.Run("GetSubscriptionPersonal", func(t *testing.T) {
		// Create a subscription to retrieve
		createReq := &booking.CreateSubscriptionPersonalRequest{
			SubscriptionPersonal: &booking.SubscriptionPersonal{
				GymId:       gymID,
				Type:        "Personal Training",
				Description: "One-on-one training sessions",
				Price:       100,
				Duration:    60, // In minutes
				Count:       10,
			},
		}
		createdSubscription, err := subscriptionRepo.CreateSubscriptionPersonal(context.Background(), createReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription)

		getReq := &booking.GetSubscriptionPersonalRequest{
			Id: createdSubscription.Id,
		}

		retrievedSubscription, err := subscriptionRepo.GetSubscriptionPersonal(context.Background(), getReq)
		assert.NoError(t, err)
		assert.NotNil(t, retrievedSubscription)
		assert.Equal(t, createdSubscription.Id, retrievedSubscription.Id)
		assert.Equal(t, createdSubscription.GymId, retrievedSubscription.GymId)

		// Cleanup
		defer deleteSubscriptionPersonal(t, db, createdSubscription.Id)
	})
	t.Run("UpdateSubscriptionPersonal", func(t *testing.T) {
		createReq := &booking.CreateSubscriptionPersonalRequest{
			SubscriptionPersonal: &booking.SubscriptionPersonal{
				GymId:       gymID,
				Type:        "Personal Training",
				Description: "One-on-one training sessions",
				Price:       100,
				Duration:    60, // In minutes
				Count:       10,
			},
		}
		createdSubscription, err := subscriptionRepo.CreateSubscriptionPersonal(context.Background(), createReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription)

		updateReq := &booking.UpdateSubscriptionPersonalRequest{
			SubscriptionPersonal: &booking.SubscriptionPersonal{
				Id:          createdSubscription.Id,
				GymId:       gymID,
				Type:        "Updated Training",
				Description: "Updated description",
				Price:       150,
				Duration:    90, // In minutes
				Count:       15,
			},
		}

		updatedSubscription, err := subscriptionRepo.UpdateSubscriptionPersonal(context.Background(), updateReq)
		assert.NoError(t, err)
		assert.NotNil(t, updatedSubscription)
		assert.Equal(t, updateReq.SubscriptionPersonal.Id, updatedSubscription.Id)
		assert.Equal(t, updateReq.SubscriptionPersonal.Type, updatedSubscription.Type)
		assert.Equal(t, updateReq.SubscriptionPersonal.Description, updatedSubscription.Description)
		assert.Equal(t, updateReq.SubscriptionPersonal.Price, updatedSubscription.Price)
		assert.Equal(t, updateReq.SubscriptionPersonal.Duration, updatedSubscription.Duration)
		assert.Equal(t, updateReq.SubscriptionPersonal.Count, updatedSubscription.Count)

		// Cleanup
		defer deleteSubscriptionPersonal(t, db, createdSubscription.Id)
	})

	t.Run("DeleteSubscriptionPersonal", func(t *testing.T) {
		createReq := &booking.CreateSubscriptionPersonalRequest{
			SubscriptionPersonal: &booking.SubscriptionPersonal{
				GymId:       gymID,
				Type:        "Personal Training",
				Description: "One-on-one training sessions",
				Price:       100,
				Duration:    60, // In minutes
				Count:       10,
			},
		}
		createdSubscription, err := subscriptionRepo.CreateSubscriptionPersonal(context.Background(), createReq)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription)

		deleteReq := &booking.DeleteSubscriptionPersonalRequest{
			Id: createdSubscription.Id,
		}

		err = subscriptionRepo.DeleteSubscriptionPersonal(context.Background(), deleteReq)
		assert.NoError(t, err)

		// Try to get the deleted subscription
		getReq := &booking.GetSubscriptionPersonalRequest{
			Id: createdSubscription.Id,
		}
		_, err = subscriptionRepo.GetSubscriptionPersonal(context.Background(), getReq)
		assert.ErrorIs(t, err, pgx.ErrNoRows)
	})

	t.Run("ListSubscriptionPersonal", func(t *testing.T) {
		// Create multiple subscriptions for the same gym
		createReq1 := &booking.CreateSubscriptionPersonalRequest{
			SubscriptionPersonal: &booking.SubscriptionPersonal{
				GymId:       gymID,
				Type:        "Personal Training 1",
				Description: "Description 1",
				Price:       100,
				Duration:    60,
				Count:       10,
			},
		}
		createdSubscription1, err := subscriptionRepo.CreateSubscriptionPersonal(context.Background(), createReq1)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription1)

		createReq2 := &booking.CreateSubscriptionPersonalRequest{
			SubscriptionPersonal: &booking.SubscriptionPersonal{
				GymId:       gymID,
				Type:        "Personal Training 2",
				Description: "Description 2",
				Price:       150,
				Duration:    90,
				Count:       15,
			},
		}
		createdSubscription2, err := subscriptionRepo.CreateSubscriptionPersonal(context.Background(), createReq2)
		assert.NoError(t, err)
		assert.NotNil(t, createdSubscription2)

		listReq := &booking.ListSubscriptionPersonalRequest{
			GymId: gymID,
		}

		listResponse, err := subscriptionRepo.ListSubscriptionPersonal(context.Background(), listReq)
		assert.NoError(t, err)
		assert.NotNil(t, listResponse)
		assert.GreaterOrEqual(t, len(listResponse.SubscriptionPersonal), 2) // At least 2 subscriptions

		// Cleanup
		defer deleteSubscriptionPersonal(t, db, createdSubscription1.Id)
		defer deleteSubscriptionPersonal(t, db, createdSubscription2.Id)
	})
}

func deleteSubscriptionPersonal(t *testing.T, db *pgx.Conn, id string) {
	// _, err := db.Exec(context.Background(), "DELETE FROM subscription_personal WHERE id = $1", id)	assert.NoError(t, err)
}
