package postgres

import (
	"context"
	"time"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// SubscriptionGroupRepo implements the SubscriptionRepoI interface for SubscriptionGroup entities.
type SubscriptionGroupRepo struct {
	db *pgx.Conn
}

// NewSubscriptionGroupRepo creates a new SubscriptionGroupRepo.
func NewSubscriptionGroupRepo(db *pgx.Conn) *SubscriptionGroupRepo {
	return &SubscriptionGroupRepo{
		db: db,
	}
}

// CreateSubscriptionGroup creates a new subscription group record.
func (r *SubscriptionGroupRepo) CreateSubscriptionGroup(ctx context.Context, req *booking.CreateSubscriptionGroupRequest) (*booking.SubscriptionGroup, error) {
	req.SubscriptionGroup.Id = uuid.New().String()
	query := `
		INSERT INTO subscription_group (
			id,
			gym_id,
			coach_id,
			type,
			description,
			price,
			capacity,
			time,
			duration,
			count,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW())
		RETURNING id, gym_id, coach_id, type, description, price, capacity, time, duration, count, created_at, updated_at
	`

	var (
		timeT     time.Time
		createdAt time.Time
		updatedAt time.Time
	)

	err := r.db.QueryRow(ctx, query,
		req.SubscriptionGroup.Id,
		req.SubscriptionGroup.GymId,
		req.SubscriptionGroup.CoachId,
		req.SubscriptionGroup.Type,
		req.SubscriptionGroup.Description,
		req.SubscriptionGroup.Price,
		req.SubscriptionGroup.Capacity,
		req.SubscriptionGroup.Time, // No conversion needed
		req.SubscriptionGroup.Duration,
		req.SubscriptionGroup.Count,
	).Scan(
		&req.SubscriptionGroup.Id,
		&req.SubscriptionGroup.GymId,
		&req.SubscriptionGroup.CoachId,
		&req.SubscriptionGroup.Type,
		&req.SubscriptionGroup.Description,
		&req.SubscriptionGroup.Price,
		&req.SubscriptionGroup.Capacity,
		&timeT, // No conversion needed
		&req.SubscriptionGroup.Duration,
		&req.SubscriptionGroup.Count,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}
	req.SubscriptionGroup.Time = timeT.Format(time.RFC3339)
	req.SubscriptionGroup.CreatedAt = createdAt.Format(time.RFC3339)
	req.SubscriptionGroup.UpdatedAt = updatedAt.Format(time.RFC3339)

	return req.SubscriptionGroup, nil
}

// GetSubscriptionGroup retrieves a subscription group record by ID.
func (r *SubscriptionGroupRepo) GetSubscriptionGroup(ctx context.Context, req *booking.GetSubscriptionGroupRequest) (*booking.SubscriptionGroup, error) {
	query := `
		SELECT
			id,
			gym_id,
			coach_id,
			type,
			description,
			price,
			capacity,
			time,
			duration,
			count,
			created_at,
			updated_at
		FROM subscription_group
		WHERE id = $1
	`

	var (
		subscription booking.SubscriptionGroup
		timeT        time.Time
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
		&subscription.Capacity,
		&timeT, // Scan into time.Time
		&subscription.Duration,
		&subscription.Count,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	subscription.Time = timeT.Format(time.RFC3339) // Format time.Time to string
	subscription.CreatedAt = createdAt.Format(time.RFC3339)
	subscription.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &subscription, nil
}

// UpdateSubscriptionGroup updates an existing subscription group record.
func (r *SubscriptionGroupRepo) UpdateSubscriptionGroup(ctx context.Context, req *booking.UpdateSubscriptionGroupRequest) (*booking.SubscriptionGroup, error) {
	query := `
		UPDATE subscription_group
		SET
			gym_id = $1,
			coach_id = $2,
			type = $3,
			description = $4,
			price = $5,
			capacity = $6,
			time = $7,
			duration = $8,
			count = $9,
			updated_at = NOW()
		WHERE id = $10
		RETURNING id, gym_id, coach_id, type, description, price, capacity, time, duration, count, created_at, updated_at
	`

	var (
		timeT     time.Time
		createdAt time.Time
		updatedAt time.Time
	)

	err := r.db.QueryRow(ctx, query,
		req.SubscriptionGroup.GymId,
		req.SubscriptionGroup.CoachId,
		req.SubscriptionGroup.Type,
		req.SubscriptionGroup.Description,
		req.SubscriptionGroup.Price,
		req.SubscriptionGroup.Capacity,
		req.SubscriptionGroup.Time, // No conversion needed
		req.SubscriptionGroup.Duration,
		req.SubscriptionGroup.Count,
		req.SubscriptionGroup.Id,
	).Scan(
		&req.SubscriptionGroup.Id,
		&req.SubscriptionGroup.GymId,
		&req.SubscriptionGroup.CoachId,
		&req.SubscriptionGroup.Type,
		&req.SubscriptionGroup.Description,
		&req.SubscriptionGroup.Price,
		&req.SubscriptionGroup.Capacity,
		&timeT, // No conversion needed
		&req.SubscriptionGroup.Duration,
		&req.SubscriptionGroup.Count,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}
	req.SubscriptionGroup.Time = timeT.Format(time.RFC3339)
	req.SubscriptionGroup.CreatedAt = createdAt.Format(time.RFC3339)
	req.SubscriptionGroup.UpdatedAt = updatedAt.Format(time.RFC3339)

	return req.SubscriptionGroup, nil
}

// DeleteSubscriptionGroup deletes a subscription group record by ID.
func (r *SubscriptionGroupRepo) DeleteSubscriptionGroup(ctx context.Context, req *booking.DeleteSubscriptionGroupRequest) error {
	query := `
		DELETE FROM subscription_group
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

// ListSubscriptionGroup retrieves a list of subscription group records by gym ID.
func (r *SubscriptionGroupRepo) ListSubscriptionGroup(ctx context.Context, req *booking.ListSubscriptionGroupRequest) (*booking.ListSubscriptionGroupResponse, error) {
	query := `
		SELECT
			id,
			gym_id,
			coach_id,
			type,
			description,
			price,
			capacity,
			time,
			duration,
			count,
			created_at,
			updated_at
		FROM subscription_group
		WHERE gym_id = $1
	`

	rows, err := r.db.Query(ctx, query, req.GymId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []*booking.SubscriptionGroup

	for rows.Next() {
		var (
			subscription booking.SubscriptionGroup
			timeT        time.Time
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
			&subscription.Capacity,
			&timeT, // Scan into time.Time
			&subscription.Duration,
			&subscription.Count,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		subscription.Time = timeT.Format(time.RFC3339) // Format time.Time to string
		subscription.CreatedAt = createdAt.Format(time.RFC3339)
		subscription.UpdatedAt = updatedAt.Format(time.RFC3339)

		subscriptions = append(subscriptions, &subscription)
	}

	return &booking.ListSubscriptionGroupResponse{SubscriptionGroup: subscriptions}, nil
}
