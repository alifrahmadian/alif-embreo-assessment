package errors

import "errors"

var (
	ErrProposedDatesLessThanThree = errors.New("you need to propose 3 dates for this event")
	ErrEventTypeRequired          = errors.New("event type is required")
	ErrProposedDatesRequired      = errors.New("proposed dates are required")
	ErrLocationRequired           = errors.New("location is required")
	ErrCompanyRequired            = errors.New("company is required")
	ErrVendorRequired             = errors.New("vendor is requirede")
	ErrEventNotFound              = errors.New("event not found")
	ErrInvalidEventID             = errors.New("invalid event id")
	ErrEventHasBeenRejected       = errors.New("event has been rejected")
	ErrEventHasBeenApproved       = errors.New("event has been approved")
	ErrConfirmedDateRequired      = errors.New("confirmed date is required")
	ErrRejectionRemarksRequired   = errors.New("rejection remarks required")
)
