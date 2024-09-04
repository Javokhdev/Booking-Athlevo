package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Athlevo/Booking-Athlevo/config"
	"github.com/Athlevo/Booking-Athlevo/storage"
	"github.com/jackc/pgx/v5"
)

// StorageP implements the storage.StorageI interface for PostgreSQL.
type StorageP struct {
	db                       *pgx.Conn
	bookingPersonalRepo      storage.BookingPersonalRepoI
	bookingGroupRepo         storage.BookingGroupRepoI
	bookingCoachRepo         storage.BookingCoachRepoI
	subscriptionPersonalRepo storage.SubscriptionPersonalRepoI
	subscriptionGroupRepo    storage.SubscriptionGroupRepoI
	subscriptionCoachRepo    storage.SubscriptionCoachRepoI
	accessRepo               storage.AccessRepoI
	accessBetaRepo           storage.AccessRepoBetaI
}

// NewPostgresStorage creates a new PostgreSQL storage instance.
func NewPostgresStorage(cfg config.Config) (storage.StorageI, error) {
	dbCon := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	db, err := pgx.Connect(context.Background(), dbCon)
	if err != nil {
		slog.Warn("Unable to connect to database:" + err.Error())
		return nil, err
	}

	if err := db.Ping(context.Background()); err != nil {
		slog.Warn("Unable to ping database:" + err.Error())
		return nil, err
	}

	return &StorageP{
		db:                       db,
		bookingPersonalRepo:      NewBookingPersonalRepo(db),
		bookingGroupRepo:         NewBookingGroupRepo(db),
		bookingCoachRepo:         NewBookingCoachRepo(db),
		subscriptionPersonalRepo: NewSubscriptionPersonalRepo(db),
		subscriptionGroupRepo:    NewSubscriptionGroupRepo(db),
		subscriptionCoachRepo:    NewSubscriptionCoachRepo(db),
		accessRepo:               NewAccessRepo(db),
		accessBetaRepo:           NewAccessBetaRepo(db),
	}, nil
}

// BookingPersonal returns the BookingPersonalRepoI implementation for PostgreSQL.
func (s *StorageP) BookingPersonal() storage.BookingPersonalRepoI {
	return s.bookingPersonalRepo
}

// BookingGroup returns the BookingGroupRepoI implementation for PostgreSQL.
func (s *StorageP) BookingGroup() storage.BookingGroupRepoI {
	return s.bookingGroupRepo
}

// BookingCoach returns the BookingCoachRepoI implementation for PostgreSQL.
func (s *StorageP) BookingCoach() storage.BookingCoachRepoI {
	return s.bookingCoachRepo
}

// SubscriptionPersonal returns the SubscriptionPersonalRepoI implementation for PostgreSQL.
func (s *StorageP) SubscriptionPersonal() storage.SubscriptionPersonalRepoI {
	return s.subscriptionPersonalRepo
}

// SubscriptionGroup returns the SubscriptionGroupRepoI implementation for PostgreSQL.
func (s *StorageP) SubscriptionGroup() storage.SubscriptionGroupRepoI {
	return s.subscriptionGroupRepo
}

// SubscriptionCoach returns the SubscriptionCoachRepoI implementation for PostgreSQL.
func (s *StorageP) SubscriptionCoach() storage.SubscriptionCoachRepoI {
	return s.subscriptionCoachRepo
}

// Access returns the AccessRepoI implementation for PostgreSQL.
func (s *StorageP) Access() storage.AccessRepoI {
	return s.accessRepo
}

// Access returns the AccessRepoI implementation for PostgreSQL.
func (s *StorageP) AccessBeta() storage.AccessRepoBetaI {
	return s.accessBetaRepo
}
