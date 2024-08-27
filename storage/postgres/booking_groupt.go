package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// BookingGroupRepo implements the BookingRepoI interface for BookingGroup entities.
type BookingGroupRepo struct {
	db *pgx.Conn
}

// NewBookingGroupRepo creates a new BookingGroupRepo.
func NewBookingGroupRepo(db *pgx.Conn) *BookingGroupRepo {
	return &BookingGroupRepo{
		db: db,
	}
}

// CreateBookingGroup creates a new booking group record if capacity allows.
func (r *BookingGroupRepo) CreateBookingGroup(ctx context.Context, req *booking.CreateBookingGroupRequest) (*booking.BookingGroup, error) {
	// 1. Get the subscription capacity
	var capacity int
	err := r.db.QueryRow(ctx, "SELECT capacity FROM subscription_group WHERE id = $1", req.BookingGroup.SubscriptionId).Scan(&capacity)
	if err != nil {
		return nil, fmt.Errorf("error getting subscription capacity: %w", err)
	}

	// 2. Count existing active bookings for the subscription
	var activeBookings int
	err = r.db.QueryRow(ctx, `
		SELECT COUNT(*) 
		FROM booking_group 
		WHERE subscription_id = $1 AND access_status = 'granted' AND start_date <= NOW() AND start_date + (
			SELECT duration * INTERVAL '1 hour' FROM subscription_group WHERE id = $1
		) > NOW()
	`, req.BookingGroup.SubscriptionId).Scan(&activeBookings)
	if err != nil {
		return nil, fmt.Errorf("error counting active bookings: %w", err)
	}

	// 3. Check if capacity allows new booking
	if activeBookings >= capacity {
		return nil, fmt.Errorf("group is full, capacity reached")
	}

	// 4. Create the booking if capacity allows
	req.BookingGroup.Id = uuid.New().String()
	query := `
		INSERT INTO booking_group (
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

	err = r.db.QueryRow(ctx, query,
		req.BookingGroup.Id,
		req.BookingGroup.UserId,
		req.BookingGroup.SubscriptionId,
		req.BookingGroup.Payment,
		req.BookingGroup.AccessStatus,
		req.BookingGroup.StartDate,
		req.BookingGroup.Count,
	).Scan(
		&req.BookingGroup.Id,
		&req.BookingGroup.UserId,
		&req.BookingGroup.SubscriptionId,
		&req.BookingGroup.Payment,
		&req.BookingGroup.AccessStatus,
		&startDate,
		&req.BookingGroup.Count,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	req.BookingGroup.StartDate = startDate.Format(time.RFC3339)
	req.BookingGroup.CreatedAt = createdAt.Format(time.RFC3339)
	req.BookingGroup.UpdatedAt = updatedAt.Format(time.RFC3339)

	return req.BookingGroup, nil
}

// // CreateBookingGroup creates a new booking group record.
// func (r *BookingGroupRepo) CreateBookingGroup(ctx context.Context, req *booking.CreateBookingGroupRequest) (*booking.BookingGroup, error) {
// 	req.BookingGroup.Id = uuid.New().String()
// 	query := `
// 		INSERT INTO booking_group (
// 			id,
// 			user_id,
// 			subscription_id,
// 			payment,
// 			access_status,
// 			start_date,
// 			count,
// 			created_at,
// 			updated_at
// 		) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
// 		RETURNING id, user_id, subscription_id, payment, access_status, start_date, count, created_at, updated_at
// 	`

// 	var (
// 		startDate time.Time
// 		createdAt time.Time
// 		updatedAt time.Time
// 	)

// 	err := r.db.QueryRow(ctx, query,
// 		req.BookingGroup.Id,
// 		req.BookingGroup.UserId,
// 		req.BookingGroup.SubscriptionId,
// 		req.BookingGroup.Payment,
// 		req.BookingGroup.AccessStatus,
// 		req.BookingGroup.StartDate,
// 		req.BookingGroup.Count,
// 	).Scan(
// 		&req.BookingGroup.Id,
// 		&req.BookingGroup.UserId,
// 		&req.BookingGroup.SubscriptionId,
// 		&req.BookingGroup.Payment,
// 		&req.BookingGroup.AccessStatus,
// 		&startDate,
// 		&req.BookingGroup.Count,
// 		&createdAt,
// 		&updatedAt,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	req.BookingGroup.StartDate = startDate.Format(time.RFC3339)
// 	req.BookingGroup.CreatedAt = createdAt.Format(time.RFC3339)
// 	req.BookingGroup.UpdatedAt = updatedAt.Format(time.RFC3339)

// 	return req.BookingGroup, nil
// }

// GetBookingGroup retrieves a booking group record by ID.
func (r *BookingGroupRepo) GetBookingGroup(ctx context.Context, req *booking.GetBookingGroupRequest) (*booking.BookingGroup, error) {
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
		FROM booking_group
		WHERE id = $1
	`

	var (
		booking   booking.BookingGroup
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

// UpdateBookingGroup updates an existing booking group record.
func (r *BookingGroupRepo) UpdateBookingGroup(ctx context.Context, req *booking.UpdateBookingGroupRequest) (*booking.BookingGroup, error) {
	query := `
		UPDATE booking_group
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
		req.BookingGroup.UserId,
		req.BookingGroup.SubscriptionId,
		req.BookingGroup.Payment,
		req.BookingGroup.AccessStatus,
		req.BookingGroup.StartDate,
		req.BookingGroup.Count,
		req.BookingGroup.Id,
	).Scan(
		&req.BookingGroup.Id,
		&req.BookingGroup.UserId,
		&req.BookingGroup.SubscriptionId,
		&req.BookingGroup.Payment,
		&req.BookingGroup.AccessStatus,
		&startDate,
		&req.BookingGroup.Count,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	req.BookingGroup.StartDate = startDate.Format(time.RFC3339)
	req.BookingGroup.CreatedAt = createdAt.Format(time.RFC3339)
	req.BookingGroup.UpdatedAt = updatedAt.Format(time.RFC3339)

	return req.BookingGroup, nil
}

// DeleteBookingGroup deletes a booking group record by ID.
func (r *BookingGroupRepo) DeleteBookingGroup(ctx context.Context, req *booking.DeleteBookingGroupRequest) error {
	query := `
		DELETE FROM booking_group
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

// ListBookingGroup retrieves a list of booking group records with optional filtering.
func (r *BookingGroupRepo) ListBookingGroup(ctx context.Context, req *booking.ListBookingGroupRequest) (*booking.ListBookingGroupResponse, error) {
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
		FROM booking_group
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

	var bookings []*booking.BookingGroup

	for rows.Next() {
		var (
			booking   booking.BookingGroup
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

	return &booking.ListBookingGroupResponse{BookingGroup: bookings}, nil
}
