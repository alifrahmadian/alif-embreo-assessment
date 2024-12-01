package errors

import "errors"

var (
	ErrProposedDatesLessThanThree = errors.New("you need to propose 3 dates for this event")
	ErrEventTypeRequired          = errors.New("event type is required")
	ErrProposedDatesRequired      = errors.New("proposed dates are required")
	ErrLocationRequired           = errors.New("location is required")
	ErrCompanyRequired            = errors.New("company is required")
	ErrVendorRequired             = errors.New("vendor is requirede")
)
