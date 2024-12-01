package models

type Event struct {
	ID              int64       `json:"id"`
	ProposedDates   []int64     `json:"proposed_dates"`
	ConfirmedDate   int64       `json:"confirmed_date"`
	Location        string      `json:"location"`
	RejectedRemarks string      `json:"rejected_remarks"`
	CreatedAt       int64       `json:"created_at"`
	ConfirmedAt     int64       `json:"confirmed_at"`
	EventTypeID     int64       `json:"event_type_id"`
	EventType       EventType   `json:"event_type"`
	CompanyID       int64       `json:"company_id"`
	Company         Company     `json:"company"`
	VendorID        int64       `json:"vendor_id"`
	Vendor          Vendor      `json:"vendor"`
	EventStatusID   int64       `json:"event_status_id"`
	EventStatus     EventStatus `json:"event_status"`
}
