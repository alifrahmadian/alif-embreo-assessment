package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/alif-embreo-assessment/internal/models"
	"github.com/alifrahmadian/alif-embreo-assessment/pkg/errors"
	"github.com/lib/pq"
)

type EventRepository interface {
	CreateEvent(event *models.Event) (*models.Event, error)
	GetEventByID(id int64) (*models.Event, error)
	GetAllEvents() ([]*models.Event, error)
	ApproveEvent(int64, *models.Event) error
	RejectEvent(int64, *models.Event) error
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

func (r *eventRepository) GetEventByID(id int64) (*models.Event, error) {
	query := `
		SELECT
			events.id,
			event_type_id,
			event_types.id,
			event_types.name,
			company_id,
			companies.id,
			companies.name,
			vendor_id,
			vendors.id,
			vendors.name,
			event_status_id,
			event_status.id,
			event_status.name,
			proposed_dates,
			confirmed_date,
			location,
			rejected_remarks,
			created_at
		FROM
			events
		LEFT JOIN
			event_types ON event_type_id = event_types.id
		LEFT JOIN
			companies ON company_id = companies.id
		LEFT JOIN
			vendors on vendor_id = vendors.id
		LEFT JOIN
			event_status on event_status_id = event_status.id
		WHERE
			events.id = $1;
	`

	event := &models.Event{}

	err := r.DB.QueryRow(
		query,
		id,
	).Scan(
		&event.ID,
		&event.EventTypeID,
		&event.EventType.ID,
		&event.EventType.Name,
		&event.CompanyID,
		&event.Company.ID,
		&event.Company.Name,
		&event.VendorID,
		&event.Vendor.ID,
		&event.Vendor.Name,
		&event.EventStatusID,
		&event.EventStatus.ID,
		&event.EventStatus.Name,
		pq.Array(&event.ProposedDates),
		&event.ConfirmedDate,
		&event.Location,
		&event.RejectedRemarks,
		&event.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrEventNotFound
		}

		return nil, err
	}

	return event, nil
}

func (r *eventRepository) GetAllEvents() ([]*models.Event, error) {
	query := `
		SELECT
			events.id,
			event_type_id,
			event_types.id,
			event_types.name,
			company_id,
			companies.id,
			companies.name,
			vendor_id,
			vendors.id,
			vendors.name,
			event_status_id,
			event_status.id,
			event_status.name,
			proposed_dates,
			confirmed_date,
			location,
			rejected_remarks,
			created_at
		FROM
			events
		LEFT JOIN
			event_types ON event_type_id = event_types.id
		LEFT JOIN
			companies ON company_id = companies.id
		LEFT JOIN
			vendors on vendor_id = vendors.id
		LEFT JOIN
			event_status on event_status_id = event_status.id;
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var events []*models.Event

	for rows.Next() {
		event := &models.Event{}
		err := rows.Scan(
			&event.ID,
			&event.EventTypeID,
			&event.EventType.ID,
			&event.EventType.Name,
			&event.CompanyID,
			&event.Company.ID,
			&event.Company.Name,
			&event.VendorID,
			&event.Vendor.ID,
			&event.Vendor.Name,
			&event.EventStatusID,
			&event.EventStatus.ID,
			&event.EventStatus.Name,
			pq.Array(&event.ProposedDates),
			&event.ConfirmedDate,
			&event.Location,
			&event.RejectedRemarks,
			&event.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r *eventRepository) ApproveEvent(id int64, event *models.Event) error {
	query := `
		UPDATE events
		SET event_status_id = $1, confirmed_date = $2
		WHERE id = $3;
	`
	_, err := r.DB.Exec(query, event.EventStatusID, event.ConfirmedDate, id)
	if err != nil {
		return err
	}

	return nil

}

func (r *eventRepository) RejectEvent(id int64, event *models.Event) error {
	query := `
		UPDATE events
		SET event_status_id = $1, rejected_remarks = $2
		WHERE id = $3;
	`
	_, err := r.DB.Exec(query, event.EventStatusID, event.RejectedRemarks, id)
	if err != nil {
		return err
	}

	return nil
}
