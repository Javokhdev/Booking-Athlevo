package postgres

import (
	"context"
	"time"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// SubscriptionCoachRepo implements the SubscriptionRepoI interface for SubscriptionCoach entities.
type SubscriptionCoachRepo struct {
	db *pgx.Conn
}

// NewSubscriptionCoachRepo creates a new SubscriptionCoachRepo.
func NewSubscriptionCoachRepo(db *pgx.Conn) *SubscriptionCoachRepo {
	return &SubscriptionCoachRepo{
		db: db,
	}
}

// CreateSubscriptionCoach creates a new subscription coach record.
func (r *SubscriptionCoachRepo) CreateSubscriptionCoach(ctx context.Context, req *booking.CreateSubscriptionCoachRequest) (*booking.SubscriptionCoach, error) {
	req.SubscriptionCoach.Id = uuid.New().String()
	query := `
		INSERT INTO subscription_coach (
			id,
			gym_id,
			coach_id,
			type,
			description,
			price,
			duration,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
		RETURNING id, gym_id, coach_id, type, description, price, duration, created_at, updated_at
	`

	var (
		createdAt time.Time
		updatedAt time.Time
	)

	err := r.db.QueryRow(ctx, query,
		req.SubscriptionCoach.Id,
		req.SubscriptionCoach.GymId,
		req.SubscriptionCoach.CoachId,
		req.SubscriptionCoach.Type,
		req.SubscriptionCoach.Description,
		req.SubscriptionCoach.Price,
		req.SubscriptionCoach.Duration,
	).Scan(
		&req.SubscriptionCoach.Id,
		&req.SubscriptionCoach.GymId,
		&req.SubscriptionCoach.CoachId,
		&req.SubscriptionCoach.Type,
		&req.SubscriptionCoach.Description,
		&req.SubscriptionCoach.Price,
		&req.SubscriptionCoach.Duration,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	req.SubscriptionCoach.CreatedAt = createdAt.Format(time.RFC3339)
	req.SubscriptionCoach.UpdatedAt = updatedAt.Format(time.RFC3339)

	return req.SubscriptionCoach, nil
}

// GetSubscriptionCoach retrieves a subscription coach record by ID.
func (r *SubscriptionCoachRepo) GetSubscriptionCoach(ctx context.Context, req *booking.GetSubscriptionCoachRequest) (*booking.SubscriptionCoach, error) {
	query := `
		SELECT
			id,
			gym_id,
			coach_id,
			type,
			description,
			price,
			duration,
			created_at,
			updated_at
		FROM subscription_coach
		WHERE id = $1
	`

	var (
		subscription booking.SubscriptionCoach
		createdAt    time.Time
		updatedAt    time.Time
	)

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&subscription.Id,
		&subscription.GymId,
		&subscription.CoachId,
		&subscription.Type,
		&subscription.Description,
		&subscription.Price,
		&subscription.Duration,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	subscription.CreatedAt = createdAt.Format(time.RFC3339)
	subscription.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &subscription, nil
}

// UpdateSubscriptionCoach updates an existing subscription coach record.
func (r *SubscriptionCoachRepo) UpdateSubscriptionCoach(ctx context.Context, req *booking.UpdateSubscriptionCoachRequest) (*booking.SubscriptionCoach, error) {
	query := `
		UPDATE subscription_coach
		SET
			gym_id = $1,
			coach_id = $2,
			type = $3,
			description = $4,
			price = $5,
			duration = $6,
			updated_at = NOW()
		WHERE id = $7
		RETURNING id, gym_id, coach_id, type, description, price, duration, created_at, updated_at
	`

	var (
		createdAt time.Time
		updatedAt time.Time
	)

	err := r.db.QueryRow(ctx, query,
		req.SubscriptionCoach.GymId,
		req.SubscriptionCoach.CoachId,
		req.SubscriptionCoach.Type,
		req.SubscriptionCoach.Description,
		req.SubscriptionCoach.Price,
		req.SubscriptionCoach.Duration,
		req.SubscriptionCoach.Id,
	).Scan(
		&req.SubscriptionCoach.Id,
		&req.SubscriptionCoach.GymId,
		&req.SubscriptionCoach.CoachId,
		&req.SubscriptionCoach.Type,
		&req.SubscriptionCoach.Description,
		&req.SubscriptionCoach.Price,
		&req.SubscriptionCoach.Duration,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	req.SubscriptionCoach.CreatedAt = createdAt.Format(time.RFC3339)
	req.SubscriptionCoach.UpdatedAt = updatedAt.Format(time.RFC3339)

	return req.SubscriptionCoach, nil
}

// DeleteSubscriptionCoach deletes a subscription coach record by ID.
func (r *SubscriptionCoachRepo) DeleteSubscriptionCoach(ctx context.Context, req *booking.DeleteSubscriptionCoachRequest) error {
	query := `
		DELETE FROM subscription_coach
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

// ListSubscriptionCoach retrieves a list of subscription coach records by gym ID.
func (r *SubscriptionCoachRepo) ListSubscriptionCoach(ctx context.Context, req *booking.ListSubscriptionCoachRequest) (*booking.ListSubscriptionCoachResponse, error) {
	query := `
		SELECT
			id,
			gym_id,
			coach_id,
			type,
			description,
			price,
			duration,
			created_at,
			updated_at
		FROM subscription_coach
		WHERE gym_id = $1
	`

	rows, err := r.db.Query(ctx, query, req.GymId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []*booking.SubscriptionCoach

	for rows.Next() {
		var (
			subscription booking.SubscriptionCoach
			createdAt    time.Time
			updatedAt    time.Time
		)

		err := rows.Scan(
			&subscription.Id,
			&subscription.GymId,
			&subscription.CoachId,
			&subscription.Type,
			&subscription.Description,
			&subscription.Price,
			&subscription.Duration,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		subscription.CreatedAt = createdAt.Format(time.RFC3339)
		subscription.UpdatedAt = updatedAt.Format(time.RFC3339)

		subscriptions = append(subscriptions, &subscription)
	}

	return &booking.ListSubscriptionCoachResponse{SubscriptionCoach: subscriptions}, nil
}
