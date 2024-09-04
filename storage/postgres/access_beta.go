package postgres

import (
	"context"
	"fmt"

	"github.com/Athlevo/Booking-Athlevo/genproto/booking"
	"github.com/jackc/pgx/v5"
)

// AccessRepo implements the AccessRepoI interface for Access entities.
type AccessBetaRepo struct {
	db *pgx.Conn
}

// NewAccessRepo creates a new AccessRepo.
func NewAccessBetaRepo(db *pgx.Conn) *AccessBetaRepo {
	return &AccessBetaRepo{
		db: db,
	}
}

// CheckUserAccess checks if the user has access to the sport hall for personal subscriptions.
func (r *AccessBetaRepo) CheckUserAccess(ctx context.Context, req *booking.AccessBetaPersonalRequest) (*booking.AccessBetaPersonalResponse, error) {
	// 1. Find the user's active personal booking for the sport hall
	query := `
		SELECT bp.id, bp.access_status
		FROM booking_personal bp
		JOIN subscription_personal sp ON bp.subscription_id = sp.id
		WHERE bp.user_id = $1 AND sp.gym_id = $2 AND bp.access_status = 'granted'
		AND bp.start_date <= NOW()
	`

	var bookingID string
	var accessStatus string
	err := r.db.QueryRow(ctx, query, req.UserId, req.SportHallId).Scan(&bookingID, &accessStatus)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &booking.AccessBetaPersonalResponse{Message: "denied"}, nil
		}
		return nil, fmt.Errorf("error checking user access: %w", err)
	}

	// 2. If access is granted, create an access_personal record
	if accessStatus == "granted" {
		if err := r.createAccessPersonalRecord(ctx, bookingID); err != nil {
			return nil, fmt.Errorf("error creating access personal record: %w", err)
		}
		return &booking.AccessBetaPersonalResponse{Message: "granted"}, nil
	}

	return &booking.AccessBetaPersonalResponse{Message: "denied"}, nil
}

// createAccessPersonalRecord creates a new access_personal record.
func (r *AccessBetaRepo) createAccessPersonalRecord(ctx context.Context, bookingID string) error {
	query := `
		INSERT INTO access_personal (booking_id, date)
		VALUES ($1, NOW())
	`
	_, err := r.db.Exec(ctx, query, bookingID)
	if err != nil {
		return fmt.Errorf("error creating access personal record: %w", err)
	}
	return nil
}
