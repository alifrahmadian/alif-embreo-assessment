package dtos

type CreateEventRequest struct {
	EventTypeID   int64   `json:"event_type_id" binding:"required"`
	VendorID      int64   `json:"vendor_id" binding:"required"`
	CompanyID     int64   `json:"company_id" binding:"required"`
	ProposedDates []int64 `json:"proposed_dates" binding:"required"`
	Location      string  `json:"location" binding:"required"`
}

type CreateEventResponse struct {
	ID            int64   `json:"id"`
	CompanyID     int64   `json:"company_id"`
	VendorID      int64   `json:"vendor_id"`
	EventStatusID int64   `json:"event_status_id"`
	EventTypeID   int64   `json:"event_type_id"`
	ProposedDates []int64 `json:"proposed_dates"`
	Location      string  `json:"location"`
	CreatedAt     int64   `json:"created_at"`
}
