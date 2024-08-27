package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// BookingPersonalRepo implements the BookingRepoI interface for BookingPersonal entities.
type BookingPersonalRepo struct {
	db *pgx.Conn
}

// NewBookingPersonalRepo creates a new BookingPersonalRepo.
func NewBookingPersonalRepo(db *pgx.Conn) *BookingPersonalRepo {
	return &BookingPersonalRepo{
		db: db,
	}
}

// CreateBookingPersonal creates a new booking personal record.
func (r *BookingPersonalRepo) CreateBookingPersonal(ctx context.Context, req *booking.CreateBookingPersonalRequest) (*booking.BookingPersonal, error) {
	req.BookingPersonal.Id = uuid.New().String()
	query := `
		INSERT INTO booking_personal (
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
		req.BookingPersonal.Id,
		req.BookingPersonal.UserId,
		req.BookingPersonal.SubscriptionId,
		req.BookingPersonal.Payment,
		req.BookingPersonal.AccessStatus,
		req.BookingPersonal.StartDate,
		req.BookingPersonal.Count,
	).Scan(
		&req.BookingPersonal.Id,
		&req.BookingPersonal.UserId,
		&req.BookingPersonal.SubscriptionId,
		&req.BookingPersonal.Payment,
		&req.BookingPersonal.AccessStatus,
		&startDate,
		&req.BookingPersonal.Count,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	req.BookingPersonal.StartDate = startDate.Format(time.RFC3339)
	req.BookingPersonal.CreatedAt = createdAt.Format(time.RFC3339)
	req.BookingPersonal.UpdatedAt = updatedAt.Format(time.RFC3339)

	return req.BookingPersonal, nil
}

// GetBookingPersonal retrieves a booking personal record by ID.
func (r *BookingPersonalRepo) GetBookingPersonal(ctx context.Context, req *booking.GetBookingPersonalRequest) (*booking.BookingPersonal, error) {
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
		FROM booking_personal
		WHERE id = $1
	`

	var (
		booking   booking.BookingPersonal
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

// UpdateBookingPersonal updates an existing booking personal record.
func (r *BookingPersonalRepo) UpdateBookingPersonal(ctx context.Context, req *booking.UpdateBookingPersonalRequest) (*booking.BookingPersonal, error) {
	query := `
		UPDATE booking_personal
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
		req.BookingPersonal.UserId,
		req.BookingPersonal.SubscriptionId,
		req.BookingPersonal.Payment,
		req.BookingPersonal.AccessStatus,
		req.BookingPersonal.StartDate,
		req.BookingPersonal.Count,
		req.BookingPersonal.Id,
	).Scan(
		&req.BookingPersonal.Id,
		&req.BookingPersonal.UserId,
		&req.BookingPersonal.SubscriptionId,
		&req.BookingPersonal.Payment,
		&req.BookingPersonal.AccessStatus,
		&startDate,
		&req.BookingPersonal.Count,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	req.BookingPersonal.StartDate = startDate.Format(time.RFC3339)
	req.BookingPersonal.CreatedAt = createdAt.Format(time.RFC3339)
	req.BookingPersonal.UpdatedAt = updatedAt.Format(time.RFC3339)

	return req.BookingPersonal, nil
}

// DeleteBookingPersonal deletes a booking personal record by ID.
func (r *BookingPersonalRepo) DeleteBookingPersonal(ctx context.Context, req *booking.DeleteBookingPersonalRequest) error {
	query := `
		DELETE FROM booking_personal
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

// ListBookingPersonal retrieves a list of booking personal records with optional filtering.
func (r *BookingPersonalRepo) ListBookingPersonal(ctx context.Context, req *booking.ListBookingPersonalRequest) (*booking.ListBookingPersonalResponse, error) {
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
		FROM booking_personal
		WHERE 1=1
	`

	if req.UserId != "" {
		query += fmt.Sprintf(" AND user_id = $%d", count)
		args = append(args, req.UserId)
		count++
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []*booking.BookingPersonal

	for rows.Next() {
		var (
			booking   booking.BookingPersonal
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

	return &booking.ListBookingPersonalResponse{BookingPersonal: bookings}, nil
}
