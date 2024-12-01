package models

import "database/sql"

type User struct {
	ID        int64         `json:"id"`
	Username  string        `json:"username"`
	Email     string        `json:"email"`
	Password  string        `json:"-"`
	RoleID    int64         `json:"role_id"`
	Role      Role          `json:"role"`
	CompanyID sql.NullInt64 `json:"company_id"`
	Company   Company       `json:"company"`
	VendorID  sql.NullInt64 `json:"vendor_id"`
	Vendor    Vendor        `json:"vendor"`
}
