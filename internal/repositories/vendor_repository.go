package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/alif-embreo-assessment/internal/models"
)

type VendorRepository interface {
	GetVendors() ([]*models.Vendor, error)
}

type vendorRepository struct {
	DB *sql.DB
}

func NewVendorRepository(db *sql.DB) VendorRepository {
	return &vendorRepository{
		DB: db,
	}
}

func (r *vendorRepository) GetVendors() ([]*models.Vendor, error) {
	query := "SELECT vendors.id, name FROM vendors;"

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var vendors []*models.Vendor

	for rows.Next() {
		vendor := &models.Vendor{}

		err := rows.Scan(&vendor.ID, &vendor.Name)
		if err != nil {
			return nil, err
		}

		vendors = append(vendors, vendor)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return vendors, nil
}
