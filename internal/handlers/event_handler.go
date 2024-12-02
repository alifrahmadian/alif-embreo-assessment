package handlers

import (
	"net/http"
	"strconv"
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

	roleID := c.GetInt64("role_id")
	companyID := c.GetInt64("company_id")

	if roleID == constants.RoleVendor {
		responses.ErrorResponse(c, http.StatusForbidden, errors.ErrForbidden.Error())
		return
	}

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
		CompanyID:     companyID,
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
		return
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

func (h *EventHandler) GetEventByID(c *gin.Context) {
	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrInvalidEventID.Error())
		return
	}

	event, err := h.EventService.GetEventByID(int64(eventID))
	if err != nil {
		if err == errors.ErrEventNotFound {
			responses.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := &dtos.GetEventResponse{
		ID:              event.ID,
		ProposedDates:   event.ProposedDates,
		ConfirmedDate:   event.ConfirmedDate,
		Location:        event.Location,
		RejectedRemarks: event.RejectedRemarks,
		CreatedAt:       event.CreatedAt,
		EventTypeID:     event.EventTypeID,
		EventType: dtos.EventTypeGetEventResponse{
			ID:   event.EventType.ID,
			Name: event.EventType.Name,
		},
		CompanyID: event.CompanyID,
		Company: dtos.CompanyGetEventResponse{
			ID:   event.Company.ID,
			Name: event.Company.Name,
		},
		VendorID: event.VendorID,
		Vendor: dtos.VendorGetEventResponse{
			ID:   event.Vendor.ID,
			Name: event.Vendor.Name,
		},
		EventStatusID: event.EventStatusID,
		EventStatus: dtos.EventStatusGetEventResponse{
			ID:   event.EventStatus.ID,
			Name: event.EventStatus.Name,
		},
	}

	responses.SuccessResponse(c, "get event successful", resp)
}

func (h *EventHandler) GetAllEvents(c *gin.Context) {
	events, err := h.EventService.GetAllEvents()
	if err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := make([]*dtos.GetEventResponse, len(events))

	for i, event := range events {
		responseData[i] = &dtos.GetEventResponse{
			ID:              event.ID,
			ProposedDates:   event.ProposedDates,
			ConfirmedDate:   event.ConfirmedDate,
			Location:        event.Location,
			RejectedRemarks: event.RejectedRemarks,
			CreatedAt:       event.CreatedAt,
			EventTypeID:     event.EventTypeID,
			EventType: dtos.EventTypeGetEventResponse{
				ID:   event.EventType.ID,
				Name: event.EventType.Name,
			},
			CompanyID: event.CompanyID,
			Company: dtos.CompanyGetEventResponse{
				ID:   event.Company.ID,
				Name: event.Company.Name,
			},
			VendorID: event.VendorID,
			Vendor: dtos.VendorGetEventResponse{
				ID:   event.Vendor.ID,
				Name: event.Vendor.Name,
			},
			EventStatusID: event.EventStatusID,
			EventStatus: dtos.EventStatusGetEventResponse{
				ID:   event.EventStatus.ID,
				Name: event.EventStatus.Name,
			},
		}
	}

	responses.SuccessResponse(c, "get events successful", responseData)
}

func (h *EventHandler) ApproveEvent(c *gin.Context) {
	var req dtos.ApproveEventRequest

	roleID := c.GetInt64("role_id")
	if roleID == constants.RoleHR {
		responses.ErrorResponse(c, http.StatusForbidden, errors.ErrForbidden.Error())
		return
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "ConfirmedDate":
				responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrConfirmedDateRequired.Error())
				return
			}
		}
	}

	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrInvalidEventID.Error())
		return
	}

	event := &models.Event{
		EventStatusID: constants.StatusApproved,
		ConfirmedDate: &req.ConfirmedDate,
	}

	err = h.EventService.ApproveEvent(int64(eventID), event)
	if err != nil {
		if err == errors.ErrEventNotFound {
			responses.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		if err == errors.ErrEventHasBeenApproved || err == errors.ErrEventHasBeenRejected {
			responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	responses.SuccessResponse(c, "Event approval successful!", nil)
}

func (h *EventHandler) RejectEvent(c *gin.Context) {
	var req dtos.RejectEventRequest

	roleID := c.GetInt64("role_id")
	if roleID == constants.RoleHR {
		responses.ErrorResponse(c, http.StatusForbidden, errors.ErrForbidden.Error())
		return
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "RejectedRemarks":
				responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrRejectionRemarksRequired.Error())
				return
			}
		}
	}

	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrInvalidEventID.Error())
		return
	}

	event := &models.Event{
		EventStatusID:   constants.StatusRejected,
		RejectedRemarks: &req.RejectedRemarks,
	}

	err = h.EventService.RejectEvent(int64(eventID), event)
	if err != nil {
		if err == errors.ErrEventNotFound {
			responses.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		if err == errors.ErrEventHasBeenApproved || err == errors.ErrEventHasBeenRejected {
			responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	responses.SuccessResponse(c, "Event rejection successful!", nil)
}
