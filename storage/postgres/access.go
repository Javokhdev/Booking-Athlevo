package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/jackc/pgx/v5"
)

// AccessRepo implements the AccessRepoI interface for Access entities.
type AccessRepo struct {
	db *pgx.Conn
}

// NewAccessRepo creates a new AccessRepo.
func NewAccessRepo(db *pgx.Conn) *AccessRepo {
	return &AccessRepo{
		db: db,
	}
}

// CreateAccessPersonal creates a new access record for a personal booking.
func (r *AccessRepo) CreateAccessPersonal(ctx context.Context, req *booking.CreateAccessPersonalRequest) (*booking.AccessPersonal, error) {
	// 1. Check booking access status
	if err := r.checkBookingAccessStatus(ctx, req.AccessPersonal.BookingPersonalId, "booking_personal"); err != nil {
		return nil, err
	}

	// 2. Create access record
	query := `
		INSERT INTO access_personal (
			booking_id,
			date
		) VALUES ($1, $2)
		RETURNING booking_id, date
	`

	var date time.Time

	err := r.db.QueryRow(ctx, query,
		req.AccessPersonal.BookingPersonalId,
		req.AccessPersonal.Date,
	).Scan(
		&req.AccessPersonal.BookingPersonalId,
		&date,
	)

	if err != nil {
		return nil, err
	}

	req.AccessPersonal.Date = date.Format(time.RFC3339)

	return req.AccessPersonal, nil
}

// ListAccessPersonal retrieves a list of access records for a personal booking.
func (r *AccessRepo) ListAccessPersonal(ctx context.Context, req *booking.ListAccessPersonalRequest) (*booking.ListAccessPersonalResponse, error) {
	query := `
		SELECT
			booking_id,
			date
		FROM access_personal
		WHERE booking_id = $1
	`

	rows, err := r.db.Query(ctx, query, req.BookingPersonalId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accesses []*booking.AccessPersonal

	for rows.Next() {
		var (
			access booking.AccessPersonal
			date   time.Time
		)

		err := rows.Scan(
			&access.BookingPersonalId,
			&date,
		)

		if err != nil {
			return nil, err
		}

		access.Date = date.Format(time.RFC3339)

		accesses = append(accesses, &access)
	}

	return &booking.ListAccessPersonalResponse{AccessPersonal: accesses}, nil
}

// CreateAccessGroup creates a new access record for a group booking.
func (r *AccessRepo) CreateAccessGroup(ctx context.Context, req *booking.CreateAccessGroupRequest) (*booking.AccessGroup, error) {
	// 1. Check booking access status
	if err := r.checkBookingAccessStatus(ctx, req.AccessGroup.BookingGroupId, "booking_group"); err != nil {
		return nil, err
	}

	// 2. Create access record
	query := `
		INSERT INTO access_group (
			booking_id,
			date
		) VALUES ($1, $2)
		RETURNING booking_id, date
	`

	var date time.Time

	err := r.db.QueryRow(ctx, query,
		req.AccessGroup.BookingGroupId,
		req.AccessGroup.Date,
	).Scan(
		&req.AccessGroup.BookingGroupId,
		&date,
	)

	if err != nil {
		return nil, err
	}

	req.AccessGroup.Date = date.Format(time.RFC3339)

	return req.AccessGroup, nil
}

// ListAccessGroup retrieves a list of access records for a group booking.
func (r *AccessRepo) ListAccessGroup(ctx context.Context, req *booking.ListAccessGroupRequest) (*booking.ListAccessGroupResponse, error) {
	query := `
		SELECT
			booking_id,
			date
		FROM access_group
		WHERE booking_id = $1
	`

	rows, err := r.db.Query(ctx, query, req.BookingGroupId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accesses []*booking.AccessGroup

	for rows.Next() {
		var (
			access booking.AccessGroup
			date   time.Time
		)

		err := rows.Scan(
			&access.BookingGroupId,
			&date,
		)

		if err != nil {
			return nil, err
		}

		access.Date = date.Format(time.RFC3339)

		accesses = append(accesses, &access)
	}

	return &booking.ListAccessGroupResponse{AccessGroup: accesses}, nil
}

// CreateAccessCoach creates a new access record for a coach booking.
func (r *AccessRepo) CreateAccessCoach(ctx context.Context, req *booking.CreateAccessCoachRequest) (*booking.AccessCoach, error) {
	// 1. Check booking access status
	if err := r.checkBookingAccessStatus(ctx, req.AccessCoach.BookingCoachId, "booking_coach"); err != nil {
		return nil, err
	}

	// 2. Create access record
	query := `
		INSERT INTO access_coach (
			booking_id,
			date
		) VALUES ($1, $2)
		RETURNING booking_id, date
	`

	var date time.Time

	err := r.db.QueryRow(ctx, query,
		req.AccessCoach.BookingCoachId,
		req.AccessCoach.Date,
	).Scan(
		&req.AccessCoach.BookingCoachId,
		&date,
	)

	if err != nil {
		return nil, err
	}

	req.AccessCoach.Date = date.Format(time.RFC3339)

	return req.AccessCoach, nil
}

// ListAccessCoach retrieves a list of access records for a coach booking.
func (r *AccessRepo) ListAccessCoach(ctx context.Context, req *booking.ListAccessCoachRequest) (*booking.ListAccessCoachResponse, error) {
	query := `
		SELECT
			booking_id,
			date
		FROM access_coach
		WHERE booking_id = $1
	`

	rows, err := r.db.Query(ctx, query, req.BookingCoachId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accesses []*booking.AccessCoach

	for rows.Next() {
		var (
			access booking.AccessCoach
			date   time.Time
		)

		err := rows.Scan(
			&access.BookingCoachId,
			&date,
		)

		if err != nil {
			return nil, err
		}

		access.Date = date.Format(time.RFC3339)

		accesses = append(accesses, &access)
	}

	return &booking.ListAccessCoachResponse{AccessCoach: accesses}, nil
}

// checkBookingAccessStatus checks if the booking has "granted" access status.
func (r *AccessRepo) checkBookingAccessStatus(ctx context.Context, bookingID, bookingTable string) error {
	var accessStatus string
	query := fmt.Sprintf(`SELECT access_status FROM %s WHERE id = $1`, bookingTable)
	err := r.db.QueryRow(ctx, query, bookingID).Scan(&accessStatus)
	if err != nil {
		return fmt.Errorf("error checking booking access status: %w", err)
	}
	log.Print(bookingID, accessStatus)
	if accessStatus != "granted" {
		return fmt.Errorf("access denied: booking status is not 'granted'")
	}
	return nil
}
