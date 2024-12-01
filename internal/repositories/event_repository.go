package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/alif-embreo-assessment/internal/models"
	"github.com/lib/pq"
)

type EventRepository interface {
	CreateEvent(event *models.Event) (*models.Event, error)
}

type eventRepository struct {
	DB *sql.DB
}

func NewEventRepository(db *sql.DB) EventRepository {
	return &eventRepository{
		DB: db,
	}
}

func (r *eventRepository) CreateEvent(event *models.Event) (*models.Event, error) {
	query := `
		INSERT INTO events (event_type_id, company_id, event_status_id, proposed_dates, location, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	err := r.DB.QueryRow(
		query,
		event.EventTypeID,
		event.CompanyID,
		event.EventStatusID,
		pq.Array(event.ProposedDates),
		event.Location,
		event.CreatedAt,
	).Scan(&event.ID)
	if err != nil {
		return nil, err
	}

	return event, nil
}
