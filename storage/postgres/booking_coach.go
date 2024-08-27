package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// BookingCoachRepo implements the BookingRepoI interface for BookingCoach entities.
type BookingCoachRepo struct {
	db *pgx.Conn
}

// NewBookingCoachRepo creates a new BookingCoachRepo.
func NewBookingCoachRepo(db *pgx.Conn) *BookingCoachRepo {
	return &BookingCoachRepo{
		db: db,
	}
}

// CreateBookingCoach creates a new booking coach record.
func (r *BookingCoachRepo) CreateBookingCoach(ctx context.Context, req *booking.CreateBookingCoachRequest) (*booking.BookingCoach, error) {
	req.BookingCoach.Id = uuid.New().String()
	query := `
		INSERT INTO booking_coach (
			id,
			user_id,
			subscription_id,
			payment,
			access_status,
			start_date,
			count,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
		RETURNING id, user_id, subscription_id, payment, access_status, start_date, count, created_at, updated_at
	`

	var (
		startDate time.Time
		createdAt time.Time
		updatedAt time.Time
	)

	err := r.db.QueryRow(ctx, query,
		req.BookingCoach.Id,
		req.BookingCoach.UserId,
		req.BookingCoach.SubscriptionId,
		req.BookingCoach.Payment,
		req.BookingCoach.AccessStatus,
		req.BookingCoach.StartDate,
		req.BookingCoach.Count,
	).Scan(
		&req.BookingCoach.Id,
		&req.BookingCoach.UserId,
		&req.BookingCoach.SubscriptionId,
		&req.BookingCoach.Payment,
		&req.BookingCoach.AccessStatus,
		&startDate,
		&req.BookingCoach.Count,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	req.BookingCoach.StartDate = startDate.Format(time.RFC3339)
	req.BookingCoach.CreatedAt = createdAt.Format(time.RFC3339)
	req.BookingCoach.UpdatedAt = updatedAt.Format(time.RFC3339)

	return req.BookingCoach, nil
}

// GetBookingCoach retrieves a booking coach record by ID.
func (r *BookingCoachRepo) GetBookingCoach(ctx context.Context, req *booking.GetBookingCoachRequest) (*booking.BookingCoach, error) {
	query := `
		SELECT
			id,
			user_id,
			subscription_id,
			payment,
			access_status,
			start_date,
			count,
			created_at,
			updated_at
		FROM booking_coach
		WHERE id = $1
	`

	var (
		booking   booking.BookingCoach
		startDate time.Time
		createdAt time.Time
		updatedAt time.Time
	)

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&booking.Id,
		&booking.UserId,
		&booking.SubscriptionId,
		&booking.Payment,
		&booking.AccessStatus,
		&startDate,
		&booking.Count,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	booking.StartDate = startDate.Format(time.RFC3339)
	booking.CreatedAt = createdAt.Format(time.RFC3339)
	booking.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &booking, nil
}

// UpdateBookingCoach updates an existing booking coach record.
func (r *BookingCoachRepo) UpdateBookingCoach(ctx context.Context, req *booking.UpdateBookingCoachRequest) (*booking.BookingCoach, error) {
	query := `
		UPDATE booking_coach
		SET
			user_id = $1,
			subscription_id = $2,
			payment = $3,
			access_status = $4,
			start_date = $5,
			count = $6,
			updated_at = NOW()
		WHERE id = $7
		RETURNING id, user_id, subscription_id, payment, access_status, start_date, count, created_at, updated_at
	`

	var (
		startDate time.Time
		createdAt time.Time
		updatedAt time.Time
	)

	err := r.db.QueryRow(ctx, query,
		req.BookingCoach.UserId,
		req.BookingCoach.SubscriptionId,
		req.BookingCoach.Payment,
		req.BookingCoach.AccessStatus,
		req.BookingCoach.StartDate,
		req.BookingCoach.Count,
		req.BookingCoach.Id,
	).Scan(
		&req.BookingCoach.Id,
		&req.BookingCoach.UserId,
		&req.BookingCoach.SubscriptionId,
		&req.BookingCoach.Payment,
		&req.BookingCoach.AccessStatus,
		&startDate,
		&req.BookingCoach.Count,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	req.BookingCoach.StartDate = startDate.Format(time.RFC3339)
	req.BookingCoach.CreatedAt = createdAt.Format(time.RFC3339)
	req.BookingCoach.UpdatedAt = updatedAt.Format(time.RFC3339)

	return req.BookingCoach, nil
}

// DeleteBookingCoach deletes a booking coach record by ID.
func (r *BookingCoachRepo) DeleteBookingCoach(ctx context.Context, req *booking.DeleteBookingCoachRequest) error {
	query := `
		DELETE FROM booking_coach
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

// ListBookingCoach retrieves a list of booking coach records with optional filtering.
func (r *BookingCoachRepo) ListBookingCoach(ctx context.Context, req *booking.ListBookingCoachRequest) (*booking.ListBookingCoachResponse, error) {
	var args []interface{}
	count := 1
	query := `
		SELECT
			id,
			user_id,
			subscription_id,
			payment,
			access_status,
			start_date,
			count,
			created_at,
			updated_at
		FROM booking_coach
		WHERE 1=1
	`

	if req.UserId != "" {
		query += fmt.Sprintf(" AND user_id = $%d", count)
		args = append(args, req.UserId)
		count++
	}

	if req.SubscriptionId != "" {
		query += fmt.Sprintf(" AND subscription_id = $%d", count)
		args = append(args, req.SubscriptionId)
		count++
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []*booking.BookingCoach

	for rows.Next() {
		var (
			booking   booking.BookingCoach
			startDate time.Time
			createdAt time.Time
			updatedAt time.Time
		)

		err := rows.Scan(
			&booking.Id,
			&booking.UserId,
			&booking.SubscriptionId,
			&booking.Payment,
			&booking.AccessStatus,
			&startDate,
			&booking.Count,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		booking.StartDate = startDate.Format(time.RFC3339)
		booking.CreatedAt = createdAt.Format(time.RFC3339)
		booking.UpdatedAt = updatedAt.Format(time.RFC3339)

		bookings = append(bookings, &booking)
	}

	return &booking.ListBookingCoachResponse{BookingCoach: bookings}, nil
}
