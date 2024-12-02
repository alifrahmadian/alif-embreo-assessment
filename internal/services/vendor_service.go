package services

import (
	"github.com/alifrahmadian/alif-embreo-assessment/internal/models"
	r "github.com/alifrahmadian/alif-embreo-assessment/internal/repositories"
)

type VendorService interface {
	GetVendors() ([]*models.Vendor, error)
}

type vendorService struct {
	VendorRepo r.VendorRepository
}

func NewVendorService(vendorRepo r.VendorRepository) VendorService {
	return &vendorService{
		VendorRepo: vendorRepo,
	}
}

func (s *vendorService) GetVendors() ([]*models.Vendor, error) {
	vendors, err := s.VendorRepo.GetVendors()
	if err != nil {
		return nil, err
	}

	return vendors, err
}
