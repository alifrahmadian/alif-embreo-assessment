package services

import (
	"github.com/alifrahmadian/alif-embreo-assessment/internal/models"
	r "github.com/alifrahmadian/alif-embreo-assessment/internal/repositories"
	"github.com/alifrahmadian/alif-embreo-assessment/pkg/errors"
)

type EventService interface {
	CreateEvent(event *models.Event) (*models.Event, error)
	GetEventByID(id int64) (*models.Event, error)
}

type eventService struct {
	EventRepo r.EventRepository
}

func NewEventService(eventRepo r.EventRepository) EventService {
	return &eventService{
		EventRepo: eventRepo,
	}
}

func (s *eventService) CreateEvent(event *models.Event) (*models.Event, error) {
	if len(event.ProposedDates) < 3 {
		return nil, errors.ErrProposedDatesLessThanThree
	}

	createdEvent, err := s.EventRepo.CreateEvent(event)
	if err != nil {
		return nil, err
	}

	return createdEvent, nil
}

func (s *eventService) GetEventByID(id int64) (*models.Event, error) {
	event, err := s.EventRepo.GetEventByID(id)
	if err != nil {
		return nil, err
	}

	return event, nil
}
