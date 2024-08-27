package postgres

import (
	"context"
	"time"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type SubscriptionPersonalRepo struct {
	db *pgx.Conn
}

// NewSubscriptionPersonalRepo creates a new SubscriptionPersonalRepo.
func NewSubscriptionPersonalRepo(db *pgx.Conn) *SubscriptionPersonalRepo {
	return &SubscriptionPersonalRepo{
		db: db,
	}
}
func (r *SubscriptionPersonalRepo) CreateSubscriptionPersonal(ctx context.Context, req *booking.CreateSubscriptionPersonalRequest) (*booking.SubscriptionPersonal, error) {
	req.SubscriptionPersonal.Id = uuid.New().String()
	query := `
		INSERT INTO subscription_personal (
			id,
			gym_id,
			type,
			description,
			price,
			duration,
			count,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
		RETURNING id, gym_id, type, description, price, duration, count, created_at, updated_at
	`

	var (
		createdAt time.Time
		updatedAt time.Time
	)

	err := r.db.QueryRow(ctx, query,
		req.SubscriptionPersonal.Id,
		req.SubscriptionPersonal.GymId,
		req.SubscriptionPersonal.Type,
		req.SubscriptionPersonal.Description,
		req.SubscriptionPersonal.Price,
		req.SubscriptionPersonal.Duration,
		req.SubscriptionPersonal.Count,
	).Scan(
		&req.SubscriptionPersonal.Id,
		&req.SubscriptionPersonal.GymId,
		&req.SubscriptionPersonal.Type,
		&req.SubscriptionPersonal.Description,
		&req.SubscriptionPersonal.Price,
		&req.SubscriptionPersonal.Duration,
		&req.SubscriptionPersonal.Count,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	req.SubscriptionPersonal.CreatedAt = createdAt.Format(time.RFC3339)
	req.SubscriptionPersonal.UpdatedAt = updatedAt.Format(time.RFC3339)

	return req.SubscriptionPersonal, nil
}

func (r *SubscriptionPersonalRepo) GetSubscriptionPersonal(ctx context.Context, req *booking.GetSubscriptionPersonalRequest) (*booking.SubscriptionPersonal, error) {
	query := `
		SELECT
			id,
			gym_id,
			type,
			description,
			price,
			duration,
			count,
			created_at,
			updated_at
		FROM subscription_personal
		WHERE id = $1
	`

	var (
		subscription booking.SubscriptionPersonal
		createdAt    time.Time
		updatedAt    time.Time
	)

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&subscription.Id,
		&subscription.GymId,
		&subscription.Type,
		&subscription.Description,
		&subscription.Price,
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

	subscription.CreatedAt = createdAt.Format(time.RFC3339)
	subscription.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &subscription, nil
}

func (r *SubscriptionPersonalRepo) UpdateSubscriptionPersonal(ctx context.Context, req *booking.UpdateSubscriptionPersonalRequest) (*booking.SubscriptionPersonal, error) {
	query := `
		UPDATE subscription_personal
		SET
			gym_id = $1,
			type = $2,
			description = $3,
			price = $4,
			duration = $5,
			count = $6,
			updated_at = NOW()
		WHERE id = $7
		RETURNING id, gym_id, type, description, price, duration, count, created_at, updated_at
	`

	var (
		createdAt time.Time
		updatedAt time.Time
	)

	err := r.db.QueryRow(ctx, query,
		req.SubscriptionPersonal.GymId,
		req.SubscriptionPersonal.Type,
		req.SubscriptionPersonal.Description,
		req.SubscriptionPersonal.Price,
		req.SubscriptionPersonal.Duration,
		req.SubscriptionPersonal.Count,
		req.SubscriptionPersonal.Id,
	).Scan(
		&req.SubscriptionPersonal.Id,
		&req.SubscriptionPersonal.GymId,
		&req.SubscriptionPersonal.Type,
		&req.SubscriptionPersonal.Description,
		&req.SubscriptionPersonal.Price,
		&req.SubscriptionPersonal.Duration,
		&req.SubscriptionPersonal.Count,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	req.SubscriptionPersonal.CreatedAt = createdAt.Format(time.RFC3339)
	req.SubscriptionPersonal.UpdatedAt = updatedAt.Format(time.RFC3339)

	return req.SubscriptionPersonal, nil
}

func (r *SubscriptionPersonalRepo) DeleteSubscriptionPersonal(ctx context.Context, req *booking.DeleteSubscriptionPersonalRequest) error {
	query := `
		DELETE FROM subscription_personal
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

func (r *SubscriptionPersonalRepo) ListSubscriptionPersonal(ctx context.Context, req *booking.ListSubscriptionPersonalRequest) (*booking.ListSubscriptionPersonalResponse, error) {
	query := `
		SELECT
			id,
			gym_id,
			type,
			description,
			price,
			duration,
			count,
			created_at,
			updated_at
		FROM subscription_personal
		WHERE gym_id = $1
	`

	rows, err := r.db.Query(ctx, query, req.GymId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []*booking.SubscriptionPersonal

	for rows.Next() {
		var (
			subscription booking.SubscriptionPersonal
			createdAt    time.Time
			updatedAt    time.Time
		)

		err := rows.Scan(
			&subscription.Id,
			&subscription.GymId,
			&subscription.Type,
			&subscription.Description,
			&subscription.Price,
			&subscription.Duration,
			&subscription.Count,
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

	return &booking.ListSubscriptionPersonalResponse{SubscriptionPersonal: subscriptions}, nil
}
