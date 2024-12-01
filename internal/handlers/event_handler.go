package handlers

import (
	"net/http"
	"time"

	"github.com/alifrahmadian/alif-embreo-assessment/internal/constants"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/handlers/dtos"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/handlers/responses"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/models"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/services"
	"github.com/alifrahmadian/alif-embreo-assessment/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EventHandler struct {
	EventService services.EventService
}

func NewEventHandler(eventService *services.EventService) *EventHandler {
	return &EventHandler{
		EventService: *eventService,
	}
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	var req dtos.CreateEventRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "EventTypeID":
				responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrEventTypeRequired.Error())
				return
			case "VendorID":
				responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrVendorRequired.Error())
				return
			case "CompanyID":
				responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrCompanyRequired.Error())
			case "ProposedDates":
				responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrProposedDatesRequired.Error())
				return
			case "Location":
				responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrLocationRequired.Error())
				return
			}
		}
	}

	event := &models.Event{
		EventTypeID:   req.EventTypeID,
		CompanyID:     req.CompanyID,
		ProposedDates: req.ProposedDates,
		Location:      req.Location,
		EventStatusID: constants.StatusPending,
		CreatedAt:     time.Now().Unix(),
	}

	createdEvent, err := h.EventService.CreateEvent(event)
	if err != nil {
		if err == errors.ErrProposedDatesLessThanThree {
			responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	resp := &dtos.CreateEventResponse{
		ID:            createdEvent.ID,
		CompanyID:     createdEvent.CompanyID,
		VendorID:      createdEvent.VendorID,
		EventStatusID: createdEvent.EventStatusID,
		EventTypeID:   createdEvent.EventTypeID,
		Location:      createdEvent.Location,
		ProposedDates: createdEvent.ProposedDates,
		CreatedAt:     createdEvent.CreatedAt,
	}

	responses.SuccessResponse(c, "Event created successfully!", resp)
}
